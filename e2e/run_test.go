package main

import (
	apicontext "github.com/brpaz/godog-api-context"
	"github.com/cucumber/godog"
	"testing"
)

func InitializeScenario(s *godog.ScenarioContext) {
	apiContext := apicontext.New("http://localhost:8080")

	apiContext.InitializeScenario(s)
}

func TestFeatures(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
