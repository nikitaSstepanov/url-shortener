<h1>Hi there 👋</h1>
<h1>It is url shortener</h1>

### Description

It is url shortener. It allows you to shorten links and create custom aliases.

### Launch
   
1) Set up configuration of the project (config/config.yaml) or use default values.

2) Set up configuration of docker (docker-compose.yaml) if you use it.

3) If you don`t use docker, raise the postgres and redis databases in advance (do not forget that app will try to connect to the dbs with the environment specified in config/config.yaml).
  
4) Set up the environment (create ".env" file) according to the example in the ".env.example" file or rename ".env.example" to ".env" to use default values.

5) Navigate to the root of the project in the terminal and enter the command:

    ```shell
    docker compose up
    ```
    
    If you don`t use docker, enter:

    ```shell
    go run ./cmd/url-shortener/main.go
    ```