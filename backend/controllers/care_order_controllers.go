package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCareOrder(c *gin.Context) {
    var orderID = c.Param("id")
    id, err := strconv.Atoi(orderID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID format"})
        return
    }

    // NEED TO CREATE THE MODEL FIST
    // var careOrder = models
}
