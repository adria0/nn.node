geth init /config/genesis.json
geth --networkid 34381 \
     --rpc --rpcaddr "0.0.0.0" --rpccorsdomain "*" --rpcapi eth,web,net,debug,txpool \
     --identity adria_p1 \
     --bootnodes enode://0aca71ed24b9163d9402925da7c807eab5d5e26fdc4256d9ae96d087610de22834affa2338aab86060dd6624c63973b8b44f13c79e0232b5fd0c8439b9e6e64d@104.248.135.183:30303
