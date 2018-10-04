package mediaCtx

import (
	"../../db/mysql"
	// "fmt"
	"os"
	"encoding/json"
	"net/http"

    "image"
    _ "image/jpeg"
    _ "image/png"	
)

func Api(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	type Media struct{
		Id int
		Dir, Fn string
		Lp string
		IsThumb bool
	}

	var id int
	var fn, dir string
	
    DB, _ := mysql.Connect()
	defer DB.Close()
		
	ms := []Media{}

	rows, _ := DB.Query("SELECT * FROM _media WHERE filename != '' ORDER BY id DESC LIMIT 0, 96")
	defer rows.Close()

	for rows.Next() {
		fn = ""
		dir = ""
		rows.Scan(&id, &fn, &dir)
		img := ""
		isthumb := false
		if Resize(dir, fn) {
			img = dir + "thumbnails-128/" + fn
			isthumb = true			
		}else{
			img = dir + fn
		}
		ms = append(ms, Media{Id: id, Dir: dir, Fn: fn, Lp: CheckLp("./"+ img), IsThumb: isthumb})
	}

	js, _ := json.Marshal(ms)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)	
}


func CheckLp(imagePath string) string {
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return ""
	}

    file, err := os.Open(imagePath)
    if err != nil {
        return ""
    }

    image, _, err := image.DecodeConfig(file)
    if err != nil {
		return ""
	}
	
	if image.Width > image.Height {
		return "l"
	}else{
		return "p"
	}
}