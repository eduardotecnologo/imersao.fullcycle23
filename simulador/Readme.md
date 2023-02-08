docker-compose up -d 
docker-compose ps
docker exec -it simulator bash 

go mod init github.com/eduardotecnologo/imersao.fullcycle23

kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
kafka-console-producer --bootstrap-server=localhost:9092 --topic=readtest
kafka-console-producer --bootstrap-server=localhost:9092 --topic=route.new-direction
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=route.new-position --group=terminal
