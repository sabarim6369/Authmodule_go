package routers
import (
	"github.com/gin-gonic/gin"
	"github.com/sabarim6369/Authmodule_go/controller"
)
func Authrouters(c *gin.Engine){
	c.POST("/create",controller.CreateUser)
	c.POST("/update",controller.UpdateUser)
	c.DELETE("/delete/:id", controller.DeleteUser) // pass ID in URL
	c.GET("/users", controller.GetAllUsers)
}