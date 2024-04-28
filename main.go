package main

import (
	"log"

	"github.com/ttmday/go-strategy-pattern/export"
)

type User struct {
	FirstName string `json:"firts_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Country   string `json:"country"`
	Charge    string `json:"charge"`
}

func main() {
	user := User{
		FirstName: "Eliezer",
		LastName:  "Capuano",
		Age:       27,
		Country:   "Venezuela",
		Charge:    "FrontEnd Developer",
	}

	em := export.NewExport(export.ExportCsvMethod{Data: user})

	if err := em.Export(); err != nil {
		log.Fatalf("Error exporting data %+v", err)
	}
}
