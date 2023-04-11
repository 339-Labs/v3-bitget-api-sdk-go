/**
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

import { RequestFile } from './models';
import { MarginOrderRequest } from './marginOrderRequest';

export class MarginBatchOrdersRequest {
    'channelApiCode'?: string;
    'ip'?: string;
    'orderList'?: Array<MarginOrderRequest>;
    /**
    * symbol
    */
    'symbol': string;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "channelApiCode",
            "baseName": "channelApiCode",
            "type": "string"
        },
        {
            "name": "ip",
            "baseName": "ip",
            "type": "string"
        },
        {
            "name": "orderList",
            "baseName": "orderList",
            "type": "Array<MarginOrderRequest>"
        },
        {
            "name": "symbol",
            "baseName": "symbol",
            "type": "string"
        }    ];

    static getAttributeTypeMap() {
        return MarginBatchOrdersRequest.attributeTypeMap;
    }
}

