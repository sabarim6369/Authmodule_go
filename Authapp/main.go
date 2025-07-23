package main
import (
	
	"github.com/gin-gonic/gin"
	"github.com/sabarim6369/Authmodule_go/routers"
	"github.com/sabarim6369/Authmodule_go/db"

)
func main(){
	r:=gin.Default();
	db.Dbconnection()
	routers.Authrouters(r)

	r.Run();
}