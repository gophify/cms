package dashboardCtx

import (
	"../../db/mysql"
	// "fmt"	
	"encoding/json"
	"net/http"
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	type SE struct{
		S, E string
	}	
	type Room struct{
		Rt_id, Id, Title string
	}
	type Property struct{
		Id, Title string
		Rooms []Room
	}
	type Availability struct{
		BlockByAdmin bool
		Room_id string
		Availability []SE
	}
	type Api struct{
		Properties []Property
		Availability []Availability
	}

	var id, title, rtid, rid, rtitle string
	var availability []byte
	availabilities := []Availability{}
	
    DB, _ := mysql.Connect()
	defer DB.Close()
		
	properties := []Property{}
	rows, _ := DB.Query("SELECT id, title FROM property ORDER BY id DESC LIMIT 0, 12")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&id, &title)

		//GRAB rooms
		r_rows, _ := DB.Query("SELECT rt.id, r.id, r.title, r.availability FROM roomtype rt RIGHT JOIN room r ON rt.id = r.parent WHERE rt.parent = "+ id +" ORDER BY r.title ASC")		
		defer r_rows.Close()
		rooms := []Room{}

		for r_rows.Next() {
			dates := []SE{}
			r_rows.Scan(&rtid, &rid, &rtitle, &availability)
			rooms = append(rooms, Room{Rt_id: rtid, Id: rid, Title: rtitle})
			if json.Unmarshal(availability, &dates) == nil{			
				availabilities = append(availabilities, Availability{BlockByAdmin: true, Room_id: rid, Availability: dates})
			}
		}
		if len(rooms) > 0 {
			properties = append(properties, Property{Id: id, Title: title, Rooms: rooms})
		}
	}

	js, _ := json.Marshal(Api{Properties: properties, Availability: availabilities})
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}