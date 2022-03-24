package model

type Profile struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Gender     string  `json:"gender"`
	Birthdate  string  `json:"birthdate"`
	Email      string  `json:"email"`
	Phone      float64 `json:"phone"`
	Salt       string  `json:"salt"`
	Hash       string  `json:"hash"`
	Created    string  `json:"created"`
	Intro      string  `json:"intro"`
	AvatarS    string  `json:"avatars"`
	AvatarL    string  `json:"avatarl"`
	PostCount  string  `json:"postCount"`
	PhotoCount string  `json:"photoCount"`
}

type Post struct {
	Id       int     `json:"id"`
	UserId   int     `json:"userId"`
	Created  string  `json:"created"`
	Tags     string  `json:"tags"`
	Content  string  `json:"content"`
	AtchType string  `json:"atchType"`
	AtchId   int     `json:"atchId"`
	AtchUrl  string  `json:"atchUrl"`
	Reaction []int64 `json:"reaction"`
	CmtCount int     `json:"cmtCount"`
}

type Comment struct {
	Id       int        `json:"id"`
	UserId   int        `json:"userId"`
	PostId   int        `json:"postId"`
	ParentId int        `json:"parentId"`
	Content  string     `json:"content"`
	Created  string     `json:"created"`
	Children []*Comment `json:"children,omitempty"`
}

type Reaction struct {
	UserId int    `json:"userId"`
	PostId int    `json:"postId"`
	T      string `json:"type"`
}

type Relationship struct {
	User1   int    `json:"user1"`
	User2   int    `json:"user2"`
	Created string `json:"created"`
	T       string `json:"type"`
	Other   string `json:"other"`
}

type Notification struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
	T          string `json:"type"`
	Created    string `json:"created"`
	FromUserId int    `json:"fromUserId"`
	PostId     int    `json:"postId"`
	CmtId      int    `json:"cmtId"`
}

type Album struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	Descr   string `json:"descr"`
	Created string `json:"created"`
}

type Photo struct {
	Id      int    `json:"id"`
	UserId  int    `json:"userId"`
	AlbumId int    `json:"albumId"`
	Url     string `json:"url"`
	Created string `json:"created"`
}

type ShortInfo struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	AvatarS string `json:"avatars"`
}
