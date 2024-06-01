# LEARE

This repository contains a project created for the "Software architecture" course on college. It consists on a microservice created with go (golang) for a microservices-based project that handles documents and it's storage, for this, It's divided on two different modules, one is the main server that uses REST for comunication, and it has a sidecar that uses a queue based communication system with AMQP. The documents to be procesed are written to disk in order to make the system asyncronous and scalable. All of the details on how this project was deployed can be found [here](https://github.com/dcanasp/leare_apiGateway)

# Technical Stack:
- Programming Language: Golang
- Database: DynamoDb and S3
- Queue: RabbitMQ 
- Communication: REST and AMPQ
- Contenerization: Docker

# Structure

It's important to note that i followed a Modular approach therefore each root folder consists of a different module. As i've mentiond before there are a server,a sidecar and a global module. In the last one reside the components that both of the before metnioned modules will use, and this will is bundled separately with each Module.

# Contact Information:
David Alfonso Cañas | Backend Software Developer inquiries: david.alfonso.canas@gmail.com