package reindex

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func init() {
	u := "note"
	p := "1234@note"
	url := "127.0.0.1"
	dbname := "note"
	port := "3306"

	dataSource := u + ":" + p + "@tcp(" + url + ":" + port + ")/" + dbname + "?charset=utf8"
	log.Printf("dataSource:%v", dataSource)

	e, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	err = e.Ping()
	if err != nil {
		panic(err)
	}

	engine = e

	log.Println("connect db successfully!")

	go ScanTweetsTable()
}

//获取db engine
func GetEngine() (*xorm.Engine, error) {
	return engine, engine.Ping()
}

//创建session
func NewSession() (*xorm.Session, error) {
	s := engine.NewSession()
	return s, s.Begin()
}

//session提交
func CommitAndClose(s *xorm.Session) error {
	if s == nil {
		return fmt.Errorf("session nil!")
	}

	defer s.Close()

	if err := s.Commit(); err != nil {
		return err
	}
	return nil
}

//session回滚
func RollBackAndClose(s *xorm.Session) error {
	if s == nil {
		return fmt.Errorf("session nil!")
	}

	defer s.Close()

	if err := s.Rollback(); err != nil {
		return err
	}
	return nil
}
