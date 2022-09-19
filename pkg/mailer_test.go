package pkg

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	//set test dir to project root
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Chdir(wd + "/..")
}
func TestSendMail(t *testing.T) {

	mail := &Mail{
		From:      "email test",
		To:        []string{"leon.lee2050@gmail.com"},
		Subject:   "This is Test subject",
		PlainHtml: "<h1>This is title</h1> and done",
	}
	assert.NoError(t, mail.Send())
}
