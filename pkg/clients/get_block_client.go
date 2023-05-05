package clients

import (
	"github.com/ybbus/jsonrpc/v2"
)

type GetBlockClient struct {
	jsonrpc.RPCClient
}

func NewGetBlockClient(apiKey string) *GetBlockClient {
	rpcCli := jsonrpc.NewClientWithOpts("https://eth.getblock.io/mainnet/", &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"x-api-key": apiKey,
		},
	})
	return &GetBlockClient{
		rpcCli,
	}
}
