package main

import (
	"fmt"
	"github.com/Oppodelldog/webtaskrunner/config"
	"github.com/Oppodelldog/webtaskrunner/handler/commandhandler"
	"github.com/Oppodelldog/webtaskrunner/handler/tasklisthandler"
	"github.com/Oppodelldog/webtaskrunner/handler/taskrunnerlisthandler"
	"github.com/Oppodelldog/webtaskrunner/integrations"
	"io/ioutil"
	"net/http"
	"os"
)

var frontendConfigs = []*config.FrontendInfo{}

func main() {
	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", overviewHandler)

	fs := http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// add custom integrations here, and in integrations package of course
	addIntegration(conf.Integrations.Ant.FrontendInfo, integrations.NewAntIntegration(conf.Integrations.Ant))
	addIntegration(conf.Integrations.Gradle.FrontendInfo, integrations.NewGradleIntegration(conf.Integrations.Gradle))
	addIntegration(conf.Integrations.Grunt.FrontendInfo, integrations.NewGruntIntegration(conf.Integrations.Grunt))
	addIntegration(conf.Integrations.Gulp.FrontendInfo, integrations.NewGulpIntegration(conf.Integrations.Gulp))

	http.Handle("/taskrunners", taskrunnerlisthandler.New(frontendConfigs))

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

	http.HandleFunc("/"+integrationPath, taskRunnerHandler)

	taskListHandler := tasklisthandler.New(integration)
	http.Handle("/"+integrationPath+"/tasks", taskListHandler)

	webSocketHandler := commandhandler.New(integration)
	http.Handle("/"+integrationPath+"/cmd", webSocketHandler.GetHandler())

	frontendConfigs = append(frontendConfigs, frontendConfiguration)
}

func taskRunnerHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("web/templates/taskrunner.html")
	fmt.Println(err)
	w.Write(b)
}

func overviewHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("web/templates/overview.html")
	fmt.Println(err)
	w.Write(b)
}
