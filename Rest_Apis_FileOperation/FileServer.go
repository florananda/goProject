package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

type Errors struct {
	Error string `json:"error"`
}
type Message struct {
	Message string `json:"token"`
	msg     string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	userid   string
	Token    string `json:"token"`
}

type Userfile struct {
	username      string
	userid        string
	Filename      string
	Contentlength string
	Contenttype   string
	file          []byte
}

type Filename struct {
	Filenames []string
}

var userfiles []Userfile

var users []User
var curUser int = 0

func CreateUser(res http.ResponseWriter, req *http.Request) {
	var user User
	err := new(Errors)
	_ = json.NewDecoder(req.Body).Decode(&user)
	matched, _ := regexp.MatchString("^[A-Za-z0-9]{3,20}$", user.Username)
	if len(user.Password) < 8 || matched == false {
		err.Error = "Username or Password is not right. Username has to be alphanumeric and 3-20 characters long. Password should be at least 8 characters long"
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err)
		return
	}
	user.userid = strconv.Itoa(curUser + 1)
	curUser = curUser + 1
	users = append(users, user)
	//json.NewEncoder(res).Encode(user)
	res.WriteHeader(http.StatusNoContent)
	res.Write(nil)
}

func LoginUser(res http.ResponseWriter, req *http.Request) {
	//Find the USer in the array
	var user User
	errm := new(Errors)
	msg := new(Message)
	_ = json.NewDecoder(req.Body).Decode(&user)
	//find the user in the data storage
	for idx, us := range users {
		if us.Username == user.Username && us.Password == user.Password {
			if us.Token != "" {
				msg.Message = us.Token
				break
			} else {
				token := strconv.Itoa(rand.Intn(5000)+1) + us.userid
				users[idx].Token = token
				msg.Message = token
				break
			}
		}
	}
	if msg.Message == "" {
		errm.Error = "username does not exist in the system"
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(errm)
		return
	}
	err := json.NewEncoder(res).Encode(msg)
	if err != nil {
		errm.Error = err.Error()
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(errm)
		return
	}

}

func PutFiles(res http.ResponseWriter, req *http.Request) {
	var user User
	var userfile Userfile
	errm := new(Errors)
	msg := new(Message)
	params := mux.Vars(req)
	token := req.Header.Get("X-Session")
	ctype := req.Header.Get("Content-Type")
	//find the session
	for _, us := range users {
		if us.Token == token {
			msg.msg = us.Token
			user = us
			break
		}
	}
	//check if token found
	if msg.msg == "" {
		errm.Error = "Token not matching"
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(errm)
		return
	}
	//store the data
	userfile.username = user.Username
	userfile.userid = user.userid
	userfile.Filename = params["fname"]
	userfile.Contenttype = ctype
	userfile.Contentlength = req.Header.Get("Content-Length")
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errm.Error = "File mot converted to bytes"
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusForbidden)
		json.NewEncoder(res).Encode(errm)
		return
	}
	userfile.file = body
	userfiles = append(userfiles, userfile)
	req.Body.Close()
	res.WriteHeader(http.StatusCreated)
	res.Header().Add("Location", "files/"+params["fname"])
	res.Write([]byte("Location:    " + "files/" + params["fname"]))

}

func GetFile(res http.ResponseWriter, req *http.Request) {
	var user User
	msgfound := false
	//errm := new(Errors)
	msg := new(Message)
	params := mux.Vars(req)
	token := req.Header.Get("X-Session")
	//find the session
	for _, us := range users {
		if us.Token == token {
			msg.msg = us.Token
			user = us
			break
		}
	}
	log.Print("step-1")
	//check if token found
	if msg.msg == "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	log.Println("Userfiles are:", userfiles)
	//Find the file
	for _, usfile := range userfiles {
		if usfile.username == user.Username && usfile.userid == user.userid && usfile.Filename == params["filename"] {
			msgfound = true
			res.WriteHeader(http.StatusOK)
			res.Header().Set("Content-Type", usfile.Contenttype)
			res.Header().Set("Content-Length", usfile.Contentlength)
			res.Write(usfile.file)
			return
		}
	}
	if msgfound == false {
		res.WriteHeader(http.StatusNotFound)
		return
	}
}

func GetFiles(res http.ResponseWriter, req *http.Request) {
	var user User
	//errm := new(Errors)
	var filenames Filename
	msg := new(Message)
	token := req.Header.Get("X-Session")
	//find the session
	for _, us := range users {
		if us.Token == token {
			msg.msg = us.Token
			user = us
			break
		}
	}

	//check if token found
	if msg.msg == "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	//Find the filenames
	for _, usfile := range userfiles {
		if usfile.username == user.Username && usfile.userid == user.userid {
			filenames.Filenames = append(filenames.Filenames, usfile.Filename)

		}
	}
	j, _ := json.Marshal(filenames)
	res.Write(j)
	filenames.Filenames = nil
}

func DeleteFile(res http.ResponseWriter, req *http.Request) {
	var user User
	var index int
	msgfound := false
	//errm := new(Errors)
	msg := new(Message)
	params := mux.Vars(req)
	token := req.Header.Get("X-Session")
	//find the session
	for _, us := range users {
		if us.Token == token {
			msg.msg = us.Token
			user = us
			break
		}
	}
	log.Print("step-1")
	//check if token found
	if msg.msg == "" {
		res.WriteHeader(http.StatusForbidden)
		return
	}
	//Find the file
	var duf []Userfile
	for idx, usfile := range userfiles {
		if usfile.username == user.Username && usfile.userid == user.userid && usfile.Filename == params["filename"] {

			msgfound = true
			index = idx
			res.WriteHeader(http.StatusNoContent)
			break
			//return
		}
	}
	if msgfound == false {
		res.WriteHeader(http.StatusNotFound)
		return
	} else {
		for idx, uf := range userfiles {
			if idx == index {
				continue
			} else {
				duf = append(duf, uf)
			}
		}
	}
	userfiles = nil
	userfiles = duf
	log.Println("Userfiles are:", userfiles)

}

func main() {
	port := flag.String("p", "8000", "port")
	//dir := flag.String("d", ".", "dir")
	flag.Parse()
	router := mux.NewRouter()

	http.HandleFunc("/favicon.ico", handlerFav)
	router.HandleFunc("/register", CreateUser).Methods("POST")
	router.HandleFunc("/login", LoginUser).Methods("POST")
	router.HandleFunc("/files/{fname}", PutFiles).Methods("PUT")
	router.HandleFunc("/files/{filename}", GetFile).Methods("GET")
	router.HandleFunc("/files", GetFiles).Methods("GET")
	router.HandleFunc("/files/{filename}", DeleteFile).Methods("DELETE")
	fmt.Print("Server STARTING AT PORT:", *port)
	log.Fatal(http.ListenAndServe(":"+*port, router))

}

func handlerFav(res http.ResponseWriter, req *http.Request) {}
