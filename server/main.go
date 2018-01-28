package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"loginv/server/gormdb"
	_ "loginv/server/gormdb" //init

	"loginv/server/models"
)

// "io"
// "net/http"
func main() {
	log.Println("starting server")
	http.HandleFunc("/hello", HellServer)
	http.HandleFunc("/login", LoginHandle)
	http.HandleFunc("/regist", ResitHandle)
	err := http.ListenAndServe(":8080", nil) //lsof -i 8080
	if err != nil {
		log.Fatalf("ListenAndServe:%v", err)
	}
	gormdb.Close()
}

func HellServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello vin es890123321")
}

func LoginHandle(w http.ResponseWriter, req *http.Request) {
	//no ssession
	err := req.ParseForm()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	input := req.PostForm
	fmt.Println(input)
	name := input.Get("name")
	pwd := input.Get("pwd")
	user := models.User{Username: name, Pwd: pwd}
	fmt.Print(user)
	err = models.ValidateUser(user)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "login succuss")
}

func ResitHandle(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	input := req.Form
	name := input.Get("name")
	pwd := input.Get("pwd")
	user := models.User{Username: name, Pwd: pwd}

	err = models.SaveUser(user)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	io.WriteString(w, "register succuss")
}
