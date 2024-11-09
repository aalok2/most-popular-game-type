package middleware

import (
    "log"
    "encoding/json"
        "github.com/gorilla/mux"
    "net/http"
)


func ValidateAreaCodeMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
         areaCode := mux.Vars(r)["area_code"]

            log.Printf("Lenght of area code", len(areaCode))

        if len(areaCode) != 3 {
            w.WriteHeader(http.StatusBadRequest)
            jsonError := map[string]string{"error": "area_code must be exactly three characters long"}
            json.NewEncoder(w).Encode(jsonError)
            return
        }
        next.ServeHTTP(w, r)
    })
}
