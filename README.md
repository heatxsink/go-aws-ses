go-aws-ses
==========

Couldn't find this in go-amz ... my intentions are to get it into go-amz if they'll allow it. I know that amazon has an SMTP interface, this is the HTTPS one. Pull requests welcome!

Example
-------
```go
import(
	"fmt"
	"github.com/heatxsink/go-aws-ses"
)

var (
	aws_key = "your_key_here"
	aws_secret = "your_secret_here"
	aws_verified_email_address = "your-verified-email@foo.com"
)

func main() {
	s := ses.New(aws_verified_email_address, aws_key, aws_secret)
	body, status_code, err := s.SendMessage("someones-email@foo.com", subject, message)
	fmt.Println("Status Code: ", status_code)
	fmt.Println("Body: ", string(body))
	if err != nil {
		fmt.Println(err)
	}
}
```
