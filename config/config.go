/*
Sniperkit-Bot
- Status: analyzed
*/

package config

//Config represents the configuration defined in webtaskrunner.yaml
type Config struct {
	Integrations Integrations `yaml:"integrations"`
}

//Integrations represents a static list of configured integrations
type Integrations struct {
	Ant    *AntConfig    `yaml:"ant"`
	Gradle *GradleConfig `yaml:"gradle"`
	Grunt  *GruntConfig  `yaml:"grunt"`
	Gulp   *GulpConfig   `yaml:"gulp"`
}

//FrontendInfo contains the configuration for one integration in frontend
type FrontendInfo struct {
	ImageUrl string `yaml:"imageUrl"`
	Name     string `yaml:"name"`
	Route    string `yaml:"route"`
}

//AntConfig is part of the webtaskrunner configuration and contains all ant specific settings
type AntConfig struct {
	FrontendInfo *FrontendInfo `yaml:"frontend"`
}

//GradleConfig is part of the webtaskrunner configuration and contains all gradle specific settings
type GradleConfig struct {
	FrontendInfo *FrontendInfo `yaml:"frontend"`
	ExecutionDir string        `yaml:"execution_dir"`
}

//GruntConfig is part of the webtaskrunner configuration and contains all grunt specific settings
type GruntConfig struct {
	FrontendInfo  *FrontendInfo `yaml:"frontend"`
	ExecutionDir  string        `yaml:"execution_dir"`
	GruntFilePath string        `yaml:"gruntfile_path"`
}

//GulpConfig is part of the webtaskrunner configuration and contains all gulp specific settings
type GulpConfig struct {
	FrontendInfo *FrontendInfo `yaml:"frontend"`
	ExecutionDir string        `yaml:"execution_dir"`
	GulpFilePath string        `yaml:"gulpfile_path"`
}
