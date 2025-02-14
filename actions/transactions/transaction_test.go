package transactions

import (
	"testing"

	"github.com/bitcoin-sv/spv-wallet/config"
	"github.com/bitcoin-sv/spv-wallet/tests"
	"github.com/stretchr/testify/suite"
)

// TestSuite is for testing the entire package using real/mocked services
type TestSuite struct {
	tests.TestSuite
}

// SetupSuite runs at the start of the suite
func (ts *TestSuite) SetupSuite() {
	ts.BaseSetupSuite()
}

// TearDownSuite runs after the suite finishes
func (ts *TestSuite) TearDownSuite() {
	ts.BaseTearDownSuite()
}

// SetupTest runs before each test
func (ts *TestSuite) SetupTest() {
	ts.BaseSetupTest()

	// Load the router & register routes
	oldBasicRoutes, oldAPIRoutes, callbackRoutes := OldTransactionsHandler(ts.AppConfig, ts.Services)
	routes := NewHandler(ts.AppConfig, ts.Services)
	oldBasicRoutes.RegisterOldBasicEndpoints(ts.Router.Group("/" + config.APIVersion))
	oldAPIRoutes.RegisterOldAPIEndpoints(ts.Router.Group("/" + config.APIVersion))
	routes.RegisterAPIEndpoints(ts.Router.Group("/api/" + config.APIVersion))
	routes.RegisterBasicEndpoints(ts.Router.Group("/api/" + config.APIVersion))
	callbackRoutes.RegisterCallbackEndpoints(ts.Router.Group(""))
}

// TearDownTest runs after each test
func (ts *TestSuite) TearDownTest() {
	ts.BaseTearDownTest()
}

// TestTestSuite kick-starts all suite tests
func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
