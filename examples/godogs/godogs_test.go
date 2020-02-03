/* file: $GOPATH/src/godogs/godogs_test.go */
package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
)

var opt = godog.Options{Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func thereAreGodogs(state *godog.ScenarioState, available int) error {
	(*state)["godogs"] = available
	return nil
}

func iEat(state *godog.ScenarioState, num int) error {
	if (*state)["godogs"].(int) < num {
		return fmt.Errorf("you cannot eat %d godogs, there are %d available", num, Godogs)
	}
	(*state)["godogs"] = (*state)["godogs"].(int) - num
	return nil
}

func thereShouldBeRemaining(state *godog.ScenarioState, remaining int) error {
	if (*state)["godogs"].(int) != remaining {
		return fmt.Errorf("expected %d godogs to be remaining, but there is %d", remaining, Godogs)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^there are (\d+) godogs$`, thereAreGodogs)
	s.Step(`^I eat (\d+)$`, iEat)
	s.Step(`^there should be (\d+) remaining$`, thereShouldBeRemaining)

	s.BeforeScenario(func(*godog.ScenarioState, interface{}) {
		Godogs = 0 // clean the state before every scenario
	})
}
