package api

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"

	"getBlockTest/pkg/clients"
	"getBlockTest/pkg/models"
	"getBlockTest/pkg/utils"
	pb "getBlockTest/proto/go_proto"
)

type GetBlockApi struct {
	*clients.GetBlockClient
}

func NewGetBlock(apiKey string) *GetBlockApi {
	return &GetBlockApi{
		clients.NewGetBlockClient(apiKey),
	}
}
func (gb *GetBlockApi) GetAddress(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	response, err := gb.Call("eth_blockNumber")

	if err != nil {
		log.Fatal(err)
	}

	blockNum, err := strconv.ParseInt(response.Result.(string)[2:], 16, 64)
	if err != nil {
		log.Fatal(err)
	}

	result := models.Result{}
	var wg sync.WaitGroup
	for i := blockNum - 100; i <= blockNum; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			block := models.BlockByNumberResponse{}
			numberBlock := fmt.Sprintf("0x%s", strconv.FormatInt(i, 16))
			err = gb.CallFor(&block, "eth_getBlockByNumber", numberBlock, true)
			if len(block.Transactions) == 0 {
				return
			}
			result.Lock()
			result.Transactions = append(result.Transactions, utils.GetMaxValue(block.Transactions))
			result.Unlock()
		}(i)

	}

	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
	keyAddress := utils.GetMaxValue(result.Transactions)
	return &pb.Response{Message: keyAddress.To}, nil
}
