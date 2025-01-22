package scencli

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	scenclibase "github.com/TerraDharitri/drt-go-chain-scenario/clibase"
	scenio "github.com/TerraDharitri/drt-go-chain-scenario/scenario/io"
	vm15scenario "github.com/TerraDharitri/drt-go-chain-vm/scenario"
	vm15wasmer "github.com/TerraDharitri/drt-go-chain-vm/wasmer"
	vm15wasmer2 "github.com/TerraDharitri/drt-go-chain-vm/wasmer2"
)

func resolveArgument(exeDir string, arg string) (string, bool, error) {
	fi, err := os.Stat(arg)
	if os.IsNotExist(err) {
		arg = filepath.Join(exeDir, arg)
		fmt.Println(arg)
		fi, err = os.Stat(arg)
	}
	if err != nil {
		return "", false, err
	}
	return arg, fi.IsDir(), nil
}

func parseOptionFlags() scenclibase.CLIRunOptions {
	forceTraceGas := flag.Bool("force-trace-gas", false, "overrides the traceGas option in the scenarios")
	useWasmer1 := flag.Bool("wasmer1", false, "use the wasmer1 executor")
	useWasmer2 := flag.Bool("wasmer2", false, "use the wasmer2 executor")
	flag.Parse()

	vmBuilder := vm15scenario.NewScenarioVMHostBuilder()
	if *useWasmer1 {
		vmBuilder.OverrideVMExecutor = vm15wasmer.ExecutorFactory()
	}
	if *useWasmer2 {
		vmBuilder.OverrideVMExecutor = vm15wasmer2.ExecutorFactory()
	}

	return scenclibase.CLIRunOptions{
		RunOptions: &scenio.RunScenarioOptions{
			ForceTraceGas: *forceTraceGas,
		},
		VMBuilder: vmBuilder,
	}
}

// ScenariosCLILegacy is the original CLI that was available in VM 1.5.
//
// Not currently used anywhere.
func ScenariosCLILegacy() {
	options := parseOptionFlags()

	// directory of this executable
	exeDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// argument
	args := flag.Args()
	if len(args) < 1 {
		panic("One argument expected - the path to the json test or directory.")
	}
	jsonFilePath, _, err := resolveArgument(exeDir, args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = scenclibase.RunScenariosAtPath(jsonFilePath, options)

	if err != nil {
		os.Exit(1)
	}
}

func FmtLegacyCli() {
	if len(os.Args) != 2 {
		panic("One argument expected - the root path where to search.")
	}

	_ = scenio.FormatAllInFolder(os.Args[1])
}
