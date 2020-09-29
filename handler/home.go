package handler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

func Home(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		visitorID := 0
		err := db.QueryRow(
			"INSERT  INTO visitors(user_agent,datetime) VALUES ($1, now()) RETURNING id ", r.UserAgent()).Scan(visitorID)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Internal Error"))

		}
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Println(w, fmt.Sprintf("Hello visitor %d", visitorID, visitorID))

	}
}
