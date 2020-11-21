# atomicity implementation details

## definition

during a batch query(multiple writes/reads in a single transaction), if one of the query fails due to some reason, the whole batch should abbort i.e. either
the whole batch query should execute or none. 
lets understand this with an example. Lets assume there are two friends `Foo` and `Bar`. `Foo` wants to transfer 100$ to `Bar`. So we want to write a query which
should deduct 100$ from `Foo`'s account and deposit 100$ to `Bar`'s account. This is how the query would look like.
```
BEGIN TRANSACTION
query-1 -> subtract 100$ from Foo account;
query-2 -> add 100$ to Bar account;
END TRANSACTION
```
Lets say, our kitchendb was able to execute query-1 but failed to execute query-2. In that case, the db should revert the changes done by the query-1. 

## implementation details



## notes
- https://brandur.org/postgres-atomicity
