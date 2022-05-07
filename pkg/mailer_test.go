package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	//set test dir to project root

}
func TestSendMail(t *testing.T) {

	mail := &Mail{
		From: "PMS ADMIN",
		To:   []string{"leon.lee2050@gmail.com"},
		//Cc:        []string{"6636101@qq.com"},
		Subject:   "This is Test subject",
		PlainHtml: "<h1>This is title</h1> and done",
	}
	assert.NoError(t, mail.Send())
}
