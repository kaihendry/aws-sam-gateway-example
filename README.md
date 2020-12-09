Going back to first principals https://dabase.com/blog/2020/sam-sucks/



Source: git@github.com:techjacker/go-serverless-api.git

https://andrewgriffithsonline.com/blog/180412-migrate-go-api-to-serverless-under-10-mins/

Idea is to find an Apex Up solution to prove that we have a way out if Apex Up siezes to exist.

Nonetheless whilst Apex Up is around, I suggest you use it!


Deployed: https://ap-southeast-1.console.aws.amazon.com/lambda/home?region=ap-southeast-1#/applications/aws-sam-gateway-example

# Local development

	cd local-go-serverless-api
	gin

https://github.com/codegangsta/gin

# Alternative proxy approach

https://github.com/awslabs/aws-lambda-go-api-proxy shows how to proxy the
handlers, though I find https://github.com/apex/gateway drop in replacement for
**ListenAndServe** far easier to work with.
