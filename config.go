package common

import "encoding/json"

type Config struct {
	Server ConfigServer `mapstructure:"server"`
	Client ConfigClient `mapstructure:"client"`
	Logger ConfigLogger `mapstructure:"logger"`
}

func (c *Config) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type ConfigServer struct {
	Http ConfigHttp `mapstructure:"http"`
	Repo ConfigRepo `mapstructure:"repo"`
	Grpc ConfigGrpc `mapstructure:"grpc"`
}

type ConfigHttp struct {
	Timeout        int       `mapstructure:"timeout"`
	HmacHeader     string    `mapstructure:"hmac_header"`
	HmacSecret     string    `mapstructure:"hmac_secret"`
	BearerToken    string    `mapstructure:"bearer_token"`
	Jwt            ConfigJwt `mapstructure:"jwt"`
	AllowedOrigins string    `mapstructure:"allowed_origins"`
	AllowedHeaders string    `mapstructure:"allowed_headers"`
	AllowedMethods string    `mapstructure:"allowed_methods"`
}

type ConfigJwt struct {
	Secret          string `mapstructure:"secret"`
	AccessTokenExp  int    `mapstructure:"access_token_exp"`
	RefreshToken    bool   `mapstructure:"refresh_token"`
	RefreshTokenExp int    `mapstructure:"refresh_token_exp"`
}

type ConfigRepo struct {
	Driver                 string `mapstructure:"driver"`
	ConnStr                string `mapstructure:"conn_str"`
	MaxIdleConns           int    `mapstructure:"max_idle_conns"`
	MaxOpenConns           int    `mapstructure:"max_open_conns"`
	ConnMaxLifetimeMinutes int    `mapstructure:"conn_max_lifetime_in_min"`
}

type ConfigGrpc struct {
	Timeout           int                     `mapstructure:"timeout"`
	HealthCheck       bool                    `mapstructure:"healthcheck"`
	EnforcementPolicy ConfigEnforcementPolicy `mapstructure:"enforcement_policy"`
	KeepAlive         ConfigKeepAlive         `mapstructure:"keep_alive"`
}

type ConfigEnforcementPolicy struct {
	Use                 bool `mapstructure:"use"`
	MinTime             int  `mapstructure:"min_time"`
	PermitWithoutStream bool `mapstructure:"permit_without_stream"`
}

type ConfigKeepAlive struct {
	MaxConnIdle         int  `mapstructure:"max_conn_idle"`
	MaxConnAge          int  `mapstructure:"max_conn_age"`
	MaxConnAgeGrace     int  `mapstructure:"max_conn_age_grace"`
	Time                int  `mapstructure:"time"`
	Timeout             int  `mapstructure:"timeout"`
	PermitWithoutStream bool `mapstructure:"permit_without_stream"`
}

type ConfigLogger struct {
	Json   bool       `mapstructure:"json"`
	Stdout bool       `mapstructure:"stdout"`
	File   ConfigFile `mapstructure:"file"`
	Level  string     `mapstructure:"level"`
}

type ConfigFile struct {
	Name       string `mapstructure:"name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackup  int    `mapstructure:"max_backup"`
	MaxAge     int    `mapstructure:"max_age"`
	Compressed bool   `mapstructure:"compressed"`
}

type ConfigClient struct {
	Oauth2 ConfigOauth2     `mapstructure:"oauth2"`
	Http   ConfigHttpClient `mapstructure:"http"`
	Grpc   ConfigGrpcClient `mapstructure:"grpc"`

	UserHttp ConfigHttpClient `mapstructure:"user_http"`
	UserGrpc ConfigGrpcClient `mapstructure:"user_grpc"`
}

type ConfigGrpcClient struct {
	Addr                 string          `mapstructure:"addr"`
	KeepAlive            ConfigKeepAlive `mapstructure:"keep_alive"`
	DefaultServiceConfig string          `mapstructure:"devault_service_config"`
	CaCertPem            string          `mapstructure:"ca_cert_pem"`
	CertServerName       string          `mapstructure:"cert_server_name"`
	DialBlock            bool            `mapstructure:"dial_block"`
	PermitWithoutStream  bool            `mapstructure:"permit_without_stream"`
	ResolverScheme       string          `mapstructure:"resolver_scheme"`
	ResolverServiceName  string          `mapstructure:"resolver_service_name"`
}
type ConfigHttpClient struct {
	Url         string `mapstructure:"url"`
	BearerToken string `mapstructure:"bearer_token"`
	HmacSecret  string `mapstructure:"hmac_secret"`
	HmacHeader  string `mapstructure:"hmac_header"`
}
type ConfigOauth2 struct {
	Source        string            `mapstructure:"source"`
	ClientID      string            `mapstructure:"client_id"`
	ClientSecret  string            `mapstructure:"client_secret"`
	RedirectUrl   string            `mapstructure:"redirect_url"`
	Scopes        []string          `mapstructure:"scopes"`
	AuthUrl       string            `mapstructure:"auth_url"`
	TokenUrl      string            `mapstructure:"token_url"`
	OpenIDConfUrl string            `mapstructure:"openid_conf_url"`
	AuthRequest   ConfigAuthRequest `mapstructure:"authrequest"`
}

type ConfigAuthRequest struct {
	ResponseUrl string `mapstructure:"response_url"`
	Wait        int    `mapstructure:"wait"`
}