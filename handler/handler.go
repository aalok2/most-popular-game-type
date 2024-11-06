package handler

import (
    "net/http"
    "github.com/gorilla/mux"
    pb "multiplayer-mode-usage/proto"
    "multiplayer-mode-usage/service"
       "encoding/json"
)
func GetPopularModeHandler(w http.ResponseWriter, r *http.Request) {
    areaCode := mux.Vars(r)["area_code"]
    mode, count, _ := service.GetPopularMode(areaCode)

    response := &pb.ModeUsageResponse{
        MostPlayedMode: mode,
        PlayerCount:    int32(count),
    }

    
    jsonData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData) 

}