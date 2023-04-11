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


class MarginCrossRateAndLimitResult(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        
        class properties:
            borrowAble = schemas.BoolSchema
            coin = schemas.StrSchema
            dailyInterestRate = schemas.StrSchema
            leverage = schemas.StrSchema
            maxBorrowableAmount = schemas.StrSchema
            transferInAble = schemas.BoolSchema
            
            
            class vips(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['MarginCrossVipResult']:
                        return MarginCrossVipResult
            
                def __new__(
                    cls,
                    arg: typing.Union[typing.Tuple['MarginCrossVipResult'], typing.List['MarginCrossVipResult']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'vips':
                    return super().__new__(
                        cls,
                        arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'MarginCrossVipResult':
                    return super().__getitem__(i)
            yearlyInterestRate = schemas.StrSchema
            __annotations__ = {
                "borrowAble": borrowAble,
                "coin": coin,
                "dailyInterestRate": dailyInterestRate,
                "leverage": leverage,
                "maxBorrowableAmount": maxBorrowableAmount,
                "transferInAble": transferInAble,
                "vips": vips,
                "yearlyInterestRate": yearlyInterestRate,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["borrowAble"]) -> MetaOapg.properties.borrowAble: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["coin"]) -> MetaOapg.properties.coin: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["dailyInterestRate"]) -> MetaOapg.properties.dailyInterestRate: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["leverage"]) -> MetaOapg.properties.leverage: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["maxBorrowableAmount"]) -> MetaOapg.properties.maxBorrowableAmount: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["transferInAble"]) -> MetaOapg.properties.transferInAble: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["vips"]) -> MetaOapg.properties.vips: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["yearlyInterestRate"]) -> MetaOapg.properties.yearlyInterestRate: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["borrowAble", "coin", "dailyInterestRate", "leverage", "maxBorrowableAmount", "transferInAble", "vips", "yearlyInterestRate", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["borrowAble"]) -> typing.Union[MetaOapg.properties.borrowAble, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["coin"]) -> typing.Union[MetaOapg.properties.coin, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["dailyInterestRate"]) -> typing.Union[MetaOapg.properties.dailyInterestRate, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["leverage"]) -> typing.Union[MetaOapg.properties.leverage, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["maxBorrowableAmount"]) -> typing.Union[MetaOapg.properties.maxBorrowableAmount, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["transferInAble"]) -> typing.Union[MetaOapg.properties.transferInAble, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["vips"]) -> typing.Union[MetaOapg.properties.vips, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["yearlyInterestRate"]) -> typing.Union[MetaOapg.properties.yearlyInterestRate, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["borrowAble", "coin", "dailyInterestRate", "leverage", "maxBorrowableAmount", "transferInAble", "vips", "yearlyInterestRate", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        borrowAble: typing.Union[MetaOapg.properties.borrowAble, bool, schemas.Unset] = schemas.unset,
        coin: typing.Union[MetaOapg.properties.coin, str, schemas.Unset] = schemas.unset,
        dailyInterestRate: typing.Union[MetaOapg.properties.dailyInterestRate, str, schemas.Unset] = schemas.unset,
        leverage: typing.Union[MetaOapg.properties.leverage, str, schemas.Unset] = schemas.unset,
        maxBorrowableAmount: typing.Union[MetaOapg.properties.maxBorrowableAmount, str, schemas.Unset] = schemas.unset,
        transferInAble: typing.Union[MetaOapg.properties.transferInAble, bool, schemas.Unset] = schemas.unset,
        vips: typing.Union[MetaOapg.properties.vips, list, tuple, schemas.Unset] = schemas.unset,
        yearlyInterestRate: typing.Union[MetaOapg.properties.yearlyInterestRate, str, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'MarginCrossRateAndLimitResult':
        return super().__new__(
            cls,
            *args,
            borrowAble=borrowAble,
            coin=coin,
            dailyInterestRate=dailyInterestRate,
            leverage=leverage,
            maxBorrowableAmount=maxBorrowableAmount,
            transferInAble=transferInAble,
            vips=vips,
            yearlyInterestRate=yearlyInterestRate,
            _configuration=_configuration,
            **kwargs,
        )

from bitget.model.margin_cross_vip_result import MarginCrossVipResult
