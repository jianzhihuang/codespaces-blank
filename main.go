package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//helloworld
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, world!")
	})
	r.GET("/hello/:id/:type", func(c *gin.Context) {
		id := c.Param("id")
		type_ := c.Param("type")

		var result string
		switch type_ {
		case "heart":
			result = generateEmoji(id, []string{"❤️", "♡", "💖", "💟", "🎁"})
		case "smile":
			result = generateEmoji(id, []string{"😀", "🤩", "😊", "🙂", "☺️", "😋"})
		case "cry":
			result = generateEmoji(id, []string{"😢", "😭", "😿"})
		case "cat":
			result = generateEmoji(id, []string{"🐈", "😾", "🐱", "😻", "🐱‍🚀"})
		case "dog":
			result = generateEmoji(id, []string{"🐶", "🐕", "🦮", "🐩", "🐕‍🦺"})
		case "pig":
			result = generateEmoji(id, []string{"🐷", "🐽", "🐖", "🐗"})
		default:
			result = fmt.Sprintf("Hello, world! Your ID is %s %s", id, type_)
		}

		c.String(http.StatusOK, result)
	})
	r.GET("/rand/:id", func(c *gin.Context) {
		id := c.Param("id")
		num, err := strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid ID")
			return
		}
		randNum := rand.Intn(num)
		c.String(http.StatusOK, fmt.Sprintf("Random number: %d", randNum))
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func generateEmoji(id string, emojis []string) string {
	count, err := strconv.Atoi(id)
	if err != nil {
		count = 1
	}

	idx := rand.Intn(len(emojis))
	return repeat(emojis[idx], count)
}

func repeat(s string, count int) string {
	var result string
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}
