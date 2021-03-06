package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/mateors/mcb"
)

//var db *sql.DB
var db *mcb.DB
var err error

func init() {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	//db, err = sql.Open("mysql", "root:drtareq@tcp(127.0.0.1:3306)/hosting_db")

	// if there is an error opening the connection, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer the close till after the main function has finished
	// executing
	//defer db.Close()

	// 	fmt.Println("db connection successful")
	// }

	//couchbase connection block

	db = mcb.Connect("localhost", "tareq", "drtareq")
	res, err := db.Ping()
	if err != nil {
		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)
}

func main() {
	fmt.Println(" Local server is running on port 8888")
	http.HandleFunc("/", home)
	http.HandleFunc("/request", request)
	http.HandleFunc("/features", features)
	http.HandleFunc("/docs", docs)

	//built in file server
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}
func features(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp, err = ptmp.ParseFiles("wpage/features.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//fmt.Fprintf(w, `welcome`)
}
func docs(w http.ResponseWriter, r *http.Request) {

	ptmp, err := template.ParseFiles("template/base.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}

	ptmp, err = ptmp.ParseFiles("wpage/docs.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}
	ptmp.Execute(w, nil)
	//c
}

type ReuestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Table   string `json:"table"`
	Status  int    `json:"status"`
}

func request(w http.ResponseWriter, r *http.Request) {

	//r.ParseForm()

	//method 1

	// name := r.FormValue("name")
	// company := r.FormValue("company")
	// email := r.FormValue("email")

	//fmt.Println(name, company, email)

	//fmt.Fprintf(w, `Received %s %s %s`, name, company, email) //response

	//method 2
	// r.ParseForm()

	// for key, val := range r.Form {

	// 	fmt.Println(key, val)
	// }

	// 	qs := "INSERT INTO `request` (`id`, `name`, `company`, `email`, `status`) VALUES (NULL, '%s', '%s', '%s', '1');"
	// 	sql := fmt.Sprintf(qs, name, company, email)

	// 	//fmt.Println(sql)
	// 	insert, err := db.Query(sql)

	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	defer insert.Close()

	// 	fmt.Fprintf(w, `Received`)
	// }

	r.ParseForm()
	for key, val := range r.Form {
		fmt.Println(key, val)
	}

	var reqTable ReuestTable

	r.Form.Add("bucket", "anonnota")
	r.Form.Add("aid", "request::7") //we will update later
	r.Form.Add("table", "request")
	r.Form.Add("status", "1")
	pRes := db.Insert(r.Form, &reqTable)
	fmt.Println(pRes.Status, pRes.Errors)

	fmt.Fprintf(w, `Success!`)
}
