package models

type User struct {
	Uid        int    `json:"uid" gorm:"primary_key;AUTO_INCREMENT"`
	Username   string `json:"username" gorm:"type:varchar(20);not null;unique"`
	Password   string `json:"password" gorm:"type:varchar(20);not null"`
	Email      string `json:"email" gorm:"type:varchar(20);not null;unique"`
	HomeUrl    string `json:"homeUrl" gorm:"type:varchar(20);not null"`
	ScreenName string `json:"screenName" gorm:"type:varchar(20);not null"`
	Created    int    `json:"created" gorm:"type:int(11);not null"`
	Activated  int    `json:"activated" gorm:"type:int(11);not null"`
	Logged     int    `json:"logged" gorm:"type:int(11);not null"`
	GroupName  string `json:"groupName" gorm:"type:varchar(20);not null"`
}

type RelationShip struct {
	Cid int `json:"cid" gorm:"primary_key;AUTO_INCREMENT"`
	Mid int `json:"mid" gorm:"type:int(11);not null"`
}

type Options struct {
	Name        string `json:"name" gorm:"primary_key;type:varchar(20);not null"`
	Value       string `json:"value" gorm:"type:varchar(20);not null"`
	Description string `json:"description" gorm:"type:varchar(20);not null"`
}

type Meta struct {
	Mid         int    `json:"mid" gorm:"primary_key;AUTO_INCREMENT"`
	Name        string `json:"name" gorm:"type:varchar(20);not null"`
	Slug        string `json:"slug" gorm:"type:varchar(20);not null"`
	Type        string `json:"type" gorm:"type:varchar(20);not null"`
	ContentType string `json:"contentType" gorm:"type:varchar(20);not null"`
	Description string `json:"description" gorm:"type:varchar(20);not null"`
	Sort        int    `json:"sort" gorm:"type:int(11);not null"`
	Parent      int    `json:"parent" gorm:"type:int(11);not null"`
}

type Log struct {
	Id       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Action   string `json:"action" gorm:"type:varchar(20);not null"`
	Data     string `json:"data" gorm:"type:varchar(20);not null"`
	AuthorId int    `json:"authorId" gorm:"type:int(11);not null"`
	Ip       string `json:"ip" gorm:"type:varchar(20);not null"`
	Created  int    `json:"created" gorm:"type:int(11);not null"`
}

type Content struct {
	Cid          int    `json:"cid" gorm:"primary_key;AUTO_INCREMENT"`
	Title        string `json:"title" gorm:"type:varchar(20);not null"`
	TitlePic     string `json:"titlePic" gorm:"type:varchar(20);not null"`
	Slug         string `json:"slug" gorm:"type:varchar(20);not null"`
	Created      int    `json:"created" gorm:"type:int(11);not null"`
	Modified     int    `json:"modified" gorm:"type:int(11);not null"`
	Content      string `json:"content" gorm:"type:varchar(20);not null"`
	AuthorId     int    `json:"authorId" gorm:"type:int(11);not null"`
	Type         string `json:"type" gorm:"type:varchar(20);not null"`
	Status       string `json:"status" gorm:"type:varchar(20);not null"`
	Tags         string `json:"tags" gorm:"type:varchar(20);not null"`
	Categories   string `json:"categories" gorm:"type:varchar(20);not null"`
	Hits         int    `json:"hits" gorm:"type:int(11);not null"`
	CommentsNum  int    `json:"commentsNum" gorm:"type:int(11);not null"`
	AllowComment int    `json:"allowComment" gorm:"type:int(11);not null"`
	AllowPing    int    `json:"allowPing" gorm:"type:int(11);not null"`
	AllowFeed    int    `json:"allowFeed" gorm:"type:int(11);not null"`
}

type Comment struct {
	Coid     int    `json:"coid" gorm:"primary_key;AUTO_INCREMENT"`
	Cid      int    `json:"cid" gorm:"type:int(11);not null"`
	Created  int    `json:"created" gorm:"type:int(11);not null"`
	Author   string `json:"author" gorm:"type:varchar(20);not null"`
	AuthorId int    `json:"authorId" gorm:"type:int(11);not null"`
	OwnerId  int    `json:"ownerId" gorm:"type:int(11);not null"`
	Mail     string `json:"mail" gorm:"type:varchar(20);not null"`
	Url      string `json:"url" gorm:"type:varchar(20);not null"`
	Ip       string `json:"ip" gorm:"type:varchar(20);not null"`
	Agent    string `json:"agent" gorm:"type:varchar(20);not null"`
	Type     string `json:"type" gorm:"type:varchar(20);not null"`
	Status   string `json:"status" gorm:"type:varchar(20);not null"`
	Parent   int    `json:"parent" gorm:"type:int(11);not null"`
	Content  string `json:"content" gorm:"type:varchar(20);not null"`
}

type Attachment struct {
	Id       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Fname    string `json:"fname" gorm:"type:varchar(20);not null"`
	Ftype    string `json:"ftype" gorm:"type:varchar(20);not null"`
	Fkey     string `json:"fkey" gorm:"type:varchar(20);not null"`
	AuthorId int    `json:"authorId" gorm:"type:int(11);not null"`
	Created  int    `json:"created" gorm:"type:int(11);not null"`
}
