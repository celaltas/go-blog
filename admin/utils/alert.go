package utils

import (
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/sessions"
)
var store = sessions.NewCookieStore([]byte(os.Getenv("YUDUM")))


func SetAlert(w http.ResponseWriter, r *http.Request, message string) error {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		return err
	}

	session.AddFlash(message)
	return session.Save(r, w)
}
func GetAlert(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	session, err := store.Get(r, "go-alert")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	flashes := session.Flashes()
	data := make(map[string]interface{})

	if len(flashes) > 0 {
		data["is_alert"] = true
		data["message"] = flashes[0]
	}else{
		data["is_alert"] = false
		data["message"] = nil
	}


	if err := session.Save(r, w); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}