package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
)

func GETCip(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed GET CIP endpoint", "username", authCtx.Username)

		cips := database.GetAllCips(app)

		json, err := json.Marshal(cips)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(json)
	}
}

func GETCipByCode(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed GET CIP endpoint", "username", authCtx.Username)
		val := req.PathValue("name")
		code, err := strconv.ParseFloat(val, 64)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("Invalid CIP code!"))
			app.Log.Error("Invalid CIP code", "error", err)
			return
		}

		cip := database.GetCipByCode(app, code)

		json, err := json.Marshal(cip)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(json)
	}
}
