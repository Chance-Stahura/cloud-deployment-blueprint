package main

import (
	//"encoding/json" 	//uses struct(message) to encode(marshal) and decode(unmarshal)
	"time"			//getting the current system time
	"log"
	"github.com/gofiber/fiber/v3"
)

type message struct {
	Message string `json:"message"`
	Timestamp int64 `json:"timestamp"`
}

func serveMessage(c fiber.Ctx) error {
	timestamp := time.Now().Unix()
	m := message{
		Message: "My name is Chance", 
		Timestamp: timestamp,
	}

	return c.JSON(m)
}


func main() {
    app := fiber.New()

	//GET route for the endpoint
	app.Get("/", serveMessage)

	//start server on port 3000
	//listen for incoming HTTP requests
    log.Fatal(app.Listen(":80"))
}