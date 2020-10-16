# grpc-echo-service-benchmark

## Prerequisites
- [ghz](https://ghz.sh/docs/install)

## How to run
1. Build the docker images and update the docker-compose file accordingly. 
2. Run `docker-compose up` to run the gRPC service loacally.
3. run bellow commannd to send gRPC requests.
```
ghz --insecure --proto ./grpc_unary_blocking.proto --call HelloWorld.hello -c 100 -z 15m --duration-stop=wait --connections=100 -d '{"body":"Joe"}' 0.0.0.0:9090
```
