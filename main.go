package main

import (
	"log"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// create a new collector
	c := colly.NewCollector()

	//username and password
	username := os.Getenv("LOOM_USERNAME")
	password := os.Getenv("LOOM_PW")

	// authenticate
	post_err := c.Post("https://loom.com/login", map[string]string{"username": username, "password": password})
	if post_err != nil {
		log.Fatal(post_err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// start scraping
	c.Visit("http://loom.com/my-videos/youtube")
}
