# ![kitchendb](https://github.com/devenbhooshan/kitchendb/workflows/kitchendb/badge.svg?branch=master)
![alt text](logo.png "logo")

## design goals

- support postgresql wire protocol
- support for both OLTP and OLAP(i.e HTAP) workloads
- sql compatible query language
- pluggable storage engine
- acid transactions

Check out [design doc](./docs/design.md) for more deatils

### why another database?

- to understand and implement internals of a dbms 

### Todos

- [x] write first draft of design doc
- [x] setup ci
- [x] import [vitess-sqlparser](https://github.com/blastrain/vitess-sqlparser)
- [ ] storage engine
- [ ] telemetry framework
- [ ] finish the design doc of transactional layer
- [ ] setup [yahoo cloud service benchmark](https://github.com/brianfrankcooper/YCSB)
