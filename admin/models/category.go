package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Title, Slug string
}

func (c Category) Migrate() {	
	db.AutoMigrate(&c)
}

func (c Category) Add() {
	db.Create(&c)

}
func (c Category) Get(where ...interface{}) Category {
	db.First(&c, where...)
	return c
}

func (c Category) GetAll(where ...interface{}) []Category {
	var categories []Category
	db.Find(&categories, where...)
	return categories
}

func (c Category) Update(column string, value interface{}){
	db.Model(&c).Update(column, value)
}

func (c Category) Updates(data Category){
	db.Model(&c).Updates(data)
}

func (c Category) Delete(){
	db.Delete(&c, c.ID)
}