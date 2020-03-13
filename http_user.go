package main

import (
	"net/http"

	"gopkg.in/mgo.v2"
)

func handleSignup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := TEMPLATES.ExecuteTemplate(w, "signup", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSignupSuccess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := TEMPLATES.ExecuteTemplate(w, "signup-success", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSignupSubmit(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("ID")
	if id == "" {
		http.Error(w, "ID 값이 빈 문자열 입니다", http.StatusBadRequest)
		return
	}
	pw := r.FormValue("Password")
	if pw == "" {
		http.Error(w, "Password 값이 빈 문자열 입니다", http.StatusBadRequest)
		return
	}
	if pw != r.FormValue("ConfirmPassword") {
		http.Error(w, "입력받은 2개의 패스워드가 서로 다릅니다", http.StatusBadRequest)
		return
	}
	encryptedPW, err := Encrypt(pw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u := User{}
	u.AccessLevel = "default"
	u.ID = id
	u.Password = encryptedPW
	err = u.CreateToken()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session, err := mgo.Dial(*flagDBIP)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer session.Close()
	err = AddUser(session, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/signup-success", http.StatusSeeOther)
}

func handleSignin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := TEMPLATES.ExecuteTemplate(w, "signin", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
