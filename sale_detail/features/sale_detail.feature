Feature: Search information Sale Details - PNR Reminder

  Scenario: Validate response information from sale-detail/pnr_reminder
    Given I want to fetch the information sale_detail with config "resources/configFile.json"
    When call "api.url" with body "resources/request/body.json" and headers "api.headers_sale_details"
    Then the response code should be 200
    And response body equals "resources/response/response.json"