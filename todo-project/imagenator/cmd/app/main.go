package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func generateNewImage(url string, dataDir string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(dataDir + "/image.jpg")
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		return err
	}

	timestamp := time.Now().Add(time.Minute * 60).Unix()
	if err := os.WriteFile(dataDir + "/generated-timestamp.txt", []byte(fmt.Sprint(timestamp)), 0644); err != nil {
		return err
	}

	return nil
}

func getGeneratedTimestamp(dataDir string) (time.Time, error) {
	data, err := os.ReadFile(dataDir + "/generated-timestamp.txt")
	if err != nil {
		return time.Unix(0, 0), err
	}

	timestampUnix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return time.Unix(0, 0), err
	}

	return time.Unix(timestampUnix, 0), nil
}

func main() {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	logger := slog.New(jsonHandler)

	url := "https://picsum.photos/1200"
	dataDir := "/usr/src/app/data"

	router := gin.Default()
	router.StaticFile("/imagenator/image", dataDir+"/image.jpg")

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	go func() {
		for {
			generatedTimestamp, err := getGeneratedTimestamp(dataDir)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					logger.Info("No image generated yet")
				} else {
					logger.Error(err.Error())
				}
			}
			waitTime := time.Until(generatedTimestamp)
			if waitTime > 0 {
				logger.Info("Waiting for " + waitTime.String())
				time.Sleep(waitTime)
			}

			logger.Info("Generating new image")
			if err := generateNewImage(url, dataDir); err != nil {
				logger.Error(err.Error())
			}
		}
	}()

	logger.Info("Starting server at 0.0.0.0:8080")
	if err := router.Run("0.0.0.0:8080"); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
