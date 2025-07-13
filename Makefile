# Variables
BINARY_NAME=bootstrap
ZIP_FILE=x-bot-lambda.zip
LAMBDA_FUNCTION_NAME=x-bot
AWS_PROFILE=x-bot

# Build for Lambda (Linux AMD64)
.PHONY: build
build:
	@echo "Building X Bot for AWS Lambda..."
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o $(BINARY_NAME) .
	@echo "Build completed: $(BINARY_NAME)"

# Create deployment package
.PHONY: package
package: build
	@echo "Creating deployment package..."
	build-lambda-zip.exe -o $(ZIP_FILE) $(BINARY_NAME)
	@echo "Package created: $(ZIP_FILE)"

# Deploy to AWS Lambda
.PHONY: deploy
deploy: package
	@echo "Deploying to AWS Lambda..."
	aws lambda update-function-code \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--zip-file fileb://$(ZIP_FILE) \
		--profile $(AWS_PROFILE)
	@echo "Deployment completed"

# Test Lambda function
.PHONY: test-lambda
test-lambda:
	@echo "Testing Lambda function..."
	@aws lambda invoke \
		--function-name $(LAMBDA_FUNCTION_NAME) \
		--payload '{}' \
		--profile $(AWS_PROFILE) \
        --output text \
        response.json
	@echo "Test completed"

# Show Lambda logs
.PHONY: logs
logs:
	@echo "Showing Lambda logs..."
	aws logs tail /aws/lambda/$(LAMBDA_FUNCTION_NAME) --follow

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME) $(ZIP_FILE)
	@echo "Clean completed"

# Show help
.PHONY: help
help:
	@echo "X Bot Lambda Commands:"
	@echo ""
	@echo "  build        - Build for AWS Lambda"
	@echo "  package      - Create deployment ZIP package"
	@echo "  deploy       - Deploy to AWS Lambda"
	@echo "  test-lambda  - Test Lambda function"
	@echo "  logs         - Show Lambda logs"
	@echo "  clean        - Clean build artifacts"
	@echo "  help         - Show this help message"
