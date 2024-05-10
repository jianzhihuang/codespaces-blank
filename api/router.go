package handler

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Handler() *gin.Engine {
	// 初始化 Gin 引擎，并定义路由
	router := gin.Default()

	//helloworld
	// router.StaticFile("/", "./static/index.html")
	router.GET("/hello/:id/:type", func(c *gin.Context) {
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
	router.GET("/rand/:id", handleRandom)
	// 把 Gin 引擎和 HTTP Request/Response 对象传递给 Vercel
	return router
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
func handleRandom(c *gin.Context) {
	id := c.Param("id")
	fileSizeMB, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
		return
	}

	fileName := fmt.Sprintf("%d.txt", fileSizeMB)

	// Check if file already exists
	if _, err := os.Stat(fileName); err == nil {
		content, err := os.ReadFile(fileName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}

		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
		c.Data(http.StatusOK, "application/octet-stream", content)
		return
	}

	// Create new file
	file, err := os.Create(fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
		return
	}
	defer file.Close()

	for i := 0; i < fileSizeMB; i++ {
		data := generateRandomData(1000000) // Generate approximately 1MB of data
		_, err := file.WriteString(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to file"})
			return
		}
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(http.StatusOK, "application/octet-stream", content)

}
func generateRandomData(length int) string {
	b := make([]byte, length/2)
	_, _ = rand.Read(b)
	s := hex.EncodeToString(b)

	var builder strings.Builder
	for i, c := range s {
		if i > 0 && i%30 == 0 {
			builder.WriteString("\n")
		}
		builder.WriteString(string(c))
	}
	return builder.String()
}

// func Listen(w http.ResponseWriter, r *http.Request) {
// 	router := gin.Default()
// 	router.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello from ~~Go!")
// 	})
// 	router.GET("/hello/:id/:type", func(c *gin.Context) {
// 		id := c.Param("id")
// 		type_ := c.Param("type")

// 		var result string
// 		switch type_ {
// 		case "heart":
// 			result = generateEmoji(id, []string{"❤️", "♡", "💖", "💟", "🎁"})
// 		case "smile":
// 			result = generateEmoji(id, []string{"😀", "🤩", "😊", "🙂", "☺️", "😋"})
// 		case "cry":
// 			result = generateEmoji(id, []string{"😢", "😭", "😿"})
// 		case "cat":
// 			result = generateEmoji(id, []string{"🐈", "😾", "🐱", "😻", "🐱‍🚀"})
// 		case "dog":
// 			result = generateEmoji(id, []string{"🐶", "🐕", "🦮", "🐩", "🐕‍🦺"})
// 		case "pig":
// 			result = generateEmoji(id, []string{"🐷", "🐽", "🐖", "🐗"})
// 		default:
// 			result = fmt.Sprintf("Hello, world! Your ID is %s %s", id, type_)
// 		}

// 		c.String(http.StatusOK, result)
// 	})
// 	router.GET("/rand/:id", func(c *gin.Context) {
// 		id := c.Param("id")
// 		println("id:", id)
// 		fileSizeMB, err := strconv.Atoi(id)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter"})
// 			return
// 		}

// 		fileName := fmt.Sprintf("%d.txt", fileSizeMB)

// 		// Check if file already exists
// 		if _, err := os.Stat(fileName); err == nil {
// 			content, err := os.ReadFile(fileName)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
// 				return
// 			}

// 			c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
// 			c.Data(http.StatusOK, "application/octet-stream", content)
// 			return
// 		}

// 		// Create new file
// 		file, err := os.Create(fileName)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create file"})
// 			return
// 		}
// 		defer file.Close()

// 		for i := 0; i < fileSizeMB; i++ {
// 			data := generateRandomData(1000000) // Generate approximately 1MB of data
// 			_, err := file.WriteString(data)
// 			if err != nil {
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write to file"})
// 				return
// 			}
// 		}

// 		content, err := os.ReadFile(fileName)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
// 			return
// 		}

// 		c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
// 		c.Data(http.StatusOK, "application/octet-stream", content)
// 	})
// 	// 把 Gin 引擎和 HTTP Request/Response 对象传递给 Vercel

// 	router.ServeHTTP(w, r)
// }
