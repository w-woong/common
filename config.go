package common

import (
	"encoding/json"

	"github.com/w-woong/common/dto"
)

type Config struct {
	Server ConfigServer `mapstructure:"server"`
	Client ConfigClient `mapstructure:"client"`
	Logger ConfigLogger `mapstructure:"logger"`

	Partner ConfigPartner `mapstructure:"partner"`
}

func (c *Config) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}

type ConfigServer struct {
	Http ConfigHttp `mapstructure:"http"`
	Repo ConfigRepo `mapstructure:"repo"`
	Grpc ConfigGrpc `mapstructure:"grpc"`

	Security SecurityConfig `mapstructure:"security"`
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
	LogLevel               string `mapstructure:"log_level"`
}

type ConfigGrpc struct {
	Timeout           int                     `mapstructure:"timeout"`
	HealthCheck       bool                    `mapstructure:"healthcheck"`
	EnforcementPolicy ConfigEnforcementPolicy `mapstructure:"enforcement_policy"`
	KeepAlive         ConfigKeepAlive         `mapstructure:"keep_alive"`
}

type SecurityConfig struct {
	Jwks                JwksConfig     `mapstructure:"jwks"`
	OpenIDConfiguration dto.OpenIDConf `mapstructure:"openid_configuration"`
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
	// Oauth2 ConfigOauth2            `mapstructure:"oauth2"`
	IDTokenCookie string                  `mapstructure:"id_token_cookie"`
	OAuth2        map[string]OAuth2Config `mapstructure:"oauth2"`
	Http          ConfigHttpClient        `mapstructure:"http"`
	Grpc          ConfigGrpcClient        `mapstructure:"grpc"`

	UserHttp ConfigHttpClient `mapstructure:"user_http"`
	UserGrpc ConfigGrpcClient `mapstructure:"user_grpc"`

	Payment ConfigPayment `mapstructure:"payment"`
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

// OAuth2 client's configs
type OAuth2Config struct {
	ClientID      string            `mapstructure:"client_id"`
	ClientSecret  string            `mapstructure:"client_secret"`
	RedirectUrl   string            `mapstructure:"redirect_url"`
	Scopes        []string          `mapstructure:"scopes"`
	AuthUrl       string            `mapstructure:"auth_url"`
	TokenUrl      string            `mapstructure:"token_url"`
	OpenIDConfUrl string            `mapstructure:"openid_conf_url"`
	AuthRequest   ConfigAuthRequest `mapstructure:"authrequest"`
	RegisterUser  bool              `mapstructure:"register_user"`
}

type JwksConfig struct {
	PrivateKid     string   `mapstructure:"private_kid"`
	PrivateKeyPath string   `mapstructure:"private_key_path"`
	PublicKids     []string `mapstructure:"public_kids"`
	PublicKeyPaths []string `mapstructure:"public_key_paths"`
}

// Deprecated
// ConfigOauth2
// type ConfigOauth2 struct {
// 	Token             ConfigToken              `mapstructure:"token"`
// 	ClientID          string                   `mapstructure:"client_id"`
// 	ClientSecret      string                   `mapstructure:"client_secret"`
// 	RedirectUrl       string                   `mapstructure:"redirect_url"`
// 	Scopes            []string                 `mapstructure:"scopes"`
// 	AuthUrl           string                   `mapstructure:"auth_url"`
// 	TokenUrl          string                   `mapstructure:"token_url"`
// 	OpenIDConfUrl     string                   `mapstructure:"openid_conf_url"`
// 	AuthRequest       ConfigAuthRequest        `mapstructure:"authrequest"`
// 	IDTokenValidators []ConfigIDTokenValidator `mapstructure:"id_token_validators"`
// }

// Deprecated
// type ConfigIDTokenValidator struct {
// 	Type          string      `mapstructure:"type"`
// 	Token         ConfigToken `mapstructure:"token"`
// 	OpenIDConfUrl string      `mapstructure:"openid_conf_url"`
// }

// Deprecated
// type ConfigToken struct {
// 	Source             string `mapstructure:"source"`
// 	IDKeyName          string `mapstructure:"id_key_name"`
// 	IDTokenKeyName     string `mapstructure:"id_token_key_name"`
// 	TokenSourceKeyName string `mapstructure:"token_source_key_name"`
// }

type ConfigAuthRequest struct {
	ResponseUrl string `mapstructure:"response_url"`
	AuthUrl     string `mapstructure:"auth_url"`
	Wait        int    `mapstructure:"wait"`
}

// ConfigPayment is PG payment configuration.
type ConfigPayment struct {
	// pg client(kcp, nice...)
	PgType string `mapstructure:"pg_type"`
	// mobile, pc
	ClientType string   `mapstructure:"client_type"`
	PG         ConfigPG `mapstructure:"pg"`
}

type ConfigPG struct {
	Url                  string `mapstructure:"url"`
	ClientID             string `mapstructure:"client_id"`
	RawCertificate       string `mapstructure:"raw_certificate"`
	ReturnUrl            string `mapstructure:"return_url"`
	PrivateKeyFileToSign string `mapstructure:"private_key_file_to_sign"`
	TradeRequestHtmlFile string `mapstructure:"trade_request_html_file"`
	AllowedPayMethods    string `mapstructure:"allowed_pay_methods"`
	ShopName             string `mapstructure:"shop_name"`
}

type ConfigPartner struct {
	Address ConfigAddress `mapstructure:"address"`
}

type ConfigAddress struct {
	Type     string `mapstructure:"type"`
	HtmlFile string `mapstructure:"html_file"`
}
