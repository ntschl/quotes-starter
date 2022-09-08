package main

import {
	"github.com/gin-gonic/gin"
	"net/http"
}

func main() {
	r := gin.Default()
	r.GET("/kanye", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Inspirational Wisdom": getQuote(),
			"Team":                 "Jessie & Nate - Gold Team Rules!",
		})
	})
	r.Run() //default's on 8080
}

