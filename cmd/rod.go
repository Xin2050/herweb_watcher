package main

import (
	"fmt"
	"github.com/Xin2050/web_overwatcher/config"
	"github.com/Xin2050/web_overwatcher/pkg"
	"github.com/go-rod/rod"
	"math/rand"
	"strings"
	"time"
)

var page *rod.Page

var (
	url          = "https://www.hermes.com/ca/en/category/women/bags-and-small-leather-goods/bags-and-clutches/#|"
	sendToList   = []string{"leon.lee2050@gmail.com"}
	watchingList = []string{"Trim 31", "24/24 - 21", "Lindy mini"} // all the watching list
	interval     = 30                                              // for page reload interval (seconds)
)

func main() {
	serverConfig := config.New().Server
	page = rod.New().ControlURL(
		serverConfig.Chrome,
	).MustConnect().MustPage(url)
	page.MustWaitLoad()

	for true {
		isFound := readProductList()
		if isFound {
			break
		}

		//get a random int between 50 and 100
		roundNumber := rand.Intn(interval) + 10

		time.Sleep(time.Second * time.Duration(roundNumber))
		page.Reload()
		page.MustWaitLoad()
		if !page.MustHas("span[class='product-item-name']") { // this path frequently changes
			fmt.Println("You need to fix the html path!")
			break
		}
	}

}

// if website changes you need to change the code to locate the title
func readProductList() bool {
	fmt.Printf("I am trying to find the product list on %s \n", time.Now().String())
	elems := page.MustElements("span[class='product-item-name']")
	if len(elems) > 0 {
		for _, el := range elems {
			//search for the product name is included in the watching list
			for _, watch := range watchingList {
				if el.MustText() == "" {
					continue
				}
				if strings.Contains(strings.ToUpper(el.MustText()), strings.ToUpper(watch)) {
					fmt.Println("I found the product: ", el.MustText(), "for the watch: ", watch)
					linkElem := el.MustParent().MustParent()
					link, err := linkElem.Attribute("href")
					if err != nil {
						fmt.Println("I can't find the link for the product: ", el.MustText())
					}

					fmt.Println("I Found the link for the watch: ", el.MustText(), link)

					sendMail(el.MustText(), link)
					//!!!!!!!! after send email the program will be end !!!!!
					return true
				}
			}
		}
	}
	fmt.Println("I did not find it")
	return false
}

func sendMail(name string, link *string) {
	linkText := ""
	if link != nil {
		linkText = *link
	}
	mail := &pkg.Mail{
		From:    "Go-Mailer",
		To:      sendToList, //[]string{"leon.lee2050@gmail.com"},
		Subject: "Your Bag is INCOMING!",
		PlainHtml: `<h1>Hi,</h1>
					<p>Your bag: ` + name + ` is INCOMING!</p>
					<p><a href="https://www.hermes.com` + linkText + `">click here</a></p>
					
					<p>Best Regards,</p>
					<p>Go-Mailer</p>`,
	}

	err := mail.Send()
	if err != nil {
		fmt.Println("Mail not sent :", err)
		return
	}
	fmt.Println("Mail sent :", name)
}
