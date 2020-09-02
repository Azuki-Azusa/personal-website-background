package model

import (
	"time"
)

// Blog model
type Lyric struct {
	ID       int
	Title    string
	Html     string
	Views    int
	Datetime time.Time
}

func CommitNewLyric(title, html string) error {
	var lyric = Lyric{Title: title, Html: html, Datetime: time.Now()}
	return Db.Create(&lyric).Error
}

func EditLyric(id int, title, html string) error {
	var lyric Lyric
	Db.First(&lyric, id)
	lyric.Title = title
	lyric.Html = html
	return Db.Save(&lyric).Error
}

func ReturnLyricList() ([]Lyric, error) {
	var lyrics []Lyric
	err := Db.Find(&lyrics).Error
	return lyrics, err
}
