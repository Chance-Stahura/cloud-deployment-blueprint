package main

import (
	"encoding/json" 	//uses struct(message) to encode(marshal) and decode(unmarshal)
	"time"			//getting the current system time
	"log"
	"github.com/gofiber/fiber/v3"
)

type message struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

func serveMessage(c fiber.Ctx) error {
	m := message{
		Message: "My name is Chance", 
		Timestamp: time.Now().Unix(),
	}

	c.Set("Content-Type", "application/json")
	data, _ := json.MarshalIndent(m, "", "  ")
	return c.SendString(string(data))
}


func main() {
    app := fiber.New()

	//GET route for the endpoint
	app.Get("/", serveMessage)

	//start server on port 3000
	//listen for incoming HTTP requests
    log.Fatal(app.Listen(":3000"))
}