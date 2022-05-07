package main

import (
	"fmt"
	"github.com/Xin2050/web_overwatcher/pkg"
	"github.com/go-rod/rod"
	"math/rand"
	"strings"
	"time"
)

var page *rod.Page

var (
	watchingList = []string{"24/24 - 21", "Lindy mini"}
)

func main() {

	const url = "https://www.hermes.com/ca/en/category/women/bags-and-small-leather-goods/bags-and-clutches/#|"
	page = rod.New().ControlURL(
		"ws://127.0.0.1:9222/devtools/browser/736aa646-776a-41b0-b888-7e712448f28b",
	).MustConnect().MustPage(url)
	page.MustWaitLoad()

	for true {
		isFound := readProductList()
		if isFound {
			break
		}

		//get a random int between 50 and 100
		roundNumber := rand.Intn(50) + 10

		time.Sleep(time.Second * time.Duration(roundNumber))
		page.Reload()
		page.MustWaitLoad()
		if !page.MustHas("h4[class='product-item-name']") {
			fmt.Println("need to fix this!")
			break
		}
	}

}

func readProductList() bool {
	fmt.Printf("I am trying to find the product list on %s \n", time.Now().String())
	elems := page.MustElements("h4[class='product-item-name']")
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
					sendMail(el.MustText(), link)
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
		From: "Go-Mailer",
		To:   []string{"leon.lee2050@gmail.com", "belle.z2050@gmail.com"}, //[]string{"leon.lee2050@gmail.com"},

		//Cc:        []string{"6636101@qq.com"},
		Subject: "Your Bag is INCOMING!",
		PlainHtml: `<h1>Hi,</h1>
					<p>Your bag: ` + name + ` is INCOMING!</p>
					<p><a href="https://www.hermes.com` + linkText + `">click here</a></p>
					
					<p>Best Regards,</p>
					<p>Go-Mailer</p>`,
	}
	mail.Send()
}
