package main
 
import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type myData struct {
    name string 
}
 
func main() {
 
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/calc", calculate).Methods("POST")
    log.Fatal(http.ListenAndServe(":8090", router))
}
 
func calculate(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)

    var data myData
    err := decoder.Decode(&data)
    if err != nil {
        panic(err)
    }

    fmt.Println(data.name)

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(data); err != nil {
        panic(err)
    }
}