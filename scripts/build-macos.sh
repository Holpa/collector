WORKDIR=$(pwd)
cd ..
go build -o bin/hopper-shopper-collector.macos main.go
cd $WORKDIR