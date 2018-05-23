package main

import (
	"net/http"
)

/*
[2]
- hello() uses the http.ResponseWriter to write a response to the client
*/
func hello(w http.ResponseWriter, r *http.Request) {
	/*
	   [3]
	   .Write() takes the more general []byte, or byte-slice, as a parameter, we do a simple type conversion of our string.
	*/
	w.Write([]byte("hello"))
}
func main() {
	/*
	   [1]
	   - syntax : http.HandleFunc("/some/path/", insertFunctionName)
	   - install a handler function at the root path of our webserver.
	   - So, Every time a new request comes into the HTTP server matching the root path, the server will spawn a new goroutine executing the hello function
	*/
	http.HandleFunc("/", hello)
	/*
	   [4]
	   - start the HTTP server on port 8080
	   - Note: This is a synchronous, or blocking, call, which will keep the program alive until interrupted
	*/
	http.ListenAndServe(":8080", nil)
}
