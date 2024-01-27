package config

import "common-web-framework/search"

type MeiliSearchConfig struct {
	BlogIndex string `yaml:"blogIndex"`
	Host      string `yaml:"host" json:"host"`
	ApiKey    string `yaml:"apiKey" json:"apiKey"`
}

var SEARCH *search.MeiliSearchClient

func LoadSearchConfig(conf MeiliSearchConfig) {
	SEARCH = search.NewMeiliSearchClient(conf.Host, conf.ApiKey)
}
