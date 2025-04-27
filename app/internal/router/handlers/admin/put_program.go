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

func PUTProgram(app *internal.App) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		authCtx := auth.Context(req.Context())
		app.Log.Info("user accessed PUT Program endpoint", "username", authCtx.Username)
		id, err := strconv.ParseInt(req.PathValue("id"), 10, 32)
		if err != nil {
			http.Error(res, "Error parsing ID", http.StatusBadRequest)
			app.Log.Error("Error parsing ID", "error", err)
			return
		}

		req.ParseMultipartForm(10 << 20)

		var currentProgram models.Program
		err = app.DB.Find(&currentProgram, models.Program{ID: int32(id)}).Error
		if err != nil {
			http.Error(res, "Error getting Program", http.StatusInternalServerError)
			app.Log.Error("Error getting Program", "error", err)
			return
		}

		nameString := req.FormValue("name")
		if nameString == "" {
			nameString = currentProgram.Name
		}

		programTypeString := req.FormValue("program_type")
		if programTypeString == "" {
			programTypeString = currentProgram.ProgramType
		}

		majorCodeString := req.FormValue("major_code")
		if majorCodeString == "" {
			majorCodeString = currentProgram.MajorCode
		}

		gradLevelString := req.FormValue("grad_level")
		schoolString := req.FormValue("school")
		onlineString := req.FormValue("online")
		campusString := req.FormValue("campus")

		descriptionString := req.FormValue("description")
		if descriptionString == "" {
			descriptionString = currentProgram.Description
		}

		cipString := req.FormValue("cip")
		if cipString == "" {
			cipString = string(currentProgram.Cip)
		}

		var gradLevel int64
		if gradLevelString == "" {
			gradLevel = int64(currentProgram.GradLevel.ID)
		} else {
			gradLevel, err = strconv.ParseInt(gradLevelString, 10, 32)
			if gradLevel < 0 || err != nil {
				http.Error(res, "Grad Level is required", http.StatusBadRequest)
				return
			}
		}

		var school int64
		if schoolString == "" {
			school = int64(currentProgram.School.ID)
		} else {
			school, err = strconv.ParseInt(schoolString, 10, 32)
			if school < 0 || err != nil {
				http.Error(res, "School is required", http.StatusBadRequest)
				return
			}
		}
		var online int8
		if onlineString == "" {
			online = currentProgram.Online
		} else {
			if onlineString == "true" || onlineString == "1" {
				online = 1
			} else if onlineString == "false" || onlineString == "0" {
				online = 0
			} else {
				http.Error(res, "Online is required", http.StatusBadRequest)
				return
			}
		}

		var campus int64
		if campusString == "" {
			campus = int64(currentProgram.Campus)
		} else {
			campus, err = strconv.ParseInt(campusString, 10, 32)
			if campus < 0 || err != nil {
				http.Error(res, "Campus is required", http.StatusBadRequest)
				return
			}
		}

		cip, err := strconv.ParseInt(cipString, 10, 32)
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

		updatedProgram := models.Program{
			ID:          int32(id),
			Name:        nameString,
			GradLevelID: int32(gradLevel),
			ProgramType: programTypeString,
			SchoolID:    int32(school),
			MajorCode:   majorCodeString,
			Online:      int8(online),
			Campus:      int8(campus),
			Description: descriptionString,
			Cip:         int32(cip),
			CreatedBy:   int32(createdBy.ID),
		}

		p, err := database.UpdateProgram(app, updatedProgram)
		if err != nil {
			http.Error(res, "Error updating Program", http.StatusInternalServerError)
			app.Log.Error("Error updating Program", "error", err)
			return
		}
		json, err := json.Marshal(p)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Error Marshalling JSON!"))
			app.Log.Error("Error marshalling JSON", "error", err)
			return
		}

		res.WriteHeader(http.StatusOK)
		res.Header().Set("Content-Type", "application/json")
		res.Write(json)
	}
}
