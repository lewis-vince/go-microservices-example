package main

import "github.com/cucumber/godog"

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^http get pets$`, httpGetPets)
	ctx.Step(`^http status code (\d+)$`, httpStatusCode)
}
