# Bamboo Blinker

A build monitor light for Atlassian Bamboo using a [BlyncLight](http://www.embrava.com/) device. 
The [bamboo-build-bunny plugin](https://bitbucket.org/eddiewebb/bamboo-build-bunny/wiki/Home) is 
used to fetch build status(es) from Bamboo.

## Download

Pre-built binaries are available [on the release pages](https://github.com/davidehringer/bamboo-blinker/releases).

## Running

```
./bamboo-blinker https://your-bamboo.com/rest/build-bunny/1.0/summary/120edef1-a493-4813-82db-edb6b078cf6b.json
```
Or to check health at a different interval (60 seconds in this case):

```
./bamboo-blinker https://your-bamboo.com/rest/build-bunny/1.0/summary/120edef1-a493-4813-82db-edb6b078cf6b.json 60
```

### Backoff Logic
To prevent dozens or hundreds of clients from beating up bamboo by checking too frequently, the application will double the delay interval up to 10 minutes when the returned timeToEvaluate > 150ms.  The interval will be reset to argument #2 (default 10s) as soon as the first response is within limits.

To override the acceptable limit for this value, pass a 3rd argument.
```
./bamboo-blinker https://your-bamboo.com/rest/build-bunny/1.0/summary/120edef1-a493-4813-82db-edb6b078cf6b.json 60 200
```

## Building

This project depends on [goblync](https://github.com/davidehringer/goblync).

```
go get github.com/boombuler/hid
go get github.com/davidehringer/goblync
go build
```

See notes for https://github.com/boombuler/hid if you have issues building for different platforms.