WORKDIR=$(pwd)
cd ..
go build -o bin/hopper-analytics-collector.macos main.go
cd $WORKDIR