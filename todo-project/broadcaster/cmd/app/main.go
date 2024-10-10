package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"net/http"
	"crypto/tls"

	"github.com/disgoorg/disgo/webhook"
	"github.com/nats-io/nats.go"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	var err error
	var natsConnection *nats.Conn

	natsUrl := os.Getenv("NATS_URL")
	onlyUseStdout := os.Getenv("ONLY_USE_STDOUT") == "true"
	discordWebhookUrl := os.Getenv("DISCORD_WEBHOOK_URL")

    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	var discordWebhookClient webhook.Client
	if onlyUseStdout {
		fmt.Println("Only using stdout")
	} else {
		fmt.Println("Sending messages to Discord")
		discordWebhookClient, err = webhook.NewWithURL(discordWebhookUrl)
		if err != nil {
			panic(err)
		}
	}

	for {
		natsConnection, err = nats.Connect(natsUrl)
		if err == nil {
			break
		}
		fmt.Println("Failed to connect to NATS, retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}

	fmt.Println("Connected to NATS")

	natsConnection.QueueSubscribe("todos", "todos_handlers", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))

		var messageTodo Todo
		if err := json.Unmarshal(m.Data, &messageTodo); err != nil {
			panic(err)
		}

		fmt.Printf("Todo: %s\n", messageTodo.Title)

		if !onlyUseStdout {
			_, err := discordWebhookClient.CreateContent(fmt.Sprintf("A todo was created: %s", messageTodo.Title))
			if err != nil {
				panic(err)
			}
			fmt.Println("Sent message to Discord")
		}
	})

	fmt.Println("Listening for messages...")
	for {
		time.Sleep(5 * time.Second)
	}
}
