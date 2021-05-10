package main

import (
	"context"
	"google.golang.org/grpc"
	pb "gitlab.com/renato776/proyecto-g16/backend"
	"github.com/gomodule/redigo/redis"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)
const (
	port = ":5000"
)
var conn redis.Conn
var err error

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

func (s *server) CrearAsistencia(ctx context.Context, in *pb.CrearAsistencia) (*pb.Asistencia, error) {
  var asistencia Asistencia
  asistencia.Name = in.GetName()
  asistencia.Carnet = in.GetCarnet()
  asistencia.Event = in.GetEvent()
  asistencia.Img = in.GetImg()
  publishAsistencia
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("carnet:%v",asistencia.Carnet),"usr"))
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("name:%v",asistencia.Carnet),"usr"))
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("event:%v",asistencia.Carnet),"usr"))
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("img:%v",asistencia.Carnet),"usr"))
	if err != nil { return "",
	  errors.New("Session validation failed.") }

	return &pb.Asistencia{IndividualAsistencia: asistencia}, nil
}

func (s *server) CrearReporte(ctx context.Context, in *pb.CrearReporte) (*pb.Reporte, error) {
  var reporte Report
  reporte.Name = in.GetName()
  reporte.Carnet = in.GetCarnet()
  reporte.ID = in.GetID()
  reporte.Body = in.GetBody()
  publishReporte
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("carnet:%v",reporte.Carnet),"usr"))
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("name:%v",reporte.Carnet),"usr"))
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("id:%v",reporte.Carnet),"usr"))
	_,err = redis.String(conn.Do("HSET",fmt.Sprintf("body:%v",reporte.Carnet),"usr"))
	if err != nil { return "",
	  errors.New("Session validation failed.") }

	return &pb.Reporte{IndividualReporte: reporte}, nil
}

func (s *server) Obtener(ctx context.Context, in *pb.CrearReporte) (*pb.Reporte, error) {
	resp,err = redis.String(conn.Do("HGETALL",in.GetFiltro()))
	if err != nil { return "",
	  errors.New("Session validation failed.") }

	return &pb.Reporte{ReporteCollection: resp}, nil
}


func Reporte(w http.ResponseWriter, r *http.Request) {
	var Reporte renato

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Detalle de casos inválido")
	}

	json.Unmarshal(reqBody, &Reporte)
	enviar(Reporte)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Reporte)
}


func main() {
	conn,err = redis.Dial("tcp",
	os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"))
	if err != nil {
		panic(err)
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.AsistenciaServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
