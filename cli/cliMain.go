package scencli

import (
	scenclibase "github.com/TerraDharitri/drt-go-chain-scenario/clibase"
)

func ScenariosCLI(version string) {
	scenclibase.ScenariosCLI(version, &runConfig{})
}
