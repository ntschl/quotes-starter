package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type quote struct {
	ID     string `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type id struct {
	ID string `json:"id"`
}

var pool *sql.DB

var quotesMap = map[string]quote{
	"ab4b7a65-d2a4-4eb7-a2db-c15bade7bb26": {ID: "ab4b7a65-d2a4-4eb7-a2db-c15bade7bb26", Quote: "Clear is better than clever.", Author: "Ronald McDonald"},
	"84ca5b5f-38f0-4e00-bcf5-ae916e887690": {ID: "84ca5b5f-38f0-4e00-bcf5-ae916e887690", Quote: "Empty string check!", Author: "Squidward Tentacles"},
	"b23071f5-e4bf-41a3-b3b1-ed232fa0ffe2": {ID: "b23071f5-e4bf-41a3-b3b1-ed232fa0ffe2", Quote: "Don't panic.", Author: "Oprah Winfrey"},
	"99fae00d-c5d4-4575-ba59-7e79efaff603": {ID: "99fae00d-c5d4-4575-ba59-7e79efaff603", Quote: "A little copying is better than a little dependency.", Author: "Chris Pratt"},
	"5441f417-1379-4997-80bc-e2eac7523133": {ID: "5441f417-1379-4997-80bc-e2eac7523133", Quote: "The bigger the interface, the weaker the abstraction.", Author: "Mary Poppins"},
	"fc27cfd6-8f29-437f-b951-0a527fa2f7d3": {ID: "fc27cfd6-8f29-437f-b951-0a527fa2f7d3", Quote: "With the unsafe package there are no guarantees.", Author: "Rob Dyrdek"},
	"f05da4ce-398c-48fb-9a54-009ec3304319": {ID: "f05da4ce-398c-48fb-9a54-009ec3304319", Quote: "Reflection is never clear.", Author: "Bobby Hill"},
	"f947a1cf-8d33-4d6b-b898-5a8bfd5a6dd4": {ID: "f947a1cf-8d33-4d6b-b898-5a8bfd5a6dd4", Quote: "Don't just check errors, handle them gracefully.", Author: "Shrek"},
	"1627de76-c799-4b18-80c7-6151baf0f585": {ID: "1627de76-c799-4b18-80c7-6151baf0f585", Quote: "Documentation is for users.", Author: "Hermione Granger"},
	"7ee7ccc8-21f0-4bea-af55-97553cb0d4d4": {ID: "7ee7ccc8-21f0-4bea-af55-97553cb0d4d4", Quote: "Errors are values.", Author: "Clark Kent"},
	"9d17a91b-3525-4bae-9a34-c4de8155767a": {ID: "9d17a91b-3525-4bae-9a34-c4de8155767a", Quote: "Make the zero value useful.", Author: "Drake"},
	"dd815990-0875-48f4-bf78-98ff8397dbed": {ID: "dd815990-0875-48f4-bf78-98ff8397dbed", Quote: "Channels orchestrate; mutexes serialize.", Author: "Yo-Yo Ma"},
	"e3669a09-3d4b-4aec-8b51-a3b3412e0603": {ID: "e3669a09-3d4b-4aec-8b51-a3b3412e0603", Quote: "Don't communicate by sharing memory, share memory by communicating.", Author: "Prince"},
	"1a10287d-e83b-45c5-ba94-f290710da7eb": {ID: "1a10287d-e83b-45c5-ba94-f290710da7eb", Quote: "Concurrency is not parallelism.", Author: "Lao Tzu"},
	"2c371688-f482-4c77-943e-89937da93d27": {ID: "2c371688-f482-4c77-943e-89937da93d27", Quote: "Design the architecture, name the components, document the details.", Author: "Tony the Tiger"},
}

// u know about main ;)
func main() {
	connectUnixSocket()
	router := gin.Default()
	router.GET("/quotes", getRandomSQL)
	router.GET("/quotes/:id", getQuoteByIDSQL)
	router.POST("/quotes", postQuote)
	router.GET("/firstquote", getFirstQuote)
	router.Run("0.0.0.0:8080")
}

// create new quote, generate uuid, add to map, return only the id
func postQuote(c *gin.Context) {
	var newQuote quote
	var newId id

	if err := c.BindJSON(&newQuote); err != nil {
		return
	}

	// only do the good stuff if the quote and author are valid, otherwise throw 400
	if validateQuote(newQuote) && authenticate(c) {
		newKey := uuid.New()
		newQuote.ID = newKey.String()
		newId.ID = newKey.String()
		quotesMap[newKey.String()] = newQuote
		c.JSON(http.StatusCreated, newId)
	} else if !validateQuote(newQuote) {
		c.JSON(http.StatusBadRequest, "quote and author must be at least 3 characters")
	} else if !authenticate(c) {
		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
	}
}

// check for api key
func authenticate(c *gin.Context) bool {
	headers := c.Request.Header
	headerStrings, exists := headers["X-Api-Key"]
	if exists && headerStrings[0] == "COCKTAILSAUCE" {
		return true
	} else {
		return false
	}
}

// check quote and author for at least 3 chars
func validateQuote(quote quote) bool {
	if len(quote.Author) >= 3 && len(quote.Quote) >= 3 {
		return true
	}
	return false
}

// connectUnixSocket initializes a Unix socket connection pool for
// a Cloud SQL instance of Postgres.
func connectUnixSocket() error {
	mustGetenv := func(k string) string {
		v := os.Getenv(k)
		if v == "" {
			log.Printf("Warning: %s environment variable not set.\n", k)
		}
		return v
	}

	var (
		dbUser         = mustGetenv("DB_USER")              // e.g. 'my-db-user'
		dbPwd          = mustGetenv("NATE_PASSWORD")        // e.g. 'my-db-password'
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance'
		dbName         = mustGetenv("DB_NAME")              // e.g. 'my-database'
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s",
		dbUser, dbPwd, dbName, unixSocketPath)

	// dbPool is the pool of database connections.
	var err error
	pool, err = sql.Open("pgx", dbURI)
	if err != nil {
		return fmt.Errorf("sql.Open: %v", err)
	}

	// ...

	return err
}

func getFirstQuote(c *gin.Context) {
	row := pool.QueryRow("SELECT ID, quote, author FROM quotes LIMIT 1")
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, q)
}

// use id param to search quote map because the keys are ids
func getQuoteByIDSQL(c *gin.Context) {
	id := c.Param("id")
	row := pool.QueryRow(fmt.Sprintf("SELECT id, quote, author FROM quotes where id = '%s'", id))
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, q)
}

// use random order clause in sql statement to randomize select order then limit 1 to grab only the first row
func getRandomSQL(c *gin.Context) {
	row := pool.QueryRow("SELECT ID, quote, author FROM quotes ORDER BY RANDOM() LIMIT 1")
	q := &quote{}
	err := row.Scan(&q.ID, &q.Quote, &q.Author)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, q)
}

// get quote with randomized key and turn into JSON
// func getRandomQuote(c *gin.Context) {
// 	if authenticate(c) {
// 		quote := quotesMap[getRandomKey()]
// 		c.JSON(http.StatusOK, quote)
// 	} else {
// 		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
// 	}
// }

// use id param to search quote map because the keys are ids
// func getQuoteByID(c *gin.Context) {
// 	id := c.Param("id")
// 	quote, exists := quotesMap[id]
// 	if exists && authenticate(c) {
// 		c.JSON(http.StatusOK, quote)
// 	} else if !exists {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "matching id not found"})
// 	} else if !authenticate(c) {
// 		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
// 	}
// }

// make array of keys and pull random key thru random index
// func getRandomKey() string {
// 	keyArray := []string{}
// 	for k := range quotesMap {
// 		keyArray = append(keyArray, k)
// 	}
// 	fmt.Println(keyArray)
// 	index := rand.Intn(len(keyArray))
// 	key := keyArray[index]
// 	return key
// }
