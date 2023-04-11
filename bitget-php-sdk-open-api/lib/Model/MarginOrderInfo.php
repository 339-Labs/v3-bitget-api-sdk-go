<?php
/**
 * MarginOrderInfo
 *
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

namespace Bitget\Model;

use \ArrayAccess;
use \Bitget\ObjectSerializer;

/**
 * MarginOrderInfo Class Doc Comment
 *
 * @category Class
 * @package  Bitget
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<string, mixed>
 */
class MarginOrderInfo implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = null;

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'MarginOrderInfo';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'base_quantity' => 'string',
        'client_oid' => 'string',
        'ctime' => 'string',
        'fill_price' => 'string',
        'fill_quantity' => 'string',
        'fill_total_amount' => 'string',
        'loan_type' => 'string',
        'order_id' => 'string',
        'order_type' => 'string',
        'price' => 'string',
        'quote_amount' => 'string',
        'side' => 'string',
        'source' => 'string',
        'status' => 'string',
        'symbol' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'base_quantity' => null,
        'client_oid' => null,
        'ctime' => null,
        'fill_price' => null,
        'fill_quantity' => null,
        'fill_total_amount' => null,
        'loan_type' => null,
        'order_id' => null,
        'order_type' => null,
        'price' => null,
        'quote_amount' => null,
        'side' => null,
        'source' => null,
        'status' => null,
        'symbol' => null
    ];

    /**
      * Array of nullable properties. Used for (de)serialization
      *
      * @var boolean[]
      */
    protected static array $openAPINullables = [
        'base_quantity' => false,
		'client_oid' => false,
		'ctime' => false,
		'fill_price' => false,
		'fill_quantity' => false,
		'fill_total_amount' => false,
		'loan_type' => false,
		'order_id' => false,
		'order_type' => false,
		'price' => false,
		'quote_amount' => false,
		'side' => false,
		'source' => false,
		'status' => false,
		'symbol' => false
    ];

    /**
      * If a nullable field gets set to null, insert it here
      *
      * @var boolean[]
      */
    protected array $openAPINullablesSetToNull = [];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of nullable properties
     *
     * @return array
     */
    protected static function openAPINullables(): array
    {
        return self::$openAPINullables;
    }

    /**
     * Array of nullable field names deliberately set to null
     *
     * @return boolean[]
     */
    private function getOpenAPINullablesSetToNull(): array
    {
        return $this->openAPINullablesSetToNull;
    }

    /**
     * Setter - Array of nullable field names deliberately set to null
     *
     * @param boolean[] $openAPINullablesSetToNull
     */
    private function setOpenAPINullablesSetToNull(array $openAPINullablesSetToNull): void
    {
        $this->openAPINullablesSetToNull = $openAPINullablesSetToNull;
    }

    /**
     * Checks if a property is nullable
     *
     * @param string $property
     * @return bool
     */
    public static function isNullable(string $property): bool
    {
        return self::openAPINullables()[$property] ?? false;
    }

    /**
     * Checks if a nullable property is set to null.
     *
     * @param string $property
     * @return bool
     */
    public function isNullableSetToNull(string $property): bool
    {
        return in_array($property, $this->getOpenAPINullablesSetToNull(), true);
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'base_quantity' => 'baseQuantity',
        'client_oid' => 'clientOid',
        'ctime' => 'ctime',
        'fill_price' => 'fillPrice',
        'fill_quantity' => 'fillQuantity',
        'fill_total_amount' => 'fillTotalAmount',
        'loan_type' => 'loanType',
        'order_id' => 'orderId',
        'order_type' => 'orderType',
        'price' => 'price',
        'quote_amount' => 'quoteAmount',
        'side' => 'side',
        'source' => 'source',
        'status' => 'status',
        'symbol' => 'symbol'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'base_quantity' => 'setBaseQuantity',
        'client_oid' => 'setClientOid',
        'ctime' => 'setCtime',
        'fill_price' => 'setFillPrice',
        'fill_quantity' => 'setFillQuantity',
        'fill_total_amount' => 'setFillTotalAmount',
        'loan_type' => 'setLoanType',
        'order_id' => 'setOrderId',
        'order_type' => 'setOrderType',
        'price' => 'setPrice',
        'quote_amount' => 'setQuoteAmount',
        'side' => 'setSide',
        'source' => 'setSource',
        'status' => 'setStatus',
        'symbol' => 'setSymbol'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'base_quantity' => 'getBaseQuantity',
        'client_oid' => 'getClientOid',
        'ctime' => 'getCtime',
        'fill_price' => 'getFillPrice',
        'fill_quantity' => 'getFillQuantity',
        'fill_total_amount' => 'getFillTotalAmount',
        'loan_type' => 'getLoanType',
        'order_id' => 'getOrderId',
        'order_type' => 'getOrderType',
        'price' => 'getPrice',
        'quote_amount' => 'getQuoteAmount',
        'side' => 'getSide',
        'source' => 'getSource',
        'status' => 'getStatus',
        'symbol' => 'getSymbol'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->setIfExists('base_quantity', $data ?? [], null);
        $this->setIfExists('client_oid', $data ?? [], null);
        $this->setIfExists('ctime', $data ?? [], null);
        $this->setIfExists('fill_price', $data ?? [], null);
        $this->setIfExists('fill_quantity', $data ?? [], null);
        $this->setIfExists('fill_total_amount', $data ?? [], null);
        $this->setIfExists('loan_type', $data ?? [], null);
        $this->setIfExists('order_id', $data ?? [], null);
        $this->setIfExists('order_type', $data ?? [], null);
        $this->setIfExists('price', $data ?? [], null);
        $this->setIfExists('quote_amount', $data ?? [], null);
        $this->setIfExists('side', $data ?? [], null);
        $this->setIfExists('source', $data ?? [], null);
        $this->setIfExists('status', $data ?? [], null);
        $this->setIfExists('symbol', $data ?? [], null);
    }

    /**
    * Sets $this->container[$variableName] to the given data or to the given default Value; if $variableName
    * is nullable and its value is set to null in the $fields array, then mark it as "set to null" in the
    * $this->openAPINullablesSetToNull array
    *
    * @param string $variableName
    * @param array  $fields
    * @param mixed  $defaultValue
    */
    private function setIfExists(string $variableName, array $fields, $defaultValue): void
    {
        if (self::isNullable($variableName) && array_key_exists($variableName, $fields) && is_null($fields[$variableName])) {
            $this->openAPINullablesSetToNull[] = $variableName;
        }

        $this->container[$variableName] = $fields[$variableName] ?? $defaultValue;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets base_quantity
     *
     * @return string|null
     */
    public function getBaseQuantity()
    {
        return $this->container['base_quantity'];
    }

    /**
     * Sets base_quantity
     *
     * @param string|null $base_quantity base_quantity
     *
     * @return self
     */
    public function setBaseQuantity($base_quantity)
    {

        if (is_null($base_quantity)) {
            throw new \InvalidArgumentException('non-nullable base_quantity cannot be null');
        }

        $this->container['base_quantity'] = $base_quantity;

        return $this;
    }

    /**
     * Gets client_oid
     *
     * @return string|null
     */
    public function getClientOid()
    {
        return $this->container['client_oid'];
    }

    /**
     * Sets client_oid
     *
     * @param string|null $client_oid client_oid
     *
     * @return self
     */
    public function setClientOid($client_oid)
    {

        if (is_null($client_oid)) {
            throw new \InvalidArgumentException('non-nullable client_oid cannot be null');
        }

        $this->container['client_oid'] = $client_oid;

        return $this;
    }

    /**
     * Gets ctime
     *
     * @return string|null
     */
    public function getCtime()
    {
        return $this->container['ctime'];
    }

    /**
     * Sets ctime
     *
     * @param string|null $ctime ctime
     *
     * @return self
     */
    public function setCtime($ctime)
    {

        if (is_null($ctime)) {
            throw new \InvalidArgumentException('non-nullable ctime cannot be null');
        }

        $this->container['ctime'] = $ctime;

        return $this;
    }

    /**
     * Gets fill_price
     *
     * @return string|null
     */
    public function getFillPrice()
    {
        return $this->container['fill_price'];
    }

    /**
     * Sets fill_price
     *
     * @param string|null $fill_price fill_price
     *
     * @return self
     */
    public function setFillPrice($fill_price)
    {

        if (is_null($fill_price)) {
            throw new \InvalidArgumentException('non-nullable fill_price cannot be null');
        }

        $this->container['fill_price'] = $fill_price;

        return $this;
    }

    /**
     * Gets fill_quantity
     *
     * @return string|null
     */
    public function getFillQuantity()
    {
        return $this->container['fill_quantity'];
    }

    /**
     * Sets fill_quantity
     *
     * @param string|null $fill_quantity fill_quantity
     *
     * @return self
     */
    public function setFillQuantity($fill_quantity)
    {

        if (is_null($fill_quantity)) {
            throw new \InvalidArgumentException('non-nullable fill_quantity cannot be null');
        }

        $this->container['fill_quantity'] = $fill_quantity;

        return $this;
    }

    /**
     * Gets fill_total_amount
     *
     * @return string|null
     */
    public function getFillTotalAmount()
    {
        return $this->container['fill_total_amount'];
    }

    /**
     * Sets fill_total_amount
     *
     * @param string|null $fill_total_amount fill_total_amount
     *
     * @return self
     */
    public function setFillTotalAmount($fill_total_amount)
    {

        if (is_null($fill_total_amount)) {
            throw new \InvalidArgumentException('non-nullable fill_total_amount cannot be null');
        }

        $this->container['fill_total_amount'] = $fill_total_amount;

        return $this;
    }

    /**
     * Gets loan_type
     *
     * @return string|null
     */
    public function getLoanType()
    {
        return $this->container['loan_type'];
    }

    /**
     * Sets loan_type
     *
     * @param string|null $loan_type loan_type
     *
     * @return self
     */
    public function setLoanType($loan_type)
    {

        if (is_null($loan_type)) {
            throw new \InvalidArgumentException('non-nullable loan_type cannot be null');
        }

        $this->container['loan_type'] = $loan_type;

        return $this;
    }

    /**
     * Gets order_id
     *
     * @return string|null
     */
    public function getOrderId()
    {
        return $this->container['order_id'];
    }

    /**
     * Sets order_id
     *
     * @param string|null $order_id order_id
     *
     * @return self
     */
    public function setOrderId($order_id)
    {

        if (is_null($order_id)) {
            throw new \InvalidArgumentException('non-nullable order_id cannot be null');
        }

        $this->container['order_id'] = $order_id;

        return $this;
    }

    /**
     * Gets order_type
     *
     * @return string|null
     */
    public function getOrderType()
    {
        return $this->container['order_type'];
    }

    /**
     * Sets order_type
     *
     * @param string|null $order_type order_type
     *
     * @return self
     */
    public function setOrderType($order_type)
    {

        if (is_null($order_type)) {
            throw new \InvalidArgumentException('non-nullable order_type cannot be null');
        }

        $this->container['order_type'] = $order_type;

        return $this;
    }

    /**
     * Gets price
     *
     * @return string|null
     */
    public function getPrice()
    {
        return $this->container['price'];
    }

    /**
     * Sets price
     *
     * @param string|null $price price
     *
     * @return self
     */
    public function setPrice($price)
    {

        if (is_null($price)) {
            throw new \InvalidArgumentException('non-nullable price cannot be null');
        }

        $this->container['price'] = $price;

        return $this;
    }

    /**
     * Gets quote_amount
     *
     * @return string|null
     */
    public function getQuoteAmount()
    {
        return $this->container['quote_amount'];
    }

    /**
     * Sets quote_amount
     *
     * @param string|null $quote_amount quote_amount
     *
     * @return self
     */
    public function setQuoteAmount($quote_amount)
    {

        if (is_null($quote_amount)) {
            throw new \InvalidArgumentException('non-nullable quote_amount cannot be null');
        }

        $this->container['quote_amount'] = $quote_amount;

        return $this;
    }

    /**
     * Gets side
     *
     * @return string|null
     */
    public function getSide()
    {
        return $this->container['side'];
    }

    /**
     * Sets side
     *
     * @param string|null $side side
     *
     * @return self
     */
    public function setSide($side)
    {

        if (is_null($side)) {
            throw new \InvalidArgumentException('non-nullable side cannot be null');
        }

        $this->container['side'] = $side;

        return $this;
    }

    /**
     * Gets source
     *
     * @return string|null
     */
    public function getSource()
    {
        return $this->container['source'];
    }

    /**
     * Sets source
     *
     * @param string|null $source source
     *
     * @return self
     */
    public function setSource($source)
    {

        if (is_null($source)) {
            throw new \InvalidArgumentException('non-nullable source cannot be null');
        }

        $this->container['source'] = $source;

        return $this;
    }

    /**
     * Gets status
     *
     * @return string|null
     */
    public function getStatus()
    {
        return $this->container['status'];
    }

    /**
     * Sets status
     *
     * @param string|null $status status
     *
     * @return self
     */
    public function setStatus($status)
    {

        if (is_null($status)) {
            throw new \InvalidArgumentException('non-nullable status cannot be null');
        }

        $this->container['status'] = $status;

        return $this;
    }

    /**
     * Gets symbol
     *
     * @return string|null
     */
    public function getSymbol()
    {
        return $this->container['symbol'];
    }

    /**
     * Sets symbol
     *
     * @param string|null $symbol symbol
     *
     * @return self
     */
    public function setSymbol($symbol)
    {

        if (is_null($symbol)) {
            throw new \InvalidArgumentException('non-nullable symbol cannot be null');
        }

        $this->container['symbol'] = $symbol;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset): bool
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    #[\ReturnTypeWillChange]
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value): void
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset): void
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    #[\ReturnTypeWillChange]
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


