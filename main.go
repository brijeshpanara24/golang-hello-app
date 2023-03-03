package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	defaultPort = "8080"
)

const (
	CONN_PORTS_ENV    = "CONN_PORTS"
	APP_NAME_ENV      = "APP_NAME"
	APP_VERSION_ENV   = "APP_VERSION"
	APP_INSTANCE_ENV  = "APP_INSTANCE"
	URL_ENV           = "URL"
)

// Main function
func main() {
	finish := make(chan bool)

	ports := getPorts()
	for _, port := range ports {
		startServer(port)
	}

	<-finish
}

func getPorts() []string {
	ports := []string{defaultPort}

	portsEnv, ok := os.LookupEnv(CONN_PORTS_ENV)
	if ok && portsEnv != "" {
		ports = strings.Split(portsEnv, ",")
	}

	return ports
}

func startServer(port string) {
	server := http.NewServeMux()
	server.HandleFunc("/", serverHandlerFunc)

	go func() {
		fmt.Printf("Starting server at port %s\n", port)
		http.ListenAndServe(fmt.Sprintf(":%s", port), server)
	}()
}

func serverHandlerFunc(w http.ResponseWriter, r *http.Request) {
	str := "Hi User !!!\n"
	str += fmt.Sprintf("Host: %s\n", r.Host)

	path := html.EscapeString(r.URL.Path)
	if path != "/" {
		str += fmt.Sprintf("Path: %s\n", path)
	}

	appName, ok := os.LookupEnv(APP_NAME_ENV)
	if ok {
		str += fmt.Sprintf("App Name: %s\n", appName)
	}

	appVersion, ok := os.LookupEnv(APP_VERSION_ENV)
	if ok {
		str += fmt.Sprintf("App Version: %s\n", appVersion)
	}

	appInstance, ok := os.LookupEnv(APP_INSTANCE_ENV)
	if ok {
		str += fmt.Sprintf("App Instance: %s\n", appInstance)
	}

	url, ok := os.LookupEnv(URL_ENV)
	if ok && url != "" {
		str += fmt.Sprintf("Url: %s\n", url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Get Error: %+v\n", err)
		}

		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Read Error: %+v\n", err)
		}

		respBody := string(body)
		respBody = strings.TrimSuffix(respBody, "\n")

		//Convert the body to type string
		str += fmt.Sprintf("Response: \n")
		str += fmt.Sprintf("------------------------------ \n")
		str += fmt.Sprintf("%s\n", respBody)
		str += fmt.Sprintf("------------------------------ \n")
	}

	w.Write([]byte(str))
}