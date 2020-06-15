package main
/*
	Archivo para el web server, este estara en el servidor de node js
*/
import (
	"net/http"
	"encoding/json"
	"html/template"
	"log"
	//importaciones para leer archivos
	"fmt"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
	"time"
	"bytes"
	"os/exec"
)

/*
	Definicion de estructuras de cada servidor
*/
type DatosServidor struct {
	Servidor1 string
	Servidor2 string
}

type Datos struct {
	RAM 	string `json:ram`
	USOR 	string `json:usor`
	OCUR	string `json:ocur`
	CPU		string `json:cpu`
}

type Process struct {
	p_pid     int
	p_usuario string
	p_estado  string
	p_mem     float64
	p_nombre  string
}

func prueba(){
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		
		log.Fatal(err)
	}

func cpuInfo(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		log.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		nombre := strings.Replace(ft[10], " ", "", 0)
		usuario := strings.Replace(ft[0], " ", "", 0)
		estado := strings.Replace(ft[7], " ", "", 0)
		mem, err := strconv.ParseFloat(ft[3], 64)
		if err != nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, usuario, estado, mem, nombre})
	}
	var cadena string
	var cadena3 string
	cadena += "<table >\n"
	cadena += "<tr>\n"
	cadena += "<th>PID</th>\n"
	cadena += "<th>Nombre</th>\n"
	cadena += "<th>Usuario</th>\n"
	cadena += "<th>Estado</th>\n"
	cadena += "<th>Porcentaje RAM</th>\n"
	cadena += "<th>Detener</th>\n"
	cadena += "</tr>"
	for _, p := range processes {
		cadena += "<tr>\n"
		log.Println("PID", p.p_pid, "usuario", p.p_usuario, "estado", p.p_estado, p.p_mem, " % de RAM", "proceso", p.p_nombre)
		cadena2 := fmt.Sprintf("%s%d%s","<td>",p.p_pid,"</td>")
		cadena += cadena2
		cadena2 = fmt.Sprintf("%s%s%s","<td>",p.p_nombre,"</td>")
		cadena += cadena2
		cadena2 = fmt.Sprintf("%s%s%s","<td>",p.p_usuario,"</td>")
		cadena += cadena2
		cadena2 =fmt.Sprintf("%s%s%s","<td>",p.p_estado,"</td>")
		cadena += cadena2
		cadena2 = fmt.Sprintf("%s%.2f%s","<td>",p.p_mem,"</td>")
		cadena += cadena2
		cadena2 = "<td><input onclick="+"\"call('2');\""+" type="+"\"button\""+" value="+"Detener"+" id="+"\"myButton1\""+"/></td>"
		cadena += cadena2
		cadena += "</tr>\n"
	}
	cadena += "</table>\n"
	fmt.Fprintf(w, cadena)
}


func monitoreo(w http.ResponseWriter, r *http.Request){
	//reviso que si la peticion es post 
	if r.Method == "POST" {
		ajax_post_data := r.FormValue("ajax_post_data")
		fmt.Println("Receive ajax post data string ", ajax_post_data)
		fmt.Println();
		w.Write([]byte(datosServ()))
	}
	
	//muestro la pagina original
	
	t, err := template.ParseFiles("index.html") //parse the html file homepage.html
	if err != nil { // if there is an error
		log.Print("Error de parsing de la plantilla: ", err) // log it
	}
	cpuInfo(w , r)	
	//Creo las variables
	datos := DatosServidor {}

	err = t.Execute(w, datos) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it	
	}
	
}

//main
func main() {
	//router := mux.NewRouter()

	//http.Handle("/", http.FileServer(http.Dir("style.css")))
	http.HandleFunc("/", monitoreo)
	//vamos a mostrarlo en el puerto 8080
	http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func datosServ() string{
	return string_servidor1() + "?" +" "
}

//aqui regreso mi string del servidor 1 osea, es donde esta ubicado mi programa de go
func string_servidor1() string{
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

//funcion para el string del servidor 2
func string_servidor2() string{
	//aqui esta la api, solo tengo que cambiar de donde conseguirla
	response, err := http.Get("http://192.168.1.34:5000/serv2")

	//aqui tenemos el error
	if err != nil {
        fmt.Print(err.Error())
       // os.Exit(1)
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
	return string(str + "%" + str1 + "%" + str2 + "%" + str3 + "%")
}