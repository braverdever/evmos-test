# Challenge

The goal of this challenge is to see if you can interact with on-chain data using simple Rest-API calls, and see original and efficient ways of solving a problem.

## Instructions

### Setup:

1. Install Evmosv18.1.0
    
    ```bash
    git clone https://github.com/evmos/evmos --branch v18.1.0 --depth 1
    cd evmos
    make install
    ```
    
2. Download the database folder:
    
    ```bash
    git clone https://github.com/evmos/challenge_db.git
    mkdir ~/.evmosd
    cp -r challenge_db/* ~/.evmosd
    ```
    
3. Run Evmosd to start the Rest-API:
    
    ```bash
    evmosd start
    ```
    

### For this assignment, you will:

1. Create an open-source (i.e public) GitHub repository to host your project.
2. Create a project that get on-chain statistics using the data between block 100 and 200.
    1. List all the smart contracts that were used and sort them by amount of interactions.
    2. Of all the wallets that interacted with the network, sort them by balance to find the richest user.
3. Add tests for the functionality that you crated.
4. Add to the `README` file:
    1. Instructions to run the code.
    2. What were the main technical decisions you made and why you made them.
    3. Relevant comments about your project and how each of the steps were performed.

## Extra Points

- Implement the solution using GoLang.
- Create tests using [Behaviour Driven Development](https://semaphoreci.com/community/tutorials/getting-started-with-bdd-in-go-using-ginkgo)
- Add a GitHub Action to run a linter (i.e, golang-ci) and tests on pull-requests.
- Save the stats to `sqlite` or a `csv` files.

## Notes:

Evmos is a Cosmos chain, so it supports Cosmos and Ethereum transactions, for this assignment you can focus just on the Ethereum transactions.

## Reference:

- Evmos Documentation: https://docs.evmos.org/
- Evmos Repository: https://github.com/evmos/evmos
- **API Queries**, this is a list of Rest-API endpoints that can be used to get the information, feel free to use any other query that can help you with the problem:
    - **GetBlockNumber**: Returns the latest block height
    
    ```bash
    curl http://localhost:8545/ \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"method":"eth_blockNumber","params":[],"id":1,"jsonrpc":"2.0"}'
    ```
    
    - **GetBlock**: The first parameter is the block height in hex encoding. Sending the second parameter as true, returns the list of transactions included in the block.
    
    ```bash
    curl http://localhost:8545/ \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"method":"eth_getBlockByNumber","params":["0xc5043f",true],"id":1,"jsonrpc":"2.0"}'
    ```
    
    - **GetTransactionReceipt**: The first parameter is the transaction hash
    
    ```bash
    curl http://localhost:8545/ \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"method":"eth_getTransactionReceipt","params":["0x3a74cb4fd48a988ad977a17cc39bec605b0aef4d556f06793b5835d92742e8a8"],"id":1,"jsonrpc":"2.0"}'
    ```
    
    - **GetTransactionTrace**: The first parameter is the transaction hash, the second one is the tracer. Using the callTracer will return the list of contracts that were used in the internals calls of the transaction.
    
    ```bash
    curl http://localhost:8545/ \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"method":"debug_traceTransaction","params":["0x609b2eb0b1750d21523b4aab93ca27b5c694d25363a82369ae18efeafc091c2c", {"tracer": "callTracer"}], "id":1,"jsonrpc":"2.0"}'
    ```
    
    - **GetBalance**: the first parameter is the address and second parameter is the height at the point where you want to query the balance.
    
    ```bash
    curl http://localhost:8545/ \
    -X POST \
    -H "Content-Type: application/json" \
    --data '{"method":"eth_getBalance","params":["0x8D97689C9818892B700e27F316cc3E41e17fBeb9", "latest"],"id":1,"jsonrpc":"2.0"}'
    ```
    

## Question to be asked after the presentation

1. If we realize that the bottleneck of getting the data is in the `debug_ethtrace` call and we know that we are going to need to run this process several times until we get all the requirements correctly, what changes in code or infrastructure will you do?
2. Instead of reading the data using http requests, we want to read the data from json files stored in disk, what changes are needed in the code to implement this change?
