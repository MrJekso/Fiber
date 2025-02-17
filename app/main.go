package main

import (
	"log"
	"os"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/MrJekso/test/hundler"
)

func main() {
	
	if len(os.Args) != 4 {
		log.Fatal("\n[ * ] error: please enter your login, password and database name\n[ * ] example run: app user password dbname")
		panic("")
	}

	var client hundler.Client
	client.Init(os.Args[1],os.Args[2],os.Args[3])
	client.InitConnectDB()
	//client.CreateTable()
	
	app := fiber.New()

	sms := fmt.Sprintf("%s:%s:%s",client.USER,client.PASS,client.DBNAME)
	//init
	app.Get("/init", func (c *fiber.Ctx) error {
		return c.SendString(sms)
	})
	//get
	app.Get("/tasks", func (c *fiber.Ctx) error {
		return c.SendString(sms)
	})
	
	/*
	//post
	app.Post("/tasks", func (c *fiber.Ctx) error {
        return c.SendString(sms)
    })

	//update
	app.Put("/tasks/:id", func (c *fiber.Ctx) error {
        return c.SendString(sms)
    })

	//delete
	app.Delete("/tasks/:id", func (c *fiber.Ctx) error {
        return c.SendString(sms)
    })
	*/

    log.Fatal(app.Listen(":3000"))
}