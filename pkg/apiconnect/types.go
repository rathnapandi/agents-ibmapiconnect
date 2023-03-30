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

type OauthToken struct {
	AccessToken        string `json:"access_token"`
	TokenType          string `json:"token_type"`
	ExpiresIn          int    `json:"expires_in"`
	RefreshToken       string `json:"refresh_token"`
	refresh_expires_in int    `json:"refresh_expires_in"`
}

type Credential struct {
	ClientId     string
	ClientSecret string
	ApiKey       string
}

type ListCatalogResponse struct {
	Total_results string    `json:"total_results"`
	Results       []Results `json:"results"`
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
}

// {
//     "total_results": 1,
//     "results": [
//         {
//             "type": "api",
//             "api_type": "rest",
//             "api_version": "2.0.0",
//             "id": "daea6c77-ceef-4c41-b55f-acea0204789e",
//             "name": "calculator_x0020_web_x0020_service",
//             "version": "1.0.0",
//             "title": "Calculator_x0020_Web_x0020_Service",
//             "state": "online",
//             "scope": "catalog",
//             "gateway_type": "datapower-api-gateway",
//             "oai_version": "openapi2",
//             "document_specification": "openapi2",
//             "base_paths": [
//                 "/calculator"
//             ],
//             "enforced": true,
//             "gateway_service_urls": [
//                 "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a"
//             ],
//             "user_registry_urls": [],
//             "oauth_provider_urls": [],
//             "tls_client_profile_urls": [],
//             "extension_urls": [],
//             "policy_urls": [
//                 "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/b2248832-8bc7-4d0b-8384-acace3d90761",
//                 "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/31bc61a7-422b-4c69-acad-27e957e7000c",
//                 "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/01a0c2f2-7afb-40b4-befa-8a78333816c5",
//                 "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/3e8a2e1e-a96e-4fba-a7c4-fba06e1ed253"
//             ],
//             "created_at": "2023-03-30T16:57:52.000Z",
//             "updated_at": "2023-03-30T17:01:32.000Z",
//             "org_url": "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/orgs/ac63f85d-5c08-41ef-b762-7daeceacc191",
//             "catalog_url": "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547",
//             "url": "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/apis/daea6c77-ceef-4c41-b55f-acea0204789e"
//         }
//     ]
// }
