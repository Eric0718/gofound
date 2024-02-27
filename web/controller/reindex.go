package controller

import (
	"fmt"
	"log"
	"time"

	"github.com/sea-team/gofound/global"
	"github.com/sea-team/gofound/searcher/model"
)

var InsertUrl = "http://127.0.0.1:5678/api/index?database=default"        //"https://www.petcat.pro/searcher/api/index?database=default"
var RemoveUrl = "http://127.0.0.1:5678/api/index/remove?database=default" //"https://www.petcat.pro/searcher/api/index/remove?database=default"
var database = "default"

func ScanTweetsTable() {
	for {
		time.Sleep(time.Minute * 2)
		log.Println("\nScanTweetsTable start>>>>>>>>>>")
		session, err := global.NewSession()
		if err != nil {
			panic("nil mysql session")
		}
		var resets []global.SearchaCalls
		session.Where("state = 0 or state = 2").Find(&resets)
		for _, res := range resets {
			var ctx global.Tweets
			if res.State == 0 {
				if ok, _ := session.Where("id = ?", res.NoteId).Get(&ctx); !ok {
					continue
				}
			} else {
				ctx.Id = res.NoteId
			}

			err := updateSearcherDB(ctx, res.State, database)
			if err != nil {
				log.Println("ScanResetTable:", err)
				continue
			}

			if res.State == 0 {
				res.State = 1
			} else if res.State == 2 {
				res.State = 3
			}
			res.UpdateAt = global.NowTimeToString()
			session.Where("id = ?", res.Id).Cols("state,update_at").Update(&res)
		}
		session.Commit()
		session.Close()
		log.Println("ScanTweetsTable end<<<<<<<<<<<<\n")
	}
}

func updateSearcherDB(ctx global.Tweets, option int, dbName string) error {
	//url := InsertUrl
	srv := GetSrv()

	if option == 0 {
		var doc model.IndexDoc
		doc.Id = uint32(ctx.Id)
		ws := srv.Word.WordCut(ctx.Context)
		var txt string
		for _, s := range ws {
			if len(s) > 0 {
				txt = txt + s
			}
		}
		log.Println("index====", ctx.Id, "txt======", txt)
		if len(txt) > 0 {
			doc.Text = txt
			data := make(map[string]interface{})
			data["context"] = ctx.Context
			doc.Document = data
			return srv.Index.AddIndex(dbName, &doc)
		}
	} else if option == 2 {
		//url = RemoveUrl
		log.Println("remove id====", ctx.Id)
		rmd := &model.RemoveIndexModel{Id: uint32(ctx.Id)}
		return srv.Index.RemoveIndex(dbName, rmd)
	}

	return fmt.Errorf("Unkonwn record option %v!", option)
}
