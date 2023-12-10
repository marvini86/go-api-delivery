package delivery

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marvini86/go-api-delivery/internal/commom/entity"
)

func GetAllDeliveries(c *gin.Context) {
	deliveries, err := GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, deliveries)
}

func GetDelivery(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	delivery, err := Get(int64(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, delivery)
}

func SaveDelivery(c *gin.Context) {
	var delivery entity.Delivery

	if err := c.BindJSON(&delivery); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	saved, err := Save(delivery)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, saved)
}
