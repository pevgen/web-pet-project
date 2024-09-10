package config

import "testing"

func TestReadYaml(t *testing.T) {
	type args struct {
		configName string
	}
	tests := []struct {
		name string
		args args
	}{
		{"configYamlFileName", args{"../../config/config.yaml"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := ReadYaml(tt.args.configName)
			if cfg.Path != "C:/Temp" {
				t.Errorf("got %v expected %v", true, cfg.Enabled)
			}
		})
	}
}
