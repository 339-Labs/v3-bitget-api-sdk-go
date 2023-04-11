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


import localVarRequest from 'request';
import http from 'http';

/* tslint:disable:no-unused-locals */
import { ApiResponseResultOfMarginBatchCancelOrderResult } from '../model/apiResponseResultOfMarginBatchCancelOrderResult';
import { ApiResponseResultOfMarginBatchPlaceOrderResult } from '../model/apiResponseResultOfMarginBatchPlaceOrderResult';
import { ApiResponseResultOfMarginOpenOrderInfoResult } from '../model/apiResponseResultOfMarginOpenOrderInfoResult';
import { ApiResponseResultOfMarginPlaceOrderResult } from '../model/apiResponseResultOfMarginPlaceOrderResult';
import { ApiResponseResultOfMarginTradeDetailInfoResult } from '../model/apiResponseResultOfMarginTradeDetailInfoResult';
import { ApiResponseResultOfVoid } from '../model/apiResponseResultOfVoid';
import { MarginBatchCancelOrderRequest } from '../model/marginBatchCancelOrderRequest';
import { MarginBatchOrdersRequest } from '../model/marginBatchOrdersRequest';
import { MarginCancelOrderRequest } from '../model/marginCancelOrderRequest';
import { MarginOrderRequest } from '../model/marginOrderRequest';

import { ObjectSerializer, Authentication, VoidAuth, Interceptor } from '../model/models';
import { HttpBasicAuth, HttpBearerAuth, ApiKeyAuth, OAuth } from '../model/models';

import { HttpError, RequestFile } from './apis';

let defaultBasePath = 'https://api.bitget.com';

// ===============================================
// This file is autogenerated - Please do not edit
// ===============================================

export enum MarginCrossOrderApiApiKeys {
    ACCESS_KEY,
    ACCESS_PASSPHRASE,
    ACCESS_SIGN,
    ACCESS_TIMESTAMP,
    SECRET_KEY,
}

export class MarginCrossOrderApi {
    protected _basePath = defaultBasePath;
    protected _defaultHeaders : any = {};
    protected _useQuerystring : boolean = false;

    protected authentications = {
        'default': <Authentication>new VoidAuth(),
        'ACCESS_KEY': new ApiKeyAuth('header', 'ACCESS-KEY'),
        'ACCESS_PASSPHRASE': new ApiKeyAuth('header', 'ACCESS-PASSPHRASE'),
        'ACCESS_SIGN': new ApiKeyAuth('header', 'ACCESS-SIGN'),
        'ACCESS_TIMESTAMP': new ApiKeyAuth('header', 'ACCESS-TIMESTAMP'),
        'SECRET_KEY': new ApiKeyAuth('header', 'SECRET-KEY'),
    }

    protected interceptors: Interceptor[] = [];

    constructor(basePath?: string);
    constructor(basePathOrUsername: string, password?: string, basePath?: string) {
        if (password) {
            if (basePath) {
                this.basePath = basePath;
            }
        } else {
            if (basePathOrUsername) {
                this.basePath = basePathOrUsername
            }
        }
    }

    set useQuerystring(value: boolean) {
        this._useQuerystring = value;
    }

    set basePath(basePath: string) {
        this._basePath = basePath;
    }

    set defaultHeaders(defaultHeaders: any) {
        this._defaultHeaders = defaultHeaders;
    }

    get defaultHeaders() {
        return this._defaultHeaders;
    }

    get basePath() {
        return this._basePath;
    }

    public setDefaultAuthentication(auth: Authentication) {
        this.authentications.default = auth;
    }

    public setApiKey(key: MarginCrossOrderApiApiKeys, value: string) {
        (this.authentications as any)[MarginCrossOrderApiApiKeys[key]].apiKey = value;
    }

    public addInterceptor(interceptor: Interceptor) {
        this.interceptors.push(interceptor);
    }

    /**
     * Margin Cross BatchCancelOrder
     * @summary batchCancelOrder
     * @param marginBatchCancelOrderRequest marginBatchCancelOrderRequest
     */
    public async marginCrossBatchCancelOrder (marginBatchCancelOrderRequest: MarginBatchCancelOrderRequest, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginBatchCancelOrderResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/batchCancelOrder';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'marginBatchCancelOrderRequest' is not null or undefined
        if (marginBatchCancelOrderRequest === null || marginBatchCancelOrderRequest === undefined) {
            throw new Error('Required parameter marginBatchCancelOrderRequest was null or undefined when calling marginCrossBatchCancelOrder.');
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'POST',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
            body: ObjectSerializer.serialize(marginBatchCancelOrderRequest, "MarginBatchCancelOrderRequest")
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginBatchCancelOrderResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginBatchCancelOrderResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
    /**
     * Margin Cross PlaceOrder
     * @summary batchPlaceOrder
     * @param marginOrderRequest marginOrderRequest
     */
    public async marginCrossBatchPlaceOrder (marginOrderRequest: MarginBatchOrdersRequest, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginBatchPlaceOrderResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/batchPlaceOrder';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'marginOrderRequest' is not null or undefined
        if (marginOrderRequest === null || marginOrderRequest === undefined) {
            throw new Error('Required parameter marginOrderRequest was null or undefined when calling marginCrossBatchPlaceOrder.');
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'POST',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
            body: ObjectSerializer.serialize(marginOrderRequest, "MarginBatchOrdersRequest")
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginBatchPlaceOrderResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginBatchPlaceOrderResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
    /**
     * Margin Cross CancelOrder
     * @summary cancelOrder
     * @param marginCancelOrderRequest marginCancelOrderRequest
     */
    public async marginCrossCancelOrder (marginCancelOrderRequest: MarginCancelOrderRequest, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginBatchCancelOrderResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/cancelOrder';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'marginCancelOrderRequest' is not null or undefined
        if (marginCancelOrderRequest === null || marginCancelOrderRequest === undefined) {
            throw new Error('Required parameter marginCancelOrderRequest was null or undefined when calling marginCrossCancelOrder.');
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'POST',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
            body: ObjectSerializer.serialize(marginCancelOrderRequest, "MarginCancelOrderRequest")
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginBatchCancelOrderResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginBatchCancelOrderResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
    /**
     * Margin Cross Fills
     * @summary fills
     * @param symbol symbol
     * @param startTime startTime
     * @param source source
     * @param endTime endTime
     * @param orderId orderId
     * @param lastFillId lastFillId
     * @param pageSize pageSize
     */
    public async marginCrossFills (symbol: string, startTime: string, source?: string, endTime?: string, orderId?: string, lastFillId?: string, pageSize?: string, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginTradeDetailInfoResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/fills';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'symbol' is not null or undefined
        if (symbol === null || symbol === undefined) {
            throw new Error('Required parameter symbol was null or undefined when calling marginCrossFills.');
        }

        // verify required parameter 'startTime' is not null or undefined
        if (startTime === null || startTime === undefined) {
            throw new Error('Required parameter startTime was null or undefined when calling marginCrossFills.');
        }

        if (symbol !== undefined) {
            localVarQueryParameters['symbol'] = ObjectSerializer.serialize(symbol, "string");
        }

        if (source !== undefined) {
            localVarQueryParameters['source'] = ObjectSerializer.serialize(source, "string");
        }

        if (startTime !== undefined) {
            localVarQueryParameters['startTime'] = ObjectSerializer.serialize(startTime, "string");
        }

        if (endTime !== undefined) {
            localVarQueryParameters['endTime'] = ObjectSerializer.serialize(endTime, "string");
        }

        if (orderId !== undefined) {
            localVarQueryParameters['orderId'] = ObjectSerializer.serialize(orderId, "string");
        }

        if (lastFillId !== undefined) {
            localVarQueryParameters['lastFillId'] = ObjectSerializer.serialize(lastFillId, "string");
        }

        if (pageSize !== undefined) {
            localVarQueryParameters['pageSize'] = ObjectSerializer.serialize(pageSize, "string");
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'GET',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginTradeDetailInfoResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginTradeDetailInfoResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
    /**
     * Margin Cross historyOrders
     * @summary history
     * @param symbol symbol
     * @param startTime startTime
     * @param source source
     * @param endTime endTime
     * @param orderId orderId
     * @param clientOid clientOid
     * @param minId minId
     * @param pageSize pageSize
     */
    public async marginCrossHistoryOrders (symbol: string, startTime: string, source?: string, endTime?: string, orderId?: string, clientOid?: string, minId?: string, pageSize?: string, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginOpenOrderInfoResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/history';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'symbol' is not null or undefined
        if (symbol === null || symbol === undefined) {
            throw new Error('Required parameter symbol was null or undefined when calling marginCrossHistoryOrders.');
        }

        // verify required parameter 'startTime' is not null or undefined
        if (startTime === null || startTime === undefined) {
            throw new Error('Required parameter startTime was null or undefined when calling marginCrossHistoryOrders.');
        }

        if (symbol !== undefined) {
            localVarQueryParameters['symbol'] = ObjectSerializer.serialize(symbol, "string");
        }

        if (source !== undefined) {
            localVarQueryParameters['source'] = ObjectSerializer.serialize(source, "string");
        }

        if (startTime !== undefined) {
            localVarQueryParameters['startTime'] = ObjectSerializer.serialize(startTime, "string");
        }

        if (endTime !== undefined) {
            localVarQueryParameters['endTime'] = ObjectSerializer.serialize(endTime, "string");
        }

        if (orderId !== undefined) {
            localVarQueryParameters['orderId'] = ObjectSerializer.serialize(orderId, "string");
        }

        if (clientOid !== undefined) {
            localVarQueryParameters['clientOid'] = ObjectSerializer.serialize(clientOid, "string");
        }

        if (minId !== undefined) {
            localVarQueryParameters['minId'] = ObjectSerializer.serialize(minId, "string");
        }

        if (pageSize !== undefined) {
            localVarQueryParameters['pageSize'] = ObjectSerializer.serialize(pageSize, "string");
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'GET',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginOpenOrderInfoResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginOpenOrderInfoResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
    /**
     * Margin Cross openOrders
     * @summary openOrders
     * @param symbol symbol
     * @param startTime startTime
     * @param endTime endTime
     * @param orderId orderId
     * @param clientOid clientOid
     * @param pageSize pageSize
     */
    public async marginCrossOpenOrders (symbol: string, startTime: string, endTime?: string, orderId?: string, clientOid?: string, pageSize?: string, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginOpenOrderInfoResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/openOrders';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'symbol' is not null or undefined
        if (symbol === null || symbol === undefined) {
            throw new Error('Required parameter symbol was null or undefined when calling marginCrossOpenOrders.');
        }

        // verify required parameter 'startTime' is not null or undefined
        if (startTime === null || startTime === undefined) {
            throw new Error('Required parameter startTime was null or undefined when calling marginCrossOpenOrders.');
        }

        if (symbol !== undefined) {
            localVarQueryParameters['symbol'] = ObjectSerializer.serialize(symbol, "string");
        }

        if (startTime !== undefined) {
            localVarQueryParameters['startTime'] = ObjectSerializer.serialize(startTime, "string");
        }

        if (endTime !== undefined) {
            localVarQueryParameters['endTime'] = ObjectSerializer.serialize(endTime, "string");
        }

        if (orderId !== undefined) {
            localVarQueryParameters['orderId'] = ObjectSerializer.serialize(orderId, "string");
        }

        if (clientOid !== undefined) {
            localVarQueryParameters['clientOid'] = ObjectSerializer.serialize(clientOid, "string");
        }

        if (pageSize !== undefined) {
            localVarQueryParameters['pageSize'] = ObjectSerializer.serialize(pageSize, "string");
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'GET',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginOpenOrderInfoResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginOpenOrderInfoResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
    /**
     * Margin Cross PlaceOrder
     * @summary placeOrder
     * @param marginOrderRequest marginOrderRequest
     */
    public async marginCrossPlaceOrder (marginOrderRequest: MarginOrderRequest, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginPlaceOrderResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/order/placeOrder';
        let localVarQueryParameters: any = {};
        let localVarHeaderParams: any = (<any>Object).assign({}, this._defaultHeaders);
        const produces = ['application/json'];
        // give precedence to 'application/json'
        if (produces.indexOf('application/json') >= 0) {
            localVarHeaderParams.Accept = 'application/json';
        } else {
            localVarHeaderParams.Accept = produces.join(',');
        }
        let localVarFormParams: any = {};

        // verify required parameter 'marginOrderRequest' is not null or undefined
        if (marginOrderRequest === null || marginOrderRequest === undefined) {
            throw new Error('Required parameter marginOrderRequest was null or undefined when calling marginCrossPlaceOrder.');
        }

        (<any>Object).assign(localVarHeaderParams, options.headers);

        let localVarUseFormData = false;

        let localVarRequestOptions: localVarRequest.Options = {
            method: 'POST',
            qs: localVarQueryParameters,
            headers: localVarHeaderParams,
            uri: localVarPath,
            useQuerystring: this._useQuerystring,
            json: true,
            body: ObjectSerializer.serialize(marginOrderRequest, "MarginOrderRequest")
        };

        let authenticationPromise = Promise.resolve();
        if (this.authentications.ACCESS_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_KEY.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_PASSPHRASE.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_PASSPHRASE.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_SIGN.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_SIGN.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.ACCESS_TIMESTAMP.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.ACCESS_TIMESTAMP.applyToRequest(localVarRequestOptions));
        }
        if (this.authentications.SECRET_KEY.apiKey) {
            authenticationPromise = authenticationPromise.then(() => this.authentications.SECRET_KEY.applyToRequest(localVarRequestOptions));
        }
        authenticationPromise = authenticationPromise.then(() => this.authentications.default.applyToRequest(localVarRequestOptions));

        let interceptorPromise = authenticationPromise;
        for (const interceptor of this.interceptors) {
            interceptorPromise = interceptorPromise.then(() => interceptor(localVarRequestOptions));
        }

        return interceptorPromise.then(() => {
            if (Object.keys(localVarFormParams).length) {
                if (localVarUseFormData) {
                    (<any>localVarRequestOptions).formData = localVarFormParams;
                } else {
                    localVarRequestOptions.form = localVarFormParams;
                }
            }
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginPlaceOrderResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginPlaceOrderResult");
                            resolve({ response: response, body: body });
                        } else {
                            reject(new HttpError(response, body, response.statusCode));
                        }
                    }
                });
            });
        });
    }
}
