package main

import (
	"fmt"
	"os"

	// Required Import to setup factory for traceability transport
	_ "github.com/Axway/agent-sdk/pkg/traceability"

	"github.com/rathnapandi/agents-ibmapiconnect/pkg/cmd/traceability"
)

func main() {
	if err := traceability.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
