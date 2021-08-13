package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Rate        float32 `json:"rate"`
	Image       string  `json:"image"`
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/products", func(c *gin.Context) {
		err := initDatabase()

		if err != nil {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Initializing DB is complete"})
	})

	router.GET("/products", func(c *gin.Context) {
		products, err := getAllData()

		if err != nil {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, products)
	})

	router.GET("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		products, err := getDataById(id)

		if err != nil {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, products[0])
	})

	router.DELETE("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		err := deleteDataById(id)

		if err != nil {
			c.IndentedJSON(http.StatusForbidden, gin.H{"message": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusCreated, gin.H{"message": "Done"})
	})

	router.PUT("/products/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))

		var replacementProduct product
		putError := c.BindJSON(&replacementProduct)

		if putError != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong information"})
			return
		}

		products, err := updateProductById(id, replacementProduct)

		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "wrong information"})
			return
		}

		c.IndentedJSON(http.StatusOK, products[0])
	})

	router.Run()
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
