package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type datosRam struct {
	Total      int32
	Usado      int32
	Libre      int32
	Compartido int32
	Cache      int32
	Buffer     int32
	Porcentaje int32
}

type datosCpu struct {
	Procesos    []procesos
	Informacion []informacion
	Porcentaje  []porcentaje
}

type procesos struct {
	Pid    int32
	Nombre string
	Estado string
	User   int32
	Ram    int32
	Hijos  []hijos
}

type hijos struct {
	Pid    int32
	Nombre string
}

type porcentaje struct {
	PorcentajeUso int32
}

type informacion struct {
	ProcesosEjecucion    int32
	ProcesosSuspendidos  int32
	ProcesosDetenidos    int32
	ProcesosZombies      int32
	ProcesosDesconocidos int32
	TotalProcesos        int32
}

var dataram datosRam

var datacpu datosCpu
var contador int = 1

func main() {
	fmt.Println("starts")
	bandera := true
	infoCpu()

	for bandera {
		delaySecond(2)
		infoRam()
		useCpu()
	}
}

func infoRam() {
	fmt.Println("DATOS OBTENIDOS DESDE EL MODULO RAM :")

	cmd := exec.Command("sh", "-c", "cat /proc/ram_201712289")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])

	errors := json.Unmarshal([]byte(output), &dataram)
	if errors != nil {
		fmt.Println(errors)
	}

	fmt.Println(dataram)
	fmt.Println("")

	conexionBD, err := crearConexionBd()
	if err != nil {
		panic(err)
	}

	hora := time.Now().Hour()
	minuto := time.Now().Minute()
	segundo := time.Now().Second()

	tiempoActual := strconv.Itoa(hora) + ":" + strconv.Itoa(minuto) + ":" + strconv.Itoa(segundo)

	query := "insert into Ram(Total, Usado, Libre, Compartido, Cachee, Buffer, Porcentaje, Tiempo) values(" + strconv.Itoa(int(dataram.Total)) + "," + strconv.Itoa(int(dataram.Usado)) + "," + strconv.Itoa(int(dataram.Libre)) + "," + strconv.Itoa(int(dataram.Compartido)) + "," + strconv.Itoa(int(dataram.Cache)) + "," + strconv.Itoa(int(dataram.Buffer)) + "," + strconv.Itoa(int(dataram.Porcentaje)) + ",'" + tiempoActual + "')"

	insertarRegistro, err := conexionBD.Prepare(query)
	if err != nil {
		panic(err)
	}
	insertarRegistro.Exec()
}

func infoCpu() {
	fmt.Println("DATOS OBTENIDOS DESDE EL MODULO CPU :")
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_201712289")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])

	split1 := strings.Split(output, "\"Informacion\"")

	split2 := strings.Split(split1[0], "\"Hijos\"")

	tempsplit := ""

	for i := 0; i < len(split2); i++ {
		split3 := strings.Split(split2[i], "]")
		for j := 0; j < len(split3); j++ {
			if split3[j][0] == 58 {
				if len(split3[j]) != 4 {
					for z := 0; z < len(split3[j])-3; z++ {
						tempsplit = tempsplit + string([]byte{split3[j][z]})
					}
					for p := len(split3[j]) - 2; p < len(split3[j]); p++ {
						tempsplit = tempsplit + string([]byte{split3[j][p]})
					}
					tempsplit = tempsplit + "]"
				} else {
					tempsplit = tempsplit + split3[j] + "]"
				}
			} else {
				tempsplit = tempsplit + split3[j]
			}
		}
		tempsplit = tempsplit + "\"Hijos\""
	}

	tempres := ""

	for i := 0; i < len(tempsplit)-11; i++ {
		tempres = tempres + string([]byte{tempsplit[i]})
	}

	tempres = tempres + "],\"Informacion\"" + split1[1]

	errors := json.Unmarshal([]byte(tempres), &datacpu)
	if errors != nil {
		fmt.Println(errors)
	}

	fmt.Println(datacpu)
	fmt.Println("")

	conexionBD, err := crearConexionBd()
	if err != nil {
		panic(err)
	}

	querydelete := "truncate table Cpu"
	eliminarRegistroCpu, err := conexionBD.Prepare(querydelete)
	if err != nil {
		panic(err)
	}
	eliminarRegistroCpu.Exec()

	querydeleteHijos := "truncate table Hijos"
	eliminarRegistroHijos, err := conexionBD.Prepare(querydeleteHijos)
	if err != nil {
		panic(err)
	}
	eliminarRegistroHijos.Exec()

	querydeleteinfo := "truncate table InfoCpu"
	eliminarRegistroInfo, err := conexionBD.Prepare(querydeleteinfo)
	if err != nil {
		panic(err)
	}
	eliminarRegistroInfo.Exec()

	for i := 0; i < len(datacpu.Procesos); i++ {

		querycpu := "insert into Cpu(Id_Cpu, Pidp, Nombrep, Estado, Usuario, Ram) values(" + strconv.Itoa(int(contador)) + "," + strconv.Itoa(int(datacpu.Procesos[i].Pid)) + ",'" + datacpu.Procesos[i].Nombre + "','" + datacpu.Procesos[i].Estado + "'," + strconv.Itoa(int(datacpu.Procesos[i].User)) + "," + strconv.Itoa(int(datacpu.Procesos[i].Ram)) + ")"
		insertarRegistrocpu, err := conexionBD.Prepare(querycpu)
		if err != nil {
			panic(err)
		}
		insertarRegistrocpu.Exec()

		if (len(datacpu.Procesos[i].Hijos)) != 0 {
			for j := 0; j < len(datacpu.Procesos[i].Hijos); j++ {
				queryhijos := "insert into Hijos(Id_Cpu, Pidh, Nombreh) values(" + strconv.Itoa(int(contador)) + "," + strconv.Itoa(int(datacpu.Procesos[i].Hijos[j].Pid)) + ",'" + datacpu.Procesos[i].Hijos[j].Nombre + "')"
				insertarRegistrohijo, err := conexionBD.Prepare(queryhijos)
				if err != nil {
					panic(err)
				}
				insertarRegistrohijo.Exec()
			}
		}

		if contador < 2000 {
			contador = contador + 1
		} else {
			contador = 1
		}
	}

	for i := 0; i < len(datacpu.Informacion); i++ {
		queryInfo := "insert into InfoCpu(ProcesosEjecucion, ProcesosSuspendidos, ProcesosDetenidos, ProcesosZombies, ProcesosDesconocidos, TotalProcesos) values(" + strconv.Itoa(int(datacpu.Informacion[i].ProcesosEjecucion)) + "," + strconv.Itoa(int(datacpu.Informacion[i].ProcesosSuspendidos)) + "," + strconv.Itoa(int(datacpu.Informacion[i].ProcesosDetenidos)) + "," + strconv.Itoa(int(datacpu.Informacion[i].ProcesosZombies)) + "," + strconv.Itoa(int(datacpu.Informacion[i].ProcesosDesconocidos)) + "," + strconv.Itoa(int(datacpu.Informacion[i].TotalProcesos)) + ")"
		insertarRegistroInfo, err := conexionBD.Prepare(queryInfo)
		if err != nil {
			panic(err)
		}
		insertarRegistroInfo.Exec()
	}
}

func useCpu() {

	fmt.Println("DATOS OBTENIDOS DESDE EL MODULO CPU :")
	cmd := exec.Command("sh", "-c", "cat /proc/cpu_201712289")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])

	split1 := strings.Split(output, "\"Informacion\"")

	split2 := strings.Split(split1[0], "\"Hijos\"")

	tempsplit := ""

	for i := 0; i < len(split2); i++ {
		split3 := strings.Split(split2[i], "]")
		for j := 0; j < len(split3); j++ {
			if split3[j][0] == 58 {
				if len(split3[j]) != 4 {
					for z := 0; z < len(split3[j])-3; z++ {
						tempsplit = tempsplit + string([]byte{split3[j][z]})
					}
					for p := len(split3[j]) - 2; p < len(split3[j]); p++ {
						tempsplit = tempsplit + string([]byte{split3[j][p]})
					}
					tempsplit = tempsplit + "]"
				} else {
					tempsplit = tempsplit + split3[j] + "]"
				}
			} else {
				tempsplit = tempsplit + split3[j]
			}
		}
		tempsplit = tempsplit + "\"Hijos\""
	}

	tempres := ""

	for i := 0; i < len(tempsplit)-11; i++ {
		tempres = tempres + string([]byte{tempsplit[i]})
	}

	tempres = tempres + "],\"Informacion\"" + split1[1]

	errors := json.Unmarshal([]byte(tempres), &datacpu)
	if errors != nil {
		fmt.Println(errors)
	}

	fmt.Println(datacpu)
	fmt.Println("")

	hora := time.Now().Hour()
	minuto := time.Now().Minute()
	segundo := time.Now().Second()

	tiempoActual := strconv.Itoa(hora) + ":" + strconv.Itoa(minuto) + ":" + strconv.Itoa(segundo)

	conexionBD, err := crearConexionBd()
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(datacpu.Porcentaje); i++ {
		queryUso := "insert into UsoCpu(Porcentaje, Tiempo) values(" + strconv.Itoa(int(datacpu.Porcentaje[i].PorcentajeUso)) + ",'" + tiempoActual + "')"
		insertarRegistroUso, err := conexionBD.Prepare(queryUso)
		if err != nil {
			panic(err)
		}
		insertarRegistroUso.Exec()
	}
}

func delaySecond(n time.Duration) {
	time.Sleep(n * time.Second)
}

func crearConexionBd() (*sql.DB, error) {
	//conexionbd := "root:familia@tcp(localhost:3306)/Practica2SO1"
	conexionbd := "root:familia@tcp(34.27.243.144:3306)/Prac2SO1"
	db, err := sql.Open("mysql", conexionbd)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
