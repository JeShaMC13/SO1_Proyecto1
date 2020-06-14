package main
import(
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"

	/*
		recorda importar a go con el comando
		go get -u github.com/gorilla/mux
	*/
	"github.com/gorilla/mux"
)

//json para mandar
type Datos struct {
	RAM 	string `json:ram`
	USOR 	string `json:usor`
	OCUR	string `json:ocur`
	CPU		string `json:cpu`
}

//cosos para mandar
var serv []Datos

func main(){
	//inicio mi router
	router := mux.NewRouter()

	//solo queremos obtener informacion entonces no importa
	router.HandleFunc("/serv2", getInfo).Methods("GET")
	http.ListenAndServe(":5000", router)
	log.Fatal(http.ListenAndServe(":5000", router))
}

func getInfo(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Dando datos")

	//obtengo el string
	str := strings.Split(string_servidor(), "%")
	ramT := str[0]
	libre := str[1]
	ocupada := str[2]
	cpu1 := str[3]

	serv = nil
	serv = append(serv, Datos{RAM:ramT, USOR:libre,OCUR:ocupada,CPU:cpu1})

	json.NewEncoder(w).Encode(serv)
}

//aqui van los datos 
func string_servidor() string {
	//veo la informacion de la memoria
	nombreArchivo := "/proc/meminfo"
	content, err := ioutil.ReadFile(nombreArchivo) 
	retorno := ""
	if err != nil { }

	//leo linea por linea
	lines := strings.Split(string(content), "\n")
	re := regexp.MustCompile("[0-9]+") //regex

	//en esta parte encuentro el valor numerico y lo transformo en numero, luego lo transformo a mb
	match := re.FindStringSubmatch(lines[0])
	mem_total := match[0]
	mem_t2, err := strconv.Atoi(mem_total)
	if err == nil {}
	retorno += fmt.Sprintf("%f",float32(mem_t2)*0.001) + "%"
	
	match = re.FindStringSubmatch(lines[1])
	mem_libre := match[0]
	mem_l2, err := strconv.Atoi(mem_libre)
	if err == nil {}

	match = re.FindStringSubmatch(lines[3])
	mem_libre3 := match[0]
	mem_l3, err := strconv.Atoi(mem_libre3)
	if err == nil {}

	match = re.FindStringSubmatch(lines[4])
	mem_libre4 := match[0]
	mem_l4, err := strconv.Atoi(mem_libre4)
	if err == nil {}
	
	//procentaje
	porcentaje1 := (1 -((float32(mem_l2)+float32(mem_l3)+float32(mem_l4)) / float32(mem_t2)))*100
	porcentaje2 := ((float32(mem_l2)+float32(mem_l3)+float32(mem_l4)) / float32(mem_t2))*100
	retorno += fmt.Sprintf("%f",porcentaje1) + "%" + fmt.Sprintf("%f",porcentaje2) + "%"

	//retorno
	return retorno + cpu_info_1()
}

//este me regresa la info del cpu
func cpu_info_1() string{
	//leo el archivo stat y veo la primera linea
	nombreArchivo := "/proc/stat"
	content, err := ioutil.ReadFile(nombreArchivo) 

	if err != nil { }
	lines := strings.Split(string(content), "\n")

	//obtengo los valores del cpu
	info := strings.Split(lines[0], " ")

	val1,err1 := strconv.Atoi(info[2])
	if err1 == nil {}

	val2,err2 := strconv.Atoi(info[3])
	if err2 == nil {}

	val3,err3 := strconv.Atoi(info[4])
	if err3 == nil {}

	val4,err4 := strconv.Atoi(info[5])
	if err4 == nil {}

	val5,err5 := strconv.Atoi(info[6])
	if err5 == nil {}

	val6,err6 := strconv.Atoi(info[7])
	if err6 == nil {}

	val7,err7 := strconv.Atoi(info[8])
	if err7 == nil {}

	val8,err8 := strconv.Atoi(info[9])
	if err8 == nil {}

	//val9,err9 := strconv.Atoi(info[10])
	//if err9 == nil {}

	//val10,err10 := strconv.Atoi(info[11])
	//if err10 == nil {}

	//actuales
	idle := float32(val4) + float32(val5)
	nonidle := float32(val1) + float32(val2) + float32(val3) + float32(val6) + float32(val7) + float32(val8)
	totalid := idle + nonidle
	
	//lo pongo a dormir
	time.Sleep(3*time.Second)

	return cpu_info_2(idle, nonidle, totalid)
}

//lo uso para calcular por 2da vez el estado del cpu
func cpu_info_2(old_idle float32, old_nonidle float32, old_totalid float32) string{
	//leo el archivo stat y veo la primera linea
	nombreArchivo := "/proc/stat"
	content, err := ioutil.ReadFile(nombreArchivo) 

	if err != nil { }
	lines := strings.Split(string(content), "\n")

	//obtengo los valores del cpu
	info := strings.Split(lines[0], " ")

	val1,err1 := strconv.Atoi(info[2])
	if err1 == nil {}

	val2,err2 := strconv.Atoi(info[3])
	if err2 == nil {}

	val3,err3 := strconv.Atoi(info[4])
	if err3 == nil {}

	val4,err4 := strconv.Atoi(info[5])
	if err4 == nil {}

	val5,err5 := strconv.Atoi(info[6])
	if err5 == nil {}

	val6,err6 := strconv.Atoi(info[7])
	if err6 == nil {}

	val7,err7 := strconv.Atoi(info[8])
	if err7 == nil {}

	val8,err8 := strconv.Atoi(info[9])
	if err8 == nil {}

	//val9,err9 := strconv.Atoi(info[10])
	//if err9 == nil {}

	//val10,err10 := strconv.Atoi(info[11])
	//if err10 == nil {}

	//actuales
	idle := float32(val4) + float32(val5)
	nonidle := float32(val1) + float32(val2) + float32(val3) + float32(val6) + float32(val7) + float32(val8)
	totalid := idle + nonidle

	//saco las diferencias
	diff_idle := idle - old_idle
	//diff_non := nonidle - old_nonidle
	diff_total := totalid - old_totalid

	//calculo el procentaje
	porcentaje := ((diff_total-diff_idle)/diff_total)*100
	return fmt.Sprintf("%f",porcentaje)
}