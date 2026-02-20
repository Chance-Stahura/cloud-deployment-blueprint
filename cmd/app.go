package main

import (
	//"encoding/json" 	//uses struct(message) to encode(marshal) and decode(unmarshal)
	//"net/http"		//functionality for both HTTP clients/servers
	//"time"			//getting the current system time
	"log"
	"github.com/gofiber/fiber/v3"
)


//type message struct {
//	message string `json:"message"`
//	timestamp time.Time `json:"timestamp"`
//}


func main() {
    app := fiber.New()

	//GET route for the endpoint
    app.Get("/", func(c fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

	//start server on port 3000
    log.Fatal(app.Listen(":3000"))

}