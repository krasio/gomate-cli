# Gomate CLI
CLI for [Gomate](https://github.com/krasio/gomate).

# Usage
```
$ go build

$ ./script/start-redis.sh

$ GOMATE_REDIS_URL=redis://localhost:9999/0 ./gomate-cli load suburb < data/suburbs.json 
Connecting to Redis using redis://localhost:9999/0.
Loaded a total of 3 items.

$ GOMATE_REDIS_URL=redis://localhost:9999/0 ./gomate-cli query suburb well
Connecting to Redis using redis://localhost:9999/0.
Query suburb for "well":

  Wellington
  Welly

$ ./script/stop-redis.sh
```
# License
MIT licence.
