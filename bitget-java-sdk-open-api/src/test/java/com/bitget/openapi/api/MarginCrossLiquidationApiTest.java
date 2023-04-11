/*
 * Bitget Open API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package com.bitget.openapi.api;

import com.bitget.openapi.ApiConfig;
import com.bitget.openapi.ApiException;
import com.bitget.openapi.model.ApiResponseResultOfMarginLiquidationInfoResult;
import com.bitget.openapi.model.ApiResponseResultOfVoid;
import com.bitget.openapi.model.MarginInterestInfo;
import com.bitget.openapi.model.MarginLiquidationInfo;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static org.assertj.core.api.Assertions.assertThat;

/**
 * API tests for MarginCrossLiquidationApi
 */
public class MarginCrossLiquidationApiTest {

    private final MarginCrossLiquidationApi api = new MarginCrossLiquidationApi(ApiConfig.getConfig());

    /**
     * list
     *
     * Get liquidation List
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void liquidationListTest() throws ApiException {
        String startTime = "1677274167003";
        String endTime = "1680057356760";
        String pageSize = "10";
        String pageId = null;
        ApiResponseResultOfMarginLiquidationInfoResult response = api.liquidationList(startTime, endTime, pageSize, pageId);
        assertThat(response).isNotNull();
        assertThat(response.getCode()).isEqualTo(("00000"));
        assertThat(response.getData()).isNotNull();
        if (response.getData().getResultList() != null && response.getData().getResultList().size() > 0) {
            for(MarginLiquidationInfo item : response.getData().getResultList()) {
                assertThat(item.getLiqFee()).isNotBlank();
                assertThat(item.getLiqEndTime()).isNotBlank();
                assertThat(item.getLiqStartTime()).isNotBlank();
                assertThat(item.getLiqRisk()).isNotBlank();
                assertThat(item.getTotalAssets()).isNotBlank();
                assertThat(item.getTotalDebt()).isNotBlank();
            }
        }
    }

}
