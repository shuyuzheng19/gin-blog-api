package config

type UploadConfig struct {
	MaxImageSize int    `yaml:"maxImageSize" json:"maxImageSize"`
	MaxFileSize  int    `yaml:"maxFileSize" json:"maxFileSize"`
	Uri          string `yaml:"uri" json:"uri"`
	Path         string `yaml:"path" json:"path"`
}
