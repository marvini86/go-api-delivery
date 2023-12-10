package customer

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/marvini86/go-api-delivery/internal/commom/entity"
)

func GetAllCustomers(c *gin.Context) {
	customers, err := GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, customers)
}

func GetCustomer(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	customer, err := Get(int64(id))

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func SaveCustomer(c *gin.Context) {
	var customer entity.Customer

	if err := c.BindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	saved, err := Save(customer)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, saved)
}
