package ses

import (
	"net"
	"net/http"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"
)

const (
	AWS_SES_ENDPOINT = "https://email.us-east-1.amazonaws.com"
)

type Ses struct {
	AwsAccessKey string
	AwsAccessSecret string
	AwsVerifiedEmailAddress string
}

func New(access_key string, access_secret string, aws_verified_email_address string) *Ses {
	e := new (Ses)
	e.AwsAccessKey = access_key
	e.AwsAccessSecret = access_secret
	e.AwsVerifiedEmailAddress = aws_verified_email_address
	return e
}

func (e *Ses) SendMessage(to, subject, body string) ([]byte, int, error) {
	data := make(url.Values)
	data.Add("Action", "SendEmail")
	data.Add("Source", e.AwsVerifiedEmailAddress)
	data.Add("Destination.ToAddresses.member.1", to)
	data.Add("Message.Subject.Data", subject)
	data.Add("Message.Body.Text.Data", body)
	data.Add("AWSAccessKeyId", e.AwsAccessKey)
	return ses_send_message(data, e.AwsAccessKey, e.AwsAccessSecret)
}

func ses_send_message(data url.Values, access_key string, access_secret string) ([]byte, int, error) {
	headers := http.Header{}
	now := time.Now().UTC()
	// date format: "Tue, 25 May 2010 21:20:27 +0000"
	date := now.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	headers.Set("Date", date)
	h := hmac.New(sha256.New, []uint8(access_secret))
	h.Write([]uint8(date))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	auth := fmt.Sprintf("AWS3-HTTPS AWSAccessKeyId=%s, Algorithm=HmacSHA256, Signature=%s", access_key, signature)
	headers.Set("X-Amzn-Authorization", auth)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	body := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", AWS_SES_ENDPOINT, body)
	if err != nil {
		return nil, -1, err
	}
	req.Header = headers
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, -1, err
	}
	defer response.Body.Close()
	status_code := response.StatusCode
	contents, err := ioutil.ReadAll(response.Body)
	neterr, ok := err.(net.Error)
	if ok && neterr.Timeout() {
		err = nil
	}
	if err != nil {
		return nil, status_code, err
	}
	return contents, status_code, nil
}
