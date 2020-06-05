package protobuf

import (
	"net/http"

	"github.com/golang/protobuf/jsonpb"
)

func P() {
	http.HandleFunc(("/"), p)
	http.ListenAndServe(":8080", nil)
}

func p(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	result := Http{Message: "result", Level: Level(1)}
	w.Header().Set("content-type", "application/json")
	//json.NewEncoder(w).Encode(result)
	(&jsonpb.Marshaler{OrigName: true}).Marshal(w, &result)
}
