# Build

1. Install nginx without root.
    ```
    ./install.sh
    ```

# Run

1. Create `logs` directory.
    ```
    $ mkdir logs
    ```

2. Create `error.log`.
    ```
    $ touch error.log
    ```

3. Run nginx with specified config.
    ```
    $ ~/.local/src/nginx-1.8.1/objs/nginx -c nginx.conf -p .
    ```

4. Validate nginx server is listening.
    ```
    $ netstat -na | grep 8081
    tcp        0      0 0.0.0.0:8081            0.0.0.0:*               LISTEN
    ```

# References
* [install.sh from Christopher Baek](https://gist.github.com/christopherbaek/c0feb9f54f4c57a7aee9a1efd1e5e18e)
* [nginx full example configuration](https://www.nginx.com/resources/wiki/start/topics/examples/full/)