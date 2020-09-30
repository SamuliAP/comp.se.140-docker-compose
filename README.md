# comp.se.140-docker-compose

### Running the exercise:
git clone https://github.com/SamuliAP/comp.se.140-docker-compose.git

cd comp.se.140-docker-compose/

docker-compose up --build -d (added detach to the course instructions)

curl localhost:8001

docker-compose down

### Explanation
A private network is automatically created between the services contained in docker-compose.yaml. The default network will use the `bridge` driver. This network driver does not by default have an entrypoint for the host machine or other external machines to access the network. We also could have defined this network ourselves, but the default settings already suited the use case.

As docker containers have dynamic addresses, in the proxy program we have to refer to the server container by host name "server". This is the name we have configured to be the service name in docker-compose.yaml. Docker automatically maps this host to the service we're referring. This mapping can actually be observed by viewing the hosts -file inside a container, and comparing it to the network configuration, which can be viewed by `docker inspect <network id>`. In this case we are not specifying a port in the proxy program, as the server -container alerady exposes port 80, the default port for http communication.

Inside the container network services can only be accessed provided they expose a port to the network. This is done in the docker-compose.yaml -file via keyword `expose`. This will expose a container port to the docker network, but will not expose it to machines outside the network, e.g. the host machine.

Now, we'll expose a port of the proxy container to the host machine. This is done via the `ports` -keyword, which will take a list of pairs of port values as parameter in format host:container, e.g. 8001:80. Now the host can access container port 80 though the hosts port 8001. 

### Example Result
Hello from 172.18.0.1:46540 <-- this is the gateway address from host machine to the proxy service. Here we can see we're actually not connecting directly to proxy from our host machine, but have to use kind of a reverse proxy (probably also doing some authorization) provided by Docker

to localhost:8001 <-- this is the host we connected with curl, we could've also connected to the actual address found in `docker inspect <network id>`, e.g. 172.18.0.3:80

Hello from 172.18.0.3:53912 <- this is the proxy service

to server <- and finally the host we connected from proxy, which again, could've also been the actual address of the server service. This time though, we can see we actually did directly connect from the proxy service to the server service.
