package authCtx

import (
	"../../db/mysql"
	"io/ioutil"
	"encoding/json"
	"net/http"
	// "fmt"
	"github.com/gorilla/sessions"
	"strings"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

var js []byte
var data = make(map[string]string)
var err error

type Result struct {
	Role, Firstname, Lastname, Uname, Email string
	Status string
	Err error
}
var result Result

func Api(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	result = Result{Status: "error"}
	body, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(body, &data) == nil{
		DB, _ := mysql.Connect()
		defer DB.Close()
		var role, firstname, lastname, uname, email string
		err = DB.QueryRow("SELECT role, firstname, lastname, uname, email FROM user WHERE LOWER(email) ='"+ strings.ToLower(data["email"]) +"' AND password = '"+ data["password"] +"'").Scan(&role, &firstname, &lastname, &uname, &email)		
		if err != nil {
			result = Result{Status: "invalid", Err: err}
		}else{
			Login(w, r)
			result = Result{Role: role, Firstname: firstname, Lastname: lastname, Uname: uname, Email: email, Status: "success"}
		}
	}
	js, _ = json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) Result{

	session, _ := store.Get(r, "cookie-name")

	// Check if user is authenticated
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		// http.Error(w, "Forbidden", http.StatusForbidden)
		result = Result{Status: "forbidden"}
	}else{
		result = Result{Status: "authenticated"}
	}

	return result
}

func Secret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	js, _ = json.Marshal(CheckAuth(w, r))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)		
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
}

func LogoutApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = false
	session.Save(r, w)

	type Res struct{
		Status string
	}
	js, _ = json.Marshal(Res{Status: "success"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func UserLog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	body, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(body, &data) == nil{
		DB, _ := mysql.Connect()
		defer DB.Close()
		_, err = DB.Exec("INSERT INTO _userlog (role, fullname, email, action) VALUES ('"+ data["role"] +"', '"+ data["fullname"] +"', '"+ data["email"] +"', '"+ strings.Replace(data["action"], "'", `\'`, -1) +"');")
	}
}