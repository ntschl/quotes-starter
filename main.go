package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

type quote struct {
	Id     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotesMap = map[string]quote{
	"ab4b7a65-d2a4-4eb7-a2db-c15bade7bb26": {Id: "ab4b7a65-d2a4-4eb7-a2db-c15bade7bb26", Quote: "Clear is better than clever.", Author: "Ronald McDonald"},
	"84ca5b5f-38f0-4e00-bcf5-ae916e887690": {Id: "84ca5b5f-38f0-4e00-bcf5-ae916e887690", Quote: "Empty string check!", Author: "Squidward Tentacles"},
	"b23071f5-e4bf-41a3-b3b1-ed232fa0ffe2": {Id: "b23071f5-e4bf-41a3-b3b1-ed232fa0ffe2", Quote: "Don't panic.", Author: "Oprah Winfrey"},
	"99fae00d-c5d4-4575-ba59-7e79efaff603": {Id: "99fae00d-c5d4-4575-ba59-7e79efaff603", Quote: "A little copying is better than a little dependency.", Author: "Chris Pratt"},
	"5441f417-1379-4997-80bc-e2eac7523133": {Id: "5441f417-1379-4997-80bc-e2eac7523133", Quote: "The bigger the interface, the weaker the abstraction.", Author: "Mary Poppins"},
	"fc27cfd6-8f29-437f-b951-0a527fa2f7d3": {Id: "fc27cfd6-8f29-437f-b951-0a527fa2f7d3", Quote: "With the unsafe package there are no guarantees.", Author: "Rob Dyrdek"},
	"f05da4ce-398c-48fb-9a54-009ec3304319": {Id: "f05da4ce-398c-48fb-9a54-009ec3304319", Quote: "Reflection is never clear.", Author: "Bobby Hill"},
	"f947a1cf-8d33-4d6b-b898-5a8bfd5a6dd4": {Id: "f947a1cf-8d33-4d6b-b898-5a8bfd5a6dd4", Quote: "Don't just check errors, handle them gracefully.", Author: "Shrek"},
	"1627de76-c799-4b18-80c7-6151baf0f585": {Id: "1627de76-c799-4b18-80c7-6151baf0f585", Quote: "Documentation is for users.", Author: "Hermione Granger"},
	"7ee7ccc8-21f0-4bea-af55-97553cb0d4d4": {Id: "7ee7ccc8-21f0-4bea-af55-97553cb0d4d4", Quote: "Errors are values.", Author: "Clark Kent"},
	"9d17a91b-3525-4bae-9a34-c4de8155767a": {Id: "9d17a91b-3525-4bae-9a34-c4de8155767a", Quote: "Make the zero value useful.", Author: "Drake"},
	"dd815990-0875-48f4-bf78-98ff8397dbed": {Id: "dd815990-0875-48f4-bf78-98ff8397dbed", Quote: "Channels orchestrate; mutexes serialize.", Author: "Yo-Yo Ma"},
	"e3669a09-3d4b-4aec-8b51-a3b3412e0603": {Id: "e3669a09-3d4b-4aec-8b51-a3b3412e0603", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Prince"},
	"1a10287d-e83b-45c5-ba94-f290710da7eb": {Id: "1a10287d-e83b-45c5-ba94-f290710da7eb", Quote: "Concurrency is not parallelism.", Author: "Lao Tzu"},
	"2c371688-f482-4c77-943e-89937da93d27": {Id: "2c371688-f482-4c77-943e-89937da93d27", Quote: "Design the architecture, name the components, document the details.", Author: "Tony the Tiger"},
}

func main() {
	router := gin.Default()
	router.GET("/quotes", getRandomQuote)
	router.GET("/quotes/:id", getQuoteByID)
	router.Run("0.0.0.0:8080")
}

// get quote with randomized key and turn into JSON
func getRandomQuote(c *gin.Context) {
	quote := quotesMap[getRandomKey()]
	c.JSON(http.StatusOK, quote)
}

// make array of keys and pull random key thru random index
func getRandomKey() string {
	keyArray := []string{}
	for k := range quotesMap {
		keyArray = append(keyArray, k)
	}
	fmt.Println(keyArray)
	index := rand.Intn(len(keyArray))
	key := keyArray[index]
	return key
}

// use id param to search quote map bevcause the keys are ids
func getQuoteByID(c *gin.Context) {
	id := c.Param("id")
	quote, exists := quotesMap[id]
	if exists {
		c.JSON(http.StatusOK, quote)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "matching id not found"})
}
