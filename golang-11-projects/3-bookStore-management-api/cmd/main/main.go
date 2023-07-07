package main

// go build -o main && ./main
import (
	"log"
	"net/http"

	"github.com/Mdromi/Crash-Course/golang-11-projects/3-bookStore-management-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
