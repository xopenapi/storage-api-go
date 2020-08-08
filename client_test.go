package storage_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xopenapi/storage-api-go"
	"log"
	"testing"
	"time"
)

func TestCreateCredentials(t *testing.T) {
	cfg := storage.NewConfiguration()
	cfg.BasePath = "https://api.xres.lucfish.com"
	cfg.AddDefaultHeader("Authorization", "Bearer xxxxxxxxxxxxxxxxxxx")
	secret := "xxxxxxxxxxxxxxxxxx"

	client := storage.NewAPIClient(cfg)

	req := storage.CreateUploadCredentialsReq{
		Channel: "qcloud",
		Extra: map[string]interface{}{
			"region": "ap-guangzhou",
			"bucket": "mofanshow-avatar-1252461817",
		},
	}
	bodyBuf := &bytes.Buffer{}
	json.NewEncoder(bodyBuf).Encode(req)

	timestamp := time.Now().Unix()
	noncestr := fmt.Sprintf("%d", timestamp)
	url := "/oss/credentials"

	s := fmt.Sprintf("%s%d%s%s%s", bodyBuf.String(), timestamp, noncestr, url, secret)
	signature := storage.Sign([]byte(s))

	r, _, err := client.CredentialsApi.Create(context.TODO(), fmt.Sprintf("%d", timestamp), noncestr, signature, req)
	assert.NoError(t, err, "create credentials returns error")

	log.Print(r)

}
