package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type user struct {
	Name string `json:"name"`
	Age  string `json:"age"`
	City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello Root Route")
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	// teachers/{id}
	// teachers/9
	// teachers/?key=value&sortby=email&sortorder=ASC

	switch r.Method {
	case http.MethodGet:
		fmt.Println(r.URL.Path) // /teachers/
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		// fmt.Println(path)

		userID := strings.TrimSuffix(path, "/")
		fmt.Println(userID)

		fmt.Println("Query Params", r.URL.Query())

		queryParams := r.URL.Query()
		sortby := queryParams.Get("sortby")
		key := queryParams.Get("key")
		sortorder := queryParams.Get("sortorder")

		fmt.Printf("Sortby: %v, Sort Order: %v, Key: %v", sortby, key, sortorder)

		if sortorder == "" {
			sortorder = "DESC"
		}

		w.Write([]byte("Hello GET method on Teachers Route"))
		// fmt.Println("Hello GET method on Teachers Route")
		// fmt.Println("Body: ", r.Body)
		// fmt.Println("Form: ", r.Form)
		// fmt.Println("Header:", r.Header)
		// fmt.Println("Context: ", r.Context())
		// fmt.Println("ContextLength:", r.ContentLength)
		// fmt.Println("Host: ", r.Host)
		// fmt.Println("Method: ", r.Method)
		// fmt.Println("Protocol:", r.Proto)
		// fmt.Println("Remote Addr:", r.RemoteAddr)
		// fmt.Println("Request URI:", r.RequestURI)
		// fmt.Println("TLS:", r.TLS)
		// fmt.Println("Trailer:", r.Trailer)
		// fmt.Println("Transfer Encoding:", r.TransferEncoding)
		// fmt.Println("URL:", r.URL)
		// fmt.Println("User Agent:", r.UserAgent())
		// fmt.Println("Port:", r.URL.Port())
	case http.MethodPost:
		// Parse form data (necessary for x-www-form-urlencoded)
		// err := r.ParseForm()
		// if err != nil {
		// 	http.Error(w, "Error parsing form", http.StatusBadRequest)
		// 	return
		// }

		// fmt.Println("Form:", r.Form)

		// // Prepare response data
		// response := make(map[string]interface{})

		// for key, value := range r.Form {
		// 	response[key] = value[0]
		// }
		// fmt.Println("Processed Response Map:", response)

		// // RAW Body
		// body, err := io.ReadAll(r.Body)
		// if err != nil {
		// 	return
		// }
		// defer r.Body.Close()
		// // When we are accessing a body then we need to make sure that we need to close that body

		// fmt.Println("RAW Body:", body)
		// fmt.Println("Raw Body:", string(body))

		// // If you expect a JSON then unmarshall it
		// var userInstance user
		// err = json.Unmarshal(body, &userInstance)
		// if err != nil {
		// 	return
		// }

		// fmt.Println("Unmarshalled JSON into an instance of user struct", userInstance)

		// // Prepare response data
		// response1 := make(map[string]interface{})

		// for key, value := range r.Form {
		// 	response[key] = value[0]
		// }

		// err = json.Unmarshal(body, &response1)
		// if err != nil {
		// 	return
		// }

		// fmt.Println("Unmarshalled JSON into a map", response)
		// // fmt.Println(userInstance)

		// fmt.Println("Body: ", r.Body)
		// fmt.Println("Form: ", r.Form)
		// fmt.Println("Header:", r.Header)
		// fmt.Println("Context: ", r.Context())
		// fmt.Println("ContextLength:", r.ContentLength)
		// fmt.Println("Host: ", r.Host)
		// fmt.Println("Method: ", r.Method)
		// fmt.Println("Protocol:", r.Proto)
		// fmt.Println("Remote Addr:", r.RemoteAddr)
		// fmt.Println("Request URI:", r.RequestURI)
		// fmt.Println("TLS:", r.TLS)
		// fmt.Println("Trailer:", r.Trailer)
		// fmt.Println("Transfer Encoding:", r.TransferEncoding)
		// fmt.Println("URL:", r.URL)
		// fmt.Println("User Agent:", r.UserAgent())
		// fmt.Println("Port:", r.URL.Port())

		w.Write([]byte("Hello POST method on Teachers Route"))
		fmt.Println("Hello POST method on Teachers Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Teachers Route"))
		fmt.Println("Hello PUT method on Teachers Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
		fmt.Println("Hello PATCH method on Teachers Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))
		fmt.Println("Hello DELETE method on Teachers Route")

	}

	// w.Write([]byte("Hello Teachers Route"))
	// fmt.Println("Hello Teachers Route")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		w.Write([]byte("Hello GET method on Students Route"))
		fmt.Println("Hello GET method on Students Route")

	case http.MethodPost:

		w.Write([]byte("Hello POST method on Students Route"))
		fmt.Println("Hello POST method on Students Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students Route"))
		fmt.Println("Hello PUT method on Students Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students Route"))
		fmt.Println("Hello PATCH method on Students Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students Route"))
		fmt.Println("Hello DELETE method on Students Route")

	}
	w.Write([]byte("Hello Students Route"))
	fmt.Println("Hello Students Route")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:

		w.Write([]byte("Hello GET method on Execs Route"))
		fmt.Println("Hello GET method on Execs Route")

	case http.MethodPost:

		w.Write([]byte("Hello POST method on Execs Route"))
		fmt.Println("Hello POST method on Execs Route")
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Execs Route"))
		fmt.Println("Hello PUT method on Execs Route")
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Execs Route"))
		fmt.Println("Hello PATCH method on Execs Route")
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Execs Route"))
		fmt.Println("Hello DELETE method on Execs Route")

	}
	w.Write([]byte("Hello Execs Route"))
	fmt.Println("Hello Execs Route")
}
func main() {
	port := ":3000"

	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students/", studentsHandler)

	http.HandleFunc("/execs/", execsHandler)

	fmt.Println("Server is running on port:", 3000)
	err := http.ListenAndServe(port, nil)

	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}
