package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	details "github.com/Sri-harsha99/Go-Kubernetes/details"
	"github.com/gorilla/mux"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {

	response := map[string]string{
		"status": "OK",
		"time":   time.Now().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "App is running")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	det, err := details.GetHostName()
	if err != nil {
		panic(err)
	}
	ip := details.GetOutboundIP()
	fmt.Println(det)
	fmt.Println(ip)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthHandler)
	r.HandleFunc("/", rootHandler)
	r.HandleFunc("/details", detailsHandler)

	log.Fatal(http.ListenAndServe(":80", r))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// }

// func main() {
// 	http.HandleFunc("/", rootHandler)

// 	fs := http.FileServer(http.Dir("static/"))
// 	http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	log.Println("Listening on :80...")
// 	http.ListenAndServe(":80", nil)
// }

// package main

// import (
// 	"fmt"
// 	"unsafe"

// 	geometry "github.com/Sri-harsha99/Go-Kubernetes/geometry"

// 	"rsc.io/quote"
// )

// func area(length, width int) int {
// 	return length * width
// }

// func main() {
// 	x := 10
// 	fmt.Println("Hello, World!")
// 	fmt.Println(quote.Go())
// 	fmt.Println(x)

// 	fmt.Printf("x is of type %T and size is %d\n", x, unsafe.Sizeof(x))

// 	area := area(5, 4)
// 	fmt.Printf("Area is : %d\n", area)

// 	days := map[string]int{
// 		"Monday":    1,
// 		"Tuesday":   2,
// 		"Wednesday": 3,
// 		"Thursday":  4,
// 		"Friday":    5,
// 		"Saturday":  6,
// 		"Sunday":    7,
// 	}
// 	fmt.Println(days)
// 	area = geometry.Area(5, 9)

// 	fmt.Println(area)

// }
