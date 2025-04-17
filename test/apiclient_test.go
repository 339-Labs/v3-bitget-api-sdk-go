package test

import (
	"fmt"
	"github.com/339-Labs/v3-bitget-api-sdk-go/config"
	"github.com/339-Labs/v3-bitget-api-sdk-go/internal"
	"github.com/339-Labs/v3-bitget-api-sdk-go/pkg/client"
	"github.com/339-Labs/v3-bitget-api-sdk-go/pkg/client/v1"
	v2 "github.com/339-Labs/v3-bitget-api-sdk-go/pkg/client/v2"
	"strconv"
	"testing"
	"time"
)

func Test_ContractsV1(t *testing.T) {
	config := config.NewBitgetConfig(config.BiGetApiKey, config.BiGetApiSecretKey, config.Passphrase, 1000, "")
	params := internal.NewParams()
	params["productType"] = "USDT-FUTURES"
	cl := new(v2.MixMarketClient).Init(config)
	rsp, err := cl.Tickers(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)

}

func Test_Contracts(t *testing.T) {
	config := config.NewBitgetConfig(config.BiGetApiKey, config.BiGetApiSecretKey, config.Passphrase, 1000, "")
	params := internal.NewParams()
	params["productType"] = "umcbl"

	rsp, err := new(v2.MixMarketClient).Init(config).Contracts(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)

}

func Test_WithdrawalRecords(t *testing.T) {

	config := config.NewBitgetConfig(config.BiGetApiKey, config.BiGetApiSecretKey, config.Passphrase, 1000, "")
	params := internal.NewParams()
	layout := "2006-01-02 15:04:05" // 时间格式模板
	start := "2025-04-14 18:30:00"
	ti, err := time.Parse(layout, start)
	if err != nil {
		fmt.Println("解析失败:", err)
		return
	}
	now := time.Now()
	end := now.UnixMilli()
	millis := ti.UnixMilli()
	params["startTime"] = strconv.FormatInt(millis, 10)
	params["endTime"] = strconv.FormatInt(end, 10)
	rsp, err := new(v2.SpotWalletApi).Init(config).WithdrawalRecords(params)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(rsp)

}

func Test_PlaceOrder(t *testing.T) {
	config := config.NewBitgetConfig(config.BiGetApiKey, config.BiGetApiSecretKey, config.Passphrase, 1000, "")
	client := new(v1.MixOrderClient).Init(config)

	params := internal.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.PlaceOrder(params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_post(t *testing.T) {
	config := config.NewBitgetConfig("", "", "", 1000, "")
	client := new(client.BitgetApiClient).Init(config)

	params := internal.NewParams()
	params["symbol"] = "BTCUSDT_UMCBL"
	params["marginCoin"] = "USDT"
	params["side"] = "open_long"
	params["orderType"] = "limit"
	params["price"] = "27012"
	params["size"] = "0.01"
	params["timInForceValue"] = "normal"

	resp, err := client.Post("/api/mix/v1/order/placeOrder", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get(t *testing.T) {
	config := config.NewBitgetConfig("", "", "", 1000, "")
	client := new(client.BitgetApiClient).Init(config)

	params := internal.NewParams()
	params["productType"] = "umcbl"

	resp, err := client.Get("/api/mix/v1/account/accounts", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_params(t *testing.T) {
	config := config.NewBitgetConfig("", "", "", 1000, "")
	client := new(client.BitgetApiClient).Init(config)

	params := internal.NewParams()

	resp, err := client.Get("/api/spot/v1/account/getInfo", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func Test_get_with_encode_params(t *testing.T) {
	config := config.NewBitgetConfig("", "", "", 1000, "")
	client := new(client.BitgetApiClient).Init(config)

	params := internal.NewParams()
	params["symbol"] = "$AIUSDT"
	params["businessType"] = "spot"

	resp, err := client.Get("/api/v2/common/trade-rate", params)
	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
