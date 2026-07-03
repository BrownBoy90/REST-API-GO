package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os" // Added to read environment variables
	mw "restapi/internal/api/middlewares"
)

type user struct {
    Name string `json:"name"`
    Age  string `json:"age"`
    City string `json:"city"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

    w.Write([]byte("Hello Root Route"))
    fmt.Println("Hello Root Route")
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case http.MethodGet:
        w.Write([]byte("Hello GET method on Teachers Route"))
        // fmt.Println("Hello GET method on Teachers Route")
    case http.MethodPost:
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
        w.Write([]byte("Hello PATCH  method on Execs Route"))
        fmt.Println("Hello PATCH method on Execs Route")
    case http.MethodDelete:
        w.Write([]byte("Hello DELETE method on Execs Route"))
        fmt.Println("Hello DELETE method on Execs Route")

    }
}
func main() {
    // Dynamically look for Cloud Run's port, default to "8080" if not provided
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    // Prefix with a colon for http.Server formatting (e.g., ":8080")
    addr := ":" + port

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

    // Create custom server
    server := &http.Server{
        Addr: addr, // Using the dynamic port address here
        Handler: mw.Compression(mw.ResponseTimeMiddleware(mw.SecurityHeaders(mw.Cors(mux)))),
        TLSConfig: tlsConfig,
    }

    fmt.Println("Server is running on port:", port)
    err := server.ListenAndServeTLS(cert, key)

    if err != nil {
        log.Fatalln("Error starting the server", err)
    }
}

// If server is rendering web pages, images, graphics, then in that case using a compression middleware will proove to be very efficient. But if it is a simple static website with small images, it is not a heav y load then it may not be a useful choice