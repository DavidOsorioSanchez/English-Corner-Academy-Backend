all: build test

# estoy tonteando con la posibilidad de crear comandos adicionales	
# pero mientras tanto dejo esto simple
# por ejemplo, un comando para correr migraciones

build:
	@echo "Building..."
	
	
	@go build -o main.exe cmd/api/main.go