package config

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type cliFlags struct {
	showVersion bool `mapstructure:"version"`
	showHelp    bool `mapstructure:"help"`
	dumpConfig  bool `mapstructure:"dump_config"`
}

func loadFlags() error {
	if !anyFlagsPassed() {
		return nil
	}

	cli := &cliFlags{}
	buxFlags := pflag.NewFlagSet("buxFlags", pflag.ContinueOnError)

	initFlags(buxFlags, cli)

	err := buxFlags.Parse(os.Args[1:])
	if err != nil {
		fmt.Printf("Flags can't be parsed: %v\n", err)
		os.Exit(1)
	}

	err = viper.BindPFlags(buxFlags)
	if err != nil {
		return err
	}

	parseCliFlags(buxFlags, cli)

	return nil
}

func anyFlagsPassed() bool {
	return len(os.Args) > 1
}

func initFlags(fs *pflag.FlagSet, cliFlags *cliFlags) {
	fs.StringP(ConfigFilePathKey, "C", "", "custom config file path")

	fs.BoolVarP(&cliFlags.showHelp, "help", "h", false, "show help")
	fs.BoolVarP(&cliFlags.showVersion, "version", "v", false, "show version")
	fs.BoolVarP(&cliFlags.dumpConfig, "dump_config", "d", false, "dump config to file, specified by config_file flag")
}

func parseCliFlags(fs *pflag.FlagSet, cli *cliFlags) {
	if cli.showHelp {
		fs.PrintDefaults()
		os.Exit(0)
	}

	if cli.showVersion {
		fmt.Println("bux-sever", "version", Version)
		os.Exit(0)
	}

	if cli.dumpConfig {
		configPath := viper.GetString(ConfigFilePathKey)
		if configPath == "" {
			configPath = DefaultConfigFilePath
		}

		err := viper.SafeWriteConfigAs(configPath)
		if err != nil {
			fmt.Printf("error while dumping config: %v", err.Error())
		}
		os.Exit(0)
	}
}