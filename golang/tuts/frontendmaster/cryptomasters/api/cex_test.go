package api_test

import (
	"cryptomasters/api"
	"testing"
)

func TestApiCall(t *testing.T) {
 _, err := api.GetRate("")  

  if err == nil {
    t.Errorf("Error should not be nil")
  }
}
