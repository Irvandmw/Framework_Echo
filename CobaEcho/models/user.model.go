package models

import (
	"net/http"

	"CobaEcho/db"
)

type Users struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func FetchAllUsers() (Response, error) {
	var obj Users
	var arrObj []Users
	var res Response

	con := db.CreateCon()

	sqlStatement := "SELECT id, name, age, address FROM users"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Age, &obj.Address)
		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

func StoreUsers(id int, nama string, umur int, alamat string,) (Response, error) {
	var res Response

	con := db.CreateCon()
	
	sqlStatement := "INSERT INTO users(id, name, age, address) VALUES(?,?,?,?)"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id, nama, umur, alamat)

	if err != nil {
		return res, err
	}

	lastInserted, err := result.LastInsertId()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInserted,
	}

	return res, nil

}


func UpdateUser(id int, nama string, umur int, alamat string)  (Response, error){
	var res Response

	con := db.CreateCon()

	sqlStatement := "UPDATE users SET name = ?,  age = ?, address = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, nil
	}

	result, err := stmt.Exec(nama, umur, alamat, id)
	if err != nil {
		return res, nil
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil

}

func DeleteUser(id int) (Response, error) {
	var res Response
	con := db.CreateCon()

	sqlStatement := "DELETE FROM users WHERE id = ?"

	stmt , err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected" : rowsAffected,
	}

	return res, nil

}