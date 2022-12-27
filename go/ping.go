package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func pingCommand(cfg *config, v *viper.Viper) *cobra.Command {
	// List command
	cmd := cobra.Command{
		Use:     "ping",
		Short:   "Ping the minio server",
		Example: "mcgo ping",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newMinioClient(cfg)
			if err != nil {
				return err
			}

			fmt.Printf("Successfully connected to %v\n", client.EndpointURL())

			return nil
		},
	}

	return &cmd
}
