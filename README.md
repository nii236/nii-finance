# Algotrading Platform
This document covers the plan for the collaborative algorithmic trading platform by the members of open-algot, a Slack group from Reddit's `/r/algotrading` community.

It will cover code, thoughts, services, parameters and ideas.

Once we have a more solid idea of what we are building, a block diagram will be drawn using https://www.draw.io/.

# Quickstart

Install [Docker](https://www.docker.com/). Then run:

`docker-compose up`

A simple microservice will start up. It contains:

- A greeter microservice
- A mock tick data generator
- A tick data subscriber
- Consul for registry

HTTP is still used for communication between services, but that will change when I get more time.

Practice writing a client that connects to the sidecar and sends a request to the greeter microservice. Example code for python and ruby can be found [here](https://github.com/micro/micro/tree/master/examples/greeter/client). This is an example of a synchroous call. For asynchronous try to hook into the "go.micro.srv.TickRecorder" topic, through the sidecar.

When you are done run:

`docker-compose down`

To close down the containers.

# Collaboration
The git repo is located [here](https://open-algot.servebeer.com/).

Pushing directly to `master` and `develop` is discouraged. Use the [git flow](http://danielkummer.github.io/git-flow-cheatsheet/) methodology for merging features.

I (nii236) will probably be pushing directly to `develop` in the early stages while I try to get a basic workflow happening for other developers to start working on their on features.

# Common Technologies
We all use different stacks, technologies and languages. We come from different industries. But to work together, we need to have common ground. I (nii236) propose the following common skills and technologies:

- Git
- Micro
- [NATS](nats.io)
- [Docker](https://www.docker.com/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers/)

With all of the above, our individual microservices will be able to communicate with each other without incident or cognitive overload.

# Architecture

The proposed microservice stack will use the following:

- Broker - Asynchronous messaging (pub/sub)
- Transport - Synchronous messaging (RPC)
- Registry - Service Discovery

Each service will be contained in a Docker container, and nats.io will be the message queue. The choice of technology will not be visible to the microservice.

## Sidecar
The sidecar is a concept that allows different languages to connect to the microservices swarm. It is basically an HTTP proxy which have a REST endpoint for RPC, and a websocket endpoint for publish/subscribe messaging.

More information can be found [here](https://blog.micro.mu/2016/03/20/micro.html).

Example code for python and ruby can be found [here](https://github.com/micro/micro/tree/master/examples/greeter).

At this point in time, only a Go library is available for direct connection to the microservice swarm.

# Microservices

## DataDownloader
`DataDownloader` will be a base class that will download data, people can implement it to use whatever data storage system they'd like.

## Backtester
`Backtester` is a system that feeds data to an algorithm and sends trades to a bookkeeper service (tradingBook)


## Optimiser
`Optimiser` will launch many instances of the backtester instance, and through genetic or bruteforce optimisation, it will find the best parameters of the algo for a certain time period.

## Web UI
`WebUI` will be the web frontend for this platform. It will provide live algo monitoring tools such as backtesting performance charts and an area to develop your strategy.


## TradingBook
`TradingBook` keeps track of trades that are passed to it via a messaging bus and these trades can be routed to a broker from there. Also records these trades for analysis later (equity curves, metrics) as a data dump (single file?).

## GraphAnalyser
`GraphAnalyser` is a stand alone tool that takes data dumps from other modules like the backtester which contain information about trades performed, and makes equity curves, and calculates other metrics (Alpha, Sharpe Ratio, etc.)

## TickProvider
`TickProvider` will fetch the tick data from the database and return it to the requester. It handles both ticker and OHLC requests.

```
syntax = "proto3";

package go.micro.srv.tickProvider;

service RequestData {
	rpc Tick(Request) returns (TickResponse) {}
	rpc OHLC(Request) returns (OHLCResponse) {}
}

message Request {
	string pair = 1;
	int64 from = 2;
	int64 to = 3;
	string resolution = 4;
}

message TickResponse {
	string pair = 1;
	string broker = 2;
	int64 time = 3;
	int32 bid = 4;
	int32 ask = 5;
	int32 last = 6;
}

message OHLCResponse {
	string pair = 1;
	string broker = 2;
	int64 time = 3;
	int32 open = 4;
	int32 high = 5;
	int32 low = 6;
	int32 close = 7;
}
```

## TickSubscriber
`TickSubscriber` subscribes to a broker's API, and on each tick will send the data to the database for recording. This service does not serve any clients. It is an automatic process which will publish each tick onto the message queue for a database to pick up.

There will need to be many different types of `TickSubscriber`, one for each type of forex broker.

```
syntax = "proto3";

package go.micro.srv.tickSubscriber;

message Tick {
	string pair = 1;
	int64 time = 2;
	int32 bid = 3;
	int32 ask = 4;
	int32 last = 5;
}
```

## Greeter
`Greeter` is a simple example of how to build a microservice. It will also form the basis on how to document a microservice in this document.

It receives a request and replies with `Hello` with the request string concatenated. It will show how to build and use the:

- Server
- Client
- Web API
- Serialization

Every microservice has a `.proto` file. The `.proto` file describes the accessible parts of the microservice exactly. It will be all the information you need to communicate with it.

`syntax` describes the version of protocol buffers we are using. In this case it will be version 3.

`package` describes the address to publish to if you want to send an asynchronous message.

`service` describes synchronous methods that can be called on this microservice.

`message` describes arguments that can go in and out of this microservice.
```
syntax = "proto3";

package go.micro.srv.greeter;

service Say {
	rpc Hello(Request) returns (Response) {}
}

message Request {
	optional string name = 1;
}

message Response {
	optional string msg = 1;
}
```

The above definition tells me that you can call `Say.Hello(Request)` to this microservice and I will receive a synchronous response in the form of `Response`.

## General interface to Broker APIs

## System Parameters
- Time period
- Resolution
- Symbols (the symbols that we will maintain data on.)
- Market(FX - Equities - Derivatives - Futures)
- Broker
