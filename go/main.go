package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "program exited due to %v", err)
		os.Exit(1)
	}
}

func rootCommand() *cobra.Command {
	cfg := config{}

	// Root command.
	cmd := cobra.Command{
		Use: "mcgo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	// Use viper to manage environment variables and args.
	v := viper.New()
	v.SetEnvPrefix("MG")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	// Bind the flags to config object.
	flags := cmd.PersistentFlags()

	flags.StringVar(&cfg.host, "host", "localhost:9000", "Minio api url")
	flags.StringVar(&cfg.keyID, "key-id", "admin", "Access key id to access minio")
	flags.StringVar(&cfg.secretKey, "secret-key", "password", "Secret access key to access minio")
	flags.BoolVar(&cfg.useSSL, "use-ssl", false, "Use SSL to connect to minio")

	// CLI commands.
	cmd.AddCommand(pingCommand(&cfg, v))
	cmd.AddCommand(listCommand(&cfg, v))

	return &cmd
}

func main() {
	err := rootCommand().Execute()
	checkErr(err)
}
