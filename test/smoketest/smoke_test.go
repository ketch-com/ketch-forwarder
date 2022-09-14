//go:build !unit && !integration && smoke

package smoketest

import (
	"context"
	"fmt"
	"github.com/ketch-com/ketch-forwarder/pkg/client"
	"github.com/stretchr/testify/suite"
	"go.ketch.com/lib/orlop/v2"
	"go.uber.org/fx"
	"testing"
	"time"
)

type SmokeTestSuiteParams struct {
	fx.In `ignore-unexported:"true"`

	Config    Config
	Callbacks *client.CallbackClientProvider
	Forwarder *client.ForwarderClientProvider
}

type SmokeTestSuite struct {
	suite.Suite

	app      *fx.App
	params   SmokeTestSuiteParams
	finished chan bool
}

func (suite *SmokeTestSuite) Finished(result bool) {
	suite.finished <- result
}

func (suite *SmokeTestSuite) WaitFinished() bool {
	return suite.WaitFinishedWithTimeout(1 * time.Minute)
}

func (suite *SmokeTestSuite) WaitFinishedWithTimeout(timeout time.Duration) bool {
	select {
	case result := <-suite.finished:
		return result

	case <-time.After(timeout):
		return false
	}
}

func (suite *SmokeTestSuite) SetupTest() {}

func (suite *SmokeTestSuite) TearDownTest() {}

func (suite *SmokeTestSuite) SetupSuite() {
	var err error
	suite.finished = make(chan bool)
	suite.app, err = orlop.TestModule("smoketest", fx.Supply(suite), Module, fx.Populate(&suite.params))
	suite.Require().NoError(err)
}

func (suite *SmokeTestSuite) TearDownSuite() {
	if suite.app != nil {
		suite.app.Stop(context.Background())
		suite.app = nil
	}
}

func (suite *SmokeTestSuite) AfterTest(suiteName, testName string) {}

func (suite *SmokeTestSuite) BeforeTest(suiteName, testName string) {}

func (suite *SmokeTestSuite) HandleStats(suiteName string, stats *suite.SuiteInformation) {
	if stats.Passed() {
		fmt.Printf("\nSmoketest successful at %v. Total duration %0.2v seconds\n\n", stats.End.Format(time.RFC1123), stats.End.Sub(stats.Start).Seconds())
	}
}

func TestSmoke(t *testing.T) {
	suite.Run(t, new(SmokeTestSuite))
}
