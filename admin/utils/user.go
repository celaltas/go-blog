package utils

import (
	"blog-app/admin/models"
	"net/http"
	"errors"
)

func SetUser(username, password string, w http.ResponseWriter, r *http.Request) error {
    session, err := store.Get(r, "go-user")
    if err != nil {
        return err
    }
    session.Values["user"] = &(models.User{Username: username, Password: password})
	err = session.Save(r, w)
    return err
}

func CheckUser(w http.ResponseWriter, r *http.Request) (*(models.User), error) {
    session, err := store.Get(r, "go-user")
    if err != nil {
        return nil, err
    }
    user, ok := session.Values["user"].(*(models.User))
    if !ok {
        return nil, errors.New("user not found in session")
    }

    return user, nil
}


func RemoveUser(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "go-user")
    if err != nil {
        return err
    }
	session.Options.MaxAge = -1
    return session.Save(r, w)
}