package storage_test

import(
	"context"
	"testing"
	"log"
	"fmt"
	"time"
	"github.com/xopenapi/storage-api-go"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestCreateCredentials(t *testing.T) {
	cfg := storage.NewConfiguration()
	cfg.BasePath = "http://192.168.0.99:10300"
	cfg.AddDefaultHeader("accesstoken", "0b88b62444694267af5f975782440843")
	// apiSecret := "59a9b6f2210340fabb280764fba6d72c"
	apiSecret := "1954acad03914a5288bf9f60f175b525"

	client := storage.NewAPIClient(cfg)

	req := storage.CreateUploadCredentialsReq{
		Channel: "qcloud",
		Extra: map[string]interface{}{
			"region": "ap-guangzhou",
			"bucket": "mofanshow-avatar-1252461817",
		},
	}

	data, err := json.Marshal(req)
	assert.NoError(t, err, "marshal returns error")

	timestamp := time.Now().Unix()
	noncestr := fmt.Sprintf("%d", timestamp)
	url := "/oss/credentials"

	s := fmt.Sprintf("%s%d%s%s%s", string(data), timestamp, noncestr, url, apiSecret)
	signature := storage.Sign([]byte(s))


	auth := context.WithValue(context.Background(), "Timestamp", timestamp)
	auth = context.WithValue(context.Background(), "Signature", signature)

	r, _, err := client.CredentialsApi.Create(auth, apiSecret, req)
	assert.NoError(t, err, "create credentials returns error")

	log.Print(r)

}
