include .env

clean:
		@rm -rf dist
		@mkdir -p dist

build: clean
		@for dir in `ls func`; do \
			GOOS=linux go build -o dist/func/$$dir github.com/sampado/aws-lambda-go-api/func/$$dir; \
		done

run:
		aws-sam-local local start-api

test:
		go test ./... --cover

package: build
		@aws cloudformation package \
			--template-file template.yml \
			--region $(AWS_REGION) \
			--output-template-file package.yml

deploy:
		@aws cloudformation deploy \
			--template-file package.yml \
			--region $(AWS_REGION) \
			--capabilities CAPABILITY_IAM \
			--stack-name $(AWS_STACK_NAME)

describe:
		@aws cloudformation describe-stacks \
			--region $(AWS_REGION) \
			--stack-name $(AWS_STACK_NAME) \

outputs:
		@make describe | jq -r '.Stacks[0].Outputs'

url:
		@make describe | jq -r ".Stacks[0].Outputs[0].OutputValue" -j