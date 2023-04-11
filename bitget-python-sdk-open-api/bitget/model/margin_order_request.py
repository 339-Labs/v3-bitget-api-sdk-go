# coding: utf-8

"""
    Bitget Open API

    No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)  # noqa: E501

    The version of the OpenAPI document: 2.0.0
    Generated by: https://openapi-generator.tech
"""

from datetime import date, datetime  # noqa: F401
import decimal  # noqa: F401
import functools  # noqa: F401
import io  # noqa: F401
import re  # noqa: F401
import typing  # noqa: F401
import typing_extensions  # noqa: F401
import uuid  # noqa: F401

import frozendict  # noqa: F401

from bitget import schemas  # noqa: F401


class MarginOrderRequest(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        required = {
            "orderType",
            "symbol",
            "side",
            "loanType",
        }
        
        class properties:
            loanType = schemas.StrSchema
            orderType = schemas.StrSchema
            side = schemas.StrSchema
            symbol = schemas.StrSchema
            baseQuantity = schemas.StrSchema
            channelApiCode = schemas.StrSchema
            clientOid = schemas.StrSchema
            ip = schemas.StrSchema
            price = schemas.StrSchema
            quoteAmount = schemas.StrSchema
            timeInForce = schemas.StrSchema
            __annotations__ = {
                "loanType": loanType,
                "orderType": orderType,
                "side": side,
                "symbol": symbol,
                "baseQuantity": baseQuantity,
                "channelApiCode": channelApiCode,
                "clientOid": clientOid,
                "ip": ip,
                "price": price,
                "quoteAmount": quoteAmount,
                "timeInForce": timeInForce,
            }
    
    orderType: MetaOapg.properties.orderType
    symbol: MetaOapg.properties.symbol
    side: MetaOapg.properties.side
    loanType: MetaOapg.properties.loanType
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["loanType"]) -> MetaOapg.properties.loanType: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["orderType"]) -> MetaOapg.properties.orderType: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["side"]) -> MetaOapg.properties.side: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["symbol"]) -> MetaOapg.properties.symbol: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["baseQuantity"]) -> MetaOapg.properties.baseQuantity: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["channelApiCode"]) -> MetaOapg.properties.channelApiCode: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["clientOid"]) -> MetaOapg.properties.clientOid: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["ip"]) -> MetaOapg.properties.ip: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["price"]) -> MetaOapg.properties.price: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["quoteAmount"]) -> MetaOapg.properties.quoteAmount: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["timeInForce"]) -> MetaOapg.properties.timeInForce: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["loanType", "orderType", "side", "symbol", "baseQuantity", "channelApiCode", "clientOid", "ip", "price", "quoteAmount", "timeInForce", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["loanType"]) -> MetaOapg.properties.loanType: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["orderType"]) -> MetaOapg.properties.orderType: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["side"]) -> MetaOapg.properties.side: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["symbol"]) -> MetaOapg.properties.symbol: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["baseQuantity"]) -> typing.Union[MetaOapg.properties.baseQuantity, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["channelApiCode"]) -> typing.Union[MetaOapg.properties.channelApiCode, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["clientOid"]) -> typing.Union[MetaOapg.properties.clientOid, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["ip"]) -> typing.Union[MetaOapg.properties.ip, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["price"]) -> typing.Union[MetaOapg.properties.price, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["quoteAmount"]) -> typing.Union[MetaOapg.properties.quoteAmount, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["timeInForce"]) -> typing.Union[MetaOapg.properties.timeInForce, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["loanType", "orderType", "side", "symbol", "baseQuantity", "channelApiCode", "clientOid", "ip", "price", "quoteAmount", "timeInForce", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        orderType: typing.Union[MetaOapg.properties.orderType, str, ],
        symbol: typing.Union[MetaOapg.properties.symbol, str, ],
        side: typing.Union[MetaOapg.properties.side, str, ],
        loanType: typing.Union[MetaOapg.properties.loanType, str, ],
        baseQuantity: typing.Union[MetaOapg.properties.baseQuantity, str, schemas.Unset] = schemas.unset,
        channelApiCode: typing.Union[MetaOapg.properties.channelApiCode, str, schemas.Unset] = schemas.unset,
        clientOid: typing.Union[MetaOapg.properties.clientOid, str, schemas.Unset] = schemas.unset,
        ip: typing.Union[MetaOapg.properties.ip, str, schemas.Unset] = schemas.unset,
        price: typing.Union[MetaOapg.properties.price, str, schemas.Unset] = schemas.unset,
        quoteAmount: typing.Union[MetaOapg.properties.quoteAmount, str, schemas.Unset] = schemas.unset,
        timeInForce: typing.Union[MetaOapg.properties.timeInForce, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'MarginOrderRequest':
        return super().__new__(
            cls,
            *args,
            orderType=orderType,
            symbol=symbol,
            side=side,
            loanType=loanType,
            baseQuantity=baseQuantity,
            channelApiCode=channelApiCode,
            clientOid=clientOid,
            ip=ip,
            price=price,
            quoteAmount=quoteAmount,
            timeInForce=timeInForce,
            _configuration=_configuration,
            **kwargs,
        )
