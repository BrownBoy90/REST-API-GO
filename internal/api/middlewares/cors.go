// CORS - Cross-Origin Resource Sharing
// It is a security feature implemented in web browsers that restricts web pages from making requests to a domain different from the one that served the web page
// THis is crucial for preventing the malicious attacks, but it can be a limitation during development or when building APIs that need to be accessed from different origins.
// These middlewares we are talking about are for production phase, in development phase, we can inactivate them while we are testing our API continuously during development

// CORS middleware allows you to configure which origins are permitted to access your resources. When a client makes a request to a resource on a different origin. For eg from HTTP localhost:3000 to http.api.example.com . The browser checks whether the server's response includes the appropriate CORS headers.

package middlewares

import (
	"fmt"
	"net/http"
)

// api is hosted at www.myapi.com
// These days we also do server side rendering so frontend is also hosted on server and that server is generating the front end from a server

// Frontend Server := www.myfrontend.com

// If you want to allow your API to be accessible by multiple domains and sometimes it does happen, may be we have 2 or 3 domains, in that case we can make a list of allowed origin

// Allowed Origins
var allowedOrigins = []string{
	"https://my-origin-url.com",
	"https:localhost:3000",
	"https://www.myfrontend.com",
}

func Cors(next http.Handler)  http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Here in course we are going to extract a header value from the request
		origin := r.Header.Get("Origin")
		fmt.Println(origin)
		
		if isOriginAllowed(origin) {
			// Usually if we are accepting origin, then we also set a header which is Access-Control-Allow-Origin
			w.Header().Set("Access-Control-Allow-Origin", origin)

		} else {
			http.Error(w, "Not Allowed by CORS", http.StatusForbidden)
			return
		}
		// w.Header().Set()
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "3600") //60sec*60mins

		// In the CORS handler, we usually also check for pre-flight requests

		if r.Method == http.MethodOptions{
			return
		}
		// We don't want to pass options method to handlers because it is just a pre-flight check it returns immidiately, allowing the request without calling the next handler, so this is just a pre-flight check performed by the browsers sometimes


		next.ServeHTTP(w,r)
	})
}

func isOriginAllowed(origin string) bool {
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin{
			return true
		}
	}
	return false
}