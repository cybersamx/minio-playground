package main

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func newMinioClient(cfg *config) (*minio.Client, error) {
	creds := credentials.NewStaticV4(cfg.keyID, cfg.secretKey, "")

	return minio.New(cfg.host, &minio.Options{
		Creds:  creds,
		Secure: false,
	})
}
