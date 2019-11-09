package main

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/phucvinh52/ldap-passwd-webui/app"
)

func main() {
	reHandler := new(app.RegexpHandler)

	reHandler.HandleFunc(".*.[js|css|png|eof|svg|ttf|woff]", "GET", app.ServeAssets)
	reHandler.HandleFunc("/", "GET", app.ServeIndex)
	reHandler.HandleFunc("/", "POST", app.ChangePassword)
	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))
	http.Handle("/", reHandler)
	fmt.Println("Starting server on port 8443")
	http.ListenAndServeTLS(":8443", "/myapp/server.crt", "/myapp/server.key", nil)
}
