package main

import(
	"net/http"
	"io/ioutil"
	"strings"
	"strconv"
	"encoding/json"
)

type memStruct struct{
	Total_mem int
	Free_mem int
}


func main(){
	http.HandleFunc("/memoria", ramInfo)
	http.ListenAndServe(":3000", nil)
}

func ramInfo(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadFile("/proc/meminfo");
	if err != nil{
		return
	}
	str := string(b)
	listaInfo := strings.Split(string(str),"\n")
	memoriaTotal := strings.Replace((listaInfo[0])[10:24], " ", "", -1)
	memorialibre := strings.Replace((listaInfo[1])[10:24], " ", "", -1)
	ramTotalKB, err1 := strconv.Atoi(memoriaTotal)
	ramFreeKB, err2 := strconv.Atoi(memorialibre)
	if err1 == nil && err2 == nil{
		ramTotalKB := ramTotalKB / 1024
		ramFreesKB := ramFreeKB / 1024
		memResponse := memStruct(ramTotalKB, ramFreeKB)
		jsonResponse, errorjson := json.Marshal(memResponse)
		if errorjson != nil{
			http.Error(w, errorjson.Error(), http.StatusInternalServerError)
			return
	
		}
	w.Header().Set("Content-Type", "application/json");
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
	}else{
		return
	}
}