WORKDIR=$(pwd)
cd ..
GOOS=linux GOARCH=amd64 go build -o bin/hopper-analytics-collector.linux main.go
cd $WORKDIR