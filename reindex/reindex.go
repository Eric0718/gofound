package reindex

import (
	"fmt"
	"log"
	"time"

	"github.com/sea-team/gofound/searcher/model"
	"github.com/sea-team/gofound/web/controller"
)

type SearchaCalls struct {
	Id       int64 `xorm:"pk"`
	NoteId   int64
	State    int //0 待入库，1 已入库，2 删
	CreateAt string
	UpdateAt string
	DeleteAt string
}
type Tweets struct {
	Id       int64 `xorm:"pk"`
	UserId   int64
	Context  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}

var InsertUrl = "http://127.0.0.1:5678/api/index?database=default"        //"https://www.petcat.pro/searcher/api/index?database=default"
var RemoveUrl = "http://127.0.0.1:5678/api/index/remove?database=default" //"https://www.petcat.pro/searcher/api/index/remove?database=default"
var database = "default"

func ScanTweetsTable() {
	for {
		time.Sleep(time.Minute * 2)
		log.Println("\nScanTweetsTable start>>>>>>>>>>")
		session, err := NewSession()
		if err != nil {
			panic("nil mysql session")
		}
		var resets []SearchaCalls
		session.Where("state = 0 or state = 2").Find(&resets)
		for _, res := range resets {
			var ctx Tweets
			if res.State == 0 {
				if ok, _ := session.Where("id = ?", res.NoteId).Get(&ctx); !ok {
					continue
				}
			} else {
				ctx.Id = res.NoteId
			}

			err := UpdateSearcherdb(ctx, res.State, database)
			if err != nil {
				log.Println("ScanResetTable:", err)
				continue
			}

			if res.State == 0 {
				res.State = 1
			} else if res.State == 2 {
				res.State = 3
			}
			res.UpdateAt = time.Now().Format("2006-01-02 15:04:05")
			session.Where("id = ?", res.Id).Cols("state,update_at").Update(&res)
		}
		session.Commit()
		session.Close()
		log.Println("ScanTweetsTable end<<<<<<<<<<<<\n")
	}
}

func UpdateSearcherdb(ctx Tweets, option int, dbName string) error {
	//url := InsertUrl
	srv := controller.GetSrv()

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

	/*
		data, _ := json.Marshal(&doc)
		request, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			return fmt.Errorf("insertContext http.NewRequest:%v", err)
		}

		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			return fmt.Errorf("insertContext http error code %v", response.StatusCode)
		}

		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(body))

		var resBody struct {
			State   bool   `json:"state"`
			Message string `json:"message"`
		}

		err = json.Unmarshal(body, &resBody)
		if err != nil {
			return err
		}
		if !resBody.State {
			return fmt.Errorf("message:%v", resBody.Message)
		}
	*/
	return fmt.Errorf("Unkonwn record option %v!", option)
}
