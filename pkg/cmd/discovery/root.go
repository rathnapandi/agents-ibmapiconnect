package discovery

import (
	"github.com/Axway/agent-sdk/pkg/agent"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/apiconnect"
	"github.com/sirupsen/logrus"

	coreapi "github.com/Axway/agent-sdk/pkg/api"
	corecmd "github.com/Axway/agent-sdk/pkg/cmd"

	"github.com/Axway/agent-sdk/pkg/cmd/service"
	corecfg "github.com/Axway/agent-sdk/pkg/config"
	subs "github.com/rathnapandi/agents-ibmapiconnect/pkg/subscription"

	"github.com/rathnapandi/agents-ibmapiconnect/pkg/config"
	"github.com/rathnapandi/agents-ibmapiconnect/pkg/discovery"
)

// RootCmd - Agent root command
var (
	RootCmd        corecmd.AgentRootCmd
	discoveryAgent *discovery.Agent
)

func init() {
	// Create new root command with callbacks to initialize the agent config and command execution.
	// The first parameter identifies the name of the yaml file that agent will look for to load the config
	RootCmd = corecmd.NewRootCmd(
		"ibmapiconnect_discovery_agent", // Name of the yaml file
		"ibmapiconnect Discovery Agent", // Agent description
		initConfig,                      // Callback for initializing the agent config
		run,                             // Callback for executing the agent
		corecfg.DiscoveryAgent,          // Agent Type (Discovery or Traceability)
	)
	config.AddConfigProperties(RootCmd.GetProperties())

	RootCmd.AddCommand(service.GenServiceCmd("pathConfig"))
}

// run Callback that agent will call to process the execution
func run() error {
	discoveryAgent.Run()
	return nil
}

// initConfig Callback that agent will call to initialize the config. CentralConfig is parsed by Agent SDK
// and passed to the callback allowing the agent code to access the central config
func initConfig(centralConfig corecfg.CentralConfig) (interface{}, error) {
	conf := &config.AgentConfig{
		CentralConfig:    centralConfig,
		ApiConnectConfig: config.NewIbmApiconnectConfig(RootCmd.GetProperties(), centralConfig.GetAgentType()),
	}

	config.SetConfig(conf)

	logger := logrus.WithFields(logrus.Fields{
		"component": "agent",
	})
	client := coreapi.NewClient(conf.ApiConnectConfig.TLS, conf.ApiConnectConfig.ProxyURL)
	gatewayClient := apiconnect.NewClient(conf.ApiConnectConfig, client, false)
	if centralConfig.IsMarketplaceSubsEnabled() {
		agent.RegisterProvisioner(subs.NewProvisioner(gatewayClient, logger))
	}

	discoveryAgent = discovery.NewAgent(conf, gatewayClient)

	return conf, nil
}
