package configs

import "github.com/spf13/viper"

func ReadConfig(configFile string) (*viper.Viper, error) {
	conf := viper.New()
	conf.SetConfigFile(configFile)
	if err := conf.ReadInConfig(); err != nil {
		return nil, err
	}

	return conf, nil
}

func ReadConfigInto(configFile string, v interface{}) error {
	conf, err := ReadConfig(configFile)
	if err != nil {
		return err
	}

	return conf.Unmarshal(v)
}
