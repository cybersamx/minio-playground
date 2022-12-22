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
   minio server --console-address :9090 ~/data/minio
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
   -v ~/data/minio:/data/minio \
   -e MINIO_ROOT_USER=minio \
   -e MINIO_ROOT_PASSWORD=minio_password \
   quay.io/minio/minio minio server /data/minio --console-address :9090
   ```
   
1. Alternatively, run the docker-compose file.

   ```shell
   docker-compose -f docker/docker-compose.yaml up
   ```

### Kubernetes (Single Node - HostPath)

This setup is a single node k8s cluster, usually a local setup.

1. Run the following manifest file to create the needed k8s storage objects.

   ```shell
   cd k8s/single-hostpath
   kubectl apply -f pod.yaml
   kc get pod -n minio
   NAME    READY   STATUS    RESTARTS   AGE
   minio   1/1     Running   0          21s
   ```

1. Run the following manifest file to expose the pod to the network.

   ```shell
   kubectl apply -f service.yaml
   kubectl get svc -n minio
   NAME            TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)          AGE
   minio-api       LoadBalancer   10.43.184.223   192.168.1.67   9000:30241/TCP   7s
   minio-console   LoadBalancer   10.43.160.222   192.168.1.67   9090:31362/TCP   7s
   ```

   Note the external ip address and navigate your browser to `http://<external-ip>:9090`.

   **Alternatively**, run `kubectl port-forward pod/minio 9000 9090 -n minio` and navigate to <http://localhost:9090/>.

### Kubernetes (Single Node - PVC)

This setup is a single node k8s cluster, usually a local setup.

1. Create directory in k8s. Since I am running Rancher Desktop, we can use `rdctl`, the equivalent of ssh to the k8s node.

   ```shell
   rdctl shell
   # Once we are in the k8s node, create the mount point for the persistent volume.
   sudo mkdir /mnt/disk1/data
   exit
   ```

1. Run the following manifest file to create the needed k8s storage objects.

   ```shell
   cd k8s/single-pvc
   kubectl apply -f storage.yaml
   kubectl get pv
   NAME       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS    REASON   AGE
   local-pv   2Gi        RWO            Retain           Available           local-storage            59s
   kubectl get pvc -n minio
   NAME        STATUS    VOLUME   CAPACITY   ACCESS MODES   STORAGECLASS    AGE
   minio-pvc   Pending                                      local-storage   51s
   # The pvc should be pending since it hasn't been claimed.
   ```

1. Run the following manifest file to run minio pod.

   ```shell
   kubectl apply -f pod.yaml
   kubectl get pod -n minio
   NAME    READY   STATUS    RESTARTS   AGE
   minio   1/1     Running   0          6m55s
   ```

1. Run the following manifest file to expose the pod to the network.

   ```shell
   kubectl apply -f service.yaml
   kubectl get svc -n minio
   NAME            TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)          AGE
   minio-api       LoadBalancer   10.43.184.223   192.168.1.67   9000:30241/TCP   7s
   minio-console   LoadBalancer   10.43.160.222   192.168.1.67   9090:31362/TCP   7s
   ```
   
   Note the external ip address and navigate your browser to `http://<external-ip>:9090`.

   **Alternatively**, run `kubectl port-forward pod/minio 9000 9090 -n minio` and navigate to <http://localhost:9090/>.

## References

[Minio: Quickstart for Kubernetes](https://min.io/docs/minio/kubernetes/upstream/index.html)
