package main
import (
  "os"
  "github.com/gofiber/fiber/v2"
)
func main(){ port:=os.Getenv("PORT"); if port==""{port="8080"}; app:=fiber.New(); app.Get("/", func(c *fiber.Ctx) error { return c.SendString("axhub go fiber ok")}); app.Listen("0.0.0.0:"+port)}
