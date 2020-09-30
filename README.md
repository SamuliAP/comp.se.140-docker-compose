# comp.se.140-docker-compose

### Running the exercise:
git clone https://github.com/SamuliAP/comp.se.140-docker-compose.git
cd comp.se.140-docker-compose/
docker-compose up --build -d (added detach to the course instructions)
curl localhost:8001
docker-compose down

### Explanation
A private network is automatically created between the services contained in docker-compose.yaml. The default network will use the "bridge" driver. This network driver does not by default have an entrypoint for the host machine or other external machines to access the containers network. We also could have defined this network ourselves, but the default settings already suited the use case.

As docker containers have dynamic addresses, in the proxy program we are referring to the server container by host name "server". This is the name we have configured to be the service name in docker-compose.yaml. Docker automatically maps this host to the service we're referring. This mapping can actually be observed by viewing the hosts -file inside a container, and comparing it to the network configuration, which can be viewed by docker inspect <network id>.

Inside the container network services can only be accessed provided they expose a port to the network. This is done in the docker-compose.yaml -file via keyword "expose". This will expose a container port to the docker network, but will not expose it to machines outside the network, e.g. the host machine.

Now, we'll finally expose a port of the proxy container to the host machine. This is done via the "ports" -keyword, which will take a list of pairs of port values as parameter in format host:container, e.g. 8001:80. Now the host can access container port 80 though the hosts port 8001. 
