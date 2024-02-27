package controller

import (
	"testing"

	"github.com/sea-team/gofound/global"
)

func Test_UpdateSearcherdb(t *testing.T) {
	var ctx global.Tweets
	ctx.Id = 12321414
	ctx.Context = "这还是个测试3"

	err := updateSearcherDB(ctx, 2, "default")
	if err != nil {
		t.Errorf("error:%v", err)
	}
}
