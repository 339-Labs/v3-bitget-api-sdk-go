package main

import (
	"fmt"
	"github.com/339-Labs/v3-bitget-api-sdk-go/config"
	"github.com/339-Labs/v3-bitget-api-sdk-go/pkg/client/ws"
	"github.com/339-Labs/v3-bitget-api-sdk-go/pkg/client/ws/model"
)

func main() {
	config := config.NewBitgetConfig(config.BiGetApiKey, config.BiGetApiSecretKey, config.Passphrase, 1000, "", "", "")
	client := new(ws.BitgetWsClient).Init(config, true, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	})

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channelsDef = append(channelsDef, subReqDef1)
	client.SubscribeDef(channelsDef)

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})
	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()
}
