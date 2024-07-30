package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	var err error

	router := gin.Default()
	requests := 0

	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	connStr := "postgres://" + dbUser + ":" + dbPass + "@" + dbHost + ":" + dbPort + "/postgres?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for {
		_, err = db.Query("CREATE TABLE IF NOT EXISTS ping_pong(requests int)")
		if err != nil {
			fmt.Println("Error connecting to database: "+ err.Error())
			fmt.Println("Retrying in 10 seconds...")
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}

	var rowCount int
	err = db.QueryRow("SELECT count(requests) FROM ping_pong").Scan(&rowCount)
	if err != nil {
		panic(err)
	}
	if rowCount == 0 {
		_, err = db.Query("INSERT INTO ping_pong(requests) VALUES(0)")
	} else {
		err = db.QueryRow("SELECT requests FROM ping_pong").Scan(&requests)
	}
	if err != nil {
		panic(err)
	}

	router.GET("/pingpong", func(ctx *gin.Context) {
		ctx.String(200, "pong "+fmt.Sprint(requests))
		if ctx.Request.Header.Get("User-Agent") != "Go-http-client/1.1" {
			err := db.QueryRow("UPDATE ping_pong SET requests=requests+1 RETURNING requests").Scan(&requests)
			if err != nil {
				ctx.String(500, err.Error())
			}
		}
	})

	router.Run("0.0.0.0:8080")
}
