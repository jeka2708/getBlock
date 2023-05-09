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

const (
	lastBlock     = "eth_blockNumber"
	fullBlockInfo = "eth_getBlockByNumber"
)

func NewGetBlock(apiKey string) *GetBlockApi {
	return &GetBlockApi{
		clients.NewGetBlockClient(apiKey),
	}
}

func (gb *GetBlockApi) getLastBlock() (string, error) {
	response, err := gb.Call(lastBlock)

	if err != nil {
		log.Println(err)
		return "", err
	}
	return response.Result.(string)[2:], nil
}

func (gb *GetBlockApi) GetAddress(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	lastBlock, err := gb.getLastBlock()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	blockNum, err := strconv.ParseInt(lastBlock, 16, 64)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := models.Result{}
	var wg sync.WaitGroup
	for i := blockNum - 100; i <= blockNum; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			block := models.BlockByNumberResponse{}
			numberBlock := fmt.Sprintf("0x%s", strconv.FormatInt(i, 16))
			err = gb.CallFor(&block, fullBlockInfo, numberBlock, true)

			if err != nil {
				log.Println(err)
				return
			}
			if len(block.Transactions) == 0 {
				return
			}
			result.Lock()
			result.Transactions = append(result.Transactions, utils.GetMaxValue(block.Transactions))
			result.Unlock()
		}(i)

	}

	wg.Wait()
	keyAddress := utils.GetMaxValue(result.Transactions)
	return &pb.Response{Message: keyAddress.To}, nil
}
