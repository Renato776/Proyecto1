package main

import (
	"log"
  "strconv"
	"fmt"
	"google.golang.org/grpc"
	pb "gitlab.com/renato776/proyecto-g16/gateway"

  "crypto/tls"
  "log"
  "net/http"
	"github.com/gorilla/mux"
	"encoding/json"

  "golang.org/x/crypto/acme/autocert"
	"io/ioutil"
)

type Report struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Carnet int `json:"carnet"`
	Course string `json:"course"`
	Body string `json:"body"`
	Server string `json:"server,omitempty"`
}

type Asistencia struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Event string `json:"event"`
	Carnet int `json:"carnet"`
	Img string `json:"img"`
	Server string `json:"server,omitempty"`
}
type ReportsCollection struct {
	Reports []Report `json:"reports"`
	Attended string `json:"attended"`
}

type IndividualReport struct {
	Report
	Attended string `json:"attended"`
}


type AsistenciasCollection struct {
	Asistencias []Asistencia `json:"reports"`
	Attended string `json:"attended"`
}

type IndividualAsistencia struct {
	Asistencia
	Attended string `json:"attended"`
}

func OptionsReq(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,POST,GET")
	w.Header().Set("Allow", "OPTIONS,POST,GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
}
func crearAsistencia(w http.ResponseWriter, r *http.Request) {
	var nuevoAsistencia renato

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Asistencia inválido")
	}

	json.Unmarshal(reqBody, &nuevoAsistencia)
	asistir(nuevoAsistencia)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(renato)
}
func crearReporte(w http.ResponseWriter, r *http.Request) {
	var nuevoReporte renato

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Reporte inválido")
	}

	json.Unmarshal(reqBody, &nuevoReporte)
	enviar(nuevoReporte)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(renato)
}


func asistir(nuevoAsistencia Asistencia) {
	addr := "10.0.30.56:5000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err, "Failed to stablish connection")
	}
	defer conn.Close()

	client := pb.NewAsistenciaClient(conn)
	req := &pb.Asistencia{
		Carnet: nuevoAsistencia.Carnet,
		Name: nuevoAsistencia.Name,
		Event: nuevoAsistencia.Event,
		Img: nuevoAsistencia.Img
	}

	resp, err := client.CrearAsistencia(context.Background(), req)
	if err != nil {
		panic(err, "Failed to send data")
	}
}

func enviar(nuevoReporte Report) {
	addr := "10.0.30.56:5000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err, "Failed to stablish connection")
	}
	defer conn.Close()

	client := pb.NewReporteClient(conn)
	req := &pb.Report{
		Carnet: nuevoReporte.Carnet,
		Name: nuevoReporte.Name,
		Course: nuevoReporte.Course,
		Body: nuevoReporte.Body
	}

	resp, err := client.CrearReporte(context.Background(), req)
	if err != nil {
		panic(err, "Failed to send data")
	}
}

func obtenerReporte() {
	addr := "10.0.30.56:5000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err, "Failed to stablish connection")
	}
	defer conn.Close()

	vars := mux.Vars(r)
	if _, ok := vars["id"]; ok {
    id,err := strconv.Atoi(vars["id"])
		if err != nil {
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Invalid format for ID"})
			return
		}
    client := pb.NewReporteClient(conn)
    req := &pb.Filtro{ Parametro: id }

    resp, err := client.ObtenerReporte(context.Background(), req)
    if err != nil {
      panic(err, "Failed to send data")
    }
    genericResponse(w, http.StatusOK, IndividualReport{resp, "Served by: "+ served})
  } else {
    genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Failed to connect"})
  }
}

func obtenerAsistenciaCarnet() {
	addr := "10.0.30.56:5000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err, "Failed to stablish connection")
	}
	defer conn.Close()

	vars := mux.Vars(r)
	if _, ok := vars["carnet"]; ok {
    carnet,err := strconv.Atoi(vars["carnet"])
		if err != nil {
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Invalid format for ID"})
			return
		}
    client := pb.NewAsistenciaClient(conn)
    req := &pb.Filtro{ Parametro: carnet }

    resp, err := client.ObtenerAsistencia(context.Background(), req)
    if err != nil {
      panic(err, "Failed to send data")
    }
    genericResponse(w, http.StatusOK, IndividualReport{resp, "Served by: "+ served})
  } else {
    genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Failed to connect"})
  }
}
func obtenerReporteCarnet() {
	addr := "10.0.30.56:5000"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err, "Failed to stablish connection")
	}
	defer conn.Close()

	vars := mux.Vars(r)
	if _, ok := vars["carnet"]; ok {
    carnet,err := strconv.Atoi(vars["carnet"])
		if err != nil {
			genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Invalid format for ID"})
			return
		}
    client := pb.NewReporteClient(conn)
    req := &pb.Filtro{ Parametro: carnet }

    resp, err := client.ObtenerReporte(context.Background(), req)
    if err != nil {
      panic(err, "Failed to send data")
    }
    genericResponse(w, http.StatusOK, IndividualReport{resp, "Served by: "+ served})
  } else {
    genericResponse(w, http.StatusInternalServerError, ErrorResponse{http.StatusInternalServerError, "Failed to connect"})
  }
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/grpc", nuevoReporte).Methods("POST")

  router := mux.NewRouter()
  router.HandleFunc("/reports", CreateNewReport).Methods(http.MethodPost)
  router.HandleFunc("/reports", OptionsReq).Methods(http.MethodOptions)
  router.HandleFunc("/reports", GetReports).Methods(http.MethodGet)
  router.HandleFunc("/reports/{id}", GetReports).Methods(http.MethodGet)
  router.HandleFunc("/reports/carnet/{carnet:[0-9]*}", GetReportsByCarnet).Methods(http.MethodGet)

  router.HandleFunc("/asistencia", CreateNewAsistencia).Methods(http.MethodPost)
  router.HandleFunc("/asistencia", OptionsReq).Methods(http.MethodOptions)
  router.HandleFunc("/asistencias", GetAsistenciasByEvent).Methods(http.MethodGet)
  router.HandleFunc("/asistencias/{id}", GetAsistenciasByEvent).Methods(http.MethodGet)
  router.HandleFunc("/asistencias/carnet/{carnet:[0-9]*}", GetAsistenciasByCarnet).Methods(http.MethodGet)
  http.Handle("/", router)

  server := &http.Server{
      Addr: ":https",
      TLSConfig: &tls.Config{
          GetCertificate: certManager.GetCertificate,
      },
  }

  go http.ListenAndServe(":http", certManager.HTTPHandler(nil))

  fmt.Println("Server ready on port 443")
  log.Fatal(server.ListenAndServeTLS("", ""))
}
