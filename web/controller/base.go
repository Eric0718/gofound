package controller

import (
	"log"

	"github.com/sea-team/gofound/global"

	"github.com/sea-team/gofound/note"
	"github.com/sea-team/gofound/searcher/model"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	ResponseSuccessWithData(c, "Welcome to GoFound")
}

// Query 查询
func Query(c *gin.Context) {
	var request = &model.SearchRequest{
		Database: c.Query("database"),
	}
	if err := c.ShouldBind(&request); err != nil {
		ResponseErrorWithMsg(c, err.Error())
		return
	}

	//添加搜索词热度
	if err := note.InsertHotWords(request.Query); err != nil {
		log.Println("InsertHotWords err:", err)
	}

	//调用搜索
	r, err := srv.Base.Query(request)
	if err != nil {
		ResponseErrorWithMsg(c, err.Error())
	} else {
		ResponseSuccessWithData(c, r)
	}
}

// GC 释放GC
func GC(c *gin.Context) {
	srv.Base.GC()
	ResponseSuccess(c)
}

// Status 获取服务器状态
func Status(c *gin.Context) {
	r := srv.Base.Status()
	ResponseSuccessWithData(c, r)
}

func GetHotWords(c *gin.Context) {
	session, _ := global.NewSession()

	hotwords := make([]global.HotWords, 0)

	session.Where("created_at > ?", global.TodayMidnight().Format(global.TimeLayout)).
		Desc("num").Limit(10, 0).Find(&hotwords)

	ResponseSuccessWithData(c, hotwords)
}
