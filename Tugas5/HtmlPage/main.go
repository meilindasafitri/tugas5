package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	em "Tugas5/HtmlPage/common"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type yamlconfig struct {
	Connection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		User     string `yaml:"user"`
		Database string `yaml:"database"`
	}
}

var db *sql.DB
var err error

func index(w http.ResponseWriter, r *http.Request) {

	var employees []em.Employees

	sql := `SELECT
			EmployeeID,
			IFNULL(LastName,'') LastName,
			IFNULL(FirstName,'') FirstName,
			IFNULL(Title,'') Title,
			IFNULL(TitleOfCourtesy,'') TitleOfCourtesy,
			IFNULL(BirthDate,'') BirthDate,
			IFNULL(HireDate,'') HireDate,
			IFNULL(Address,'') Address,
			IFNULL(City,'') City,
			IFNULL(Region,'') Region,
			IFNULL(PostalCode,'') PostalCode,
			IFNULL(Country,'') Country,
			IFNULL(HomePhone,'') HomePhone,
			IFNULL(Extension,'') Extension,
			IFNULL(Photo,'') Photo,
			IFNULL(Notes,'') Notes,
			IFNULL(ReportsTo,'') ReportsTo,
			IFNULL(ProvinceName,'') ProvinceName
		FROM employees ORDER BY EmployeeID`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var employee em.Employees
		err := result.Scan(&employee.EmployeeID, &employee.LastName, &employee.FirstName,
			&employee.Title, &employee.TitleOfCourtesy, &employee.BirthDate, &employee.HireDate,
			&employee.Address, &employee.City, &employee.Region, &employee.PostalCode, &employee.Country, &employee.HomePhone, &employee.Extension,
			&employee.Photo, &employee.Notes, &employee.ReportsTo, &employee.ProvinceName)

		if err != nil {
			panic(err.Error())
		}
		employees = append(employees, employee)
	}

	t, err := template.ParseFiles("index.html")
	t.Execute(w, employees)

	if err != nil {
		panic(err.Error())
	}

}

func main() {
	yamlFile, err := ioutil.ReadFile("../Yaml/config.yml")
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
		return
	}
	var yamlConfig yamlconfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	host := yamlConfig.Connection.Host
	port := yamlConfig.Connection.Port
	user := yamlConfig.Connection.User
	pass := yamlConfig.Connection.Password
	data := yamlConfig.Connection.Database

	var (
		//<user>:<passwprd>@tcp<IP address>/<Password>
		mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)
	)
	db, err = sql.Open("mysql", mySQL)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	log.Println("Server started on: http://localhost:8081")
	http.HandleFunc("/", index)
	http.ListenAndServe(":8081", nil)

}
