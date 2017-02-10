package config

//Config represents the configuration defined in webtaskrunner.yaml
type Config struct {
	Grunt GruntConfig `yaml:"grunt"`
}

//GruntConfig is part of the webtaskrunner configuration and contains all grunt specific settings
type GruntConfig struct {
	ExecutionDir  string `yaml:"execution_dir"`
	GruntFilePath string `yaml:"gruntfile_path"`
}
