package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// BindEnv Bind each cobra flag to its associated viper configuration (config file and environment variable)
func BindEnv(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := strings.ReplaceAll(f.Name, ".", "_")
		configValue := os.Getenv(strings.ToUpper(configName))

		if !f.Changed && configValue != "" {
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", configValue))
		}
	})
}
