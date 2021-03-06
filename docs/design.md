## goals

- support postgresql wire protocol
- support for both OLTP and OLAP(i.e HTAP) workloads
- sql compatible query language
- pluggable storage engine
- acid transactions
- replicate data across multiple nodes to provide HA


## Database Components

- [storage engine](./layers/storage/storage.md)
- [networking](./layers/networking/networking.md)
- [sql layer](./layers/sql-layer/sql.md)
- [transactional layer](./layers/transactional/transaction.md)
- [replication layer](./layers/replication/replication.md)
- misc
  - telemetry
  - config management
  - [benchmarks](https://github.com/brianfrankcooper/YCSB)


## links

- [cockroachdb design doc](https://github.com/cockroachdb/cockroach/blob/master/docs/design.md)
- [DB_Engines](https://db-engines.com/en/)
- https://github.com/alex-dukhno/database
- https://github.com/benstopford/awesome-db-benchmarks


## distributed sql
- https://github.com/citusdata/citus
