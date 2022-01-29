package testhelper

import "github.com/ms-url-shortner/app/config"

func MockConfig() {
	cfg := &config.Configuration{
		MSName:      "test-ms-url-shortnet",
		Environment: "test",
	}
	config.SetConfig(cfg)
}
