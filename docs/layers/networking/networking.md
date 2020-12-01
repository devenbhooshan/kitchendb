## networking

- a goroutine(nicknamed mg) accepts incoming connection request on default port 1206
- when a client tries to make a connection, mg spawns another go-routine. All the future requests from the client is managed by that go-routine
- goroutine dies once the connection is closed by the client