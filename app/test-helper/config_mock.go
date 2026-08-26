package testhelper

import "github.com/ms-url-shortner/app/config"

// MockConfig sets up a test configuration with default values
func MockConfig() {
	cfg := &config.Configuration{
		MSName:      "test-ms-url-shortnet",
		Environment: "test",
	}
	config.SetConfig(cfg)
}
