# Services Example

## Build

### Build script
```sh
go build -ldflags "-X main.version="$(<VERSION)" -X main.date=`date "+%Y%m%d-%H%M%S"`"  ./cmd/service_example/.
```

## Database 

### Create and start database
```sh
docker run -P --publish 127.0.0.1:5432:5432 --name service-example-db -e POSTGRES_USER=service-user -e POSTGRES_PASSWORD=service-password -e POSTGRES_DB=service-example -d library/postgres
```

### Start
```sh
docker start service-example-db
```

### Stop
```sh
docker stop service-example-db
```

### Remove database
```sh
docker rm -f stop service-example-db
```