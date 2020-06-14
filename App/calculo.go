//esto era solo para probar unas funciones

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
	"os"
	"encoding/json"
)

//json para mandar
type Datos struct {
	RAM 	string `json:ram`
	USOR 	string `json:usor`
	OCUR	string `json:ocur`
	CPU		string `json:cpu`
}

//funcion para el string del servidor 2
func main(){
	//aqui esta la api, solo tengo que cambiar de donde conseguirla
	response, err := http.Get("http://192.168.1.34:5000/serv2")

	//aqui tenemos el error
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
	}
	
	responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
	}
	
	var responseObj []Datos
	json.Unmarshal(responseData, &responseObj)

	str := responseObj[0].RAM
	str1 := responseObj[0].USOR
	str2 := responseObj[0].OCUR
	str3 := responseObj[0].CPU
	fmt.Println(string(str + "%" + str1 + "%" + str2 + "%" + str3 + "%"))
}