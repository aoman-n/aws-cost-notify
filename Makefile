.PHONY: build deploy start

build:
	sam build

deploy:
	sam build
	sam deploy --guided \
		--profile lambda-user

start:
	sam local invoke --env-vars env.json