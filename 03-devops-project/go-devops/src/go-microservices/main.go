package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	det "go-devops/details"

	"github.com/gorilla/mux"
)

// healthHandler, this will throw a json response
func healthHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Checking application health at http://localhost/health")
	response := map[string]string{
		"status":    "UP",
		"timestamp": time.Now().String(),
	}
	json.NewEncoder(w).Encode(response)

}

// rootHandler, this will throw a json response
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Homepage available at http://localhost/")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Application is up and running")

}

// detailsHandler
func detailsHandler(w http.ResponseWriter, r *http.Request) {

	// picking up the details from "deatils" package
	log.Println("Fetching the details http://localhost/details")
	hostname, err := det.GetHostName()
	if err != nil {
		panic(err)
	}
	IP, _ := det.GetIP()
	fmt.Println(hostname, IP)

	// printing the details as a JSON response
	response := map[string]string{
		"hostname": hostname,
		"ip":       IP.String(),
	}
	json.NewEncoder(w).Encode(response)

}

func main() {
	// resgistering new router
	r := mux.NewRouter()

	// handling function using API call
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Println("Web server has started at http://localhost")
	// to print out any sort of fatal error
	log.Fatal(http.ListenAndServe(":80", r))
}

// Simple web server to practice

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// // rootHandler
// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token %s\n", r.URL.Path, r.URL.Query().Get("token"))
// 	// try checking if it's working or not
// 	// enter http://localhost?token=12345
// }
// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	// once this is done a static folder needs to be created in go-microservices directory

// 	log.Println("Web server has started")
// 	http.ListenAndServe(":80", nil)
// }

// Simple code to practice
// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	geo "go-devops/geometry"

// 	"rsc.io/quote"
// )

// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = 2 * (length + width)
// 	return
// }

// func main() {
// 	x := 10
// 	name := "Asutosh Panda"
// 	isWorking := false

// 	fmt.Println("\nHello World!")
// 	fmt.Println(quote.Go())
// 	fmt.Println(x, name, isWorking)
// 	fmt.Printf("\nType of name %T and size is %d", name, unsafe.Sizeof(name))

// 	a, p := rectProps(2, 4)
// 	fmt.Printf("\nArea is %f and Perimeter is %f", a, p)

// 	var daysOfTheMonth = map[string]int{"Jan": 31, "Feb": 28}
// 	fmt.Println(daysOfTheMonth)

// 	area := geo.Area(1, 2)
// 	diag := geo.Diagonal(1, 2)
// 	fmt.Println(area, diag)
// }
