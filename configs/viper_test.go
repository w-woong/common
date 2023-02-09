package configs_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/w-woong/common"
	"github.com/w-woong/common/configs"
)

type MyConfig struct {
	Name string       `mapstructure:"name"`
	Age  int          `mapstructure:"age"`
	Http MyConfigHttp `mapstructure:"http"`
}

type MyConfigHttp struct {
	Context MyConfigContext `mapstructure:"context"`
	Header  MyConfigHeader  `mapstructure:"header"`
}

type MyConfigContext struct {
	Timeout int      `mapstructure:"timeout"`
	Value   []string `mapstructure:"value"`
}

type MyConfigHeader struct {
	Auth MyConfigAuth `mapstructure:"auth"`
}

type MyConfigAuth struct {
	Token string `mapstructure:"token"`
}

func (d *MyConfig) String() string {
	b, _ := json.Marshal(d)
	return string(b)
}

func TestReadConfigInto(t *testing.T) {
	conf := MyConfig{}
	err := configs.ReadConfigInto("./tests/config.yml", &conf)
	assert.Nil(t, err)
	fmt.Println(conf.String())
}

func TestConfig(t *testing.T) {
	conf := common.Config{}
	err := configs.ReadConfigInto("./tests/config.yml", &conf)
	assert.Nil(t, err)
	fmt.Println(conf.String())
	for k, v := range conf.Client.OAuth2 {
		fmt.Println(k, v)
	}
}
