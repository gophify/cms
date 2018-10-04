package formGeneratorCtx

import (
	"../../db/mysql"
	"io/ioutil"
	"encoding/json"
	"net/http"
	// "fmt"	
	"strconv"
)

var err error
var js []byte
var data = make(map[string]string)

type Result struct{
	Id, Status string
	Err error
	Data interface{}
}
var result Result

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(body, &data) == nil{
		key := data["key"]
		switch key {
			case "getformjson":
				js = getformjson(data["formId"], data["formType"])
			case "listallforms":
				js = listAllForms(data["formType"])
			case "savejson":
				js = saveJson(data)
			case "createtable":
				js = createTable(data["sql"], data["type"], data["id"])
			default:
				panic("error")
		}		
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

type FG struct{
	Name, Title, Parent, Child, ViewChildLbl, Json string
}
func getformjson(formId string, formType string) []byte {
	DB, _ := mysql.Connect()
	defer DB.Close()

	var name, title, parent, child, viewChildLbl, where string
	if formType == "independent" || formType == "i" {
		where = "id = "+ formId
	}else{
		where = "name = '"+ formId +"'"
	}
	err = DB.QueryRow("SELECT name, title, parent, child, viewChildLbl, json FROM _generator WHERE "+ where).Scan(&name, &title, &parent, &child, &viewChildLbl, &js)

	if err != nil{
		result.Status = "failed"
		result.Err = err
	}else{
		result.Status = "success"
		result.Data = FG{Name: name, Title: title, Parent: parent, Child: child, ViewChildLbl: viewChildLbl, Json: string(js)}
	}
	
	js, _ = json.Marshal(result)
	return js
}

func listAllForms(formType string) []byte {
	type FormGenerator struct{
		Id int
		Name, Icon, Title, Type, Parent string
	}

	FGs := []FormGenerator{}
	DB, _ := mysql.Connect()
	defer DB.Close()

	where := ""
	if formType != "" {
		where = " WHERE type = '"+ formType +"'"
	}
	rows, err2 := DB.Query("SELECT id, name, icon, title, type, parent FROM _generator"+ where +" LIMIT 0, 24")
    if err2 != nil {
        // fmt.Println(err2.Error())
    }	
	defer rows.Close()

	for rows.Next() {
		id, name, icon, title, ftype, parent := 0, "", "", "", "", ""
		rows.Scan(&id, &name, &icon, &title, &ftype, &parent)
		FGs = append(FGs, FormGenerator{Id: id, Name: name, Icon: icon, Title: title, Type: ftype, Parent: parent})
	}

	js, _ = json.Marshal(FGs)
	return js
}

func saveJson(data map[string]string) []byte{
	var formId, where string

	DB, _ := mysql.Connect()
	defer DB.Close()

	if data["formId"] == "0" {
		formId = strconv.FormatInt(mysql.GetLastId("_generator"), 10)
		where = "id = "+ formId
		if data["type"] == "i" {
			data["name"] = ""
			data["parent"] = ""
			result.Id = formId
		}else{
			result.Id = data["name"]
		}		
	}else{
		if data["type"] == "i" {
			data["name"] = ""
			data["parent"] = ""
			result.Id = data["formId"]
			where = "id = "+ data["formId"]
		}else{
			result.Id = data["name"]
			where = "name = '"+ data["name"] +"'"
		}
	}
	_, err = DB.Exec("UPDATE _generator SET name = '"+ data["name"] +"', icon = '"+ data["icon"] +"', title = '"+ data["title"] +"', type = '"+ data["type"] +"', parent = '"+ data["parent"] +"', child = '"+ data["child"] +"', viewChildLbl = '"+ data["viewChildLbl"] +"', json = '"+ data["json"] +"' WHERE "+ where)

	if data["parent"] != "" {
		DB.Exec("UPDATE _generator SET child = '"+ data["name"] +"' WHERE name = '"+ data["parent"] + "'")
	}

	if err != nil{
		result.Status = "failed"
		result.Err = err
	}else{
		result.Status = "success"
	}
	js, _ = json.Marshal(result)	
	return js
}

func createTable(sql string, fType string, formId string) []byte{
	type Result struct{
		Status, Msg string
		Error error
	}
	var result Result

	DB, _ := mysql.Connect()
	defer DB.Close()
	
	if fType != "i" {
		DB, err2 := DB.Query(sql)
		defer DB.Close()
		
		if err2 != nil {
			err = err2
		}
		result.Msg = "table created succesfully"
	}else{
		_, err2 := DB.Exec("INSERT IGNORE INTO _independent (form_id) VALUES ('"+ formId +"');")

		if err2 != nil {
			err = err2
		}else{
			result.Msg = "migrate succesfully"
		}		
	}

	if err != nil {
		result.Status = "error"
		result.Error = err
	}else{
		result.Status = "success"		
	}

	js, _ = json.Marshal(result)
	return js
}