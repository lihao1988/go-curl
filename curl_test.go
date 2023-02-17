package curl

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Get("/iaas/appio/ab", data)
	fmt.Println("Get: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", Get, data, JsonType)
	fmt.Println("curl-get: ", string(dataBytes), err)
}

func TestPostJson(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.PostByForm("/iaas/appio/ab", data)
	fmt.Println("Post: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", Post, data, FormType)
	fmt.Println("curl-post: ", string(dataBytes), err)
}

func TestPutJson(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Put("/iaas/appio/ab", data)
	fmt.Println("Put: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", Put, data, JsonType)
	fmt.Println("curl-put: ", string(dataBytes), err)
}

func TestPatchJson(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Patch("/iaas/appio/ab", data)
	fmt.Println("Patch: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio/ab", Patch, data, JsonType)
	fmt.Println("curl-patch: ", string(dataBytes), err)
}

func TestDelete(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}

	dataBytes, err := client.Curl("/iaas/appio", Delete, data, JsonType)
	fmt.Println("Delete: ", string(dataBytes), err)

	dataBytes, err = client.Curl("/iaas/appio", Delete, data, JsonType)
	fmt.Println("curl-delete: ", string(dataBytes), err)
}
