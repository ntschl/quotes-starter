package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type quote struct {
	Id     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotesMap = map[int]quote{
	0:  {Id: uuid.New().String(), Quote: "Clear is better than clever.", Author: "Ronald McDonald"},
	1:  {Id: uuid.New().String(), Quote: "Empty string check!", Author: "Squidward Tentacles"},
	2:  {Id: uuid.New().String(), Quote: "Don't panic.", Author: "Oprah Winfrey"},
	3:  {Id: uuid.New().String(), Quote: "A little copying is better than a little dependency.", Author: "Chris Pratt"},
	4:  {Id: uuid.New().String(), Quote: "The bigger the interface, the weaker the abstraction.", Author: "Mary Poppins"},
	5:  {Id: uuid.New().String(), Quote: "With the unsafe package there are no guarantees.", Author: "Rob Dyrdek"},
	6:  {Id: uuid.New().String(), Quote: "Reflection is never clear.", Author: "Bobby Hill"},
	7:  {Id: uuid.New().String(), Quote: "Don't just check errors, handle them gracefully.", Author: "Shrek"},
	8:  {Id: uuid.New().String(), Quote: "Documentation is for users.", Author: "Hermione Granger"},
	9:  {Id: uuid.New().String(), Quote: "Errors are values.", Author: "Clark Kent"},
	10: {Id: uuid.New().String(), Quote: "Make the zero value useful.", Author: "Drake"},
	11: {Id: uuid.New().String(), Quote: "Channels orchestrate; mutexes serialize.", Author: "Yo-Yo Ma"},
	12: {Id: uuid.New().String(), Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Prince"},
	13: {Id: uuid.New().String(), Quote: "Concurrency is not parallelism.", Author: "Lao Tzu"},
	14: {Id: uuid.New().String(), Quote: "Design the architecture, name the components, document the details.", Author: "Tony the Tiger"},
}

func main() {
	router := gin.Default()
	router.GET("/quotes", GetQuote)
	router.Run("0.0.0.0:8080")
}

func GetRandom() quote {
	key := rand.Intn(len(quotesMap))
	quote := quotesMap[key]
	return quote
}

func GetQuote(c *gin.Context) {
	c.JSON(http.StatusOK, GetRandom())
}

// var quotes = []quote{
// 	{Quote: "Clear is better than clever.", Author: "Ronald McDonald"},
// 	{Quote: "Empty string check!", Author: "Squidward Tentacles"},
// 	{Quote: "Don't panic.", Author: "Oprah Winfrey"},
// 	{Quote: "A little copying is better than a little dependency.", Author: "Chris Pratt"},
// 	{Quote: "The bigger the interface, the weaker the abstraction.", Author: "Mary Poppins"},
// 	{Quote: "With the unsafe package there are no guarantees.", Author: "Rob Dyrdek"},
// 	{Quote: "Reflection is never clear.", Author: "Bobby Hill"},
// 	{Quote: "Don't just check errors, handle them gracefully.", Author: "Shrek"},
// 	{Quote: "Documentation is for users.", Author: "Hermione Granger"},
// 	{Quote: "Errors are values.", Author: "Clark Kent"},
// 	{Quote: "Make the zero value useful.", Author: "Drake"},
// 	{Quote: "Channels orchestrate; mutexes serialize.", Author: "Yo-Yo Ma"},
// 	{Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Prince"},
// 	{Quote: "Concurrency is not parallelism.", Author: "Lao Tzu"},
// 	{Quote: "Design the architecture, name the components, document the details.", Author: "Tony the Tiger"},
// }

// func getQuotes(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, quotes)
// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/quotes", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"quote": getQuote(),
// 			"author": "Jessie & Nate - Gold Team Rules!",
// 		})
// 	})
// 	r.Run() //default's on 8080
// }
