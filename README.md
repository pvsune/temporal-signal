### Steps to run this sample:
1) You need a Temporal service running. See details in [README.md](https://github.com/temporalio/samples-go)
2) Run the following command to start the worker
```
go run worker/main.go
```
3) Run the following command to start the example
```
go run starter/main.go
```
4) Run the following command to change the message
```
go run ./signal/main.go -s '{"message": "Philamer"}'
```
5) Run the following command to cancel the workflow
```
go run ./signal/main.go -s '{"type": "cancel"}'
```
6) Run the following command to extend the workflow
```
go run ./signal/main.go -s '{"type": "extend", "duration": 5}'
```