.PHONY: all build-server build-sidecar

all: build-server build-sidecar
# Build commands
build-server:
#	@echo "Building the server application..."
	start cmd /c "cd server && make"  

build-sidecar:
# 	@echo "Building the sidecar application..."
	start cmd /c "cd sidecar && make" 

# run-server:
# 	@echo "Running the server application..."
# 	start cmd /k ".\bin\server.exe"

# run-sidecar:
# 	@echo "Running the sidecar application..."
# 	start cmd /k ".\bin\sidecar.exe"

