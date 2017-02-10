package main

import (
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/config"
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"github.com/Oppodelldog/webtaskrunner/webhandler/ajaxhandler"
	"github.com/Oppodelldog/webtaskrunner/webhandler/websockethandler"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// add custom integrations here, and in integrations package of course
	addIntegration("ant", integrations.NewAntIntegration())
	addIntegration("gradle", integrations.NewGradleIntegration())
	addIntegration("grunt", integrations.NewGruntIntegration(conf.Grunt))

	http.ListenAndServe(":"+getPort(), nil)
}

func getPort() string {
	sPort := os.Getenv("PORT")
	if sPort == "" {
		sPort = "8080"
	}
	return sPort
}

func addIntegration(integrationPath string, integration integrations.Integration) {

	http.HandleFunc("/"+integrationPath, indexHandler)

	taskListHandler := ajaxhandler.New(integration)
	http.Handle("/"+integrationPath+"/tasks", taskListHandler)

	webSocketHandler := websockethandler.New(integration)
	http.Handle("/"+integrationPath+"/cmd", webSocketHandler.GetHandler())
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("web/templates/index.html")
	fmt.Println(err)
	w.Write(b)
}
