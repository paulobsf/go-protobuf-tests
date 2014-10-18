go-protobuf-tests
=================

Testing google protocol buffers (using Go)

Reference: https://code.google.com/p/goprotobuf/


## Installing protobuf compiler for go (Ubuntu)

```
# apt-get install golang-goprotobuf-dev
```

## Generating .go from .proto files

```
$ cd dir/with/proto/files
$ protoc --go_out=. *.proto
```