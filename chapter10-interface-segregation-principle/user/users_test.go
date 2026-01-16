package user

import (
	"testing"

	"siuyunyip.io/isp-demo/operation"
)

func TestUser(t *testing.T) {
	operation := operation.Operation{}

	u1 := NewUser1(operation)
	u2 := NewUser2(operation)
	u3 := NewUser3(operation)

	u1.DoWork()
	u2.DoWork()
	u3.DoWork()
}
