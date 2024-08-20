APP_NAME=github-profile


clean:
	rm -rf ./bin/*


build:clean
	@go build -o ./bin/$(APP_NAME) 

run:build
	@./bin/$(APP_NAME)


PHONY: clean build run