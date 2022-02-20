# run

```shell
# start server
go run server/server.go
# deploy task
go run client/main.go --taskType 0 --count 1000
# cancel task
go run client/main.go --inspect --op c --queue bundle --taskId 9ad5cc5b-cf5b-4c65-b586-013467b3f1f8
```