## Components

- networking
- sql layer
  - parser
  - query planner
  - executor
- transactional layer
- storage engine
- misc
  - telemetry
  - config management

## networking

- a goroutine(nicknamed mg) accepts incoming connection request on default port 1206
- when a client tries to make a connection, mg spawns another go-routine. All the future requests from the client is managed by that go-routine
- goroutine dies once the connection is closed by the client

## Database transaction guarantees

 - [atomicity i.e abortability](./atomicity.md)
 - snapshot isolation(implementaion same as postgresql) [cockroachdb details](https://www.cockroachlabs.com/blog/serializable-lockless-distributed-isolation-cockroachdb/)
 - durability*

# links

- [cockroachdb design doc](https://github.com/cockroachdb/cockroach/blob/master/docs/design.md)