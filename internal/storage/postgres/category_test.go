package postgres

import (
	"testing"
	// "github.com/OurLuv/prefood/internal/model"
)

func TestGetCategoryById(t *testing.T) {
	r := NewCategoryStorage(pool)

	var err error
	if _, err = r.GetById(10, 5); err != nil {
		t.Error(err)
	}
}
