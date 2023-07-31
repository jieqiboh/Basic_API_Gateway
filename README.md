## Basic API Gateway
This gateway was generated from our code generator with slight modifications, and supports 3 microservices that are registered to an etcd cluster.
The services are EchoService, BizService, HelloService on ports 8000, 9001, 9000 respectively.
For more information on the exact addresses, refer to the /constants/constants.go file.


## Running the Gateway

Set up the microservices and etcd cluster first before running the gateway, then do:

```bash
  go build gateway
```
Once the gateway binary executable is generated, do:

```bash
  ./gateway [port number]
```
This runs the gateway on the specified port number.