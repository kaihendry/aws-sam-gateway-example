STACK = aws-sam-gateway-example

.PHONY: build deploy validate destroy

deploy: build
	AWS_PROFILE=mine sam deploy --stack-name $(STACK)

build:
	CGO_ENABLED=0 sam build

validate:
	aws cloudformation validate-template --template-body file://template.yaml

destroy:
	AWS_PROFILE=mine aws cloudformation delete-stack --stack-name $(STACK)
