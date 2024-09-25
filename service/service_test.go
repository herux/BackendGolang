package service_test

import (
	"errors"
	"testing"

	"github.com/herux/indegooweather/config"
	"github.com/herux/indegooweather/db"
	"github.com/herux/indegooweather/service"
	"github.com/stretchr/testify/assert"
)

func TestFetchAndStoreIndegoData_Success(t *testing.T) {
	_ = config.Load("../config.yaml")
	db.Init(true)

	err := service.FetchAndStoreIndegoData()
	assert.NoError(t, err)
}

func mockFetchAndStoreIndegoData_Failure() error {
	// Simulate an error condition (e.g., failed to fetch data)
	return errors.New("failed to fetch data from Indego API")
}
