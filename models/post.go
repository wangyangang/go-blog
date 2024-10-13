package models

import (
	"go-blog/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`        // 文章id
	Title      string    `json:"title"`      // 文章标题
	Slug       string    `json:"slug"`       // 自定义页面path
	Content    string    `json:"content"`    // 文章的html
	Markdown   string    `json:"markdown"`   // 文章的markdown
	CategoryId int       `json:"categoryId"` // 分类id
	UserId     int       `json:"userId"`     // 用户id
	ViewCount  int       `json:"viewCount"`  // 查看次数
	Type       int       `json:"type"`       // 文章类型：0 普通 1 自定义文章
	CreateAt   time.Time `json:"createAt"`   // 创建时间
	UpdateAt   time.Time `json:"updateAt"`   // 更新时间
}

type PostMore struct {
	Pid          int           `json:"pid"`
	Title        string        `json:"title"`
	Slug         string        `json:"slug"`
	Content      template.HTML `json:"content"`
	CategoryId   int           `json:"categoryId"`
	CategoryName string        `json:"categoryName"`
	UserId       int           `json:"userId"`
	UserName     string        `json:"userName"`
	ViewCount    int           `json:"viewCount"`
	Type         int           `json:"type"`
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}
type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
