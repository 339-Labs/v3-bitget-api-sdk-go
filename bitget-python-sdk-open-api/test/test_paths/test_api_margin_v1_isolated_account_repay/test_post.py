# coding: utf-8

"""


    Generated by: https://openapi-generator.tech
"""

import unittest
from unittest.mock import patch

import urllib3

import bitget
from bitget.paths.api_margin_v1_isolated_account_repay import post  # noqa: E501
from bitget import configuration, schemas, api_client

from .. import ApiTestMixin


class TestApiMarginV1IsolatedAccountRepay(ApiTestMixin, unittest.TestCase):
    """
    ApiMarginV1IsolatedAccountRepay unit test stubs
        repay  # noqa: E501
    """
    _configuration = configuration.Configuration()

    def setUp(self):
        used_api_client = api_client.ApiClient(configuration=self._configuration)
        self.api = post.ApiForpost(api_client=used_api_client)  # noqa: E501

    def tearDown(self):
        pass

    response_status = 200

    def testApi(self):
        configuration = ApiTestMixin.get_default_configuration()
        # Enter a context with an instance of the API client
        with bitget.ApiClient(configuration) as api_client:
            # Create an instance of the API class
            from bitget.apis.tags import margin_isolated_account_api
            api_instance = margin_isolated_account_api.MarginIsolatedAccountRepay(api_client)
            try:
                from bitget.model.margin_isolated_repay_request import MarginIsolatedRepayRequest
                req = MarginIsolatedRepayRequest(
                    coin="USDT",
                    repayAmount="1",
                    symbol="BTCUSDT"
                )
                api_response = api_instance.margin_isolated_account_repay(req)
                print(api_response)
                self.assertIsNotNone(api_response)
                self.assertIsNotNone(api_response.body)
                self.assertEqual(api_response.body['code'], '00000')
                self.assertEqual(api_response.body['msg'], 'success')
                self.assertEqual(api_response.body['data']['coin'], 'USDT')
                self.assertIsNotNone(api_response.body['data']['repayAmount'])
                self.assertIsNotNone(api_response.body['data']['remainDebtAmount'])
            except bitget.ApiException as e:
                print("Exception when calling place_order: %s\n" % e)




if __name__ == '__main__':
    unittest.main()
