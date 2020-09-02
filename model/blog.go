package model

import (
	"time"

	"../database"
)

var Db = database.GetInstance()

// Blog model
type Blog struct {
	ID       int
	Title    string
	Markdown string
	Html     string
	Views    int
	Datetime time.Time
}

func CommitNewBlog(title, markdown, html string) error {
	var blog = Blog{Title: title, Markdown: markdown, Html: html, Datetime: time.Now()}
	return Db.Create(&blog).Error
}

func EditBlog(id int, title, markdown, html string) error {
	var blog Blog
	Db.First(&blog, id)
	blog.Title = title
	blog.Markdown = markdown
	blog.Html = html
	return Db.Save(&blog).Error
}

func ReturnBlogList() ([]Blog, error) {
	var blogs []Blog
	err := Db.Find(&blogs).Error
	return blogs, err
}
