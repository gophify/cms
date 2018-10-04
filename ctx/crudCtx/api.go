package crudCtx

import (
	"../../db/mysql"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"fmt"
	// "../authCtx"
	"strconv"
	"strings"
)

var js []byte
var data = make(map[string]string)

type F struct{
	T, M, V string
}
var fields []F
var err error

type Result struct {
	Id string
	Status string
	Err error
}
var result Result

func Api(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	body, _ := ioutil.ReadAll(r.Body)

	// auth := authCtx.CheckAuth(w, r)
	// if auth.Status == "forbidden" {
	// 	js, _ = json.Marshal(auth)
		
	// }else 
	if json.Unmarshal(body, &data) == nil{

		key := data["key"]
		switch key {
			case "save":
				if data["formType"] == "independent" {
					js = saveIndependent(data)
				}else{					
					js = save(data["cparent"], data)
				}
			case "read":
				js = read(data["formId"], data["contentId"])
			case "readIndependent":
				js = readIndependent(data["formId"])
			case "list":
				js = list(w, r, data["formId"], data["cparent"], data["cparent_tblname"])
			case "delete":
				// js = delete(data["formId"], data["contentId"])
			default:
				panic("unrecognized escape character")
		}
		
	}else{
		js, _ = json.Marshal(Result{Status: "error"})
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}

func readIndependent(formId string) []byte {
	DB, _ := mysql.Connect()
	defer DB.Close()
	_ = DB.QueryRow("SELECT json FROM _independent WHERE form_id = "+ formId).Scan(&js)
	return js
}

func read(formId string, id string) []byte {
	DB, _ := mysql.Connect()
	defer DB.Close()
	rows, _ := DB.Query("SELECT * FROM " + formId + " WHERE id = "+ id)
	defer rows.Close()
	
	columns, _ := rows.Columns()
    count := len(columns)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)
    data := make(map[string]interface{})

    for rows.Next() {
        for i, _ := range columns {
            valuePtrs[i] = &values[i]
        }
        rows.Scan(valuePtrs...)
        for i, col := range columns {
		    var inputs []string
            var v interface{}
            val := values[i]
            b, ok := val.([]byte)
            if (ok) {
            	v = b
                if json.Unmarshal(b, &inputs) == nil{
                	v = inputs
                }else{
            		v = string(b)
            	}
            }
            data[col] = v
        }
	}
	js, _ := json.Marshal(data)
	return js
}

type FG struct{
	Icon, Title, Child, ViewChildLbl string
	Col4list []byte
}
func getCol4List(formId string) FG {
	DB, _ := mysql.Connect()
	defer DB.Close()

	var icon, title, child, viewChildLbl string
	_ = DB.QueryRow("SELECT icon, title, child, viewChildLbl, col4list FROM _generator WHERE name = '"+ formId +"'").Scan(&icon, &title, &child, &viewChildLbl, &js)
	return FG{Icon: icon, Title: title, Child: child, ViewChildLbl: viewChildLbl, Col4list: js}
}

func list(w http.ResponseWriter, r *http.Request, formId string, cparent string, cparent_tblname string) []byte {

	type Col struct{
		M, L, Class, Type string
	}
	var cols []Col
	var select4list_arr []string
	var select4list string
	var parent_id, parent_id_tblname string

	fg := getCol4List(formId)

	if json.Unmarshal(fg.Col4list, &cols) == nil{
	}else{
		cols = append(cols, Col{M: "title", L: "Title", Class: "xs10", Type: ""})
	}

	for _, col := range cols {
		select4list_arr = append(select4list_arr, col.M)
	}
	select4list = strings.Join(select4list_arr, ", ")

	DB, _ := mysql.Connect()
	defer DB.Close()
	rows, _ := DB.Query("SELECT id, "+ select4list +" FROM "+ formId +" WHERE parent = '"+ cparent +"' LIMIT 0, 24")
	defer rows.Close()

	columns, _ := rows.Columns()
    count := len(columns)
    values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	
	type Data map[string]interface{}
	var datas []Data

    for rows.Next() {
        for i, _ := range columns {
            valuePtrs[i] = &values[i]
        }
		rows.Scan(valuePtrs...)
		
		data := make(Data)
        for i, col := range columns {
		    var inputs []string
            var v interface{}
            val := values[i]
            b, ok := val.([]byte)
            if (ok) {
            	v = b
                if json.Unmarshal(b, &inputs) == nil{
                	v = inputs
                }else{
            		v = string(b)
            	}
            }
			data[col] = v
		}
		datas = append(datas, Data(data))
	}

	type List struct{
		Icon, Title, Parent_id, Parent_tblname, Child, ViewChildLbl string
		Cols []Col
		Datas []Data
	}
	var list List

	//find parent of parent on current list
	_ = DB.QueryRow("SELECT parent FROM "+ cparent_tblname +" WHERE id = "+ cparent).Scan(&parent_id)
	if parent_id != "" {
		_ = DB.QueryRow("SELECT parent FROM _generator WHERE name = '"+ cparent_tblname +"'").Scan(&parent_id_tblname)
	}
	
	list = List{Icon: fg.Icon, Title: fg.Title, Parent_id: parent_id, Parent_tblname: parent_id_tblname, Child: fg.Child, ViewChildLbl: fg.ViewChildLbl, Cols: cols, Datas: datas}

	js, _ := json.Marshal(list)
	return js	
}

func saveIndependent(data map[string]string) []byte {
	DB, _ := mysql.Connect()
	defer DB.Close()
	replacer := strings.NewReplacer("\\n", "<br />", "'", `\'`, "\\", "/")
	fmt.Println("UPDATE _independent SET json = '"+ replacer.Replace(data["content"]) +"' WHERE form_id="+ data["formId"] + ";")
	_, err = DB.Exec("UPDATE _independent SET json = '"+ replacer.Replace(data["content"]) +"' WHERE form_id="+ data["formId"] + ";")

	if err != nil {
		result = Result{Err: err, Status: "failed"}
	}else{
		result = Result{Id: data["formId"], Status: "success"}
	}

	js, _ = json.Marshal(result)
	return js
}

func save(cparent string, data map[string]string) []byte {
	DB, _ := mysql.Connect()
	defer DB.Close()
	var sql_i_fv []string
	if json.Unmarshal([]byte(data["content"]), &fields) == nil{
		for _, e := range fields {
			if e.V != "" {
				sql_i_fv = append(sql_i_fv, (e.M + "='" + strings.Replace(e.V, "'", `\'`, -1) + "'"))	
			}
		}
	}
	var contentId string
	if data["contentId"] == "0" {
		contentId = strconv.FormatInt(mysql.GetLastId(data["formId"]), 10)
	}else{
		contentId = data["contentId"]
	}
	_, err = DB.Exec("UPDATE "+ data["formId"] +" SET parent = '"+ cparent +"', "+ strings.Join(sql_i_fv, ", ") + " WHERE id="+ contentId + ";")
	
	if err != nil {
		result = Result{Err: err, Status: "failed"}
	}else{
		result = Result{Id: contentId, Status: "success"}
	}

	js, _ = json.Marshal(result)
	return js
}