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
	punctuationLevel := c.DefaultPostForm("punctuation_level", "1")

	if num, err := strconv.Atoi(emojiLevel); err != nil || 0 > num || num > 10 {
		emojiLevel = "4"
	}

	if num, err := strconv.Atoi(punctuationLevel); err != nil || 0 > num || num > 3 {
		punctuationLevel = "0"
	}

	config := generator.Config{}
	config.TargetName = name
	config.EmojiNum, _ = strconv.Atoi(emojiLevel)
	// Typo? Punctiuation -> PunctuationLevel
	config.PunctiuationLebel, _ = strconv.Atoi(punctuationLevel)

	result, err := generator.Start(config)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, gin.H{
		"name":              name,
		"emoji_level":       emojiLevel,
		"punctuation_level": punctuationLevel,
		"message":           result,
	})
}
