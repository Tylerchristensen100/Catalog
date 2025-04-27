package router

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"catalog.tylerChristensen/internal"
	"catalog.tylerChristensen/internal/auth"
	"catalog.tylerChristensen/internal/constants"
	db "catalog.tylerChristensen/internal/database"
	docs_handlers "catalog.tylerChristensen/internal/docs/handlers"
	"catalog.tylerChristensen/internal/models"
	"catalog.tylerChristensen/internal/router/components"
	"catalog.tylerChristensen/internal/router/handlers"
	"catalog.tylerChristensen/internal/router/handlers/admin"
	"catalog.tylerChristensen/internal/router/handlers/api"
	auth_handlers "catalog.tylerChristensen/internal/router/handlers/auth"
	"catalog.tylerChristensen/internal/router/middleware"

	"github.com/a-h/templ"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	middlewarez "github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

func Router(app *internal.App) http.Handler {

	ctx := context.Background()
	authZ, err := authorization.New(ctx, zitadel.New(app.Config.AuthDomain), oauth.DefaultAuthorization(app.Config.AuthKey))
	if err != nil {
		slog.Error("zitadel sdk could not initialize", "error", err)
		os.Exit(1)
	}

	mw := middlewarez.New(authZ)
	app.Auth = mw

	server := http.NewServeMux()

	//PUBLIC API
	server.HandleFunc("GET /api/programs", http.HandlerFunc(api.GETPrograms(app)))
	server.HandleFunc("GET /api/courses", http.HandlerFunc(api.GETCourses(app)))
	server.HandleFunc("GET /api/cips", http.HandlerFunc(api.GETCips(app)))
	server.HandleFunc("GET /api/schools", http.HandlerFunc(api.GETSchools(app)))
	server.HandleFunc("GET /api/grad-levels", http.HandlerFunc(api.GETGradLevels(app)))
	server.HandleFunc("GET /api/", http.HandlerFunc(api.GETGeneric(app)))

	//PUBLIC API DOCS
	server.Handle("GET /api/docs", http.HandlerFunc(docs_handlers.SwaggerHTML(app)))
	server.Handle("GET /api/docs/openapi.yaml", http.HandlerFunc(docs_handlers.OpenAPI(app)))

	//PUBLIC ROUTES
	server.HandleFunc("GET /components/courses", templ.Handler(components.Courses(db.GetAllCourses(app))).ServeHTTP)
	server.HandleFunc("GET /components/programs", templ.Handler(components.Programs(db.GetAllPrograms(app))).ServeHTTP)

	//ADMIN ROUTES   ----- REQUIRES AUTH ------

	server.Handle("GET /api/admin/program", auth.RequireAuthorization(models.FacultyRole)(admin.GETProgram(app)))
	server.Handle("POST /api/admin/program", auth.RequireAuthorization(models.FacultyRole)(admin.POSTProgram(app)))
	server.Handle("PUT /api/admin/program/{id}", auth.RequireAuthorization(models.FacultyRole)(admin.PUTProgram(app)))

	server.Handle("GET /api/admin/course", auth.RequireAuthorization(models.FacultyRole)(admin.GETCourse(app)))
	server.Handle("POST /api/admin/course", auth.RequireAuthorization(models.FacultyRole)(admin.POSTCourse(app)))
	server.Handle("PUT /api/admin/course/{id}", auth.RequireAuthorization(models.FacultyRole)(admin.PUTCourse(app)))

	server.Handle("GET /api/admin/cip", auth.RequireAuthorization(models.FacultyRole)(admin.GETCip(app)))
	server.Handle("POST /api/admin/cip", auth.RequireAuthorization(models.FacultyRole)(admin.POSTCip(app)))
	server.Handle("PUT /api/admin/cip/{id}", auth.RequireAuthorization(models.FacultyRole)(admin.PUTCip(app)))

	server.Handle("GET /api/admin/school", auth.RequireAuthorization(models.FacultyRole)(admin.GETSchools(app)))
	server.Handle("POST /api/admin/school", auth.RequireAuthorization(models.FacultyRole)(admin.POSTSchool(app)))
	server.Handle("PUT /api/admin/school/{code}", auth.RequireAuthorization(models.FacultyRole)(admin.PUTSchool(app)))

	server.Handle("GET /api/admin/", auth.RequireAuthorization(models.FacultyRole)(admin.Generic(app)))
	server.Handle("GET /api/admin/user", auth.RequireAuthorization(models.FacultyRole, models.AdminRole)(auth_handlers.User(app)))

	//HEALTH CHECK
	server.HandleFunc("GET /healthz", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("healthy"))
	})

	//STATIC FILES
	server.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir(constants.PublicDir))))

	//WEB ROUTES
	server.HandleFunc("GET /404", http.HandlerFunc(handlers.NotFound(app)))

	server.HandleFunc("GET /programs/", http.HandlerFunc(handlers.GetPrograms(app)))
	server.HandleFunc("GET /programs/{name}", http.HandlerFunc(handlers.GetProgramsByName(app)))

	server.HandleFunc("GET /courses/", http.HandlerFunc(handlers.GetCourses(app)))
	server.HandleFunc("GET /courses/{courseCode}", http.HandlerFunc(handlers.GetCoursesByCourseCode(app)))

	server.HandleFunc("GET /schools/", http.HandlerFunc(handlers.GetSchools(app)))
	server.HandleFunc("GET /schools/{code}", http.HandlerFunc(handlers.GetSchoolsByCode(app)))

	server.HandleFunc("GET /", http.HandlerFunc(handlers.GetIndex(app)))

	//SSO ROUTES
	server.HandleFunc("GET /login", http.HandlerFunc(auth_handlers.Login(app)))
	server.HandleFunc("GET /callback", http.HandlerFunc(auth_handlers.Callback(app)))
	server.Handle("POST /api/logout", http.HandlerFunc(auth_handlers.Logout(app)))

	return middleware.Logger(app, middleware.Headers(app, server))
}
