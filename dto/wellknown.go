package dto

type OpenIDConf struct {
	Issuer                string `json:"issuer,omitempty" mapstructure:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint,omitempty" mapstructure:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint,omitempty" mapstructure:"token_endpoint"`
	JwksUri               string `json:"jwks_uri,omitempty" mapstructure:"jwks_uri"`
	UserinfoEndpoint      string `json:"userinfo_endpoint,omitempty" mapstructure:"userinfo_endpoint"`
	RevocationEndpoint    string `json:"revocation_endpoint,omitempty" mapstructure:"revocation_endpoint"`
	EndSessionEndpoint    string `json:"end_session_endpoint,omitempty" mapstructure:"end_session_endpoint"`

	IntrospectEndpoint string `json:"introspect_endpoint,omitempty" mapstructure:"introspect_endpoint"`
}
