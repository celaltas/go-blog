package models

import (
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username, Password string
}


func (u User) Migrate() {
	db.AutoMigrate(&u)
}

func (u User) Add() {
	db.Create(&u)

}
func (u User) Get(where ...interface{}) User {
	db.First(&u, where...)
	return u
}

func (u User) GetAll(where ...interface{}) []User {
	var users []User
	db.Find(&users, where...)
	return users
}

func (u User) Update(column string, value interface{}){
	db.Model(&u).Update(column, value)
}

func (u User) Updates(data User){
	db.Model(&u).Updates(data)
}

func (u User) Delete(){
	db.Delete(&u, u.ID)
}