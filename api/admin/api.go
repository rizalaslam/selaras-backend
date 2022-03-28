// Package admin ties together administration resources and handlers.
package admin

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg"

	"github.com/rizalaslam/selaras-backend/auth/authorize"
	"github.com/rizalaslam/selaras-backend/database"
	"github.com/rizalaslam/selaras-backend/logging"
)

const (
	roleAdmin = "admin"
)

type ctxKey int

const (
	ctxAccount ctxKey = iota
)

// API provides admin application resources and handlers.
type API struct {
	Accounts *AccountResource
}

// NewAPI configures and returns admin application API.
func NewAPI(db *pg.DB) (*API, error) {

	accountStore := database.NewAdmAccountStore(db)
	accounts := NewAccountResource(accountStore)

	api := &API{
		Accounts: accounts,
	}
	return api, nil
}

// Router provides admin application routes.
func (a *API) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(authorize.RequiresRole(roleAdmin))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Admin"))
	})

	r.Mount("/accounts", a.Accounts.router())
	return r
}

func log(r *http.Request) logrus.FieldLogger {
	return logging.GetLogEntry(r)
}
