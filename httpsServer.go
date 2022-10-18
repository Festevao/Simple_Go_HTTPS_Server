package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

var ENV_PORT string = ":443"
var httpsServerCodeFilePath string = getThisCodeFilePath()

func getThisCodeFilePath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("unable to get the current filename")
	}
	path := filepath.Dir(filename)

	return path
}

func init() { //Check the environment vars: PORT
	portFromEnv := os.Getenv("PORT")
	if intPort, err := strconv.Atoi(portFromEnv); portFromEnv != "" {
		if err != nil || !(intPort >= 1 && intPort <= 65535) {
			log.Fatal("%PORT% need to be a integer between 1-65535", err)
		}
		ENV_PORT = ":" + strconv.Itoa(intPort)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("This is an example server.\n"))
	})
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}
	srv := &http.Server{
		Addr:         ENV_PORT,
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	log.Fatal(srv.ListenAndServeTLS(filepath.Join(httpsServerCodeFilePath, "/TLS/server.crt"), filepath.Join(httpsServerCodeFilePath, "/TLS/server.key")))
}
