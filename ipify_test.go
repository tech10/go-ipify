package ipify

import (
	"fmt"
	"testing"
)

func Test_getIp(t *testing.T) {
	ApiUri := API_URI
	ip, err := getIp(ApiUri)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
}

func Test_getIpFailure(t *testing.T) {
	ApiUri := "https://api.ipifyyyyyyyyyyyy.org"
	ip, err := getIp(ApiUri)
	if err == nil || ip != "" {
		t.Error("Request to " + ApiUri + " should have failed, but succeeded.")
	} else {
		fmt.Println(err)
	}
}
