# coding: utf-8

"""


    Generated by: https://openapi-generator.tech
"""

import unittest
from unittest.mock import patch

import urllib3

import bitget
from bitget.paths.api_margin_v1_isolated_order_history import get  # noqa: E501
from bitget import configuration, schemas, api_client

from .. import ApiTestMixin


class TestApiMarginV1IsolatedOrderHistory(ApiTestMixin, unittest.TestCase):
    """
    ApiMarginV1IsolatedOrderHistory unit test stubs
        history  # noqa: E501
    """
    _configuration = configuration.Configuration()

    def setUp(self):
        used_api_client = api_client.ApiClient(configuration=self._configuration)
        self.api = get.ApiForget(api_client=used_api_client)  # noqa: E501

    def tearDown(self):
        pass

    response_status = 200

    def testApi(self):
        configuration = ApiTestMixin.get_default_configuration()
        # Enter a context with an instance of the API client
        with bitget.ApiClient(configuration) as api_client:
            # Create an instance of the API class
            from bitget.apis.tags import margin_isolated_order_api
            api_instance = margin_isolated_order_api.MarginIsolatedOrderApi(api_client)
            try:
                req = {
                    'symbol':"BTCUSDT",
                    'startTime': "1679133422000",
                }
                api_response = api_instance.margin_isolated_history_orders(req)
                print(api_response)
                self.assertIsNotNone(api_response)
                self.assertIsNotNone(api_response.body)
                self.assertEqual(api_response.body['code'], '00000')
                self.assertEqual(api_response.body['msg'], 'success')
                self.assertIsNotNone(api_response.body['data']['orderList'])
                for item in api_response.body['data']['orderList']:
                    print(item)
                    self.assertEqual(item['symbol'], 'BTCUSDT')
                    self.assertIsNotNone(item['orderType'])
                    self.assertIsNotNone(item['source'])
                    self.assertIsNotNone(item['orderId'])
                    self.assertIsNotNone(item['loanType'])
                    self.assertIsNotNone(item['price'])
                    self.assertIsNotNone(item['side'])
                    self.assertIsNotNone(item['status'])
                    self.assertIsNotNone(item['baseQuantity'])
                    self.assertIsNotNone(item['quoteAmount'])
                    self.assertIsNotNone(item['fillQuantity'])
                    self.assertIsNotNone(item['fillTotalAmount'])
            except bitget.ApiException as e:
                print("Exception when calling place_order: %s\n" % e)


if __name__ == '__main__':
    unittest.main()
