<?php
/**
 * MarginIsolatedRepayApi
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
 * Do not edit the class manually.
 */

namespace Bitget\Api;

use GuzzleHttp\Client;
use GuzzleHttp\ClientInterface;
use GuzzleHttp\Exception\ConnectException;
use GuzzleHttp\Exception\RequestException;
use GuzzleHttp\Psr7\MultipartStream;
use GuzzleHttp\Psr7\Request;
use GuzzleHttp\RequestOptions;
use Bitget\ApiException;
use Bitget\Configuration;
use Bitget\HeaderSelector;
use Bitget\ObjectSerializer;

/**
 * MarginIsolatedRepayApi Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */
class MarginIsolatedRepayApi
{
    /**
     * @var ClientInterface
     */
    protected $client;

    /**
     * @var Configuration
     */
    protected $config;

    /**
     * @var HeaderSelector
     */
    protected $headerSelector;

    /**
     * @var int Host index
     */
    protected $hostIndex;

    /** @var string[] $contentTypes **/
    public const contentTypes = [
        'isolateRepayList' => [
            'application/json',
        ],
    ];

/**
     * @param ClientInterface $client
     * @param Configuration   $config
     * @param HeaderSelector  $selector
     * @param int             $hostIndex (Optional) host index to select the list of hosts if defined in the OpenAPI spec
     */
    public function __construct(
        ClientInterface $client = null,
        Configuration $config = null,
        HeaderSelector $selector = null,
        $hostIndex = 0
    ) {
        $this->client = $client ?: new Client();
        $this->config = $config ?: new Configuration();
        $this->headerSelector = $selector ?: new HeaderSelector();
        $this->hostIndex = $hostIndex;
    }

    /**
     * Set the host index
     *
     * @param int $hostIndex Host index (required)
     */
    public function setHostIndex($hostIndex): void
    {
        $this->hostIndex = $hostIndex;
    }

    /**
     * Get the host index
     *
     * @return int Host index
     */
    public function getHostIndex()
    {
        return $this->hostIndex;
    }

    /**
     * @return Configuration
     */
    public function getConfig()
    {
        return $this->config;
    }

    /**
     * Operation isolateRepayList
     *
     * list
     *
     * @param  string $symbol symbol (required)
     * @param  string $start_time startTime (required)
     * @param  string $coin coin (optional)
     * @param  string $repay_id repayId (optional)
     * @param  string $end_time endTime (optional)
     * @param  string $page_size pageSize (optional)
     * @param  string $page_id pageId (optional)
     * @param  string $contentType The value for the Content-Type header. Check self::contentTypes['isolateRepayList'] to see the possible values for this operation
     *
     * @throws \Bitget\ApiException on non-2xx response
     * @throws \InvalidArgumentException
     * @return \Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult|\Bitget\Model\ApiResponseResultOfVoid|\Bitget\Model\ApiResponseResultOfVoid|\Bitget\Model\ApiResponseResultOfVoid
     */
    public function isolateRepayList($symbol, $start_time, $coin = null, $repay_id = null, $end_time = null, $page_size = null, $page_id = null, string $contentType = self::contentTypes['isolateRepayList'][0])
    {
        list($response) = $this->isolateRepayListWithHttpInfo($symbol, $start_time, $coin, $repay_id, $end_time, $page_size, $page_id, $contentType);
        return $response;
    }

    /**
     * Operation isolateRepayListWithHttpInfo
     *
     * list
     *
     * @param  string $symbol symbol (required)
     * @param  string $start_time startTime (required)
     * @param  string $coin coin (optional)
     * @param  string $repay_id repayId (optional)
     * @param  string $end_time endTime (optional)
     * @param  string $page_size pageSize (optional)
     * @param  string $page_id pageId (optional)
     * @param  string $contentType The value for the Content-Type header. Check self::contentTypes['isolateRepayList'] to see the possible values for this operation
     *
     * @throws \Bitget\ApiException on non-2xx response
     * @throws \InvalidArgumentException
     * @return array of \Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult|\Bitget\Model\ApiResponseResultOfVoid|\Bitget\Model\ApiResponseResultOfVoid|\Bitget\Model\ApiResponseResultOfVoid, HTTP status code, HTTP response headers (array of strings)
     */
    public function isolateRepayListWithHttpInfo($symbol, $start_time, $coin = null, $repay_id = null, $end_time = null, $page_size = null, $page_id = null, string $contentType = self::contentTypes['isolateRepayList'][0])
    {
        $request = $this->isolateRepayListRequest($symbol, $start_time, $coin, $repay_id, $end_time, $page_size, $page_id, $contentType);

        try {
            $options = $this->createHttpClientOption();
            try {
                $response = $this->client->send($request, $options);
            } catch (RequestException $e) {
                throw new ApiException(
                    "[{$e->getCode()}] {$e->getMessage()}",
                    (int) $e->getCode(),
                    $e->getResponse() ? $e->getResponse()->getHeaders() : null,
                    $e->getResponse() ? (string) $e->getResponse()->getBody() : null
                );
            } catch (ConnectException $e) {
                throw new ApiException(
                    "[{$e->getCode()}] {$e->getMessage()}",
                    (int) $e->getCode(),
                    null,
                    null
                );
            }

            $statusCode = $response->getStatusCode();

            if ($statusCode < 200 || $statusCode > 299) {
                throw new ApiException(
                    sprintf(
                        '[%d] Error connecting to the API (%s)',
                        $statusCode,
                        (string) $request->getUri()
                    ),
                    $statusCode,
                    $response->getHeaders(),
                    (string) $response->getBody()
                );
            }

            switch($statusCode) {
                case 200:
                    if ('\Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult' === '\SplFileObject') {
                        $content = $response->getBody(); //stream goes to serializer
                    } else {
                        $content = (string) $response->getBody();
                        if ('\Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult' !== 'string') {
                            $content = json_decode($content);
                        }
                    }

                    return [
                        ObjectSerializer::deserialize($content, '\Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                case 400:
                    if ('\Bitget\Model\ApiResponseResultOfVoid' === '\SplFileObject') {
                        $content = $response->getBody(); //stream goes to serializer
                    } else {
                        $content = (string) $response->getBody();
                        if ('\Bitget\Model\ApiResponseResultOfVoid' !== 'string') {
                            $content = json_decode($content);
                        }
                    }

                    return [
                        ObjectSerializer::deserialize($content, '\Bitget\Model\ApiResponseResultOfVoid', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                case 429:
                    if ('\Bitget\Model\ApiResponseResultOfVoid' === '\SplFileObject') {
                        $content = $response->getBody(); //stream goes to serializer
                    } else {
                        $content = (string) $response->getBody();
                        if ('\Bitget\Model\ApiResponseResultOfVoid' !== 'string') {
                            $content = json_decode($content);
                        }
                    }

                    return [
                        ObjectSerializer::deserialize($content, '\Bitget\Model\ApiResponseResultOfVoid', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                case 500:
                    if ('\Bitget\Model\ApiResponseResultOfVoid' === '\SplFileObject') {
                        $content = $response->getBody(); //stream goes to serializer
                    } else {
                        $content = (string) $response->getBody();
                        if ('\Bitget\Model\ApiResponseResultOfVoid' !== 'string') {
                            $content = json_decode($content);
                        }
                    }

                    return [
                        ObjectSerializer::deserialize($content, '\Bitget\Model\ApiResponseResultOfVoid', []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
            }

            $returnType = '\Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult';
            if ($returnType === '\SplFileObject') {
                $content = $response->getBody(); //stream goes to serializer
            } else {
                $content = (string) $response->getBody();
                if ($returnType !== 'string') {
                    $content = json_decode($content);
                }
            }

            return [
                ObjectSerializer::deserialize($content, $returnType, []),
                $response->getStatusCode(),
                $response->getHeaders()
            ];

        } catch (ApiException $e) {
            switch ($e->getCode()) {
                case 200:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 400:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\Bitget\Model\ApiResponseResultOfVoid',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 429:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\Bitget\Model\ApiResponseResultOfVoid',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
                case 500:
                    $data = ObjectSerializer::deserialize(
                        $e->getResponseBody(),
                        '\Bitget\Model\ApiResponseResultOfVoid',
                        $e->getResponseHeaders()
                    );
                    $e->setResponseObject($data);
                    break;
            }
            throw $e;
        }
    }

    /**
     * Operation isolateRepayListAsync
     *
     * list
     *
     * @param  string $symbol symbol (required)
     * @param  string $start_time startTime (required)
     * @param  string $coin coin (optional)
     * @param  string $repay_id repayId (optional)
     * @param  string $end_time endTime (optional)
     * @param  string $page_size pageSize (optional)
     * @param  string $page_id pageId (optional)
     * @param  string $contentType The value for the Content-Type header. Check self::contentTypes['isolateRepayList'] to see the possible values for this operation
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Promise\PromiseInterface
     */
    public function isolateRepayListAsync($symbol, $start_time, $coin = null, $repay_id = null, $end_time = null, $page_size = null, $page_id = null, string $contentType = self::contentTypes['isolateRepayList'][0])
    {
        return $this->isolateRepayListAsyncWithHttpInfo($symbol, $start_time, $coin, $repay_id, $end_time, $page_size, $page_id, $contentType)
            ->then(
                function ($response) {
                    return $response[0];
                }
            );
    }

    /**
     * Operation isolateRepayListAsyncWithHttpInfo
     *
     * list
     *
     * @param  string $symbol symbol (required)
     * @param  string $start_time startTime (required)
     * @param  string $coin coin (optional)
     * @param  string $repay_id repayId (optional)
     * @param  string $end_time endTime (optional)
     * @param  string $page_size pageSize (optional)
     * @param  string $page_id pageId (optional)
     * @param  string $contentType The value for the Content-Type header. Check self::contentTypes['isolateRepayList'] to see the possible values for this operation
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Promise\PromiseInterface
     */
    public function isolateRepayListAsyncWithHttpInfo($symbol, $start_time, $coin = null, $repay_id = null, $end_time = null, $page_size = null, $page_id = null, string $contentType = self::contentTypes['isolateRepayList'][0])
    {
        $returnType = '\Bitget\Model\ApiResponseResultOfMarginIsolatedRepayInfoResult';
        $request = $this->isolateRepayListRequest($symbol, $start_time, $coin, $repay_id, $end_time, $page_size, $page_id, $contentType);

        return $this->client
            ->sendAsync($request, $this->createHttpClientOption())
            ->then(
                function ($response) use ($returnType) {
                    if ($returnType === '\SplFileObject') {
                        $content = $response->getBody(); //stream goes to serializer
                    } else {
                        $content = (string) $response->getBody();
                        if ($returnType !== 'string') {
                            $content = json_decode($content);
                        }
                    }

                    return [
                        ObjectSerializer::deserialize($content, $returnType, []),
                        $response->getStatusCode(),
                        $response->getHeaders()
                    ];
                },
                function ($exception) {
                    $response = $exception->getResponse();
                    $statusCode = $response->getStatusCode();
                    throw new ApiException(
                        sprintf(
                            '[%d] Error connecting to the API (%s)',
                            $statusCode,
                            $exception->getRequest()->getUri()
                        ),
                        $statusCode,
                        $response->getHeaders(),
                        (string) $response->getBody()
                    );
                }
            );
    }

    /**
     * Create request for operation 'isolateRepayList'
     *
     * @param  string $symbol symbol (required)
     * @param  string $start_time startTime (required)
     * @param  string $coin coin (optional)
     * @param  string $repay_id repayId (optional)
     * @param  string $end_time endTime (optional)
     * @param  string $page_size pageSize (optional)
     * @param  string $page_id pageId (optional)
     * @param  string $contentType The value for the Content-Type header. Check self::contentTypes['isolateRepayList'] to see the possible values for this operation
     *
     * @throws \InvalidArgumentException
     * @return \GuzzleHttp\Psr7\Request
     */
    public function isolateRepayListRequest($symbol, $start_time, $coin = null, $repay_id = null, $end_time = null, $page_size = null, $page_id = null, string $contentType = self::contentTypes['isolateRepayList'][0])
    {

        // verify the required parameter 'symbol' is set
        if ($symbol === null || (is_array($symbol) && count($symbol) === 0)) {
            throw new \InvalidArgumentException(
                'Missing the required parameter $symbol when calling isolateRepayList'
            );
        }

        // verify the required parameter 'start_time' is set
        if ($start_time === null || (is_array($start_time) && count($start_time) === 0)) {
            throw new \InvalidArgumentException(
                'Missing the required parameter $start_time when calling isolateRepayList'
            );
        }







        $resourcePath = '/api/margin/v1/isolated/repay/list';
        $formParams = [];
        $queryParams = [];
        $headerParams = [];
        $httpBody = '';
        $multipart = false;

        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $symbol,
            'symbol', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            true // required
        ) ?? []);
        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $coin,
            'coin', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            false // required
        ) ?? []);
        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $repay_id,
            'repayId', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            false // required
        ) ?? []);
        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $start_time,
            'startTime', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            true // required
        ) ?? []);
        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $end_time,
            'endTime', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            false // required
        ) ?? []);
        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $page_size,
            'pageSize', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            false // required
        ) ?? []);
        // query params
        $queryParams = array_merge($queryParams, ObjectSerializer::toQueryValue(
            $page_id,
            'pageId', // param base name
            'string', // openApiType
            '', // style
            false, // explode
            false // required
        ) ?? []);




        $headers = $this->headerSelector->selectHeaders(
            ['application/json', ],
            $contentType,
            $multipart
        );

        // for model (json/xml)
        if (count($formParams) > 0) {
            if ($multipart) {
                $multipartContents = [];
                foreach ($formParams as $formParamName => $formParamValue) {
                    $formParamValueItems = is_array($formParamValue) ? $formParamValue : [$formParamValue];
                    foreach ($formParamValueItems as $formParamValueItem) {
                        $multipartContents[] = [
                            'name' => $formParamName,
                            'contents' => $formParamValueItem
                        ];
                    }
                }
                // for HTTP post (form)
                $httpBody = new MultipartStream($multipartContents);

            } elseif (stripos($headers['Content-Type'], 'application/json') !== false) {
                # if Content-Type contains "application/json", json_encode the form parameters
                $httpBody = \GuzzleHttp\json_encode($formParams);
            } else {
                // for HTTP post (form)
                $httpBody = ObjectSerializer::buildQuery($formParams);
            }
        }

        // this endpoint requires API key authentication
        $apiKey = $this->config->getApiKeyWithPrefix('ACCESS-KEY');
        if ($apiKey !== null) {
            $headers['ACCESS-KEY'] = $apiKey;
        }
        // this endpoint requires API key authentication
        $apiKey = $this->config->getApiKeyWithPrefix('ACCESS-PASSPHRASE');
        if ($apiKey !== null) {
            $headers['ACCESS-PASSPHRASE'] = $apiKey;
        }
        // this endpoint requires API key authentication
        $apiKey = $this->config->getApiKeyWithPrefix('ACCESS-SIGN');
        if ($apiKey !== null) {
            $headers['ACCESS-SIGN'] = $apiKey;
        }
        // this endpoint requires API key authentication
        $apiKey = $this->config->getApiKeyWithPrefix('ACCESS-TIMESTAMP');
        if ($apiKey !== null) {
            $headers['ACCESS-TIMESTAMP'] = $apiKey;
        }
        // this endpoint requires API key authentication
        $apiKey = $this->config->getApiKeyWithPrefix('SECRET-KEY');
        if ($apiKey !== null) {
            $headers['SECRET-KEY'] = $apiKey;
        }

        $defaultHeaders = [];
        if ($this->config->getUserAgent()) {
            $defaultHeaders['User-Agent'] = $this->config->getUserAgent();
        }

        $headers = array_merge(
            $defaultHeaders,
            $headerParams,
            $headers
        );

        $operationHost = $this->config->getHost();
        $query = ObjectSerializer::buildQuery($queryParams);
        return \Bitget\Utils::getAutoSignWarpHttpRequest(
            'GET',
            $operationHost . $resourcePath . ($query ? "?{$query}" : ''),
            $headers,
            $httpBody
        );
    }

    /**
     * Create http client option
     *
     * @throws \RuntimeException on file opening failure
     * @return array of http client options
     */
    protected function createHttpClientOption()
    {
        $options = [];
        if ($this->config->getDebug()) {
            $options[RequestOptions::DEBUG] = fopen($this->config->getDebugFile(), 'a');
            if (!$options[RequestOptions::DEBUG]) {
                throw new \RuntimeException('Failed to open the debug file: ' . $this->config->getDebugFile());
            }
        }

        return $options;
    }
}
