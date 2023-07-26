package customer_app

import (
	"github.com/PYxy/go-web/internal/customer-app/controller/v1/customerInfo"
	"github.com/PYxy/go-web/internal/customer-app/store/mysql"
	"github.com/gin-gonic/gin"
)

// 定义路由
func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installController(g)
}

func installMiddleware(g *gin.Engine) {

}

func installController(g *gin.Engine) {
	storeMysql, _ := mysql.GetMySQLFactoryOr(nil)
	v1 := g.Group("/v1")
	{
		info := v1.Group("/info")
		{
			infoController := customerInfo.NewCustomerInfoController(storeMysql)
			//{"status":0,"name":"小白","password":"123456","email":"as1206159854@163.com","phone":"13719088025","totalPolicy":0,"hobbySlice":["rap","sing"],"hobby":""}
			info.POST("/create", infoController.Create)
		}
	}
}
