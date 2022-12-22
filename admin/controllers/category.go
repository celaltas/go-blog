package controllers

import (
	"blog-app/admin/models"
	"blog-app/admin/utils"
	"fmt"
	"html/template"
	"net/http"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
)

type Category struct{}

func (c Category) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view, err := template.ParseFiles(utils.Include("categories/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"]  = models.Category{}.GetAll()
	data["Alert"],_ = utils.GetAlert(w,r)
	view.ExecuteTemplate(w, "index", data)
}


func (c Category) Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	title := r.FormValue("title")
	slug := slug.Make(title)
	models.Category{
		Title: title,
		Slug: slug,
	}.Add()
	utils.SetAlert(w,r,"The category has been added successfully")
	http.Redirect(w,r, "/admin/categories", http.StatusSeeOther)
}
func (c Category) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := models.Category{}.Get(ps.ByName("id"))
	category.Delete()
	utils.SetAlert(w,r,"The category has been deleted successfully")
	http.Redirect(w,r, "/admin/categories", http.StatusSeeOther)
}