package customerInfo

import (
	"context"
	"github.com/PYxy/go-web/internal/customer-app/store"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

func (i *InfoController) Create(c *gin.Context) {
	var info *store.Customer

	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "400",
			"message": "参数异常",
		})
		return
	}
	v := validator.New()
	err := v.Struct(info)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"code":    "400",
			"message": "参数不合法",
		})
		return
	}
	ctx, canncel := context.WithTimeout(context.Background(), time.Second*5)
	defer canncel()
	//操作数据库
	if err = i.srv.CustomerInfoSrv().Create(ctx, info); err == nil {
		c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "数据保存成功",
		})
	} else {
		c.JSON(http.StatusOK, map[string]string{
			"code":    "0",
			"message": "数据保存失败",
		})
	}

}
