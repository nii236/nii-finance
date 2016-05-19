# Spinup Instructions
## Start the Message Queue

Download the latest nats server from [here](https://github.com/nats-io/gnatsd/releases/latest). Extract the binary into your PATH.

Run the nats server.
```bash
gnatsd
```
## Compiling the Greeter Microservice

### Standard
Compile the project. Make sure you have [Go 1.6](golang.org/) installed.
```bash
cd services/greeter
go build
```


### Docker
Or compile it with Docker.

Set GOOS and GOARCH to your environment of choice. GOOS can be `windows`, `linux` or `osx`. GOARCH can be `386` or `amd64`.
```bash
cd services/greeter
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp  GOOS=linux -e GOARCH=386 golang:1.6 go build -v
```

## Running the Microservice
Run the microservice.
```bash
./greeter \
--broker nats \
--broker_address localhost:4222 \
--transport nats \
--transport_address localhost:4222 \
--registry nats \
--registry_address localhost:4222
```

Run the client.
```bash
./greeter \
--broker nats \
--broker_address localhost:4222 \
--transport nats \
--transport_address localhost:4222 \
--registry nats \
--registry_address localhost:4222 \
--client
```

This outputs `Hello ${NAME}`.
