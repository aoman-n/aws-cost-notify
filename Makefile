.PHONY: build dev

build:
	sam build

deploy:
	sam build
	sam deploy --guided --env-vars env.json

start:
	sam local invoke --env-vars env.json