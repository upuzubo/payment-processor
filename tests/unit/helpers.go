package payment_processor

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

func NewRouter(db *sqlx.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create_payment", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		payment := Payment{
			Amount:       r.FormValue("amount"),
			Currency:     r.FormValue("currency"),
			Description:  r.FormValue("description"),
		}

		if _, err = db.Exec(`INSERT INTO payments (amount, currency, description) VALUES ($1, $2, $3) RETURNING id`, payment.Amount, payment.Currency, payment.Description); err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.HandleFunc("/get_payment/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var payment Payment
		if err := db.Get(&payment, `SELECT * FROM payments WHERE id = $1`, id); err != nil {
			if err.(*pq.Error).Code == "23503" {
				http.Error(w, "not found", http.StatusNotFound)
			} else {
				log.Println(err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
			return
		}

		json.NewEncoder(w).Encode(payment)
	})

	return r
}