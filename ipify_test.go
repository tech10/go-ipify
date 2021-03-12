package ipify

import (
	"testing"
)

func Test_GetIp(t *testing.T) {
	t.Log("Testing both IPV4 and IPV6 resolution.")
	ip, err := GetIp()
	if err != nil {
		t.Fatal(err)
	}
	if ip == "" {
		t.Fatal("No error was encountered, but no IP address was retrieved.")
	} else {
		t.Log("IP address for either IPV4 or IPV6 resolution has succeeded. Result: " + ip)
	}
}

func Test_GetIp4(t *testing.T) {
	ip, err := GetIp4()
	if err != nil {
		t.Log("IPV4 address retrieval has failed.")
		t.Log(err)
		t.FailNow()
	}
	if ip == "" {
		t.Fatal("No error was encountered, but no IP address was retrieved.")
	}
	t.Log("IPV4 address retrieval has succeeded. Result: " + ip)
}

func Test_GetIp6(t *testing.T) {
	ip, err := GetIp6()
	if err != nil {
		t.Log("IPV6 address retrieval has failed.")
		t.Log(err)
		t.FailNow()
	}
	if ip == "" {
		t.Fatal("No error was encountered, but no IP address was retrieved.")
	}
	t.Log("IPV6 address retrieval has succeeded. Result: " + ip)
}

func Test_getIpFailure(t *testing.T) {
	ApiUri := "https://api.ipifyyyyyyyyyyyy.org"
	ip, err := getIp(ApiUri)
	if err == nil || ip != "" {
		t.Fatal("Request to " + ApiUri + " should have failed, but succeeded.")
	} else {
		t.Log(err)
	}
}
