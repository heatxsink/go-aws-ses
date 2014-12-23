package ses

import (
	"testing"
	"fmt"
)

var (
	aws_key = "your_key_here"
	aws_secret = "your_secret_here"
	aws_verified_email_address = "your-verified-email@foo.com"
	test_to = "someones-email@foo.com"
)

func TestSendMessage(t *testing.T) {
	fmt.Println("ses.SendMessage()")
	subject := "Test Subject"
	message := "Test Message\nTest Message\tTest Message"
	sss := New(aws_key, aws_secret, aws_verified_email_address)
	body, status_code, err := sss.SendMessage(test_to, subject, message)
	fmt.Println("Status Code: ", status_code)
	fmt.Println("Body: ", string(body))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}