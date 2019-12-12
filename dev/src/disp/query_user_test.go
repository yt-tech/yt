package disp

import "testing"

func TestQueryUser(t *testing.T) {
	var user = userSignUpInfo{
		account:  "user1",
		password: "abc",
	}
	var userError = userSignUpInfo{
		account:  "user1",
		password: "abcd",
	}
	mlog.Println(queryUser(user))
	mlog.Println(queryUser(userError))
}
