package main

import (
	"fmt"
	"github.com/gobs/pretty"
	"github.com/raff/godet"
	"os"
	"time"
)

const url = "https://www.hermes.com/ca/en/category/women/bags-and-small-leather-goods/bags-and-clutches/#|"

func deleteFiles() {
	err := os.Remove("page.pdf")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("screenshot.png")
	if err != nil {
		fmt.Println(err)
	}

}
func main() {
	//deleteFiles()
	// connect to Chrome instance
	remote, err := godet.Connect("localhost:9222", false)
	if err != nil {
		fmt.Println("cannot connect to Chrome instance:", err)
		return
	}

	// disconnect when done
	defer remote.Close()
	// get list of open tabs

	tabs, _ := remote.TabList("")
	for index, tab := range tabs {
		if index == 0 {
			continue
		}
		remote.CloseTab(tab)
	}
	/*
		// install some callbacks
		remote.CallbackEvent(godet.EventClosed, func(params godet.Params) {
			fmt.Println("RemoteDebugger connection terminated.")
		})

		remote.CallbackEvent("Network.requestWillBeSent", func(params godet.Params) {
			fmt.Println("requestWillBeSent",
				params["type"],
				params["documentURL"],
				params["request"].(map[string]interface{})["url"])
		})

		remote.CallbackEvent("Network.responseReceived", func(params godet.Params) {
			fmt.Println("responseReceived",
				params["type"],
				params["response"].(map[string]interface{})["url"])
		})

		remote.CallbackEvent("Log.entryAdded", func(params godet.Params) {
			entry := params["entry"].(map[string]interface{})
			fmt.Println("LOG", entry["type"], entry["level"], entry["text"])
		})
	*/

	tab, err := remote.NewTab(url)
	if err != nil {
		fmt.Println("cannot navigate to url:", err)
		return
	}
	// enable event processing
	remote.RuntimeEvents(true)
	remote.NetworkEvents(true)
	remote.PageEvents(true)
	remote.DOMEvents(true)
	remote.LogEvents(true)

	remote.ActivateTab(tab)

	remote.AllEvents(true)

	//print tab all info
	fmt.Printf("tab: %v\n", tab)
	time.Sleep(3 * time.Second)
	// get table DOM
	//for i := 0; i < 5; i++ {
	//	html, err := remote.GetOuterHTML(i)
	//	if err != nil {
	//		continue
	//	} else {
	//		fmt.Println(html)
	//		break
	//	}
	//}
	dom, err := remote.QuerySelectorAll(5, "div[class='grid-result']")

	if err != nil {
		fmt.Println("cannot get document:", err)
		return
	}
	pretty.PrettyPrint(dom)

	//
	//// navigate in existing tab
	////_ = remote.ActivateTab(tabs[0])
	////remote.CloseTab(tabs[0])
	//
	//// re-enable events when changing active tab
	//remote.AllEvents(true) // enable all events

}

func PrintMap(m map[string]interface{}) {
	for key, value := range m {
		// check value type, if it's a map, print it
		if valueMap, ok := value.(map[string]interface{}); ok {
			PrintMap(valueMap)
		} else {
			fmt.Printf("unknow - %s: %v\n", key, value)
		}
	}
}
