# storage engine

the responsibility of storage engine is to manage how data is stored both in disk or memory. There are multiple storage engines available in the market(some are b-tree based and some are lsm tree based) eg:  levelsdb, rocksdb, hyperleveldb, lmdb etc. mongodb uses wiredTiger storage engine. mongodb also supports in-memory storage engines. some of the databases support pluggable storage engine where one can pick the best engine depending on the workloads. 

### [pebble](https://github.com/cockroachdb/pebble) as the first storage engine, why
- it is inspired by rocksdb
- it is written in go(so no interop overheads)
- it is being used in production in cockroach DB

### pluggable storage engine
- we want kitchendb to be extensible in terms of the type of storage engine it supports
- Add implementation details here

### links

- https://www.influxdata.com/blog/benchmarking-leveldb-vs-rocksdb-vs-hyperleveldb-vs-lmdb-performance-for-influxdb/#:~:text=LevelDB%20is%20the%20winner%20on,metrics%20except%20for%20disk%20size.
- https://github.com/EnterpriseDB/zheap
- https://wiki.postgresql.org/wiki/Future_of_storage
- https://israelg99.github.io/2017-06-07-Introduction-to-Database-Storage-Engines/
- https://blog.yugabyte.com/a-busy-developers-guide-to-database-storage-engines-the-basics/
- http://www.benstopford.com/2015/02/14/log-structured-merge-trees/

