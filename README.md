## DateTime-Server-Mariam_Mahrous

A basic HTTP server that returns the current date and time. This server is implemented using multiple web frameworks (`http` and `gin`) and dockerized using Docker and Docker Compose.

### How to Install

1. Run Docker Compose to generate Docker images. This will produce `http-server` and `gin-server` images.
    ```sh
    docker-compose up
    ```

2. To run the images, use:
    ```sh 
    #run http-server image
    docker run http-server
    #run http-server image
    docker run gin-server
    ```

### How to Use

1. First, figure out the ID of the container:
    ```sh
    docker ps -a
    ```
    Example output:

| CONTAINER ID | IMAGE        | COMMAND                 | CREATED         | STATUS               | PORTS | NAMES                                      |
|--------------|--------------|-------------------------|-----------------|----------------------|-------|---------------------------------------------|
| 971870b28d4f | http-server  | "/bin/sh -c 'go run …"  | 23 minutes ago  | Exited (137) 24 secs |       | datetime-server-mariam_mahrous-http-1       |
| 004aa8381192 | gin-server   | "/bin/sh -c 'go run …"  | 23 minutes ago  | Exited (137) 24 secs |       | datetime-server-mariam_mahrous-gin-server   |```

2. Get the IP address of the currently running container:
    ```sh
    docker inspect 004aa8381192
    # docker inspect [container id] 
    ```

3. Run the `curl` command:
    ```sh
    # For gin server
    curl http://[ipaddress]:8000/datetime
    # For http server
    curl http://[ipaddress]:8080/datetime
    ```

