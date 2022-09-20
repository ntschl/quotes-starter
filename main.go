package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quote struct {
	Id     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotesMap = map[int]quote{
	0:  {Id: "ab4b7a65-d2a4-4eb7-a2db-c15bade7bb26", Quote: "Clear is better than clever.", Author: "Ronald McDonald"},
	1:  {Id: "84ca5b5f-38f0-4e00-bcf5-ae916e887690", Quote: "Empty string check!", Author: "Squidward Tentacles"},
	2:  {Id: "b23071f5-e4bf-41a3-b3b1-ed232fa0ffe2", Quote: "Don't panic.", Author: "Oprah Winfrey"},
	3:  {Id: "99fae00d-c5d4-4575-ba59-7e79efaff603", Quote: "A little copying is better than a little dependency.", Author: "Chris Pratt"},
	4:  {Id: "5441f417-1379-4997-80bc-e2eac7523133", Quote: "The bigger the interface, the weaker the abstraction.", Author: "Mary Poppins"},
	5:  {Id: "fc27cfd6-8f29-437f-b951-0a527fa2f7d3", Quote: "With the unsafe package there are no guarantees.", Author: "Rob Dyrdek"},
	6:  {Id: "f05da4ce-398c-48fb-9a54-009ec3304319", Quote: "Reflection is never clear.", Author: "Bobby Hill"},
	7:  {Id: "f947a1cf-8d33-4d6b-b898-5a8bfd5a6dd4", Quote: "Don't just check errors, handle them gracefully.", Author: "Shrek"},
	8:  {Id: "1627de76-c799-4b18-80c7-6151baf0f585", Quote: "Documentation is for users.", Author: "Hermione Granger"},
	9:  {Id: "7ee7ccc8-21f0-4bea-af55-97553cb0d4d4", Quote: "Errors are values.", Author: "Clark Kent"},
	10: {Id: "9d17a91b-3525-4bae-9a34-c4de8155767a", Quote: "Make the zero value useful.", Author: "Drake"},
	11: {Id: "dd815990-0875-48f4-bf78-98ff8397dbed", Quote: "Channels orchestrate; mutexes serialize.", Author: "Yo-Yo Ma"},
	12: {Id: "e3669a09-3d4b-4aec-8b51-a3b3412e0603", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Prince"},
	13: {Id: "1a10287d-e83b-45c5-ba94-f290710da7eb", Quote: "Concurrency is not parallelism.", Author: "Lao Tzu"},
	14: {Id: "2c371688-f482-4c77-943e-89937da93d27", Quote: "Design the architecture, name the components, document the details.", Author: "Tony the Tiger"},
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
