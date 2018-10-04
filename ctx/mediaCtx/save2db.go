package mediaCtx

import (
	"../../db/mysql"
)

func save2db(id string, fn string, path string) error{
	DB, _ := mysql.Connect()
	defer DB.Close()

	_, err := DB.Exec("UPDATE _media SET filename = '"+ fn +"', path = '"+ path +"' WHERE id = "+ id)	
	return err
}