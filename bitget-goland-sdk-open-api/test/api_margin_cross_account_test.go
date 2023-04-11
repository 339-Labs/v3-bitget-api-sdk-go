/*
Bitget Open API

Testing MarginCrossAccountApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	openapiclient "github.com/bitget/openapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_openapi_MarginCrossAccountApiService(t *testing.T) {
	apiClient := openapiclient.NewAPIClient(openapiclient.NewDefaultConfiguration())

	t.Run("Test MarginCrossAccountApiService MarginCrossAccountAssets", func(t *testing.T) {

		resp, httpRes, err := apiClient.MarginCrossAccountApi.MarginCrossAccountAssets(context.Background()).Coin("USDT").Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		assert.NotNil(t, resp.GetData())

		for i, item := range resp.GetData() {
			fmt.Printf("%d %v\n", i, item)
			assert.NotEmpty(t, item.GetAvailable())
			assert.NotEmpty(t, item.GetBorrow())
			assert.NotEmpty(t, item.GetCoin())
			assert.NotEmpty(t, item.GetInterest())
			assert.NotEmpty(t, item.GetTotalAmount())
		}

	})

	t.Run("Test MarginCrossAccountApiService MarginCrossAccountBorrow", func(t *testing.T) {

		param := *openapiclient.NewMarginCrossLimitRequestWithDefaults() // NewMarginCrossLimitRequestWithDefaults | param
		param.SetCoin("USDT")
		param.SetBorrowAmount("1")
		resp, httpRes, err := apiClient.MarginCrossAccountApi.MarginCrossAccountBorrow(context.Background()).MarginCrossLimitRequest(param).Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		data := resp.GetData()
		assert.NotNil(t, data.GetBorrowAmount())
		assert.NotNil(t, data.Coin)

	})

	t.Run("Test MarginCrossAccountApiService MarginCrossAccountMaxBorrowableAmount", func(t *testing.T) {

		param := *openapiclient.NewMarginCrossMaxBorrowRequestWithDefaults() // NewMarginCrossMaxBorrowRequestWithDefaults | param
		param.SetCoin("USDT")
		resp, httpRes, err := apiClient.MarginCrossAccountApi.MarginCrossAccountMaxBorrowableAmount(context.Background()).MarginCrossMaxBorrowRequest(param).Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		data := resp.GetData()
		assert.NotNil(t, data.GetMaxBorrowableAmount())
		assert.NotNil(t, data.Coin)

	})

	t.Run("Test MarginCrossAccountApiService MarginCrossAccountMaxTransferOutAmount", func(t *testing.T) {

		resp, httpRes, err := apiClient.MarginCrossAccountApi.MarginCrossAccountMaxTransferOutAmount(context.Background()).Coin("USDT").Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		data := resp.GetData()
		assert.NotNil(t, data.GetMaxTransferOutAmount())
		assert.NotNil(t, data.Coin)

	})

	t.Run("Test MarginCrossAccountApiService MarginCrossAccountRepay", func(t *testing.T) {

		param := *openapiclient.NewMarginCrossRepayRequestWithDefaults() // NewMarginCrossRepayRequestWithDefaults | param
		param.SetCoin("USDT")
		param.SetRepayAmount("1")
		resp, httpRes, err := apiClient.MarginCrossAccountApi.MarginCrossAccountRepay(context.Background()).MarginCrossRepayRequest(param).Execute()
		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "success", resp.GetMsg())
		data := resp.GetData()
		assert.NotNil(t, data.GetRepayAmount())
		assert.NotNil(t, data.Coin)

	})

	t.Run("Test MarginCrossAccountApiService MarginCrossAccountRiskRate", func(t *testing.T) {

		resp, httpRes, err := apiClient.MarginCrossAccountApi.MarginCrossAccountRiskRate(context.Background()).Execute()

		bs, _ := json.Marshal(resp)
		var out bytes.Buffer
		json.Indent(&out, bs, "", "\t")
		fmt.Printf("result=%v\n", out.String())

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
		assert.Equal(t, "00000", resp.GetCode())
		assert.Equal(t, "success", resp.GetMsg())
		data := resp.GetData()
		assert.NotNil(t, data.GetRiskRate())

	})

	t.Run("Test MarginCrossAccountApiService Void", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.MarginCrossAccountApi.Void(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
