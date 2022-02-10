# meminero - Scraper for EVM-based chains

## Features

- works with any web3 compatible node (including providers like Infura, Chainstack etc)
- reorgs are handled automatically
- supports multiple chains
- includes an integrity checker that makes sure the data is consistent

## Configuration

For all the available config options please see the [configuration file](./config-generated.yml). A sample configuration
file can be found [here](./config-sample.yml).

The tool also includes rich help functions. To access them, use the `--help` flag.

```shell
./meminero --help
```

## Running

The scraper can run in two different modes:

### 1. queue

In this mode, the scraper will wait for tasks in the redis queue and process them. To enable live chain tracking, see
the `queuekeeper` options in the config file.

```shell
./meminero scrape queue
```

### 2. single mode - the scraper will run for a single block and exit

```shell
./meminero scrape single --block=<block-number>
```

## Sync files
In order to extract relevant information for each product, some more data is required. 
For example, the scraper must know all the SMART Alpha pools it should monitor. 

This information is stored in sync files, grouped by network in the [sync-files folder](./sync-files).

**These files must be treated with extreme caution because they are also served to the frontend and can result in loss of funds if not handled properly.** 

> **Note:** The sync files allow for insert and update of information but not deletion. If you want to delete information, you must do it manually.

### Syncing the information
This requires a proper configuration so the scraper can connect to the database where the information should be stored.
```shell
./meminero sync --syncer.network=<network> --syncer.datasets=<dataset1,dataset2...>
```

## Database schema

The database schema is built using the migrations in this repo.

Each product [governance, smart_alpha, smart_exposure, smart_yield, yield_farming] have their own set of migrations that
are run in a different schema.

Please see the [migrations](./db/migrations) for more information.

## Data extraction

Each product is grouped into a set of Storables. See storable interface [here](./types/Storable.go)

All Storables can be found in the [storables](./processor/storables) folder.

Each storable is responsible for extracting data from the blockchain and storing it in the database. Everything else
from scraping to decoding the information and putting it into known data types is handled by the other parts of the
architecture and should be considered a given in the regular development process.

For example, all the events from SMART Alpha are captured using [this storable](./processor/storables/smartalpha/events)
.

