package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"strconv"
	"strings"
	"sync"
)

type Teacher struct {
	ID        int
	FirstName string
	LastName  string
	Class     string
	Subject   string
}

var (
	teachers = make(map[int]Teacher)
	mutex    = &sync.Mutex{}
	nextID   = 1
)

// Initialize some dummy data
func init() {
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "John",
		LastName:  "Doe",
		Class:     "9A",
		Subject:   "Math",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Smith",
		Class:     "10A",
		Subject:   "Algebra",
	}
	nextID++
	teachers[nextID] = Teacher{
		ID:        nextID,
		FirstName: "Jane",
		LastName:  "Doe",
		Class:     "11A",
		Subject:   "Biology",
	}

}

func getTeachershandler(w http.ResponseWriter, r *http.Request) {

	// fmt.Println(r.URL.Path)
	path := strings.TrimPrefix(r.URL.Path, "/teachers/")
	idStr := strings.TrimSuffix(path, "/")
	fmt.Println(idStr)

	if idStr == "" {
		firstName := r.URL.Query().Get("first_name")
		lastName := r.URL.Query().Get("last_name")

		teacherList := make([]Teacher, 0, len(teachers))
		for _, teacher := range teachers {
			if (firstName == "" || teacher.FirstName == firstName) && (lastName == "" || teacher.LastName == lastName) {
				teacherList = append(teacherList, teacher)
			}
		}
		response := struct {
			Status string    `json:"status"`
			Count  int       `json:"count"`
			Data   []Teacher `json:"data"`
		}{
			Status: "success",
			Count:  len(teacherList),
			Data:   teacherList,
		}

		w.Header().Set("Conten-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

	// Handle Path Parameter
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	teacher, exists := teachers[id]
	if !exists {
		http.Error(w, "Teacher not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(teacher)
}

func addTeacherHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newTeachers []Teacher
	err := json.NewDecoder(r.Body).Decode(&newTeachers)
	if err!=nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
	}

	addedTeachers := make([]Teacher, len(newTeachers))
	for i, newTeacher := range newTeachers {
		newTeacher.ID = nextID
		teachers[nextID] = newTeacher
		addedTeachers[i] = newTeacher
		nextID++
	}
	w.Header().Set("Content-Type", "application/json")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		getTeachershandler(w, r)
	case http.MethodPost:
		// Post request handler
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Teachers Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Teachers Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Teachers Route"))

	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// Call GET method handler function
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Students Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Students Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method on Students Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Students Route"))

	}
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method on Execs Route"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method on Execs Route"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method on Execs Route"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH  method on Execs Route"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method on Execs Route"))

	}
}
func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	// Multiplexer refers to a request multiplexer which is a router that matches incomming HTTP requests to their respective handlers based on request URL and method

	// http.ServeMUX is the default HTTP request mutiplexer provided by the Go standard library

	// Why use MUX?
	// Mux allows you to define multiple routes, that is multiple end points for your API.

	// Each route can have its own handler function, enabling you to organize your api better.

	// Also MUX helps separating the logic for different routes, making the code cleaner and more maintainable

	// MUX is helful only when we have multiple routes(alot)

	// As number of endpoints increase, then MUX will play a very significant role in organizing and managing our code, managing our routes.

	// So if we want to group related routes or apply a middleware to a specific set of routes, using a mux makes this easier

	// Also when we want to use custom handlers or middlewares having mux allows you to easily apply those to a specific set of routes or requests.

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// rl := mw.NewRateLimiter(5, time.Minute)

	// hppOptions := mw.HPPOptions{
	// 	CheckQuery:                  true,
	// 	CheckBody:                   true,
	// 	CheckBodyOnlyForContentType: "application/x-www-form-urlencoded",
	// 	Whitelist:                   []string{"sortBy", "sortOrder", "name", "age", "class"},
	// }

	// secureMux := mw.Cors(rl.Middleware(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Compression(mw.Hpp(hppOptions)((mux)))))))
	// secureMux := applyMiddleware(mux, mw.Hpp(hppOptions), mw.Compression, mw.SecurityHeaders, mw.ResponseTimeMiddleware, rl.Middleware, mw.Cors)
	secureMux := mw.SecurityHeaders(mux)

	// Create custom server
	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on port:", 3000)
	err := server.ListenAndServeTLS(cert, key)

	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
}

// If server is rendering web pages, images, graphics, then in that case using a compression middleware will proove to be very efficient. But if it is a simple static website with small images, it is not a heav y load then it may not be a useful choice

// Middleware is a function that wraps an http.handler with additional functionality
type Middleware func(http.Handler) http.Handler

// Appply Middlewares
func ApplyMiddleware(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)

	}
	return handler
}
