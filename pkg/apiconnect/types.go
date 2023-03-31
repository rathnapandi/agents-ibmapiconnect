package apiconnect

// API -
type API struct {
	ApiSpec       []byte
	ID            string
	Name          string
	Description   string
	Version       string
	Url           string
	Documentation []byte
	AuthPolicy    string
	ComponentId   string
	EnvironmentId string
}

type AuthnRequest struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	ApiKey       string `json:"api_key,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	GrantType    string `json:"grant_type"`
}

type OauthToken struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	ExpiresIn        int    `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
}

type Credential struct {
	ClientId     string
	ClientSecret string
	ApiKey       string
}

type ListCatalogResponse struct {
	TotalResults int       `json:"total_results"`
	Results      []Results `json:"results"`
}

type Results struct {
	Type                  string   `json:"type"`
	ApiType               string   `json:"api_type"`
	ApiVersion            string   `json:"api_version"`
	Id                    string   `json:"id"`
	Name                  string   `json:"name"`
	Version               string   `json:"version"`
	Title                 string   `json:"title"`
	State                 string   `json:"state"`
	Scope                 string   `json:"scope"`
	GatewayType           string   `json:"gateway_type"`
	OaiVersion            string   `json:"oai_version"`
	DocumentSpecification string   `json:"document_specification"`
	BasePaths             []string `json:"base_paths"`
	Enforced              bool     `json:"enforced"`
	GatewayServiceUrls    []string `json:"gateway_service_urls"`
}

type Wsdl struct {
	ContentType string `json:"content_type"`
	Content     string `json:"content"`
}

type DatapowerGateway struct {
	CatalogBase string `json:"catalog_base"`
}
