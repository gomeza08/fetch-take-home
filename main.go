package main

import (
	"fetch-assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = []models.Receipt{}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", processReciepts)
	router.GET("/receipts/:id/points", getReceptPointsByReceiptID)

	router.Run("localhost:8080")

}

func processReciepts(c *gin.Context) {
	var newReceipt models.Receipt

	if err := c.BindJSON(&newReceipt); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "The receipt is invalid"})
		return
	}
	newReceipt.ID = uuid.New()
	newReceipt.Points = newReceipt.CalculatePoints()
	receipts = append(receipts, newReceipt)
	c.IndentedJSON(http.StatusOK, gin.H{"id": newReceipt.ID})
}

func getReceptPointsByReceiptID(c *gin.Context) {
	id := c.Param("id")

	for _, r := range receipts {
		if r.ID.String() == id {
			c.IndentedJSON(http.StatusOK, gin.H{"points": r.Points})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No receipt found for that id"})
}
