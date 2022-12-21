# Minio Playground

This is a playground for [minio](https://min.io), a programmatic middleware for providing a S3-compatible API to many storage solutions from local filesystems to network storage to public Cloud storage like S3, Azure Blob, or Google Cloud Storage (GCS). This means a seamless integration and deployment of code regardless what the underlying storage solution is.

## Context and Use Case

This project was created as part of a spike in exploring a reliable data layer for an ML model training workload. During model development/experimenting, we want to use the local filesystem to read/write the datasets. For training or serving the trained model to production, we will deploy the model to the public Cloud to be executed. We want to leverage object storage solutions on the Cloud providers for the model's inputs and outputs. We want to use minio to cleanly abstract the code from the underlying data storage, enabling portability at least to start, with the local filesystem and the public Cloud providers.

## Setup

### MacOS

1. Install minio via Homebrew.

   ```shell
   brew install minio/stable/minio
   ```
   
1. Run minio server.

   ```shell
   mkdir -p ~/minio/data
   export MINIO_ROOT_USER=minio
   export MINIO_ROOT_PASSWORD=minio_password
   minio server --console-address :9090 ~/minio/data
   ```

### Docker

1. Create a new directory.

   ```shell
   mkdir -p ~/minio/data
   ```

1. Pull and run minio.

   ```shell
   docker run \
   -p 9000:9000 \
   -p 9090:9090 \
   --name minio \
   -v ~/minio/data:/data \
   -e MINIO_ROOT_USER=minio \
   -e MINIO_ROOT_PASSWORD=minio_password \
   quay.io/minio/minio minio server /data --console-address :9090
   ```
   
1. Alternatively, run the docker-compose file.

   ```shell
   docker-compose -f docker/docker-compose.yaml up
   ```

### Kubernetes

