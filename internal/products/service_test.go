package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllBySeller(t *testing.T) {

	//arrange
	sm := NewStorageMock()
	sv := NewService(sm)

	//act

	t.Run("Happy Path", func(t *testing.T) {

		//arrange

		db := []Product{{
			ID:          "1",
			SellerID:    "ABC001",
			Description: "generic",
			Price:       30.0,
		}}

		sm.Data = db
		sm.Err = nil

		//act

		//act

		sellers, err := sv.GetAllBySeller("2")

		//assert
		assert.NoError(t, err)
		assert.Equal(t, 1, len(sellers))
		assert.Equal(t, "generic", sellers[0].Description)
		assert.True(t, sm.Spy)

	})

	t.Run("Error Path", func(t *testing.T) {

		//arrange

		db := []Product{{
			ID:          "1",
			SellerID:    "ABC001",
			Description: "generic",
			Price:       30.0,
		}}

		sm.Data = db
		e := errors.New("error internal")
		sm.Err = e

		//act

		sellers, err := sv.GetAllBySeller("2")

		//assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, e)
		assert.Empty(t, sellers)
		assert.True(t, sm.Spy)

	})

}
