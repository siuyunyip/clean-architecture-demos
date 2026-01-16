package user

import "siuyunyip.io/isp-demo/operation"

type User1 struct {
	Operation operation.U1Ops
}

func NewUser1(ops operation.U1Ops) *User1 {
	return &User1{
		Operation: ops,
	}
}

func (u *User1) DoWork() {
	u.Operation.Ops1()
}

type User2 struct {
	Operation operation.U2Ops
}

func NewUser2(ops operation.U2Ops) *User2 {
	return &User2{
		Operation: ops,
	}
}

func (u *User2) DoWork() {
	u.Operation.Ops2()
}

type User3 struct {
	Operation operation.U3Ops
}

func NewUser3(ops operation.U3Ops) *User3 {
	return &User3{
		Operation: ops,
	}
}

func (u *User3) DoWork() {
	u.Operation.Ops3()
}
