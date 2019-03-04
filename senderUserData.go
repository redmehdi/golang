package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/csv"
    "os"
)

type User struct {
	Nombre string   `json: nombre`
	Apellido string   `json: apellido`
    Edad string `json: edad`
    InscriptionDate string `json: inscriptionDate`
    Telefono string `json: telefono`
    Email string `json: email`
	ChangeBy string `json: changeBy`
}

func client() {
	
	
	csvFile, err := os.Open("./mockup.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer csvFile.Close()
 
    reader := csv.NewReader(csvFile)
    reader.FieldsPerRecord = -1
 
    csvData, err := reader.ReadAll()
	
	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
	
	for _, each := range csvData {
		user := User{
		Nombre 			: each[1],
		Apellido 		: each[0],
		Edad        	: each[2],
		InscriptionDate : each[3],
		Telefono        : each[5],
		Email 			: each[4],
		ChangeBy 		:"GO_PROCESS",
		}
		log.Printf("Sending data User: %v\n", user)
		userJson, err := json.Marshal(user)
		log.Printf("data User to json: %v\n", bytes.NewBuffer(userJson))
		
		req, err := http.NewRequest("POST", "http://localhost:8080/add", bytes.NewBuffer(userJson))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		body, err := ioutil.ReadAll(resp.Body)

		fmt.Println("Response: ", string(body))
		resp.Body.Close()
	}
}

func main() {
	client()
}