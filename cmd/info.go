/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type CLIConfig struct {
	project string
}

var (
	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --namespace is bound to CLI_NAMESPACE.
	envPrefix = "CLI"

	// Replace hyphenated flag names with camelCase in the config file
	replaceHyphenWithCamelCase = false

	debugEnabled bool
	namespace    string
	cliConfig    CLIConfig
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Debug("info called")
		logrus.Infof("Debug Enabled from flag: %v", debugEnabled)
		logrus.Infof("Debug Enabled from viper: %v", viper.GetString("debugEnabled"))
		logrus.Infof("Namespace from flag: %s", namespace)
		logrus.Infof("Namespace from viper: %s", viper.GetString("namespace"))
		logrus.Infof("Project from flag: %s", cliConfig.project)
		logrus.Infof("Project from viper: %s", viper.GetString("project"))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	infoCmd.PersistentFlags().BoolVarP(&debugEnabled, "debugEnabled", "d", false, "Debug enabled")
	infoCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Current namespace")
	infoCmd.Flags().StringVarP(&cliConfig.project, "project", "p", "", "Current project")

	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
	viper.SetDefault("namespace", "fromSetDefault")
	viper.SetDefault("project", "fromSetDefault")

	// Bind flags
	viper.BindPFlags(infoCmd.Flags())
	viper.BindPFlags(infoCmd.PersistentFlags())
	// viper.BindPFlag("namespace", infoCmd.Flags().Lookup("namespace"))
	// infoCmd.Flags().VisitAll(func(f *pflag.Flag) {
	// 	viper.BindPFlag(f.Name, infoCmd.Flags().Lookup(f.Name))
	// })

	// Bind viper
	bindViper(infoCmd.Flags())
	bindViper(infoCmd.PersistentFlags())
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindViper(flagSet *pflag.FlagSet) {
	flagSet.VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name
		logrus.Infof("Config key name: %s", configName)

		// If using camelCase in the config file, replace hyphens with a camelCased string.
		// Since viper does case-insensitive comparisons, we don't need to bother fixing the case, and only need to remove the hyphens.
		if replaceHyphenWithCamelCase {
			configName = strings.ReplaceAll(f.Name, "-", "")
		}

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && viper.IsSet(configName) {
			val := viper.Get(configName)
			flagSet.Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
