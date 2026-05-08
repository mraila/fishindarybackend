package main

import (
	"fishindary/handler"
	"fishindary/service"
	"fishindary/store"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	st := store.NewStore()
	spotSvc := service.NewSpotService(st)
	catchSvc := service.NewCatchService(st, spotSvc)

	catchHandler := handler.NewCatchHandler(catchSvc)
	spotHandler := handler.NewSpotHandler(spotSvc, catchSvc)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ALLES GUT MEINE FREUNDE")
	})

	http.HandleFunc("/catches", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			catchHandler.CreateCatch(w, r)
		case http.MethodGet:
			catchHandler.GetCatches(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/spots", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			spotHandler.CreateSpot(w, r)
		case http.MethodGet:
			spotHandler.GetSpots(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/spots/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/catches") {
			spotHandler.GetCatchesBySpotID(w, r)
			return
		}
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
