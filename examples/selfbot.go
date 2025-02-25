package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/girlpaws/fenrir"
	"github.com/joho/godotenv"
)

func main() {

	var (
		session *fenrir.Session
		err     error
	)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("USER_TOKEN")
	if token == "" {
		data := fenrir.LoginData{
			Email:        os.Getenv("EMAIL"),
			Password:     os.Getenv("PASSWORD"),
			FriendlyName: os.Getenv("FRIENDLY_NAME"),
		}

		session, _, err = fenrir.NewWithLogin(data)
		if err != nil {
			panic(err)
		}

		write := map[string]string{"USER_TOKEN": session.Token}
		err = godotenv.Write(write, ".env")
		if err != nil {
			panic(err)
		}
	} else {
		session = fenrir.New(token)
	}

	// Append a function that handles ready events.
	// We will print some details from the event to the console when we receive EventReady.
	session.AddHandler(func(session *fenrir.Session, r *fenrir.EventReady) {
		fmt.Printf("Ready to process commands from %d user(s) across %d server(s)\n", len(r.Users), len(r.Servers))
	})

	// Append a function that handles message events. We will process any message that is "!ping"
	// and respond with the latency of the websocket connection, if possible.
	session.AddHandler(func(session *fenrir.Session, m *fenrir.EventMessage) {

		// If the message content is not "!ping", ignore the message.
		if m.Content != "!ping" {
			return
		}

		// Simulate a typing event for a second
		err := session.ChannelBeginTyping(m.Channel)
		if err != nil {
			fmt.Println("Error sending typing event: ", err)
		}

		time.Sleep(1 * time.Second)

		// Construct a message to send back to the channel.
		var send fenrir.MessageSend

		// If the last heartbeat ack is zero, we can't do maths to get the latency.
		if !session.LastHeartbeatAck.IsZero() {
			latency := session.LastHeartbeatAck.Sub(session.LastHeartbeatSent)
			send.Content = fmt.Sprintf("Latency: %s", latency)
		} else {
			send.Content = "Latency data unavailable"
		}

		// Send the message to the channel.
		message, err := session.ChannelMessageSend(m.Channel, send)
		if err != nil {
			fmt.Println("Error sending message: ", err)
			return
		}

		fmt.Println("Sent message:", message.Content)
	})

	// Open the session.
	err = session.Open()
	if err != nil {
		panic(err)
	}

	// Wait for a signal; keep the bot running
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close session.
	err = session.Close()
}
