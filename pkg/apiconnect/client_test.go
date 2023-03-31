package apiconnect

import (
	"fmt"
	"testing"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestListApi(t *testing.T) {

	response := `{
		"total_results": 1,
		"results": [
			{
				"type": "api",
				"api_type": "rest",
				"api_version": "2.0.0",
				"id": "daea6c77-ceef-4c41-b55f-acea0204789e",
				"name": "calculator_x0020_web_x0020_service",
				"version": "1.0.0",
				"title": "Calculator_x0020_Web_x0020_Service",
				"state": "online",
				"scope": "catalog",
				"gateway_type": "datapower-api-gateway",
				"oai_version": "openapi2",
				"document_specification": "openapi2",
				"base_paths": [
					"/calculator"
				],
				"enforced": true,
				"gateway_service_urls": [
					"https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a"
				],
				"user_registry_urls": [],
				"oauth_provider_urls": [],
				"tls_client_profile_urls": [],
				"extension_urls": [],
				"policy_urls": [
					"https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/b2248832-8bc7-4d0b-8384-acace3d90761",
					"https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/31bc61a7-422b-4c69-acad-27e957e7000c",
					"https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/01a0c2f2-7afb-40b4-befa-8a78333816c5",
					"https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/configured-gateway-services/44aa500d-7b95-4b1d-afd3-1a0b0fab968a/policies/3e8a2e1e-a96e-4fba-a7c4-fba06e1ed253"
				],
				"created_at": "2023-03-30T16:57:52.000Z",
				"updated_at": "2023-03-30T17:01:32.000Z",
				"org_url": "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/orgs/ac63f85d-5c08-41ef-b762-7daeceacc191",
				"catalog_url": "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547",
				"url": "https://platform-api.us-east-a.apiconnect.automation.ibm.com/api/catalogs/ac63f85d-5c08-41ef-b762-7daeceacc191/bb314580-777e-4e55-8429-f3e8be310547/apis/daea6c77-ceef-4c41-b55f-acea0204789e"
			}
		]
	}
	`
	mc := &MockClient{}
	cfg := &config.ApiConnectConfig{
		IbmApiConnectUrl: "",
		CachePath:        "/tmp",
		ClientId:         "abc",
		ClientSecret:     "xyz",
		ApiKey:           "xyz",
		ProxyURL:         "",
	}
	apiConnectClient := NewClient(cfg, mc, true)
	apiConnectClient.auth.SetToken("xyz")

	mc.SendFunc = func(request coreapi.Request) (*coreapi.Response, error) {
		return &coreapi.Response{
			Code: 200,
			Body: []byte(response),
		}, nil
	}

	apiResponse, err := apiConnectClient.ListAPIs()

	assert.Nil(t, err)
	fmt.Println(apiResponse)

}

func TestGetWsdl(t *testing.T) {

	response := `{
		"content_type": "application/wsdl",
		"content": "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0idXRmLTgiPz4KPHdzZGw6ZGVmaW5pdGlvbnMgeG1sbnM6dG09Imh0dHA6Ly9taWNyb3NvZnQuY29tL3dzZGwvbWltZS90ZXh0TWF0Y2hpbmcvIiB4bWxuczpzb2FwZW5jPSJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy9zb2FwL2VuY29kaW5nLyIgeG1sbnM6bWltZT0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3NkbC9taW1lLyIgeG1sbnM6dG5zPSJodHRwOi8vdGVtcHVyaS5vcmcvIiB4bWxuczpzb2FwPSJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93c2RsL3NvYXAvIiB4bWxuczpzPSJodHRwOi8vd3d3LnczLm9yZy8yMDAxL1hNTFNjaGVtYSIgeG1sbnM6c29hcDEyPSJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93c2RsL3NvYXAxMi8iIHhtbG5zOmh0dHA9Imh0dHA6Ly9zY2hlbWFzLnhtbHNvYXAub3JnL3dzZGwvaHR0cC8iIHRhcmdldE5hbWVzcGFjZT0iaHR0cDovL3RlbXB1cmkub3JnLyIgeG1sbnM6d3NkbD0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3NkbC8iPgogICAgPHdzZGw6ZG9jdW1lbnRhdGlvbiB4bWxuczp3c2RsPSJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93c2RsLyI+UGVyZm9ybXMgc2ltcGxlIG1hdGggb3ZlciB0aGUgV2ViPC93c2RsOmRvY3VtZW50YXRpb24+CiAgICA8d3NkbDp0eXBlcz4KICAgICAgICA8czpzY2hlbWEgZWxlbWVudEZvcm1EZWZhdWx0PSJxdWFsaWZpZWQiIHRhcmdldE5hbWVzcGFjZT0iaHR0cDovL3RlbXB1cmkub3JnLyI+CiAgICAgICAgICAgIDxzOmVsZW1lbnQgbmFtZT0iQWRkIj4KICAgICAgICAgICAgICAgIDxzOmNvbXBsZXhUeXBlPgogICAgICAgICAgICAgICAgICAgIDxzOnNlcXVlbmNlPgogICAgICAgICAgICAgICAgICAgICAgICA8czplbGVtZW50IG1pbk9jY3Vycz0iMSIgbWF4T2NjdXJzPSIxIiBuYW1lPSJhIiB0eXBlPSJzOmludCIgLz4KICAgICAgICAgICAgICAgICAgICAgICAgPHM6ZWxlbWVudCBtaW5PY2N1cnM9IjEiIG1heE9jY3Vycz0iMSIgbmFtZT0iYiIgdHlwZT0iczppbnQiIC8+CiAgICAgICAgICAgICAgICAgICAgPC9zOnNlcXVlbmNlPgogICAgICAgICAgICAgICAgPC9zOmNvbXBsZXhUeXBlPgogICAgICAgICAgICA8L3M6ZWxlbWVudD4KICAgICAgICAgICAgPHM6ZWxlbWVudCBuYW1lPSJBZGRSZXNwb25zZSI+CiAgICAgICAgICAgICAgICA8czpjb21wbGV4VHlwZT4KICAgICAgICAgICAgICAgICAgICA8czpzZXF1ZW5jZT4KICAgICAgICAgICAgICAgICAgICAgICAgPHM6ZWxlbWVudCBtaW5PY2N1cnM9IjEiIG1heE9jY3Vycz0iMSIgbmFtZT0iQWRkUmVzdWx0IiB0eXBlPSJzOmludCIgLz4KICAgICAgICAgICAgICAgICAgICA8L3M6c2VxdWVuY2U+CiAgICAgICAgICAgICAgICA8L3M6Y29tcGxleFR5cGU+CiAgICAgICAgICAgIDwvczplbGVtZW50PgogICAgICAgICAgICA8czplbGVtZW50IG5hbWU9IlN1YnRyYWN0Ij4KICAgICAgICAgICAgICAgIDxzOmNvbXBsZXhUeXBlPgogICAgICAgICAgICAgICAgICAgIDxzOnNlcXVlbmNlPgogICAgICAgICAgICAgICAgICAgICAgICA8czplbGVtZW50IG1pbk9jY3Vycz0iMSIgbWF4T2NjdXJzPSIxIiBuYW1lPSJhIiB0eXBlPSJzOmludCIgLz4KICAgICAgICAgICAgICAgICAgICAgICAgPHM6ZWxlbWVudCBtaW5PY2N1cnM9IjEiIG1heE9jY3Vycz0iMSIgbmFtZT0iYiIgdHlwZT0iczppbnQiIC8+CiAgICAgICAgICAgICAgICAgICAgPC9zOnNlcXVlbmNlPgogICAgICAgICAgICAgICAgPC9zOmNvbXBsZXhUeXBlPgogICAgICAgICAgICA8L3M6ZWxlbWVudD4KICAgICAgICAgICAgPHM6ZWxlbWVudCBuYW1lPSJTdWJ0cmFjdFJlc3BvbnNlIj4KICAgICAgICAgICAgICAgIDxzOmNvbXBsZXhUeXBlPgogICAgICAgICAgICAgICAgICAgIDxzOnNlcXVlbmNlPgogICAgICAgICAgICAgICAgICAgICAgICA8czplbGVtZW50IG1pbk9jY3Vycz0iMSIgbWF4T2NjdXJzPSIxIiBuYW1lPSJTdWJ0cmFjdFJlc3VsdCIgdHlwZT0iczppbnQiIC8+CiAgICAgICAgICAgICAgICAgICAgPC9zOnNlcXVlbmNlPgogICAgICAgICAgICAgICAgPC9zOmNvbXBsZXhUeXBlPgogICAgICAgICAgICA8L3M6ZWxlbWVudD4KICAgICAgICA8L3M6c2NoZW1hPgogICAgPC93c2RsOnR5cGVzPgogICAgPHdzZGw6bWVzc2FnZSBuYW1lPSJBZGRTb2FwSW4iPgogICAgICAgIDx3c2RsOnBhcnQgbmFtZT0icGFyYW1ldGVycyIgZWxlbWVudD0idG5zOkFkZCIgLz4KICAgIDwvd3NkbDptZXNzYWdlPgogICAgPHdzZGw6bWVzc2FnZSBuYW1lPSJBZGRTb2FwT3V0Ij4KICAgICAgICA8d3NkbDpwYXJ0IG5hbWU9InBhcmFtZXRlcnMiIGVsZW1lbnQ9InRuczpBZGRSZXNwb25zZSIgLz4KICAgIDwvd3NkbDptZXNzYWdlPgogICAgPHdzZGw6bWVzc2FnZSBuYW1lPSJTdWJ0cmFjdFNvYXBJbiI+CiAgICAgICAgPHdzZGw6cGFydCBuYW1lPSJwYXJhbWV0ZXJzIiBlbGVtZW50PSJ0bnM6U3VidHJhY3QiIC8+CiAgICA8L3dzZGw6bWVzc2FnZT4KICAgIDx3c2RsOm1lc3NhZ2UgbmFtZT0iU3VidHJhY3RTb2FwT3V0Ij4KICAgICAgICA8d3NkbDpwYXJ0IG5hbWU9InBhcmFtZXRlcnMiIGVsZW1lbnQ9InRuczpTdWJ0cmFjdFJlc3BvbnNlIiAvPgogICAgPC93c2RsOm1lc3NhZ2U+CiAgICA8d3NkbDpwb3J0VHlwZSBuYW1lPSJDYWxjdWxhdG9yX3gwMDIwX1dlYl94MDAyMF9TZXJ2aWNlU29hcCI+CiAgICAgICAgPHdzZGw6b3BlcmF0aW9uIG5hbWU9IkFkZCI+CiAgICAgICAgICAgIDx3c2RsOmRvY3VtZW50YXRpb24geG1sbnM6d3NkbD0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3NkbC8iPkNvbXB1dGVzIHRoZSBzdW0gb2YgdHdvIGludGVnZXJzPC93c2RsOmRvY3VtZW50YXRpb24+CiAgICAgICAgICAgIDx3c2RsOmlucHV0IG1lc3NhZ2U9InRuczpBZGRTb2FwSW4iIC8+CiAgICAgICAgICAgIDx3c2RsOm91dHB1dCBtZXNzYWdlPSJ0bnM6QWRkU29hcE91dCIgLz4KICAgICAgICA8L3dzZGw6b3BlcmF0aW9uPgogICAgICAgIDx3c2RsOm9wZXJhdGlvbiBuYW1lPSJTdWJ0cmFjdCI+CiAgICAgICAgICAgIDx3c2RsOmRvY3VtZW50YXRpb24geG1sbnM6d3NkbD0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvd3NkbC8iPkNvbXB1dGVzIHRoZSBkaWZmZXJlbmNlIGJldHdlZW4gdHdvIGludGVnZXJzPC93c2RsOmRvY3VtZW50YXRpb24+CiAgICAgICAgICAgIDx3c2RsOmlucHV0IG1lc3NhZ2U9InRuczpTdWJ0cmFjdFNvYXBJbiIgLz4KICAgICAgICAgICAgPHdzZGw6b3V0cHV0IG1lc3NhZ2U9InRuczpTdWJ0cmFjdFNvYXBPdXQiIC8+CiAgICAgICAgPC93c2RsOm9wZXJhdGlvbj4KICAgIDwvd3NkbDpwb3J0VHlwZT4KICAgIDx3c2RsOmJpbmRpbmcgbmFtZT0iQ2FsY3VsYXRvcl94MDAyMF9XZWJfeDAwMjBfU2VydmljZVNvYXAiIHR5cGU9InRuczpDYWxjdWxhdG9yX3gwMDIwX1dlYl94MDAyMF9TZXJ2aWNlU29hcCI+CiAgICAgICAgPHNvYXA6YmluZGluZyB0cmFuc3BvcnQ9Imh0dHA6Ly9zY2hlbWFzLnhtbHNvYXAub3JnL3NvYXAvaHR0cCIgLz4KICAgICAgICA8d3NkbDpvcGVyYXRpb24gbmFtZT0iQWRkIj4KICAgICAgICAgICAgPHNvYXA6b3BlcmF0aW9uIHNvYXBBY3Rpb249Imh0dHA6Ly90ZW1wdXJpLm9yZy9BZGQiIHN0eWxlPSJkb2N1bWVudCIgLz4KICAgICAgICAgICAgPHdzZGw6aW5wdXQ+CiAgICAgICAgICAgICAgICA8c29hcDpib2R5IHVzZT0ibGl0ZXJhbCIgLz4KICAgICAgICAgICAgPC93c2RsOmlucHV0PgogICAgICAgICAgICA8d3NkbDpvdXRwdXQ+CiAgICAgICAgICAgICAgICA8c29hcDpib2R5IHVzZT0ibGl0ZXJhbCIgLz4KICAgICAgICAgICAgPC93c2RsOm91dHB1dD4KICAgICAgICA8L3dzZGw6b3BlcmF0aW9uPgogICAgICAgIDx3c2RsOm9wZXJhdGlvbiBuYW1lPSJTdWJ0cmFjdCI+CiAgICAgICAgICAgIDxzb2FwOm9wZXJhdGlvbiBzb2FwQWN0aW9uPSJodHRwOi8vdGVtcHVyaS5vcmcvU3VidHJhY3QiIHN0eWxlPSJkb2N1bWVudCIgLz4KICAgICAgICAgICAgPHdzZGw6aW5wdXQ+CiAgICAgICAgICAgICAgICA8c29hcDpib2R5IHVzZT0ibGl0ZXJhbCIgLz4KICAgICAgICAgICAgPC93c2RsOmlucHV0PgogICAgICAgICAgICA8d3NkbDpvdXRwdXQ+CiAgICAgICAgICAgICAgICA8c29hcDpib2R5IHVzZT0ibGl0ZXJhbCIgLz4KICAgICAgICAgICAgPC93c2RsOm91dHB1dD4KICAgICAgICA8L3dzZGw6b3BlcmF0aW9uPgogICAgPC93c2RsOmJpbmRpbmc+CiAgICA8d3NkbDpiaW5kaW5nIG5hbWU9IkNhbGN1bGF0b3JfeDAwMjBfV2ViX3gwMDIwX1NlcnZpY2VTb2FwMTIiIHR5cGU9InRuczpDYWxjdWxhdG9yX3gwMDIwX1dlYl94MDAyMF9TZXJ2aWNlU29hcCI+CiAgICAgICAgPHNvYXAxMjpiaW5kaW5nIHRyYW5zcG9ydD0iaHR0cDovL3NjaGVtYXMueG1sc29hcC5vcmcvc29hcC9odHRwIiAvPgogICAgICAgIDx3c2RsOm9wZXJhdGlvbiBuYW1lPSJBZGQiPgogICAgICAgICAgICA8c29hcDEyOm9wZXJhdGlvbiBzb2FwQWN0aW9uPSJodHRwOi8vdGVtcHVyaS5vcmcvQWRkIiBzdHlsZT0iZG9jdW1lbnQiIC8+CiAgICAgICAgICAgIDx3c2RsOmlucHV0PgogICAgICAgICAgICAgICAgPHNvYXAxMjpib2R5IHVzZT0ibGl0ZXJhbCIgLz4KICAgICAgICAgICAgPC93c2RsOmlucHV0PgogICAgICAgICAgICA8d3NkbDpvdXRwdXQ+CiAgICAgICAgICAgICAgICA8c29hcDEyOmJvZHkgdXNlPSJsaXRlcmFsIiAvPgogICAgICAgICAgICA8L3dzZGw6b3V0cHV0PgogICAgICAgIDwvd3NkbDpvcGVyYXRpb24+CiAgICAgICAgPHdzZGw6b3BlcmF0aW9uIG5hbWU9IlN1YnRyYWN0Ij4KICAgICAgICAgICAgPHNvYXAxMjpvcGVyYXRpb24gc29hcEFjdGlvbj0iaHR0cDovL3RlbXB1cmkub3JnL1N1YnRyYWN0IiBzdHlsZT0iZG9jdW1lbnQiIC8+CiAgICAgICAgICAgIDx3c2RsOmlucHV0PgogICAgICAgICAgICAgICAgPHNvYXAxMjpib2R5IHVzZT0ibGl0ZXJhbCIgLz4KICAgICAgICAgICAgPC93c2RsOmlucHV0PgogICAgICAgICAgICA8d3NkbDpvdXRwdXQ+CiAgICAgICAgICAgICAgICA8c29hcDEyOmJvZHkgdXNlPSJsaXRlcmFsIiAvPgogICAgICAgICAgICA8L3dzZGw6b3V0cHV0PgogICAgICAgIDwvd3NkbDpvcGVyYXRpb24+CiAgICA8L3dzZGw6YmluZGluZz4KICAgIDx3c2RsOnNlcnZpY2UgbmFtZT0iQ2FsY3VsYXRvcl94MDAyMF9XZWJfeDAwMjBfU2VydmljZSI+CiAgICAgICAgPHdzZGw6ZG9jdW1lbnRhdGlvbiB4bWxuczp3c2RsPSJodHRwOi8vc2NoZW1hcy54bWxzb2FwLm9yZy93c2RsLyI+UGVyZm9ybXMgc2ltcGxlIG1hdGggb3ZlciB0aGUgV2ViPC93c2RsOmRvY3VtZW50YXRpb24+CiAgICAgICAgPHdzZGw6cG9ydCBuYW1lPSJDYWxjdWxhdG9yX3gwMDIwX1dlYl94MDAyMF9TZXJ2aWNlU29hcCIgYmluZGluZz0idG5zOkNhbGN1bGF0b3JfeDAwMjBfV2ViX3gwMDIwX1NlcnZpY2VTb2FwIj4KICAgICAgICAgICAgPHNvYXA6YWRkcmVzcyBsb2NhdGlvbj0iaHR0cHM6Ly9lY3Muc3lyLmVkdS9mYWN1bHR5L2Zhd2NldHQvSGFuZG91dHMvY3NlNzc1L2NvZGUvY2FsY1dlYlNlcnZpY2UvQ2FsYy5hc214IiAvPgogICAgICAgIDwvd3NkbDpwb3J0PgogICAgICAgIDx3c2RsOnBvcnQgbmFtZT0iQ2FsY3VsYXRvcl94MDAyMF9XZWJfeDAwMjBfU2VydmljZVNvYXAxMiIgYmluZGluZz0idG5zOkNhbGN1bGF0b3JfeDAwMjBfV2ViX3gwMDIwX1NlcnZpY2VTb2FwMTIiPgogICAgICAgICAgICA8c29hcDEyOmFkZHJlc3MgbG9jYXRpb249Imh0dHBzOi8vZWNzLnN5ci5lZHUvZmFjdWx0eS9mYXdjZXR0L0hhbmRvdXRzL2NzZTc3NS9jb2RlL2NhbGNXZWJTZXJ2aWNlL0NhbGMuYXNteCIgLz4KICAgICAgICA8L3dzZGw6cG9ydD4KICAgIDwvd3NkbDpzZXJ2aWNlPgo8L3dzZGw6ZGVmaW5pdGlvbnM+"
	}
	`
	mc := &MockClient{}
	cfg := &config.ApiConnectConfig{
		IbmApiConnectUrl: "",
		CachePath:        "/tmp",
		ClientId:         "abc",
		ClientSecret:     "xyz",
		ApiKey:           "xyz",
		ProxyURL:         "",
	}
	apiConnectClient := NewClient(cfg, mc, true)
	apiConnectClient.auth.SetToken("xyz")

	mc.SendFunc = func(request coreapi.Request) (*coreapi.Response, error) {
		return &coreapi.Response{
			Code: 200,
			Body: []byte(response),
		}, nil
	}

	wsdl, err := apiConnectClient.DownloadWsdl("daea6c77-ceef-4c41-b55f-acea0204789e")

	assert.Nil(t, err)
	assert.Equal(t, wsdl.ContentType, "application/wsdl")
}

func TestGetWsdlFailure(t *testing.T) {

	response := `{
		"status": 404,
		"message": [
			"Not found"
		]
	}
	`
	mc := &MockClient{}
	cfg := &config.ApiConnectConfig{
		IbmApiConnectUrl: "",
		CachePath:        "/tmp",
		ClientId:         "abc",
		ClientSecret:     "xyz",
		ApiKey:           "xyz",
		ProxyURL:         "",
	}
	apiConnectClient := NewClient(cfg, mc, true)
	apiConnectClient.auth.SetToken("xyz")

	mc.SendFunc = func(request coreapi.Request) (*coreapi.Response, error) {
		return &coreapi.Response{
			Code: 404,
			Body: []byte(response),
		}, nil
	}

	wsdl, err := apiConnectClient.DownloadWsdl("daea6c77-ceef-4c41-b55f-acea0204789e")

	assert.Nil(t, err)

	assert.Equal(t, wsdl.ContentType, "")
}
