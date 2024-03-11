.PHONY: start
start:
	@echo "Building the application..."
	@go build -o ../bin/ 
	@echo "Running the application..."
	@../bin/stream.exe

# Compilado con gnuWin32
#start cmd /c "docker start some-rabbit && timeout /t 60 && cd HighQualityMicroservice && make"
# docker run -d --name video_queue -p 5672:5672 -p 15672:15672 -e RABBITMQ_DEFAULT_USER=arquisoft -e RABBITMQ_DEFAULT_PASS=password rabbitmq:management