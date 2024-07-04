# DateTime-Server-Mariam_Mahrous

A basic HTTP server that returns the current date and time. This server is implemented using multiple web frameworks (`http` and `gin`).

## How to Install

1. To produce `http-server` and `gin-server` images run one of the following commands.
    ```sh
    make images #2
    ```

2. To run a container 
    ```sh 
    # For gin server
    make gin
    # For http server
    make http
    ```
    ## faster way
    ```sh 
    docker-compose up #This will produce the images and run both containers
    ```

## How to Use

1. Using google chrome or your preferred browser type
    <br>
    ```
    #for gin server
    http://localhost:8000/datetime
    #for http server
    http://localhost:8080/datetime
    ```
### Alternative Usage using curl command

2. Run using `curl` command:
    ```sh
    # For gin server
    curl http://localhost:8000/datetime
    # For http server
    curl http://localhost:8080/datetime
    ```

## How to test
To be able to run unit tests run: <br>
```sh
    make test
```
<br>