package config

//Config represents the configuration defined in webtaskrunner.yaml
type Config struct {
	Grunt GruntConfig `yaml:"grunt"`
	Gulp  GulpConfig  `yaml:"gulp"`
}

//GruntConfig is part of the webtaskrunner configuration and contains all grunt specific settings
type GruntConfig struct {
	ExecutionDir  string `yaml:"execution_dir"`
	GruntFilePath string `yaml:"gruntfile_path"`
}

//GulpConfig is part of the webtaskrunner configuration and contains all gulp specific settings
type GulpConfig struct {
	ExecutionDir string `yaml:"execution_dir"`
	GulpFilePath string `yaml:"gulpfile_path"`
}
