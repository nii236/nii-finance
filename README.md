# Nii Finance Trading Swarm
## Installation

### Dependencies
- [Go 1.6](https://golang.org/)
- [Glide](https://glide.sh/)
- [NATS](http://nats.io/)
- [Docker](https://www.docker.com/)

### Spinup
```
$ go get -d github.com/nii236/nii-finance
$ cd $GOPATH/src/github.com/nii236/nii-finance
$ glide install
$ go build
$ docker-compose up
```

## Usage

In its current state, the swarm will:
- Subscribe to the IB Gateway
- Pull out USDJPY ticker data
- Publish onto the NATS queue
- Be received by the tickRecorder service
- Write the data point into InfluxDB

## Contributing

1. Fork it!
2. Create your feature branch using `git flow feature start my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request

### Ideas

- A service that writes into MongoDB instead of InfluxDB
- A service that calculates indicators based on drip feed data on NATS
- A service that opens and closes trades for you
- A service that pulls historical data from InfluxDB in a defined time resolution
- Using go-micro's sidecar to allow writing algorithms in any language

## TODO
- Use ENV vars for everything instead of hard coding strings
- Make services more durable (when dependent services such as IB Gateway go down)
