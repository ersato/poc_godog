Feature: [MLB] Search PNR information Claim State

  Scenario: Validate Sale Detail response information from pnr_dispute_elected_action_reminder
    Given I want to fetch the information claim state with config "resources/config.json"
    When call "sale-detail" "api.url" more "pnr_dispute_elected_action_reminder" with body "resources/request_body.json" and headers "api.headers"
    Then the response code should be 200
    And response body equals "resources/claim_state/pnr_dispute_elected_action_reminder/response.json"