package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	p_pid     int
	p_usuario string
	p_estado  string
	p_mem     float64
	p_nombre  string
}

func main() {
	//http.HandleFunc("/memoria", ramInfo)
	fmt.Println("El servidor esta a la escucha en el puerto 3000")
	http.HandleFunc("/cpu", cpuInfo)
	http.ListenAndServe(":3000", nil)
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
		cadena2 = fmt.Sprintf("%s%s%s","<td>","preba","</td>")
		cadena += cadena2
		cadena += "</tr>\n"
	}
	cadena += "<\table>\n"
	fmt.Fprintf(w, cadena)
}
