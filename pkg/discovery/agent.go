package discovery

import (
	"os"
	"os/signal"
	"syscall"

	coreagent "github.com/Axway/agent-sdk/pkg/agent"
	"github.com/Axway/agent-sdk/pkg/apic/provisioning"
	"github.com/Axway/agent-sdk/pkg/cache"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/apiconnect"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/config"
)

type Repeater interface {
	Loop()
	OnConfigChange(cfg *config.IbmApiConnectConfig)
	Stop()
}

// Agent -
type Agent struct {
	client    apiconnect.Client
	stopAgent chan bool
	discovery Repeater
	publisher Repeater
}

// NewAgent creates a new agent
func NewAgent(cfg *config.AgentConfig, client apiconnect.Client) (agent *Agent) {
	buffer := 5
	apiChan := make(chan *ServiceDetail, buffer)

	pub := &publisher{
		apiChan:     apiChan,
		stopPublish: make(chan bool),
		publishAPI:  coreagent.PublishAPI,
	}

	c := cache.New()

	svcHandler := &serviceHandler{
		client: client,
		cache:  c,
	}

	svcHandler.mode = marketplace

	disc := &discovery{
		apiChan:           apiChan,
		cache:             c,
		client:            client,
		centralClient:     coreagent.GetCentralClient(),
		discoveryPageSize: 50,
		pollInterval:      cfg.IbmApiConnectConfig.PollInterval,
		stopDiscovery:     make(chan bool),
		serviceHandler:    svcHandler,
	}

	return newAgent(client, disc, pub)
}

func newAgent(
	client apiconnect.Client,
	discovery Repeater,
	publisher Repeater,
) *Agent {
	return &Agent{
		client:    client,
		discovery: discovery,
		publisher: publisher,
		stopAgent: make(chan bool),
	}
}

// onConfigChange apply configuration changes
func (a *Agent) onConfigChange() {
	cfg := config.GetConfig()

	// Stop Discovery & Publish
	a.discovery.Stop()
	a.publisher.Stop()

	a.client.OnConfigChange(cfg.IbmApiConnectConfig)
	a.discovery.OnConfigChange(cfg.IbmApiConnectConfig)

	// Restart Discovery & Publish
	go a.discovery.Loop()
	go a.publisher.Loop()
}

// Run the agent loop
func (a *Agent) Run() {
	coreagent.OnConfigChange(a.onConfigChange)
	if config.GetConfig().CentralConfig.IsMarketplaceSubsEnabled() {
		// register request
		regHrdProp := provisioning.NewSchemaPropertyBuilder().
			SetName("requestHeader").
			SetLabel("Request Header").
			SetReadOnly().
			IsString().
			SetDefaultValue("x-api-key")

		// api key
		coreagent.NewAPIKeyCredentialRequestBuilder(
			coreagent.WithCRDName(provisioning.APIKeyCRD),
			coreagent.WithCRDTitle("API Key"),
			coreagent.WithCRDRequestSchemaProperty(regHrdProp),
		).Register()
		coreagent.NewAPIKeyAccessRequestBuilder().Register()
	}

	// // basic auth
	// coreagent.NewBasicAuthCredentialRequestBuilder(
	// 	coreagent.WithCRDIsSuspendable(),
	// ).Register()
	// coreagent.NewBasicAuthAccessRequestBuilder().Register()

	go a.discovery.Loop()
	go a.publisher.Loop()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM, os.Interrupt)

	<-gracefulStop
	a.Stop()
}

// Stop stops the discovery agent.
func (a *Agent) Stop() {
	a.discovery.Stop()
	a.publisher.Stop()
	close(a.stopAgent)
}
