package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title, Slug, Description, Content, Image_Url string
	CategoryID                                   int
}

func (p Post) Migrate() {
	db.AutoMigrate(&p)
}

func (p Post) Add() {
	db.Create(&p)

}
func (p Post) Get(where ...interface{}) Post {
	db.First(&p, where...)
	return p
}

func (p Post) GetAll(where ...interface{}) []Post {
	var posts []Post
	db.Find(&posts, where...)
	return posts
}

func (p Post) Update(column string, value interface{}){
	db.Model(&p).Update(column, value)
}

func (p Post) Updates(data Post){
	db.Model(&p).Updates(data)
}

func (p Post) Delete(){
	db.Delete(&p, p.ID)
}