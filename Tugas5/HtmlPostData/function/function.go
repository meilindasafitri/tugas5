package function

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

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

func RouteIndexGet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("form").ParseFiles("index.html"))
		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func RouteSubmitPost(w http.ResponseWriter, r *http.Request) {
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

	if r.Method == "POST" {
		var tmpl = template.Must(template.New("result").ParseFiles("index.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		EmployeeID := r.FormValue("EmployeeID")
		LastName := r.FormValue("LastName")
		FirstName := r.FormValue("FirstName")
		Title := r.FormValue("Title")
		TitleOfCourtesy := r.FormValue("TitleOfCourtesy")
		BirthDate := r.FormValue("BirthDate")
		HireDate := r.FormValue("HireDate")
		Address := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		HomePhone := r.FormValue("HomePhone")
		Extension := r.FormValue("Extension")
		Photo := r.FormValue("Photo")
		Notes := r.FormValue("Notes")
		ReportsTo := r.FormValue("ReportsTo")
		ProvinceName := r.FormValue("ProvinceName")

		var data = map[string]string{"EmployeeID": EmployeeID}

		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//insert ke database ke table employees
		stmt, err := db.Prepare("INSERT INTO employees (EmployeeID,LastName,FirstName,Title,TitleOfCourtesy,BirthDate,HireDate,Address,City,Region,PostalCode,Country,HomePhone,Extension,Photo,Notes,ReportsTo,ProvinceName) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		_, err = stmt.Exec(EmployeeID, LastName, FirstName, Title, TitleOfCourtesy, BirthDate, HireDate, Address, City, Region, PostalCode, Country, HomePhone, Extension, Photo, Notes, ReportsTo, ProvinceName)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}
