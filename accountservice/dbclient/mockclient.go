package dbclient

import (
	"github.com/ryanyogan/goblog/accountservice/model"
	"github.com/stretchr/testify/mock"
)

// MockBoltClient -
type MockBoltClient struct {
	mock.Mock
}

// QueryAccount - mocking all fucntions from the Interface{}
func (m *MockBoltClient) QueryAccount(accountID string) (model.Account, error) {
	args := m.Mock.Called(accountID)
	return args.Get(0).(model.Account), args.Error(1)
}

// OpenBoltDB -
func (m *MockBoltClient) OpenBoltDB() {
	// Do nothing
}

// Seed -
func (m *MockBoltClient) Seed() {
	// Do nothing
}

// Check -
func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	return args.Get(0).(bool)
}
