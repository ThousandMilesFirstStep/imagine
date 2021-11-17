package domain

import (
	"testing"
)

func TestRegisterFilter(t *testing.T) {
	called := false

	filterName := "myNewFilter"
	newFilter := func(image Image, conf map[string]interface{}) error {
		called = true
		return nil
	}

	RegisterFilter(filterName, newFilter)

	t.Run("add the filter to the filters store", func(t *testing.T) {
		if filtersStore[filterName] == nil {
			t.Fail()
		}
	})

	t.Run("registered the appropriate function in the store", func(t *testing.T) {
		image := &TestImage{}
		conf := map[string]interface{}{}

		filtersStore[filterName](image, conf)

		if !called {
			t.Fail()
		}
	})
}
