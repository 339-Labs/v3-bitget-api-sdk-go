# coding: utf-8

"""


    Generated by: https://openapi-generator.tech
"""

import unittest
from unittest.mock import patch

import urllib3

import bitget
from bitget.paths.api_margin_v1_cross_public_tier_data import get  # noqa: E501
from bitget import configuration, schemas, api_client

from .. import ApiTestMixin


class TestApiMarginV1CrossPublicTierData(ApiTestMixin, unittest.TestCase):
    """
    ApiMarginV1CrossPublicTierData unit test stubs
        tierData  # noqa: E501
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
            from bitget.apis.tags import margin_cross_public_api
            api_instance = margin_cross_public_api.MarginCrossPublicTierData(api_client)
            try:
                req = {
                    'coin':"USDT"
                }
                api_response = api_instance.margin_cross_public_tier_data(req)
                print(api_response)
                self.assertIsNotNone(api_response)
                self.assertIsNotNone(api_response.body)
                self.assertEqual(api_response.body['code'], '00000')
                self.assertEqual(api_response.body['msg'], 'success')
                self.assertIsNotNone(api_response.body['data'])
                for item in api_response.body['data']:
                    print(item)
                    self.assertEqual(item['coin'], 'USDT')
                    self.assertIsNotNone(item['tier'])
                    self.assertIsNotNone(item['leverage'])
                    self.assertIsNotNone(item['maxBorrowableAmount'])
                    self.assertIsNotNone(item['maintainMarginRate'])
            except bitget.ApiException as e:
                print("Exception when calling place_order: %s\n" % e)



if __name__ == '__main__':
    unittest.main()
