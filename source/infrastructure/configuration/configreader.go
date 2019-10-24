package configuration

import (
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	once sync.Once
	r    *configReader
)

type AppSettings struct {
	HostInfo
}

type HostInfo struct {
	KubernetesHost string
	HostUsername   string
	HostPassword   string
}

type IConfigReader interface {
	GetAllValues() (*AppSettings, error)
}

type configReader struct {
	configFile string
	v          *viper.Viper
}

func (r *configReader) GetAllValues() (*AppSettings, error) {

	var appSettings AppSettings
	if err := r.v.ReadInConfig(); err != nil {
		return nil, err
	}
	err := r.v.Unmarshal(&appSettings)
	if err != nil {
		return nil, err
	}
	return &appSettings, err
}

func NewConfigReader(configFile string, vip *viper.Viper) IConfigReader {
	once.Do(func() {
		vip.SetConfigName(configFile)
		vip.AddConfigPath(".")
		r = &configReader{
			configFile: configFile,
			v:          vip,
		}
	})
	return r
}

func GetConfig(defaultEnv string) (*AppSettings, error) {
	var vip = viper.GetViper()
	var configReader = NewConfigReader("config."+EnvString("MACHINE_ENVIRONMENT", defaultEnv), vip)
	return configReader.GetAllValues()
}

func EnvString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
