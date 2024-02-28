package main

import (
	"net/http"

	"pythononline.usmanbala.com/ui"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {

	router := http.NewServeMux()

	fileServer := http.FileServer(http.FS(ui.Files))
	router.Handle("GET /static/", fileServer)

	router.HandleFunc("GET /ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	router.Handle("GET /", dynamic.ThenFunc(app.home))
	router.Handle("GET /ide", dynamic.ThenFunc(app.ide))
	router.Handle("GET /lessons/", dynamic.ThenFunc(app.lessons))
	router.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	router.Handle("POST /user/signup", dynamic.ThenFunc(app.userSignupPost))
	router.Handle("GET /user/login", dynamic.ThenFunc(app.userLogin))
	router.Handle("POST /user/login", dynamic.ThenFunc(app.userLoginPost))

	protected := dynamic.Append(app.requireAuthentication)
	router.Handle("GET /progress", protected.ThenFunc(app.progressMap))
	router.Handle("POST /progress", dynamic.ThenFunc(app.updateProgress))
	router.Handle("POST /user/logout", protected.ThenFunc(app.userLogoutPost))
	standard := alice.New(app.recoverPanic, app.logRequest)
	return standard.Then(router)
}
