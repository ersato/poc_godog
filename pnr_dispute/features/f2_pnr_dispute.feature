Feature: 2 - Search PNR information Dispute Opened

  Scenario: 2.1 - Validate response information from pnr_dispute_opened
    Given I want to fetch the information pnr_dispute_opened with config "resources/feature_1/scenario_1/config.json"
    When call "api.url" with body "resources/feature_1/scenario_1/request_body.json" and headers "api.headers_pnr_dispute"
    Then the response code should be 200
    And response body equals "resources/feature_1/scenario_1/response.json"

  Scenario: 2.2 - Perform attribute validation from response pnr_dispute_opened
    Given I want to separately validate values returned pnr_dispute_opened with config "resources/feature_1/scenario_2/config.json"
    When call "api.url" with body "resources/feature_1/scenario_2/request_body.json" and headers "api.headers_pnr_dispute"
    Then response equals
      | resource_type | shipment                                                                        |
      | resource_id   | 42139118821                                                                     |
      | status.type   | info                                                                            |
      | status.header | Cancelamento solicitado                                                         |
      | status.title  | Entraremos em contato antes de domingo, 31 de dezembro, para te dar uma solução |