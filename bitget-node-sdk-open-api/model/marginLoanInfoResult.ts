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
import { MarginLoanInfo } from './marginLoanInfo';

export class MarginLoanInfoResult {
    'maxId'?: string;
    'minId'?: string;
    'resultList'?: Array<MarginLoanInfo>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "maxId",
            "baseName": "maxId",
            "type": "string"
        },
        {
            "name": "minId",
            "baseName": "minId",
            "type": "string"
        },
        {
            "name": "resultList",
            "baseName": "resultList",
            "type": "Array<MarginLoanInfo>"
        }    ];

    static getAttributeTypeMap() {
        return MarginLoanInfoResult.attributeTypeMap;
    }
}

