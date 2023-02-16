package curl

import (
	"fmt"
	"testing"
)

func TestCurl(t *testing.T) {
	client := NewClient("http://localhost:8083")
	data := map[string]string{"namespace": "1"}
	dataBytes, err := client.Curl("/iaas/appio?namespace=1", Get, data, JsonType)

	fmt.Println(string(dataBytes), err)
}
