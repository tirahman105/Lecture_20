package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/mateors/mcb"
)

var db *mcb.DB

type ReuestTable struct {
	ID      string `json:"aid"`
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Table   string `json:"table"`
	Status  int    `json:"status"`
}

func init() {

	db = mcb.Connect("localhost", "tareq", "drtareq")

	res, err := db.Ping()
	if err != nil {

		fmt.Println(res)
		os.Exit(1)
	}
	fmt.Println(res, err)

}

func main() {

	//How to insert into couchbase bucket
	var myData ReuestTable

	form := make(url.Values, 0)
	form.Add("bucket", "anonnota") //bucket Name
	form.Add("aid", "d00")         //document ID
	form.Add("name", "Md. Tareq Ibna Rahman")
	form.Add("company", "Anonnota")
	form.Add("email", "tirahman105@gmail.com")
	form.Add("table", "request") //what type of data or table name in general (SQL)
	form.Add("status", "1")

	p := db.Insert(form, &myData)    //pass by reference (&myData)
	fmt.Println("Status:", p.Status) //p.Status == Success means data successfully inserted to bucket.

	//How to retrieve from couchbase bucket (selected fields only)

	// pres := db.Query("SELECT aid,name,age,profession FROM master_erp WHERE type='participant'")
	// rows := pres.GetRows()

	// fmt.Println("Total Rows:", len(rows))
	// fmt.Println(rows)

	//How to retrieve from couchbase bucket (All fields using *)

	// pres := db.Query("SELECT * FROM master_erp WHERE type='participant'")
	// rows := pres.GetBucketRows("master_erp") //bucketName as argument

	// fmt.Println("Total Rows:", len(rows))
	// fmt.Println(rows)

}
