package config

import (
	admin "blog-app/admin/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Routes() *httprouter.Router {
	r:=httprouter.New()
	r.GET("/admin",admin.Dashboard{}.Index)
	r.GET("/admin/add-new",admin.Dashboard{}.NewItem)
	r.POST("/admin/add",admin.Dashboard{}.AddItem)
	r.GET("/admin/delete/:id",admin.Dashboard{}.Delete)
	r.GET("/admin/edit/:id",admin.Dashboard{}.Edit)
	r.POST("/admin/update/:id",admin.Dashboard{}.Update)


	r.GET("/admin/categories",admin.Category{}.Index)
	r.POST("/admin/categories/add",admin.Category{}.Add)
	r.GET("/admin/categories/delete/:id",admin.Category{}.Delete)






	r.GET("/admin/login", admin.UserController{}.Index)
	r.POST("/admin/do_login", admin.UserController{}.Login)
	r.GET("/admin/logout", admin.UserController{}.Logout)



	r.ServeFiles("/admin/assets/*filepath",http.Dir("admin/assets") )
	r.ServeFiles("/uploads/*filepath",http.Dir("uploads") )
	return r
}