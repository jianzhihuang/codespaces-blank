package handler

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// func Handler() *gin.Engine {
// 	// 初始化 Gin 引擎，并定义路由
// 	router := gin.Default()

// 	//helloworld
// 	// router.StaticFile("/", "./static/index.html")
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
// 	router.GET("/rand/:id", handleRandom)
// 	// 把 Gin 引擎和 HTTP Request/Response 对象传递给 Vercel
// 	return router
// }

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

	// Send a GET request to the specified URL
	resp, err := http.Get(fmt.Sprintf("https://rustint.shuttleapp.rs/rand/%d", fileSizeMB))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}
	defer resp.Body.Close()

	// Set the headers for the response
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.DataFromReader(http.StatusOK, resp.ContentLength, resp.Header.Get("Content-Type"), resp.Body, nil)
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

// Handler 是导出的函数，将由 Vercel 执行
func Handler(w http.ResponseWriter, r *http.Request) {
	router := gin.Default()

	// Serve static files from the "public" directory
	router.StaticFile("/", "./public/index.html")

	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Vercel Gin!",
		})
	})
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
	// 将 Gin 的处理器适配到标准 net/http
	router.ServeHTTP(w, r)
}
