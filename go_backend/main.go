package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Customer struct {
	ID          int    `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	PhoneNumber string `json:"phonenumber"`
}

func main() {

	fmt.Println("Hello World HELLO")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	PORT := os.Getenv("PORT")

	app := fiber.New()
	customers := []Customer{}
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello World"})
	})

	app.Post("/api/add_customer", func(c *fiber.Ctx) error {
		customer := &Customer{}
		if err := c.BodyParser(customer); err != nil {
			return err
		}
		customer.ID = len(customers) + 1
		customers = append(customers, *customer)
		// return c.Status(201).JSON(fiber.Map{
		// 	"message": fmt.Sprintf("Customer %s created successfully", customer.FirstName),
		// })
		return c.Status(201).JSON(customer)
	})
	app.Listen(":" + PORT)
}
