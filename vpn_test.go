package vpn

import (
	"testing"
)

func TestAddVpnAccount(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &AddAccountParams{
		Name:        "sdkTest3",
		Note:        "sdk测试3",
		ParentGroup: "/软件研发中心/软件部",
		Passwd:      GetRandom(12),
		RoleName:    "软件部",
	}
	res, err := client.AddAccount(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDelVpnAccount(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &DelAccountParams{Names: "sdkTest3"}
	res, err := client.DelAccount(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
func TestGetAccountDetail(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &GetAccountDetailParams{Username: "test0917"}
	res, err := client.GetAccountDetail(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
func TestResetPwd(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &ResetPwdParams{
		NewName:     "sdkTest3",
		Note:        "sdk测试3",
		OldName:     "sdkTest3",
		ParentGroup: "/软件研发中心/软件部",
		Passwd:      "ffdfdf",
		RoleName:    "软件部",
	}
	res, err := client.ResetPwd(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestGetUserList(t *testing.T) {
	client := NewVpnClient("", "")
	params := &GetUserList{Limit: "100"}
	res, err := client.GetUserList(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestDisconnectUser(t *testing.T) {
	client := NewVpnClient("", "")
	params := &DisConnectUser{Users: ""}
	res, err := client.DisConnectUser(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestIsUserExist(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &IsUserExist{Username: ""}
	res, err := client.IsUserExist(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestIsGroupExist(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &IsGroupExist{GroupName: "/供应链"}
	res, err := client.IsGroupExist(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func TestAddGroup(t *testing.T) {
	client := NewVpnClientWithEnv()
	params := &AddGroup{
		Name:        "信息化管理部",
		ParentGroup: "/",
		Note:        "信息化管理部",
	}
	res, err := client.AddGroup(params)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
