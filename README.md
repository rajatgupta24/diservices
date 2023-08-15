# diservices: distributed-services

## Step1: 
Built a simple JSON/HTTP commit log service that accepts and responds with JSON and stores the records in those requests to an in-memory log.

Run to start the server: 

```bash
go run main.go
```

Open another tab in your terminal and run the following commands to add some records to your log:

```bash
curl -X POST localhost:8080 -d '{"record": {"value": "Test"}}'
```

Read the records back by running the following commands:

```bash
curl -X GET localhost:8080 -d '{"offset": 0}
```
