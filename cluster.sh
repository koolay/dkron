./dkron-linux agent -node-name=02 -backend=consul -backend-machine=127.0.0.1:8500 -http-addr=:8082 -rpc-port=6867 -bind-addr=127.0.0.1:8947 -join=127.0.0.1:8946 -server
./dkron-linux agent -node-name=03 -backend=consul -backend-machine=127.0.0.1:8500 -http-addr=:8083 -rpc-port=6866 -bind-addr=127.0.0.1:8948 -join=127.0.0.1:8946 -server
./dkron-linux agent -node-name=04 -backend=consul -backend-machine=127.0.0.1:8500 -http-addr=:8084 -rpc-port=6865 -bind-addr=127.0.0.1:8949 -join=127.0.0.1:8946 -server
./dkron-linux agent -node-name=05 -backend=consul -backend-machine=127.0.0.1:8500 -http-addr=:8085 -rpc-port=6864 -bind-addr=127.0.0.1:8950 -join=127.0.0.1:8946 -server
