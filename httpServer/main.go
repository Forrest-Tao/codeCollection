package httpServer

import "net/http"

func demo() {
	pattern := "/v1//api/greet"
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {

	})
	panic(http.ListenAndServe(":8080", nil))
}
