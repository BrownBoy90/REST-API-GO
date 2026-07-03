package middlewares

import "net/http"

// we are giving it the argument next so that we can pass that next handler to the function which is returned inside the security headers middleware.
func SecurityHeaders(next http.Handler) http.Handler {
	// At the end of the middleware, we are going to call next, and this is going to receive the response and the request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Now let's add some security headers to our response that we are going to send back to the client
		w.Header().Set("X-DNS-Prefetch-Control", "off") // Disables DNS Pre-fetching which is a mechanism that allows browsers to resolve domain names in the background while a user is browsing a web page. Disabling this can reduce the risk of DNS related attacks and improve privacy.

		// Internally, when this header is set, browsers respect the directive and avoid prefetching DNS requests

		// Without this header browsers may pre-fetch DNS requests, which can expose users to privacy concerns or unnecessary DNS traffic.
		w.Header().Set("X-Frame-Options", "DENY")
		// This header prevents the web page from being displayed in an iframe on other websites.

		// This is a defense against clickjaching attacks, where a user is tricked into clicking something different
		w.Header().Set("X-XSS-Protections", "1;mode=block")
		// This header enables cross-site scripting filter built into most modern web browsers and instructs the browser to block the page if an XSS attack is detected. Internally, the browser activates its cross-site scripting XSS protection mechanisms. Without this header, the browser's cross-site scripting filter might still be enabled by default, but not all browsers enforce this without the header, and some may simply sanitize this script instead of blocking the page
		w.Header().Set("X-Content-Type-Options", "nosniff")
		// This header prevents browsers from Mime sniffing a response away from the declared content type. It ensures that the files are served with the correct Mime types, reducing the risk of certain attacks example cross site scripting via content type diffusion Mime means multi-purpose internet mail Extentions and it is an extention of the original SMTP email protocol, which is simple mail transport protocol. It lets users Mime, lets users exchange different kind of data including audio, video , images and application programs over email and mime type is now called media type but also sometimes called content type and it is a string sent along with a file indicating the type of the file. It describes the content format for eg a sound file might be labelled as audio/ogg oran image file may be tagged as image/png. Now internally using this header browsers are instructed to trust the server provided content type and they do not attempt to guess the content type without this header. Browsers might try to sniff the content type and could misinterpret files, potentially leading to security vulnerabilities
		w.Header().Set("Strict-Transport-Security", "max-age=63072000;includeSubDomains;preload")
		// What this does is that it enforces HTTPS for specified max age in seconds. It tells browsers to interact with your site only over HTTPS and to remember this for the specified duration. By using header, the browsers will refuse to connect to the site via HTTP after receiving this header and will instead only connect via HTTPS. Without this header, users could inadvertently access your site via HTTP, exposing them to man in the middle attacks.
		w.Header().Set("Content-Security-Policy", "default-src 'self'")
		//  Next we have content security policy which is default src self. Now this header controls which resources can be loaded on the page by setting the default SRC self, you are telling the browser to only load resources from the same origin, effectively mitigating many types of attacks such as cross-site scripting attacks. Internally, the browser enforces these restrictions and blocks any resources that don't match the policy, without this header, the browser would allow any content to be loaded, increasing the risk of cross-site scripting, data injection, and other attacks.
		w.Header().Set("Referrer-Policy", "no-referrer")
		// Finally we have the referrer policy, this header controls how much referrer information should be included with requests made from your site. No referrer ensures that information is sent with requests. Internally, the browsers will strip out the referrer information from the request made from your site and without this header, the browsers might include full referrer information by default, potentially leaking sensitive information to third party sites.
		w.Header().Set("X-Powered-By", "Django") // Misnaming the backend stack as seen on dev tools in browser

		// UNderstand these headers HW:
		w.Header().Set("Server", "")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Permissions-Policy", "geolocation=(self), microphone=()")
		next.ServeHTTP(w, r)
	})

}

// ==== Skeleton of a Middleware ====
// func securityHeaders(next http.Handler) http.Handler {
// 	// At the end of the middleware, we are going to call next, and this is going to receive the response and the request
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
// next.ServeHTTP(w,r)
// 	})

// }
