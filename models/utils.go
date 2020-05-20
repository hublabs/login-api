package models

type ModelConfig struct {
	AppEnv       string
	ColleagueApi string
}

var modelConfig *ModelConfig

func SetModelConfig(m *ModelConfig) {
	modelConfig = m
}
