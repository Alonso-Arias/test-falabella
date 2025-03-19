# Archivo Makefile
TOOLS_DIR := tools

# Variables de entorno
ENV_VARS = MYSQL_CONNECTION="root:123456@tcp(localhost:3306)/TEST?parseTime=true" \
	       API_KEY="9efa9f3adefee7cdf3c4d5e14f3bed7b" \
		   URL_BASE="http://apilayer.net/api"
		 

# Objetivo para correr la aplicación
run:
	$(ENV_VARS) go run api/api.go

# Limpieza opcional (si necesitas eliminar binarios, etc.)
clean:
	rm -rf ./bin

tools/golangci-lint/golangci-lint:
	mkdir -p tools/
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b tools/golangci-lint latest

lint: $(TOOLS_DIR)/golangci-lint/golangci-lint
	./$(TOOLS_DIR)/golangci-lint/golangci-lint run ./...

# Pruebas con cobertura
test:
	$(ENV_VARS) go test ./... -cover -coverprofile=coverage.out && go tool cover -func=coverage.out | grep total | awk '{print "Coverage: "$3}'

# Generar un informe HTML de cobertura
coverage: test
	go tool cover -html=coverage.out -o coverage.html
	@echo "Cobertura generada en coverage.html. Puedes abrirlo en tu navegador."