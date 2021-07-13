INFO [07-12|17:24:12.261] Starting Geth on Ethereum mainnet...
INFO [07-12|17:24:12.261] Bumping default cache on mainnet         provided=1024 updated=4096
INFO [07-12|17:24:12.262] Maximum peer count                       ETH=50 LES=0 total=50
WARN [07-12|17:24:12.262] The flag --rpc is deprecated and will be removed June 2021, please use --http
WARN [07-12|17:24:12.262] The flag --rpcport is deprecated and will be removed June 2021, please use --http.port
INFO [07-12|17:24:12.262] Smartcard socket not found, disabling    err="stat /run/pcscd/pcscd.comm: no such file or directory"
WARN [07-12|17:24:12.262] Sanitizing cache to Go's GC limits       provided=4096 updated=1310
INFO [07-12|17:24:12.263] Set global gas cap                       cap=50,000,000
INFO [07-12|17:24:12.263] Allocated trie memory caches             clean=196.00MiB dirty=327.00MiB
INFO [07-12|17:24:12.263] Allocated cache and file handles         database=/home/csx/ethx/go-ethereum/data0/geth/chaindata cache=653.00MiB handles=2048
INFO [07-12|17:24:12.288] Opened ancient database                  database=/home/csx/ethx/go-ethereum/data0/geth/chaindata/ancient readonly=false
INFO [07-12|17:24:12.288] Initialised chain configuration          config="{ChainID: 1 Homestead: 1150000 DAO: 1920000 DAOSupport: true EIP150: 2463000 EIP155: 2675000 EIP158: 2675000 Byzantium: 4370000 Constantinople: 7280000 Petersburg: 7280000 Istanbul: 9069000, Muir Glacier: 9200000, Berlin: 12244000, London: <nil>, Engine: ethash}"
INFO [07-12|17:24:12.288] Disk storage enabled for ethash caches   dir=/home/csx/ethx/go-ethereum/data0/geth/ethash count=3
INFO [07-12|17:24:12.288] Disk storage enabled for ethash DAGs     dir=/home/csx/.ethash count=2
INFO [07-12|17:24:12.289] Initialising Ethereum protocol           network=1 dbversion=8
INFO [07-12|17:24:12.309] Loaded most recent local header          number=0 hash=d4e567..cb8fa3 td=17,179,869,184 age=52y3mo1w
INFO [07-12|17:24:12.309] Loaded most recent local full block      number=0 hash=d4e567..cb8fa3 td=17,179,869,184 age=52y3mo1w
INFO [07-12|17:24:12.309] Loaded most recent local fast block      number=0 hash=d4e567..cb8fa3 td=17,179,869,184 age=52y3mo1w
INFO [07-12|17:24:12.310] Loaded local transaction journal         transactions=0 dropped=0
INFO [07-12|17:24:12.310] Regenerated local transaction journal    transactions=0 accounts=0
INFO [07-12|17:24:12.310] Gasprice oracle is ignoring threshold set threshold=2
WARN [07-12|17:24:12.310] Unclean shutdown detected                booted=2021-07-12T15:16:47+0800 age=2h7m25s
WARN [07-12|17:24:12.311] Unclean shutdown detected                booted=2021-07-12T15:17:17+0800 age=2h6m55s
WARN [07-12|17:24:12.311] Unclean shutdown detected                booted=2021-07-12T15:22:30+0800 age=2h1m42s
INFO [07-12|17:24:12.311] Starting peer-to-peer node               instance=Geth/TestNode/v1.10.5-unstable-10eb654f-20210623/linux-amd64/go1.16.5
INFO [07-12|17:24:12.340] New local node record                    seq=13 id=168b4f4b91a309b4 ip=127.0.0.1 udp=0 tcp=30303
INFO [07-12|17:24:12.340] Started P2P networking                   self="enode://d1d704b83476a6e95582d08ea391c58af3b28133bcdfa1a8fcad5bce44ee36095d794ecb217465b2d8b19d5f0a76402b02d36a1928ab5e45193664d6361622ea@127.0.0.1:30303?discport=0"
INFO [07-12|17:24:12.341] IPC endpoint opened                      url=/home/csx/ethx/go-ethereum/data0/geth.ipc
INFO [07-12|17:24:12.341] HTTP server started                      endpoint=127.0.0.1:8545 prefix= cors= vhosts=localhost
INFO [07-12|17:24:12.342] WebSocket enabled                        url=ws://[::]:8547
WARN [07-12|17:24:12.377] Served eth_coinbase                      reqid=3 t="30.1µs" err="etherbase must be explicitly specified"
Welcome to the Geth JavaScript console!

instance: Geth/TestNode/v1.10.5-unstable-10eb654f-20210623/linux-amd64/go1.16.5
at block: 0 (Thu Jan 01 1970 08:00:00 GMT+0800 (CST))
datadir: /home/csx/ethx/go-ethereum/data0
modules: admin:1.0 debug:1.0 eth:1.0 ethash:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0
