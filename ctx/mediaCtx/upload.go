package mediaCtx

import (
    "../../db/mysql"
	"net/http"
	"encoding/json"
	"os"
	"fmt"
    "io"  
    "strings"
    "strconv"
)

func Upload(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")

    err := r.ParseMultipartForm(200000) // grab the multipart form
    if err != nil {
        fmt.Fprintln(w, err)
        return
    }
	formdata := r.MultipartForm // ok, no problem so far, read the Form data
    files := formdata.File["multiplefiles"] // grab the filenames
    os.OpenFile("./media/", os.O_WRONLY|os.O_CREATE, 0666)
	
	type Result struct{
        Status, Id, Dir, Fn, Lp string
        IsThumb bool
        Err error
	}
	var results []Result

    for i, _ := range files { // loop through the files one by one
        file, err := files[i].Open()
        defer file.Close()
        if err != nil {
            fmt.Fprintln(w, err)
            return
        }

        oriFn := files[i].Filename
        strMid := strconv.FormatInt(mysql.GetLastId("_media"), 10)
        fnSplit := strings.SplitN(oriFn, ".", 2)
        newFn := fnSplit[0] + "_" + strMid + "." + fnSplit[1]

        dir := "/media/images/"
        out, err := os.Create("." + dir + newFn)

        defer out.Close()
        if err != nil {
            fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
            return
        }
        _, err = io.Copy(out, file)

        if err != nil {
            fmt.Fprintln(w, err)
            return
        }

		isthumb := false
		if Resize(dir, newFn) {
			isthumb = true			
		}       
        
        errSv2db := save2db(strMid, newFn, dir)
        if errSv2db == nil {
            results = append(results, Result{Status: "success", Id: strMid, Dir: dir, Fn: newFn, Lp: CheckLp("./"+ dir + newFn), IsThumb: isthumb})
        }else{
            results = append(results, Result{Status: "failed", Fn: oriFn, Err: errSv2db})
        }
    }

    js, _ := json.Marshal(results)
 	w.Header().Set("Content-Type", "application/json")
	w.Write(js)   
}