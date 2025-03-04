package router

import (
	"github.com/gin-gonic/gin"
	"github.com/sea-team/gofound/web/controller"
)

// InitIndexRouter 索引路由
func InitIndexRouter(Router *gin.RouterGroup) {

	indexRouter := Router.Group("index")
	{
		indexRouter.POST("", controller.AddIndex)           // 添加单条索引
		indexRouter.POST("batch", controller.BatchAddIndex) // 批量添加索引
		//indexRouter.POST("remove", controller.RemoveIndex)  // 删除索引
	}
}
