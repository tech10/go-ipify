package ipify

import (
	"testing"
)

func Test_getIp(t *testing.T) {
	ApiUri := API_URI
	ip, err := getIp(ApiUri)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ip)
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
