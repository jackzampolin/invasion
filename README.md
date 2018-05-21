[![Documentation](https://godoc.org/github.com/jackzampolin/invasion?status.svg)](http://godoc.org/github.com/jackzampolin/invasion)

# Invasion

Invasion is a game that simulates the invasion of a planet specified by a file in the following format:

```
cityName1 north=cityName2 west=cityName3
cityName2 east=cityName1
...
```

Aliens are generated and placed on this "map". Every turn the aliens choose a road from the city they are currently in and move to a new city. If there are 2 or more aliens in a city at the end of a turn the city, the aliens and all the roads to that city are destroyed.

### Building/Running Invasion

To build `invasion` have a [working Golang environment](https://golang.org/doc/install) and the [`dep` package manager](https://github.com/golang/dep) installed. Just run the following:

```
$ dep ensure
$ go install main.go
```

Then you will be able to run `invasion`:

```
$ invasion
A game where aliens destroy a map while you watch!

Usage:
  invasion [command]

Available Commands:
  help        Help about any command
  newMap      Generates a new map file
  play        A brief description of your command

Flags:
  -h, --help   help for invasion

Use "invasion [command] --help" for more information about a command.
```

### Generating Maps

`invasion` contains a nifty map generator that will, by default create a world with 10 cities with randomly generated names of 16 chars each. You can also optionally pass the `--numCities` arguement to specify a larger or smaller world:

```
$ invasion newMap myMap --numCities 10
Created map file myMap.map in current directory...
```

### Playing invasion

Once you have a map to play on its as easy as running `play` and passing in the path to the map and the desired number of aliens. These aliens will have randomly generated names with 8 chars each. The game will run for 10,000 turns or until all the aliens have destroyed themselves:

```
$ invasion play myMap.map 5
iOGawnniiGXNCRJw has been destroyed by alien XJdNLntv and alien LNuTOMhb! (turn 1)
eTJOWfTUTXnOkIvI has been destroyed by alien SkUvsfzN and alien mObSJrBA! (turn 1)

ENDING MAP (8):

YhwtyKxYDXfXJwgi
nJzhgwpOuwgXTlya west=wQbObcFCTIxuPkrd north=LVJDgIicKkhcUhBG
EFBAbAQLJnQgSJlN west=YhwtyKxYDXfXJwgi
SSODSIhrFwiWnjAJ west=YhwtyKxYDXfXJwgi north=SSODSIhrFwiWnjAJ east=nJzhgwpOuwgXTlya
LVJDgIicKkhcUhBG west=GjGXJyZfOEekfcxo south=wQbObcFCTIxuPkrd
hjmlyYDgrqUGkMiL west=hjmlyYDgrqUGkMiL north=LVJDgIicKkhcUhBG
wQbObcFCTIxuPkrd south=wQbObcFCTIxuPkrd east=wQbObcFCTIxuPkrd west=hjmlyYDgrqUGkMiL
GjGXJyZfOEekfcxo north=hjmlyYDgrqUGkMiL west=hjmlyYDgrqUGkMiL

ENDING ALIENS:

UVwTfWpG city=wQbObcFCTIxuPkrd turns=10000
```
