/*
Bitget Open API

Testing MarginCrossLiquidationApiService

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

func Test_openapi_MarginCrossLiquidationApiService(t *testing.T) {

	apiClient := openapiclient.NewAPIClient(openapiclient.NewDefaultConfiguration())

	t.Run("Test MarginCrossLiquidationApiService CrossLiquidationList", func(t *testing.T) {

		resp, httpRes, err := apiClient.MarginCrossLiquidationApi.CrossLiquidationList(context.Background()).StartTime("1677274167003").EndTime("1680057356760").PageSize("10").Execute()

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
		data := resp.GetData()
		for i, item := range data.GetResultList() {
			fmt.Printf("%d %v\n", i, item)
			assert.NotEmpty(t, item.GetLiqFee())
			assert.NotEmpty(t, item.GetLiqEndTime())
			assert.NotEmpty(t, item.GetLiqStartTime())
			assert.NotEmpty(t, item.GetLiqRisk())
			assert.NotEmpty(t, item.GetTotalAssets())
			assert.NotEmpty(t, item.GetTotalDebt())
		}

	})

}
