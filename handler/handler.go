package handler

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "multiplayer-mode-usage/service"
    "net/http"
    pb "multiplayer-mode-usage/proto"
)

func GetPopularModeHandler(w http.ResponseWriter, r *http.Request) {
    areaCode := mux.Vars(r)["area_code"]

    mode, count, err := service.GetPopularMode(areaCode)

    if err != nil {
        if err.Error() == "no data found for the specified area code" {
            w.WriteHeader(http.StatusNotFound)
            jsonError := map[string]string{"error": "No data found for the specified area code"}
            json.NewEncoder(w).Encode(jsonError)
            return
        }

        w.WriteHeader(http.StatusInternalServerError)
        jsonError := map[string]string{"error": "Internal server error. Please try again later."}
        json.NewEncoder(w).Encode(jsonError)
        return
    }

    response := &pb.ModeUsageResponse{
        MostPlayedMode: mode,
        PlayerCount:    int32(count),
    }

    jsonData, err := json.Marshal(struct {
        AreaCode      string `json:"area_code"`
        MostPlayedMode string `json:"most_played_mode"`
        PlayerCount    int32  `json:"player_count"`
    }{
        AreaCode:      areaCode,
        MostPlayedMode: response.MostPlayedMode,
        PlayerCount:    response.PlayerCount,
    })

    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        jsonError := map[string]string{"error": "Internal server error. Please try again later."}
        json.NewEncoder(w).Encode(jsonError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}
