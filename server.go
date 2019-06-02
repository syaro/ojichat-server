package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/greymd/ojichat/generator"
	"google.golang.org/appengine"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/post", createMessage)

	http.Handle("/", router)
	appengine.Main()
}

func createMessage(c *gin.Context) {
	name := c.PostForm("name")
	emojiLevel := c.DefaultPostForm("emoji_level", "4")
	punctiuationLevel := c.DefaultPostForm("punctiuation_level", "1")

	if num, err := strconv.Atoi(emojiLevel); err != nil || num < 1 || num > 10 {
		emojiLevel = "4"
	}

	config := generator.Config{}
	config.TargetName = name
	config.EmojiNum, _ = strconv.Atoi(emojiLevel)
	config.punctiuationLevel, _ = strconv.Atoi(punctiuationLevel)

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"name":               name,
		"emoji_level":        emojiLevel,
		"punctiuation_level": punctiuationLevel,
		"message":            result,
	})
}
