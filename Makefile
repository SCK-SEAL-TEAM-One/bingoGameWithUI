all: analysis unit-test build-artifact accept push-artifact
	
analysis:
	cd src && golangci-lint run

unit-test:
	cd src && go test ./... --cover

build-artifact:
	# build app
	# docker image build -t golfapipol/bingo:1 .
	docker-compose build
accept:
	# run app && run test
	docker-compose up -d
	# docker container run -d --name my-bingo -p 3000:3000 golfapipol/bingo:1
	robot atdd/ui/verticalBingo.robot
	docker-compose down

down:
	docker-compose down

push-artifact:
	docker push golfapipol/bingo:1