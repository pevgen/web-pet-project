package config

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		configPath string
	}
	tests := []struct {
		name string
		args args
		want AppConfig
	}{
		{"Read config from file",
			args{configPath: "../../config/app-config-test.yaml"},
			AppConfig{
				WebServer: WebServerConfig{
					Port: 1234,
				},
				Db: DBConfig{
					Postgres: PgConfig{
						ConnectString: "pgc",
					},
					Mongodb: MongoConfig{
						ConnectString: "mgc",
						DbName:        "dbn",
					},
				},
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.configPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew_Panic_if_file_doesnt_exist(t *testing.T) {
	defer func() { _ = recover() }()
	New("/bad_file_name")
	t.Errorf("Can't read config (bad path). Should be panic !")
}
