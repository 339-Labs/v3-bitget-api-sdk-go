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


class MarginLoanInfoResult(
    schemas.DictSchema
):
    """NOTE: This class is auto generated by OpenAPI Generator.
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """


    class MetaOapg:
        
        class properties:
            maxId = schemas.StrSchema
            minId = schemas.StrSchema
            
            
            class resultList(
                schemas.ListSchema
            ):
            
            
                class MetaOapg:
                    
                    @staticmethod
                    def items() -> typing.Type['MarginLoanInfo']:
                        return MarginLoanInfo
            
                def __new__(
                    cls,
                    arg: typing.Union[typing.Tuple['MarginLoanInfo'], typing.List['MarginLoanInfo']],
                    _configuration: typing.Optional[schemas.Configuration] = None,
                ) -> 'resultList':
                    return super().__new__(
                        cls,
                        arg,
                        _configuration=_configuration,
                    )
            
                def __getitem__(self, i: int) -> 'MarginLoanInfo':
                    return super().__getitem__(i)
            __annotations__ = {
                "maxId": maxId,
                "minId": minId,
                "resultList": resultList,
            }
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["maxId"]) -> MetaOapg.properties.maxId: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["minId"]) -> MetaOapg.properties.minId: ...
    
    @typing.overload
    def __getitem__(self, name: typing_extensions.Literal["resultList"]) -> MetaOapg.properties.resultList: ...
    
    @typing.overload
    def __getitem__(self, name: str) -> schemas.UnsetAnyTypeSchema: ...
    
    def __getitem__(self, name: typing.Union[typing_extensions.Literal["maxId", "minId", "resultList", ], str]):
        # dict_instance[name] accessor
        return super().__getitem__(name)
    
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["maxId"]) -> typing.Union[MetaOapg.properties.maxId, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["minId"]) -> typing.Union[MetaOapg.properties.minId, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: typing_extensions.Literal["resultList"]) -> typing.Union[MetaOapg.properties.resultList, schemas.Unset]: ...
    
    @typing.overload
    def get_item_oapg(self, name: str) -> typing.Union[schemas.UnsetAnyTypeSchema, schemas.Unset]: ...
    
    def get_item_oapg(self, name: typing.Union[typing_extensions.Literal["maxId", "minId", "resultList", ], str]):
        return super().get_item_oapg(name)
    

    def __new__(
        cls,
        *args: typing.Union[dict, frozendict.frozendict, ],
        maxId: typing.Union[MetaOapg.properties.maxId, str, schemas.Unset] = schemas.unset,
        minId: typing.Union[MetaOapg.properties.minId, str, schemas.Unset] = schemas.unset,
        resultList: typing.Union[MetaOapg.properties.resultList, list, tuple, schemas.Unset] = schemas.unset,
        _configuration: typing.Optional[schemas.Configuration] = None,
        **kwargs: typing.Union[schemas.AnyTypeSchema, dict, frozendict.frozendict, str, date, datetime, uuid.UUID, int, float, decimal.Decimal, None, list, tuple, bytes],
    ) -> 'MarginLoanInfoResult':
        return super().__new__(
            cls,
            *args,
            maxId=maxId,
            minId=minId,
            resultList=resultList,
            _configuration=_configuration,
            **kwargs,
        )

from bitget.model.margin_loan_info import MarginLoanInfo
