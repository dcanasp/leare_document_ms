.PHONY: start
start:
	@echo "Building the application..."
	@go build -o bin/ 
	@echo "Running the application..."
	@./bin/stream.exe

# Compilado con gnuWin32