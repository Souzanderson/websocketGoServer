# websocketGoServer

- For run:
    * Select your executable by OS in machine and run, example:
        
        ```bash
        ./gows_mac_arm64
        ```

        * It`s run on port 8083;

        ```bash
        ./gows_mac_arm64 8085
        ```

        * It`s run on port 8085;


- Model url:

    ```bash
    ws://<host>:<port>?id=<iduser>&topic=<topiclogin>
    ```

- Model payload (body):

    ```json
    {
        "topic": "topic for send message or all (all receive message)",
        "payload": "JSON like (string, array, json...)"
    }
    ```