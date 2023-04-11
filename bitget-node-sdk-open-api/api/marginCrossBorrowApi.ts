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
import { ApiResponseResultOfMarginLoanInfoResult } from '../model/apiResponseResultOfMarginLoanInfoResult';
import { ApiResponseResultOfVoid } from '../model/apiResponseResultOfVoid';

import { ObjectSerializer, Authentication, VoidAuth, Interceptor } from '../model/models';
import { HttpBasicAuth, HttpBearerAuth, ApiKeyAuth, OAuth } from '../model/models';

import { HttpError, RequestFile } from './apis';

let defaultBasePath = 'https://api.bitget.com';

// ===============================================
// This file is autogenerated - Please do not edit
// ===============================================

export enum MarginCrossBorrowApiApiKeys {
    ACCESS_KEY,
    ACCESS_PASSPHRASE,
    ACCESS_SIGN,
    ACCESS_TIMESTAMP,
    SECRET_KEY,
}

export class MarginCrossBorrowApi {
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

    public setApiKey(key: MarginCrossBorrowApiApiKeys, value: string) {
        (this.authentications as any)[MarginCrossBorrowApiApiKeys[key]].apiKey = value;
    }

    public addInterceptor(interceptor: Interceptor) {
        this.interceptors.push(interceptor);
    }

    /**
     * Get Loan List
     * @summary list
     * @param startTime startTime
     * @param coin coin
     * @param endTime endTime
     * @param loanId loanId
     * @param pageSize pageSize
     * @param pageId pageId
     */
    public async crossLoanList (startTime: string, coin?: string, endTime?: string, loanId?: string, pageSize?: string, pageId?: string, options: {headers: {[name: string]: string}} = {headers: {}}) : Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginLoanInfoResult;  }> {
        const localVarPath = this.basePath + '/api/margin/v1/cross/loan/list';
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

        // verify required parameter 'startTime' is not null or undefined
        if (startTime === null || startTime === undefined) {
            throw new Error('Required parameter startTime was null or undefined when calling crossLoanList.');
        }

        if (coin !== undefined) {
            localVarQueryParameters['coin'] = ObjectSerializer.serialize(coin, "string");
        }

        if (startTime !== undefined) {
            localVarQueryParameters['startTime'] = ObjectSerializer.serialize(startTime, "string");
        }

        if (endTime !== undefined) {
            localVarQueryParameters['endTime'] = ObjectSerializer.serialize(endTime, "string");
        }

        if (loanId !== undefined) {
            localVarQueryParameters['loanId'] = ObjectSerializer.serialize(loanId, "string");
        }

        if (pageSize !== undefined) {
            localVarQueryParameters['pageSize'] = ObjectSerializer.serialize(pageSize, "string");
        }

        if (pageId !== undefined) {
            localVarQueryParameters['pageId'] = ObjectSerializer.serialize(pageId, "string");
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
            return new Promise<{ response: http.IncomingMessage; body: ApiResponseResultOfMarginLoanInfoResult;  }>((resolve, reject) => {
                localVarRequest(localVarRequestOptions, (error, response, body) => {
                    if (error) {
                        reject(error);
                    } else {
                        if (response.statusCode && response.statusCode >= 200 && response.statusCode <= 299) {
                            body = ObjectSerializer.deserialize(body, "ApiResponseResultOfMarginLoanInfoResult");
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
