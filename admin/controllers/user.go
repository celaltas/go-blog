package controllers

import (
	"blog-app/admin/models"
	"blog-app/admin/utils"
	"crypto/sha256"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)



type UserController struct {}




func (u UserController) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	view, err := template.ParseFiles(utils.Include("user/login")...)

	if err != nil {
		fmt.Println(err)
		return 
	}
	data := make(map[string]interface{})
	data["Alert"],_ = utils.GetAlert(w,r) 
	view.ExecuteTemplate(w, "login",data)
}
func (u UserController) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	username := r.FormValue("username")
	data := []byte(r.FormValue("password"))
	hash := sha256.Sum256(data)
	password := fmt.Sprintf("%x", hash)
	

	user:= models.User{}.Get("username = ? AND password = ?", username,password)
	if user == (models.User{}){
		utils.SetAlert(w,r,"Login failed.")
		http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
	}else{
		utils.SetUser(user.Username,user.Password,w,r )
		utils.SetAlert(w,r,"Welcome")
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
}
func (u UserController) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	utils.RemoveUser(w,r)
	utils.SetAlert(w,r,"Good By")
	http.Redirect(w, r, "/admin/login", http.StatusSeeOther)
}

