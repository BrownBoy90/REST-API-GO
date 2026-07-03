// using compression middleware in go can be very beneficial for improving the performance of your web applications, compression reduces the size of the response body sent over the network, which can significantly decrease loading times for your application. This is especially important for large assets like images, style sheets and Javascript files.

// By compressing responses, you can minimize the amount of data sent over the network, reducing badwidth costs and overall efficiency. This can be particularly advantageous for applications with high traffic or those serving users in bandwidth constrained environments. Faster loading time leads to better user experience, which can improve theuser satisfaction and retention. Compression helps ensure that users receive content as quickly as possible, particularly for users on slower internet connection.

// Our compression middleware can automatically handle content negotiation based on the accept encoding header sent by the client. This means that the client that support compression will receive compressed responses, while others will receive uncompressed ones
package middlewares

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func Compression(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts gzip encoding
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip"){
			next.ServeHTTP(w,r)
		}
		// Set the response header
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()
		// Whether it is a file or if it is a standard input output writer, we close that writer

		// Wrap the ResponseWriter
		w = &gzipResponseWriter{ResponseWriter: w, Writer: gz}


		
		next.ServeHTTP(w,r)
		fmt.Println("Sent the response from Compression Middleware")
	})
}

// gzipResponseWriter wraps http.ResponseWriter to write gzipped responses
// We already know it is a interface and we'll pass it not with a named field, but as an anonymous field so that we can directly access its properties, its methods and then what we can do is we will define our own method which will be under the same name as the method of the HTTP response writer 

// by including an anonymous interface we trick go compilor into thinking that this is standard http.ResponseWriter
type gzipResponseWriter struct{
	http.ResponseWriter
	Writer *gzip.Writer
}

// Overriding the write method
func (g *gzipResponseWriter) Write (b []byte) (int, error) {
	return g.Writer.Write(b)
}