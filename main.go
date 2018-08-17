/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sniperkit/snk.fork.webtaskrunner/config"
	"github.com/sniperkit/snk.fork.webtaskrunner/handler/commandhandler"
	"github.com/sniperkit/snk.fork.webtaskrunner/handler/taskshandler"
	"github.com/sniperkit/snk.fork.webtaskrunner/integrations"
)

var frontendConfigs = []*config.FrontendInfo{}
var configuredIntegrations []integrations.Integration

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", indexPageHandler)

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// add custom integrations here, and in integrations package of course
	if conf.Integrations.Ant != nil {
		addIntegration(conf.Integrations.Ant.FrontendInfo, integrations.NewAntIntegration(conf.Integrations.Ant))
	}
	if conf.Integrations.Gradle != nil {
		addIntegration(conf.Integrations.Gradle.FrontendInfo, integrations.NewGradleIntegration(conf.Integrations.Gradle))
	}
	if conf.Integrations.Grunt != nil {
		addIntegration(conf.Integrations.Grunt.FrontendInfo, integrations.NewGruntIntegration(conf.Integrations.Grunt))
	}
	if conf.Integrations.Gulp != nil {
		addIntegration(conf.Integrations.Gulp.FrontendInfo, integrations.NewGulpIntegration(conf.Integrations.Gulp))
	}

	webSocketHandler := taskshandler.New(configuredIntegrations)
	http.Handle("/tasklist", webSocketHandler.GetHandler())

	http.ListenAndServe(":"+getPort(), nil)
}

func getPort() string {
	sPort := os.Getenv("PORT")
	if sPort == "" {
		sPort = "8080"
	}
	return sPort
}

func addIntegration(frontendConfiguration *config.FrontendInfo, integration integrations.Integration) {

	integrationPath := frontendConfiguration.Route

	webSocketHandler := commandhandler.New(integration)
	http.Handle("/"+integrationPath+"/cmd", webSocketHandler.GetHandler())

	frontendConfigs = append(frontendConfigs, frontendConfiguration)

	configuredIntegrations = append(configuredIntegrations, integration)
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("web/templates/taskrunner.html")
	fmt.Println(err)
	w.Write(b)
}
