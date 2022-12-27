package main

import (
	"context"
	"fmt"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	reqTimeout = 10 * time.Second
)

func list(client *minio.Client, path string) error {
	ctx, cancel := context.WithTimeout(context.Background(), reqTimeout)
	defer cancel()

	if path == "" {
		// List buckets.
		buckets, err := client.ListBuckets(ctx)
		if err != nil {
			return err
		}

		for _, bucket := range buckets {
			fmt.Printf("%-15v %s\n", bucket.CreationDate, bucket.Name)
		}

		return nil
	}

	// List objects.
	objs := client.ListObjects(ctx, path, minio.ListObjectsOptions{
		Recursive: false,
	})

	for obj := range objs {
		if obj.Err != nil {
			return obj.Err
		}

		fmt.Printf("%-15v %s\n", obj.LastModified, obj.Key)
	}

	return nil
}

func listCommand(cfg *config, v *viper.Viper) *cobra.Command {
	// List command
	cmd := cobra.Command{
		Use:     "list",
		Short:   "List resources on minio",
		Example: "mcgo list <buckets|objects>",
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var path string
			if len(args) > 0 {
				path = args[0]
			}

			client, err := newMinioClient(cfg)
			if err != nil {
				return err
			}

			return list(client, path)
		},
	}

	return &cmd
}
