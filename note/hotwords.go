package note

import (
	"github.com/sea-team/gofound/global"
)

func InsertHotWords(words string) error {
	var ERR error
	session, _ := global.NewSession()

	var hotword global.HotWords
	ext, _ := session.Where("words = ?", words).Get(&hotword)
	if ext {
		hotword.Num += 1
		_, ERR = session.Where("id = ?", hotword.Id).Cols("num").Update(&hotword)
	} else {
		_, ERR = session.Insert(&global.HotWords{
			Words:     words,
			Num:       1,
			CreatedAt: global.NowTimeToString(),
		})
	}
	global.CommitAndClose(session, ERR)

	return ERR
}
