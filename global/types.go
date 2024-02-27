package global

//已发布twit入库标记
type SearchaCalls struct {
	Id       int64 `xorm:"pk"`
	NoteId   int64
	State    int //0 待入库，1 已入库，2 删
	CreateAt string
	UpdateAt string
	DeleteAt string
}

//Twit 数据
type Tweets struct {
	Id       int64 `xorm:"pk"`
	UserId   int64
	Context  string
	CreateAt string
	UpdateAt string
	DeleteAt string
}

type HotWords struct {
	Id        int64 `xorm:"pk"`
	Words     string
	Num       int32
	CreatedAt string
}
