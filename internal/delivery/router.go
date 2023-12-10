package delivery

import "github.com/gin-gonic/gin"

func InitRouter(r gin.RouterGroup) {

	r.GET("/deliveries", GetAllDeliveries)
	r.POST("/deliveries", SaveDelivery)
	r.GET("/deliveries/:id", GetDelivery)

}
