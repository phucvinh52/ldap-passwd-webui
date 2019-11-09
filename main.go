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
	// gen key
	/*
		# Key considerations for algorithm "RSA" ≥ 2048-bit
		openssl genrsa -out server.key 2048

		# Key considerations for algorithm "ECDSA" (X25519 || ≥ secp384r1)
		# https://safecurves.cr.yp.to/
		# List ECDSA the supported curves (openssl ecparam -list_curves)
		openssl ecparam -genkey -name secp384r1 -out server.key
		openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
	*/
	http.ListenAndServeTLS(":8443", "/myapp/server.crt", "/myapp/server.key", nil)
}
