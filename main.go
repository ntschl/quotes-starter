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

// hoisted data connection pool vartiable for accessibility
var pool *sql.DB

// u know about main ;)
func main() {
	connectUnixSocket()
	router := gin.Default()
	router.GET("/quotes", getRandomQuote)
	router.GET("/quotes/:id", getQuoteByID)
	router.DELETE("/quotes/:id", deleteQuote)
	router.POST("/quotes", postQuote)
	router.GET("/firstquote", getFirstQuote)
	router.Run("0.0.0.0:8080")
}

// connectUnixSocket initializes a Unix socket connection pool (?) for
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
		unixSocketPath = mustGetenv("INSTANCE_UNIX_SOCKET") // e.g. '/cloudsql/project:region:instance' AKA HOST NAME
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

// post new quote to database
func postQuote(c *gin.Context) {
	q := &quote{}
	var newID id
	newID.ID = uuid.New().String()

	if err := c.BindJSON(&q); err != nil {
		return
	}

	if validateQuote(*q) && authenticate(c) {
		sqlString := "INSERT INTO quotes (id, quote, author) VALUES ($1, $2, $3)"
		_, err := pool.Exec(sqlString, &newID.ID, &q.Quote, &q.Author)
		if err != nil {
			fmt.Println("Something's wrong!")
		}
		c.JSON(http.StatusCreated, newID)
	} else if !validateQuote(*q) {
		c.JSON(http.StatusBadRequest, "quote and author must be at least 3 characters")
	} else if !authenticate(c) {
		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
	}
}

// grab the first quote in the database - made for testing connection to database
func getFirstQuote(c *gin.Context) {
	if authenticate(c) {
		row := pool.QueryRow("SELECT ID, quote, author FROM quotes LIMIT 1")
		q := &quote{}
		err := row.Scan(&q.ID, &q.Quote, &q.Author)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, q)
	} else {
		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
	}
}

// use id param to search quote map because the keys are ids - invalid id throws not found
func getQuoteByID(c *gin.Context) {
	if authenticate(c) {
		id := c.Param("id")
		row := pool.QueryRow(fmt.Sprintf("SELECT id, quote, author FROM quotes where id = '%s'", id))
		q := &quote{}
		err := row.Scan(&q.ID, &q.Quote, &q.Author)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "matching id not found"})
		} else {
			c.JSON(http.StatusOK, q)
		}
	} else {
		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
	}
}

// use random order clause in sql statement to randomize select order then limit 1 to grab only the first row
func getRandomQuote(c *gin.Context) {
	if authenticate(c) {
		row := pool.QueryRow("SELECT ID, quote, author FROM quotes ORDER BY RANDOM() LIMIT 1")
		q := &quote{}
		err := row.Scan(&q.ID, &q.Quote, &q.Author)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, q)
	} else {
		c.JSON(http.StatusUnauthorized, "you ain't authorized!")
	}
}

// delete quote from database by id
func deleteQuote(c *gin.Context) {
	if authenticate(c) {
		id := c.Param("id")
		_, err := pool.Exec(fmt.Sprintf("DELETE FROM quotes WHERE id = '%s'", id))
		if err != nil {
			fmt.Println("Something's wrong!")
		}
		c.JSON(http.StatusNoContent, "(~o.o)~ you shouldn't see this! ~(o.o~)")
	} else {
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
