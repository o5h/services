#!/bin/bash

go build -ldflags "-X main.version="$(<VERSION)" -X main.date=`date "+%Y%m%d-%H%M%S"`"  ./cmd/service_example/.
