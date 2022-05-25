Feature: Pet Store Check Feature

	Scenario: Pet Store Check Scenario
	Given http get pets
	Then http status code 202

