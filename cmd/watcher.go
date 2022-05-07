package main

import (
	"fmt"
	"github.com/Xin2050/web_overwatcher/config"
	"github.com/gocolly/colly"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//load config file
	config := config.New()
	if len(config.Tasks) < 1 {
		log.Fatal("No Tasks found")
	}
	for _, task := range config.Tasks {
		go distributeTask(task)
	}
	// Wait for interrupt signal to gracefully shutdown the server with
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL}
	cn := make(chan os.Signal, 1)
	signal.Notify(cn, signals...)
	<-cn
	fmt.Println("received an exit command, start to try closing....")

}

func distributeTask(task config.Task) {
	fmt.Println("distributing task:", task.Name)
	{
		taskDoneChan := make(chan string)
		doTask(task, taskDoneChan)
		<-taskDoneChan
		fmt.Println("JOB done.")
		time.Sleep(time.Duration(task.Frequency) * time.Second)
	}

}

func doTask(task config.Task, taskChan chan string) {
	fmt.Println("Doing task:", task.Name)

	list, err := loadList(task)
	if err != nil {
		fmt.Println("Error loading list:", err)
	}
	fmt.Println(list)

	close(taskChan)
}

func loadList(task config.Task) ([]string, error) {

	c := colly.NewCollector()
	c.OnHTML("div", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://www.t66y.com/thread0806.php?fid=25")
	return nil, nil
}
