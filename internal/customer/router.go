package customer

import "github.com/gin-gonic/gin"

func InitRouter(r gin.RouterGroup) {

	r.GET("/customers", GetAllCustomers)
	r.POST("/customers", SaveCustomer)
	r.GET("/customers/:id", GetCustomer)

}
