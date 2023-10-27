package middleware

import "testing"

func TestCreateAndVarifToken(t *testing.T) {
	var expectingId uint = 5
	tkn, err := CreateToken(expectingId)
	if err != nil {
		t.Error(err)
	}
	id, err := VerifyToken(tkn)
	if err != nil {
		t.Error(err)
	}
	if id != expectingId {
		t.Errorf("Expected: %d, but got: %d", id, expectingId)
	}
}
