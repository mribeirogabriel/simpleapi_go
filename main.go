package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "hello world!")
}
func postMain(c *gin.Context) {
	jsonData, err := c.GetRawData()
	jsonString := string(jsonData)
	if err != nil {

	}
	c.IndentedJSON(http.StatusOK, jsonString)
}

func main() {
	router := gin.Default()
	router.GET("/", getMain)
	router.POST("/echo", postMain)
	router.Run()
}
