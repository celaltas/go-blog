package main

import (
	
	"blog-app/config"
	"fmt"
	"net/http"
	"encoding/gob"
	admin_models "blog-app/admin/models"
)

func main() {
	admin_models.ConnectDB()
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	admin_models.Category{}.Migrate()
	gob.Register(&(admin_models.User{}))

	fmt.Println("Server initializing at 8080...")
	if err := http.ListenAndServe(":8080",config.Routes()); err != nil {
		fmt.Printf("Error: %v\n",err)
	}
	
}

