package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
    "fmt"
    "io/ioutil"


	_ "github.com/go-sql-driver/mysql"
)


//Structure Employee
type Employee struct {
	ID        int
	FirstName string   `json:"FirstName"`
	LastName   string  `json:"LastName"`
	City      string   `json:"City"`
	Age       string   `json:"Age"`
}

//Database connection
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "tushar321"
	dbName := "employee"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//TO get all details
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Employee ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var name, city, lastname, age string
		err = selDB.Scan(&id, &name, &lastname, &age, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.FirstName = name
		emp.LastName = lastname
		emp.Age = age
		emp.City = city
		res = append(res, emp)
	}
	json.NewEncoder(w).Encode(res)
	defer db.Close()
}

//To show by one Id
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city, lastname, age string
		err = selDB.Scan(&id, &name, &lastname, &age, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.FirstName = name
		emp.LastName = lastname
		emp.Age = age
		emp.City = city

	}
	json.NewEncoder(w).Encode(emp)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	// empty function
}

//to edit the data
func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM Employee WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var name, city, lastname, age string
		err = selDB.Scan(&id, &name, &lastname, &age, &city)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.FirstName = name
		emp.LastName = lastname
		emp.Age = age
		emp.City = city
	}
	json.NewEncoder(w).Encode(emp)
	defer db.Close()
}

//to insert the data
func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		stmt, err := db.Prepare("INSERT INTO employee.employee(name, lastname, age, city) VALUES(?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
        json.Unmarshal(body, &keyVal)
		FirstName := keyVal["FirstName"]
		LastName := keyVal["LastName"]
		City := keyVal["City"]
		Age := keyVal["Age"]

		_, err = stmt.Exec(FirstName, LastName, City, Age)
         if err != nil{
			panic(err.Error())
		 }
	}
	fmt.Fprintf(w, "New post was created")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//to update the data
func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		stmt, err := db.Prepare("UPDATE Employee SET name=?, lastname=?, age=?, city=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		keyVal := make(map[string]string)
        json.Unmarshal(body, &keyVal)
		FirstName := keyVal["FirstName"]
		LastName := keyVal["LastName"]
		City := keyVal["City"]
		Age := keyVal["Age"]
		Id := keyVal["ID"]

		_, err = stmt.Exec(FirstName, LastName, City, Age, Id)
         if err != nil{
			panic(err.Error())
		}
	}
	fmt.Fprintf(w, "DataBase Updated")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//to delete the data
func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM Employee WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

//Routers
func main() {
	log.Println("Server started")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":8081", nil)
}
