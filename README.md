# Get Slowest Query DB

This Simple Api Write With Fiber Freamwork and that will return the slowest connected database queries

## Run

If you run this API first you need to import slowest_query_db.sql to postgres server and then set server option in env folder

please set PivotTimetoCheckSlowesQuery in env folder base on milisecond

```bash
go run main.go
```

## This Simple Api

Support pagination For Get List Slowest Querys

Support filtering by SELECT,INSERT,UPDATE,DELETE on Get List Slowest Querys

Support sorting by time spent Slowest Querys

## License

[MIT](https://choosealicense.com/licenses/mit/)
