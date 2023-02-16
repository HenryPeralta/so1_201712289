package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strconv"
	"context"
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)

type valCalculadora struct{
	Val1 float32 `json:"Val1"`
	Operador string `json:"Operador"`
	Val2 float32 `json:"Val2"`
	Resultado float32 `json:"Resultado"`
	Fecha string `json:"Fecha"`
	Bandera bool `json:"Bandera"`
	Mensaje string `json:"Mensaje"`
}

type CalculadoraBD struct{
	Val1 float32
	Operador string
	Val2 float32
	Resultado float32
	Fecha string
	Bandera bool
	Mensaje string
}

var calcu = []CalculadoraBD{}

var valores valCalculadora

func main(){
	ctx := context.Background()
	db, err := crearConexionBd()
	if(err != nil){
		panic(err)
	}

	err = queryCalculadora(ctx, db)
	if(err != nil){
		panic(err)
	}

	request()
}

func getOperacion(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(valores)
}

func getHistorial(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calcu)
}

func createOperacion(w http.ResponseWriter, req *http.Request){
	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(w, "Insertar operacion valida")
	}

	json.Unmarshal(reqBody, &valores)
	var result float32 = 0
	bandera := true
	mensaje := "Correcto"

	switch valores.Operador {
	case "+":
		result = valores.Val1 + valores.Val2
	case "-":
		result = valores.Val1 - valores.Val2
	case "*":
		result = valores.Val1 * valores.Val2
	case "/":
		if(valores.Val2 == 0){
			result = 0
			bandera = false
			mensaje = "Error al dividir con cero"
		}else{
			result = valores.Val1 / valores.Val2
		}
	}

	numMes := ""

	anio := time.Now().Year()
	mes := time.Now().Month()
	dia := time.Now().Day()
	hora := time.Now().Hour()
	minuto := time.Now().Minute()
	segundo := time.Now().Second()

	switch strconv.Itoa(int(mes)) {
	case "1":
		numMes = "01"
	case "2":
		numMes = "02"
	case "3":
		numMes = "03"
	case "4":
		numMes = "04"
	case "5":
		numMes = "05"
	case "6":
		numMes = "06"
	case "7":
		numMes = "07"
	case "8":
		numMes = "08"
	case "9":
		numMes = "09"
	case "10":
		numMes = "10"
	case "11":
		numMes = "11"
	case "12":
		numMes = "12"
	}

	tiempoActual := strconv.Itoa(anio) + "-" + numMes + "-" + strconv.Itoa(dia) + " " + strconv.Itoa(hora) + ":" + strconv.Itoa(minuto) + ":" + strconv.Itoa(segundo)

	respuesta := valCalculadora{
		Val1: valores.Val1,
		Operador: valores.Operador,
		Val2: valores.Val2,
		Resultado: result,
		Fecha: tiempoActual,
		Bandera: bandera,
		Mensaje: mensaje,
	}

	datosJson, err := json.Marshal(respuesta)
	if err != nil {
		fmt.Println(err)
	}

	json.Unmarshal(datosJson, &valores)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(valores)

	conexionBD, err := crearConexionBd()
	if(err != nil){
		panic(err)
	}

	strval1 := fmt.Sprintf("%v", valores.Val1)
	strval2 := fmt.Sprintf("%v", valores.Val2)
	strresult := fmt.Sprintf("%v", valores.Resultado)
	strbandera := strconv.FormatBool(valores.Bandera)

	query := "insert into calculadora values(" + strval1 + ",'" + valores.Operador + "'," + strval2 + "," + strresult + ",'" + valores.Fecha + "'," + strbandera + ",'" + valores.Mensaje + "')"

	insertarRegistro, err := conexionBD.Prepare(query)
	if(err != nil){
		panic(err)
	}
	insertarRegistro.Exec()

	temp := CalculadoraBD{
		Val1: valores.Val1,
		Operador: valores.Operador,
		Val2: valores.Val2,
		Resultado: result,
		Fecha: tiempoActual,
		Bandera: bandera,
		Mensaje: mensaje,
	}

	calcu = append(calcu, temp)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
      w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, req)
		})
}

func request(){
	router := mux.NewRouter().StrictSlash(false)
	enableCORS(router)
	router.HandleFunc("/operacion", getOperacion).Methods("GET")
	router.HandleFunc("/operacion", createOperacion).Methods("POST")
	router.HandleFunc("/historial", getHistorial).Methods("GET")

	log.Println("Escuchando en http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func crearConexionBd()(*sql.DB, error){
	conexionbd := "root:familia@tcp(localhost:3306)/EjemploDB"
	db, err := sql.Open("mysql", conexionbd)
	if(err != nil){
		return nil, err
	}

	err = db.Ping()
	if(err != nil){
		return nil, err
	}

	return db, nil
}

func queryCalculadora(ctx context.Context, db *sql.DB) error{
	qry := "select c.val1, c.operador, c.val2, c.resultado, c.fecha, c.bandera, c.mensaje from calculadora c"

	rows, err := db.QueryContext(ctx, qry)
	if(err != nil){
		return nil
	}

	for rows.Next(){
		b := CalculadoraBD{}

		err = rows.Scan(&b.Val1, &b.Operador, &b.Val2, &b.Resultado, &b.Fecha, &b.Bandera, &b.Mensaje)
		if(err != nil){
			return nil
		}

		calcu = append(calcu, b)
	}

	return nil
}