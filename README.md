# Algotrading Platform


This document covers the plan for the collaborative algorithmic trading platform by the members of open-algot, a Slack group from Reddit's `/r/algotrading` community.

It will cover code, thoughts, services, parameters and ideas.

Once we have a more solid idea of what we are building, a block diagram will be drawn using https://www.draw.io/.

# Services

## DataDownloader
This will be a base class that will download data, people can implement it to use whatever data storage system they'd like.

## Backtester
A system that feeds data to an algorithm and sends trades to a bookkeeper service (tradingBook)


## Optimiser
This microservice will launch many instances of the backtester instance, and through genetic or bruteforce optimisation, it will find the best parameters of the algo for a certain time period. Also generates a data dump that is used by graphAnalyser

## Live Algo Monitoring Tools (Web UI)



## TradingBook
A microservice that keeps track of trades that are passed to it via a messaging bus and these trades can be routed to a broker from there. Also records these trades for analysis later (equity curves, metrics)

## graphAnalyser
This module is a stand alone tool that takes data dumps from other modules like the backtester which contain information about trades performed, and makes equity curves, and calculates other metrics (Alpha, Sharpe Ratio, etc.)


## TickProvider
`TickProvider` will fetch the tick data from the database and return it to the requester.


## TickSubscribers
`TickSubscriber` subscribes to a broker's API, and on each tick will send the data to the database for recording.

### Params
- TimePeriod
- Resolution
- Symbol


## General interface to Broker APIs 

### Params (general)
- Time period
- Resolution
- Symbols
- Market(FX - Equities - Derivatives - Futures)
- Broker


- implementations to Broker APIs

## Repo
https://open-algot.servebeer.com/


