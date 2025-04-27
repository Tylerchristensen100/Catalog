package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/database"
	"catalog.tylerChristensen/internal/helpers"
	"catalog.tylerChristensen/internal/models"
)

func POSTProgram(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed POST Program endpoint", "username", authCtx.Username)

		err := req.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(res, "Error parsing form", http.StatusBadRequest)
			app.Log.Error("Error parsing form", "error", err)
			return
		}

		name := req.FormValue("name")
		if name == "" {
			http.Error(res, "Name is required", http.StatusBadRequest)
			return
		}

		gradLevel, err := strconv.ParseInt(req.FormValue("grad_level"), 10, 32)
		if gradLevel < 0 || err != nil {
			http.Error(res, "Grad Level is required", http.StatusBadRequest)
			return
		}

		programType := req.FormValue("program_type")
		if programType == "" {
			http.Error(res, "Program Type is required", http.StatusBadRequest)
			return
		}

		school, err := strconv.ParseInt(req.FormValue("school"), 10, 32)
		if school < 0 || err != nil {
			http.Error(res, "School is required", http.StatusBadRequest)
			return
		}

		majorCode := req.FormValue("major_code")
		if majorCode == "" {
			http.Error(res, "Major Code is required", http.StatusBadRequest)
			return
		}

		online, err := strconv.ParseInt(req.FormValue("online"), 10, 32)
		if online < 0 || err != nil {
			http.Error(res, "Online is required", http.StatusBadRequest)
			return
		}

		campus, err := strconv.ParseInt(req.FormValue("campus"), 10, 32)
		if campus < 0 || err != nil {
			http.Error(res, "Campus is required", http.StatusBadRequest)
			return
		}

		description := req.FormValue("description")
		if description == "" {
			http.Error(res, "Description is required", http.StatusBadRequest)
			return
		}

		cip, err := strconv.ParseFloat(req.FormValue("cip"), 32)
		if cip < 0 || err != nil {
			http.Error(res, "CIP is required", http.StatusBadRequest)
			return
		}

		createdBy, err := helpers.GetUserFromUsername(app, authCtx)
		if err != nil {
			http.Error(res, "Error getting User", http.StatusInternalServerError)
			app.Log.Error("Error getting User", "error", err)
			return
		}

		program := models.Program{
			Name:        name,
			GradLevel:   models.GradLevel{ID: int32(gradLevel)},
			ProgramType: programType,
			School:      models.School{ID: int32(school)},
			MajorCode:   majorCode,
			Online:      int8(online),
			Campus:      int8(campus),
			Description: description,
			Cip:         int32(cip),
			CreatedBy:   int32(createdBy.ID),
		}

		p, err := database.CreateProgram(app, program)
		if err != nil {
			http.Error(res, "Error creating program", http.StatusInternalServerError)
			app.Log.Error("Error creating program", "error", err)
			return
		}
		json, err := json.Marshal(p)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}

		res.WriteHeader(http.StatusCreated)
		res.Header().Set("Content-Type", "application/json")
		res.Write(json)
	}
}
