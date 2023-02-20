package dto

type OpenIDConf struct {
	Issuer                string `json:"issuer" mapstructure:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint" mapstructure:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint" mapstructure:"token_endpoint"`
	JwksUri               string `json:"jwks_uri" mapstructure:"jwks_uri"`
	UserinfoEndpoint      string `json:"userinfo_endpoint" mapstructure:"userinfo_endpoint"`
	RevocationEndpoint    string `json:"revocation_endpoint" mapstructure:"revocation_endpoint"`
	EndSessionEndpoint    string `json:"end_session_endpoint" mapstructure:"end_session_endpoint"`

	IntrospectEndpoint string `json:"introspect_endpoint" mapstructure:"introspect_endpoint"`
}
