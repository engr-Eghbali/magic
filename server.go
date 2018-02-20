package main

import (
	f "fmt"
	"log"
	"net/http"
	"runtime"
	"time"

	magic_validation "./plugins/authentication"
	magic_security "./plugins/authentication/layer2"
	magic_struct "./plugins/authentication/layer2/layer3/typedef"
	_ "github.com/go-sql-driver/mysql"
	cache "github.com/patrickmn/go-cache"
	"github.com/unrolled/secure"
	//"github.com/gocraft/dbr"
)

//////////////1/////////
//import end
func submit_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		var input magic_struct.Userdata
		magic_validation.Initialsubmit(r, &input)
		flg := magic_validation.Validation(&input, w, r)

		if flg {

			if flg = magic_security.Checkbrute(input.Uid, r); flg {
				log.Print()
				panic("++++++++brute danger++++++++\n")
			} else {
				submitFLG := magic_validation.Submit(w, r, input)
				if submitFLG {
					log.Print()
					f.Print("user submited successfully\n")
				} else {
					log.Print()
					f.Print("!!!!!!!!!!user failed to submit!!!!!!!!!!!\n")

				}
			}

		} else {
			f.Fprintf(w, "input is not valid")
		}

	}
}

func login_ctrl(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "POST" {
		var input magic_struct.Userdata
		magic_validation.Initiallogin(r, &input)
		flg := magic_validation.Validation(&input, w, r)
		if flg {
			if false { //magic_security.Checkbrute(input.Uid, r); flg {
				log.Print()
				panic("\n !++++++++brute danger++++++++!\n")
			} else {
				loginFLG := magic_validation.Login(w, r, input)
				if loginFLG {
					log.Print()
					f.Print("\n -_-_-_-_-_user logon successfull_-_-_-_-_-_-\n")

				} else {
					log.Print()
					f.Print("\n!!!!!!!!!!_user failed to login_!!!!!!!!!!!\n")

				}
			}

		} else {
			f.Fprintf(w, "input is not valid")
		}

	}
}

func courier_ctrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/javascript")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	r.ParseForm()
	id := r.Form["id"][0]
	loc := r.Form["loc"][0]
	// You can also retrieve the collection, where c == cc.

	runtime.GOMAXPROCS(runtime.NumCPU())
	s := time.Now()
	//for i := 0; i < 27000; i++ {
	c.Set(id, loc, cache.DefaultExpiration)
	//}
	f.Print("time spend:")
	f.Print(time.Since(s))

}

//-func upload_ctrl(w http.ResponseWriter, r *http.Request) {
//-	magic.Handle(w, r)
//-}

var c = cache.New(5*time.Minute, 10*time.Minute)

//\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\
//\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/

func main() {

	//initial route security
	submit_init := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	login_init := secure.New(secure.Options{
		AllowedHosts:            []string{"ssl.example.com"},                     // AllowedHosts is a list of fully qualified domain names that are allowed. Default is empty list, which allows any and all host names.
		HostsProxyHeaders:       []string{"X-Forwarded-Hosts"},                   // HostsProxyHeaders is a set of header keys that may hold a proxied hostname value for the request.
		SSLRedirect:             true,                                            // If SSLRedirect is set to true, then only allow HTTPS requests. Default is false.
		SSLTemporaryRedirect:    false,                                           // If SSLTemporaryRedirect is true, the a 302 will be used while redirecting. Default is false (301).
		SSLHost:                 "ssl.example.com",                               // SSLHost is the host name that is used to redirect HTTP requests to HTTPS. Default is "", which indicates to use the same host.
		SSLProxyHeaders:         map[string]string{"X-Forwarded-Proto": "https"}, // SSLProxyHeaders is set of header keys with associated values that would indicate a valid HTTPS request. Useful when using Nginx: `map[string]string{"X-Forwarded-Proto": "https"}`. Default is blank map.
		STSSeconds:              315360000,                                       // STSSeconds is the max-age of the Strict-Transport-Security header. Default is 0, which would NOT include the header.
		STSIncludeSubdomains:    true,                                            // If STSIncludeSubdomains is set to true, the `includeSubdomains` will be appended to the Strict-Transport-Security header. Default is false.
		STSPreload:              true,                                            // If STSPreload is set to true, the `preload` flag will be appended to the Strict-Transport-Security header. Default is false.
		ForceSTSHeader:          false,                                           // STS header is only included when the connection is HTTPS. If you want to force it to always be added, set to true. `IsDevelopment` still overrides this. Default is false.
		FrameDeny:               true,                                            // If FrameDeny is set to true, adds the X-Frame-Options header with the value of `DENY`. Default is false.
		CustomFrameOptionsValue: "SAMEORIGIN",                                    // CustomFrameOptionsValue allows the X-Frame-Options header value to be set with a custom value. This overrides the FrameDeny option. Default is "".
		ContentTypeNosniff:      true,                                            // If ContentTypeNosniff is true, adds the X-Content-Type-Options header with the value `nosniff`. Default is false.
		BrowserXssFilter:        true,                                            // If BrowserXssFilter is true, adds the X-XSS-Protection header with the value `1; mode=block`. Default is false.
		CustomBrowserXssValue:   "1; report=https://example.com/xss-report",      // CustomBrowserXssValue allows the X-XSS-Protection header value to be set with a custom value. This overrides the BrowserXssFilter option. Default is "".
		ContentSecurityPolicy:   "default-src 'self'",                            // ContentSecurityPolicy allows the Content-Security-Policy header value to be set with a custom value. Default is "".
		// PublicKey: `pin-sha256="base64+primary=="; pin-sha256="base64+backup=="; max-age=5184000; includeSubdomains; report-uri="https://www.example.com/hpkp-report"`, // PublicKey implements HPKP to prevent MITM attacks with forged certificates. Default is "".
		//ReferrerPolicy: "same-origin" // ReferrerPolicy allows the Referrer-Policy header with the value to be set with a custom value. Default is "".
		//***triger***
		IsDevelopment: true, // This will cause the AllowedHosts, SSLRedirect, and STSSeconds/STSIncludeSubdomains options to be ignored during development. When deploying to production, be sure to set this to false.
	})

	app1 := submit_init.Handler(http.HandlerFunc(submit_ctrl))
	app2 := login_init.Handler(http.HandlerFunc(login_ctrl))

	http.Handle("/submit", app1)
	http.Handle("/login", app2)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))

	http.HandleFunc("/courier", func(w http.ResponseWriter, r *http.Request) {
		courier_ctrl(w, r)
	})
	http.ListenAndServe(":9090", nil)

	//mail tool here:
	//myplug.Mail(rcp magic_struct .List, subject string, text string)

	//router := httprouter.New()
	//router.POST("/submit", submitCTRL)
	//router.POST("/login", loginCTRL)
	//router.POST("/upload", uploadCTRL)
	//router.POST("/courier", courierCRTL)
	//log.Fatal(http.ListenAndServe(":8080", router))
}
