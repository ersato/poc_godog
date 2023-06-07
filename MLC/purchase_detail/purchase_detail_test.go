package purchase_detail

import (
	"github.com/cucumber/godog"
	"poc_go/utils"
	"testing"
)

func TestFeatures(t *testing.T) {
	api := &utils.ApiFeature{}
	suite := godog.TestSuite{
		ScenarioInitializer: func(s *godog.ScenarioContext) {
			scenarioValidateResponseBodyJson(s, api)
		},
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}
	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func scenarioValidateResponseBodyJson(s *godog.ScenarioContext, api *utils.ApiFeature) {
	s.Step(`I want to fetch the information claim state with config "([^"]*)"$`, api.SetConfig)
	s.Step(`^call "([^"]*)" "([^"]*)" more "([^"]*)" with body "([^"]*)" and headers "([^"]*)"$`, api.CallService)
	s.Step(`^the response code should be (\d+)$`, api.TheResponseCodeShouldBe)
	s.Step(`^response body equals "([^"]*)"$`, api.TheResponseBodyEqual)
}
