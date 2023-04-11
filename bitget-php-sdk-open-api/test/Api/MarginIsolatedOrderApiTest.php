<?php
/**
 * MarginIsolatedOrderApiTest
 * PHP version 7.4
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Bitget Open API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 2.0.0
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 6.2.1
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Please update the test case below to test the endpoint.
 */

namespace Bitget\Test\Api;

use \Bitget\Configuration;
use \Bitget\ApiException;
use \Bitget\ObjectSerializer;
use Exception;
use PHPUnit\Framework\TestCase;

/**
 * MarginIsolatedOrderApiTest Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class MarginIsolatedOrderApiTest extends TestCase
{

    private $config;

    private $apiInstance;

    /**
     * Setup before running any test cases
     */
    public static function setUpBeforeClass(): void
    {
    }

    /**
     * Setup before running each test case
     */
    public function setUp(): void
    {
        $this->config = \Bitget\Config::getDefaultConfig();
        $this->apiInstance = new \Bitget\Api\MarginIsolatedOrderApi(
        // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
        // This is optional, `GuzzleHttp\Client` will be used as default.
            new \GuzzleHttp\Client(),
            $this->config
        );
    }

    /**
     * Clean up after running each test case
     */
    public function tearDown(): void
    {
    }

    /**
     * Clean up after running all test cases
     */
    public static function tearDownAfterClass(): void
    {
    }

    /**
     * Test case for marginIsolatedBatchCancelOrder
     *
     * batchCancelOrder.
     *
     */
    public function testMarginIsolatedBatchCancelOrder()
    {
        try {
            $req = new \Bitget\Model\MarginOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setSide("buy");
            $req->setOrderType("limit");
            $req->setTimeInForce("gtc");
            $req->setPrice("1600");
            $req->setBaseQuantity("0.625");
            $req->setQuoteAmount("1000");
            $req->setLoanType("normal");
            $result = $this->apiInstance->marginIsolatedPlaceOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderId());

            $req = new \Bitget\Model\MarginBatchCancelOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setOrderIds([$result->getData()->getOrderId()]);
            $result = $this->apiInstance->marginIsolatedBatchCancelOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getResultList()[0]->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginIsolatedBatchPlaceOrder
     *
     * batchPlaceOrder.
     *
     */
    public function testMarginIsolatedBatchPlaceOrder()
    {
        try {
            $item = new \Bitget\Model\MarginOrderRequest(); //
            $item->setSymbol("BTCUSDT");
            $item->setSide("buy");
            $item->setOrderType("market");
            $item->setTimeInForce("gtc");
            $item->setQuoteAmount("10");
            $item->setLoanType("normal");
            $req = new \Bitget\Model\MarginBatchOrdersRequest();
            $req->setSymbol("BTCUSDT");
            $req->setOrderList([$item]);
            $result = $this->apiInstance->marginIsolatedBatchPlaceOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getResultList()[0]->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginIsolatedCancelOrder
     *
     * cancelOrder.
     *
     */
    public function testMarginIsolatedCancelOrder()
    {
        try {
            $req = new \Bitget\Model\MarginOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setSide("buy");
            $req->setOrderType("limit");
            $req->setTimeInForce("gtc");
            $req->setPrice("1600");
            $req->setBaseQuantity("0.625");
            $req->setQuoteAmount("1000");
            $req->setLoanType("normal");
            $result = $this->apiInstance->marginIsolatedPlaceOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderId());

            $req = new \Bitget\Model\MarginCancelOrderRequest(); //
            $req->setSymbol("BTCUSDT");
            $req->setOrderId($result->getData()->getOrderId());
            $result = $this->apiInstance->marginIsolatedCancelOrder($req);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getResultList()[0]->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginIsolatedFills
     *
     * fills.
     *
     */
    public function testMarginIsolatedFills()
    {
        try {
            $result = $this->apiInstance->marginIsolatedFills("1679133422000", "BTCUSDT");
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getFills());
            foreach ($result->getData()->getFills() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertNotNull($item->getSide());
                $this->assertNotNull($item->getOrderType());
                $this->assertNotNull($item->getOrderId());
                $this->assertNotNull($item->getFillTotalAmount());
                $this->assertNotNull($item->getFillQuantity());
                $this->assertNotNull($item->getFillPrice());
                $this->assertNotNull($item->getFeeCcy());
                $this->assertNotNull($item->getFees());
                $this->assertNotNull($item->getFillId());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginIsolatedHistoryOrders
     *
     * history.
     *
     */
    public function testMarginIsolatedHistoryOrders()
    {
        try {
            $result = $this->apiInstance->marginIsolatedHistoryOrders("1679133422000","BTCUSDT",);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderList());
            foreach ($result->getData()->getOrderList() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertEquals($item->getSymbol(), "BTCUSDT");
                $this->assertNotNull($item->getSide());
                $this->assertNotNull($item->getSource());
                $this->assertNotNull($item->getStatus());
                $this->assertNotNull($item->getOrderId());
                $this->assertNotNull($item->getQuoteAmount());
                $this->assertNotNull($item->getBaseQuantity());
                $this->assertNotNull($item->getFillPrice());
                $this->assertNotNull($item->getFillQuantity());
                $this->assertNotNull($item->getFillTotalAmount());
                $this->assertNotNull($item->getLoanType());
                $this->assertNotNull($item->getOrderType());
                $this->assertNotNull($item->getPrice());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginIsolatedOpenOrders
     *
     * openOrders.
     *
     */
    public function testMarginIsolatedOpenOrders()
    {
        try {
            $result = $this->apiInstance->marginIsolatedOpenOrders("BTCUSDT", "1679133422000" );
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderList());
            foreach ($result->getData()->getOrderList() as $item) {
                print_r($item);
                $this->assertNotNull($item);
                $this->assertEquals($item->getSymbol(), "BTCUSDT");
                $this->assertNotNull($item->getSide());
                $this->assertNotNull($item->getSource());
                $this->assertNotNull($item->getStatus());
                $this->assertNotNull($item->getOrderId());
                $this->assertNotNull($item->getQuoteAmount());
                $this->assertNotNull($item->getBaseQuantity());
                $this->assertNotNull($item->getFillPrice());
                $this->assertNotNull($item->getFillQuantity());
                $this->assertNotNull($item->getFillTotalAmount());
                $this->assertNotNull($item->getLoanType());
                $this->assertNotNull($item->getOrderType());
                $this->assertNotNull($item->getPrice());
            }
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }

    /**
     * Test case for marginIsolatedPlaceOrder
     *
     * placeOrder.
     *
     */
    public function testMarginIsolatedPlaceOrder()
    {
        try {
            $placeOrderReq = new \Bitget\Model\MarginOrderRequest(); //
            $placeOrderReq->setSymbol("BTCUSDT");
            $placeOrderReq->setSide("buy");
            $placeOrderReq->setOrderType("market");
            $placeOrderReq->setTimeInForce("gtc");
            $placeOrderReq->setQuoteAmount("10");
            $placeOrderReq->setLoanType("normal");
            $result = $this->apiInstance->marginIsolatedPlaceOrder($placeOrderReq);
            print_r($result);

            $this->assertNotNull($result);
            $this->assertEquals($result->getCode(), "00000");
            $this->assertEquals($result->getMsg(), "success");
            $this->assertNotNull($result->getData());
            $this->assertNotNull($result->getData()->getOrderId());
        } catch (Exception $e) {
            echo 'Exception when calling : ', $e->getMessage(), PHP_EOL;
        }
    }
}
