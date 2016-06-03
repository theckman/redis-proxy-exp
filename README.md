# redis-proxy-exp
This is an experimental piece of code used for testing. This takes a key,
and returns the value of that key from the Redis instance. I'm only open
sourcing this so it's not lost, or in hopes it may provide some sort of
value to someone (however unlikely).

I needed code to test that two systems were talking properly, so I wrote
this. While it's released under the
[MIT License](https://opensource.org/licenses/MIT), it probably should not
be used in production. Or anywhere...

This works by listening on port `1620` on localhost, and responding to requests
like this:

```
curl 'http://127.0.0.1:1620/get?key=testKey'
```

## Installation
```
go get github.com/theckman/redis-proxy-exp
```

## Usage
```
Usage:
  redis-proxy-exp [OPTIONS]

Application Options:
  -p, --port=       port to bind to for HTTP interface (default: 1620)
      --redis-host= IP or hostname or Redis instance (required)
      --redis-port= port number where Redis is listening (default: 6379)
  -V, --version     print versions string and exit

Help Options:
  -h, --help        Show this help message
```
