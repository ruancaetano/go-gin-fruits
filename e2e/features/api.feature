Feature: create fruit
  Background:
    Given I set header "Content-Type" with value "application/json"


  Scenario: create fruit
    Given I set header "x-owner" with value "ruan"
    When I send "POST" request to "/fruits" with body:
      """json
      {
          "name": "test",
          "quantity": 3,
          "price": 10.0
      }
      """
    Then The response code should be 201
    Then I store the value of body path "id" as "createdFruitId" in scenario scope


  Scenario: get fruit
    When I send "GET" request to "/fruits/test-uuid"
    Then The response code should be 200

  Scenario: update fruit
    When I send "PUT" request to "/fruits/test-uuid" with body:
      """json
      {
          "quantity": 100,
          "price": 100.0
      }
      """
    Then The response code should be 200
    Then The json path "quantity" should have value "100"

  Scenario: search fruit
    Given I set query param "name" with value "te"
    And I set query param "status" with value "comestible"
    And I set query param "offset" with value "1"
    And I set query param "limit" with value "100"
    When I send "GET" request to "/fruits/search"
    Then The response code should be 200
    Then The json path "Results" should have count "1"


  Scenario: search fruit
    When I send "DELETE" request to "/fruits/test-uuid"
    Then The response code should be 200
    Then The json path "status" should have value "podrido"
