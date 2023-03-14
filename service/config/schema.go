package config

type TomlConfig struct {
	Version   string
	Addr      string
	Predictor struct {
		ModelPath  string
		FeatureLen int
	}
}
