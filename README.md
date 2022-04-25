# hopper-shopper-collector

## Requirements

-   Go v1.18 (required for of generics support)
-   MongoDB instance (Docker, local, cloud etc)

## Usage

The project is structured around individual CLI commands. It's also required to set the environment variable `MONGO_URI` which has to be a valid MongoDB URI string (like `mongodb://localhost:27017`).

This repo does not include any tools to load data from MongoDB. It's only purpose is to constantly collect data and save it into a local database to provide fast data access. On-chain contract reads can take quite some time (multiple seconds) and sometimes fails. It also prevents DDoS'ing the public Avalanche RPC through this tool.

**Base Usage**

> `go run main.go --help`

Which prints the available commands and general help:

```
Usage:
   [command]

Available Commands:
  base-shares    Load and save curent base shares for adventures
  completion     Generate the autocompletion script for the specified shell
  fly-supply     Load and save current FLY supply
  help           Help about any command
  hoppers        Load and save a snapshot of all hoppers
  markets        Load and save a snapshot of all market listings
  migrate-supply Migrate legacy supply to new schema
  prices         Load and save current crypto prices
  votes          Load and save curent votes / veShare for adventures

Flags:
  -h, --help   help for this command

Use " [command] --help" for more information about a command.
```

### Commands

**base-shares**  
`go run main.go base-shares`: Load current base-shares from on-chain contracts for all adventures and save it into MongoDB

**fly-supply**  
`go run main.go fly-supply`: Load current total FLY supply from the on-chain FLY contract and save it into MongoDB

**hoppers**  
`go run main.go hoppers`: Load all available hoppers using the official Hoppersgame Graph. Once loaded calculates a bunch of useful indicators like individual adventure ratings and base FLY / level and saves it into MongoDB.

> Note: This command will flush the `hoppers` collection each execution so the database won't grow too large to fast

**markets**  
`go run main.go markets`: Load all market hopper related listings with price, status and date

> Note: This command will flush the `listings` collection each execution so the database won't grow too large to fast

**prices**  
`go run main.go prices`: Load the current AVAX (USD + EUR) and FLY (USD + EUR) prices and save it into MongoDB

**votes**  
`go run main.go votes`: Load current votes / veShare from on-chain contracts for all adventures and save it into MongoDB
