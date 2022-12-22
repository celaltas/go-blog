package controllers

import (
	"blog-app/admin/models"
	"blog-app/admin/utils"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Dashboard struct{}

func (d Dashboard) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if _, err := utils.CheckUser(w, r); err != nil {
		utils.SetAlert(w, r, "Please login")
		http.Redirect(w, r, "admin/login", http.StatusSeeOther)
		return
	}

	view, err := template.New("index").Funcs(template.FuncMap{
		"getCategory":func (categoryID int) string {
			return models.Category{}.Get(categoryID).Title
		},
	}).ParseFiles(utils.Include("dashboard/list")...)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := make(map[string]interface{})
	data["Posts"] = models.Post{}.GetAll()
	data["Alert"], _ = utils.GetAlert(w, r)
	view.ExecuteTemplate(w, "index", data)
}

func (d Dashboard) NewItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view, err := template.ParseFiles(utils.Include("dashboard/add")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	view.ExecuteTemplate(w, "index", data)
}

func (d Dashboard) AddItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	title := r.FormValue("title")
	slug := slug.Make(title)
	description := r.FormValue("explanation")
	category, _ := strconv.Atoi(r.FormValue("category"))
	content := r.FormValue("content")

	r.ParseMultipartForm(10 << 20)
	file, header, err := r.FormFile("file")

	if err != nil {
		fmt.Println(err)

	}

	f, err := os.Create("uploads/" + header.Filename)
	if err != nil {
		fmt.Println(err)

	}
	bytes, err := io.Copy(f, file)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("The number of bytes are: %d\n", bytes)

	models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  category,
		Content:     content,
		Image_Url:   "uploads/" + header.Filename,
	}.Add()

	utils.SetAlert(w, r, "Post was successfully added!")
	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}

func (d Dashboard) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post := models.Post{}.Get(ps.ByName("id"))
	post.Delete()
	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func (d Dashboard) Edit(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view, err := template.ParseFiles(utils.Include("dashboard/edit")...)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make(map[string]interface{})
	data["Categories"] = models.Category{}.GetAll()
	data["Post"] = models.Post{}.Get(ps.ByName("id"))
	view.ExecuteTemplate(w, "index", data)
}
func (d Dashboard) Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	post := models.Post{}.Get(ps.ByName("id"))
	title := r.FormValue("title")
	slug := slug.Make(title)
	description := r.FormValue("explanation")
	category, _ := strconv.Atoi(r.FormValue("category"))
	content := r.FormValue("content")
	isSelected := r.FormValue("is_selected")

	var imageUrl string

	if isSelected == "1" {

		r.ParseMultipartForm(10 << 20)
		file, header, err := r.FormFile("file")

		if err != nil {
			fmt.Println(err)

		}

		f, err := os.Create("uploads/" + header.Filename)
		if err != nil {
			fmt.Println(err)

		}
		bytes, err := io.Copy(f, file)
		if err != nil {
			fmt.Println(err)

		}

		fmt.Printf("The number of bytes are: %d\n", bytes)
		imageUrl = "uploads/" + header.Filename
		os.Remove(post.Image_Url)

	} else {
		imageUrl = post.Image_Url

	}

	post.Updates(models.Post{
		Title:       title,
		Slug:        slug,
		Description: description,
		CategoryID:  category,
		Content:     content,
		Image_Url:   imageUrl,
	})

	http.Redirect(w, r, "/admin", http.StatusSeeOther)

}
