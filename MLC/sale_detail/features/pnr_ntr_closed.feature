Feature: [MLC] Search PNR information Claim State

  Scenario: Validate Sale Detail response information from pnr_ntr_closed
    Given I want to fetch the information claim state with config "resources/config.json"
    When call "sale-detail" "api.url" more "pnr_ntr_closed" with body "resources/request_body.json" and headers "api.headers"
    Then the response code should be 404
    And response body equals "resources/claim_state/pnr_ntr_closed/response.json"