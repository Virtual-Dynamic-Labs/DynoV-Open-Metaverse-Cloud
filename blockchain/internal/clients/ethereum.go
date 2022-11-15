package ethereum

import (
	"github.com/Virtual-Dynamic-Labs/DynoV-Open-Metaverse-Cloud/blockchain/pkg/log"
	"github.com/chenzhijie/go-web3"
	"github.com/chenzhijie/go-web3/eth"
)

type EthClient struct {
	contracts map[string]*eth.Contract
	client *web3.Web3
	logger log.Logger
}

func CreateEthereumClient(rpcProviderURL string, logger log.Logger) *EthClient {
	web3Client, err := web3.NewWeb3(rpcProviderURL)
	if err != nil {
		panic(err)
	}

	return &EthClient{
		client: web3Client,
		logger: logger,
	}
}

func (w *EthClient) AddContract(contractName, contractAddr, abiString string) {
	contract, err := w.client.Eth.NewContract(abiString, contractAddr)
	if err != nil {
		w.logger.Error(err)
	}
	w.contracts[contractName] = contract
}

func (w *EthClient) GetBlockNumber() uint64 {
	blockNumber, err := w.client.Eth.GetBlockNumber()
	if err != nil {
		w.logger.Error(err)
	}
	w.logger.Info("Current block number: ", blockNumber)
	return blockNumber
}