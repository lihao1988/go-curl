package curl

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}
	dataBytes, err := client.Curl("/iaas/appio?namespace=1", Get, data, JsonType)

	fmt.Println(string(dataBytes), err)
}

func TestPostJson(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}
	dataBytes, err := client.Curl("/iaas/appio/ab", Post, data, FormType)

	fmt.Println(string(dataBytes), err)
}

func TestPutJson(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}
	dataBytes, err := client.Curl("/iaas/appio/ab", Put, data, JsonType)

	fmt.Println(string(dataBytes), err)
}

func TestPatchJson(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}
	dataBytes, err := client.Curl("/iaas/appio/ab", Patch, data, FormType)

	fmt.Println(string(dataBytes), err)
}

func TestDelete(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}
	dataBytes, err := client.Curl("/iaas/appio", Delete, data, JsonType)

	fmt.Println(string(dataBytes), err)
}
