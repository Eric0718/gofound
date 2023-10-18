package reindex

import (
	"testing"
)

func Test_UpdateSearcherdb(t *testing.T) {
	var ctx Tweets
	ctx.Id = 12321414
	ctx.Context = "这还是个测试3"

	err := UpdateSearcherdb(ctx, 2)
	if err != nil {
		t.Errorf("error:%v", err)
	}
}
