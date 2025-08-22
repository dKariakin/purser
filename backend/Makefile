.install-lint:
	@echo "Checking if golangci-lint is installed..."
	@which golangci-lint > /dev/null 2>&1 || (curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh)

lint: .install-lint
	@echo "Cleaning unnecessary dependencies..."
	@go mod tidy
	@echo "Formatting code..."
	@golangci-lint fmt
	@echo "Running linter check..."
	@golangci-lint run