while true
do
	echo "Collecting hoppers ..."
        MONGO_URI="mongodb://localhost:27017" go run main.go
	date
	echo "DONE!"
        sleep 90 
done
