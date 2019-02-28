package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "bytes"
)
 
type User struct {
    nombre string 
    apellido  string
    edad       string
    inscriptionDate string
	telefono	string             
	email string  
	changeBy string     
}
 
func main() {
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
 
    var user User
    var users []User
	var dummy Dummy
 
    for _, each := range csvData {
		b := new(bytes.Buffer)
        user.apellido = each[0]
        user.nombre = each[1]
        user.edad = each[2]
        user.inscriptionDate = each[3]
		user.email = each[4]
		user.telefono = each[5]      
		user.changeBy = "GOTEST"
		
		 url := "http://localhost:8080/add"
	    json.NewDecoder(b).Decode(user)
		fmt.Println("URL:testttttttttttttttttttt>", b)
        err := json.Unmarshal(b, &dummy)
	
//	    var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
//		var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
//		b, err := json.Marshal(user)
	    req, err := http.NewRequest("POST", url,b)
	    req.Header.Set("X-Custom-Header", "myfirstGoLang")
	    req.Header.Set("Content-Type", "application/json")
	
	    client := &http.Client{}
	    resp, err := client.Do(req)
	    if err != nil {
	        panic(err)
	    }
	    defer resp.Body.Close()
	
	    fmt.Println("response Status:", resp.Status)
	    fmt.Println("response Headers:", resp.Header)
	    body, _ := ioutil.ReadAll(resp.Body)
	    fmt.Println("response Body:", string(body))
        users = append(users, user)
    }
 
    // Convert to JSON
    jsonData, err := json.Marshal(users)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
 
    fmt.Println(string(jsonData))
 
    jsonFile, err := os.Create("./data.json")
    if err != nil {
        fmt.Println(err)
    }
    defer jsonFile.Close()
 
    jsonFile.Write(jsonData)
    jsonFile.Close()
}