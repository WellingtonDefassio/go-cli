package utils

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strings"
	"time"
)

func GetChromeScreenShot(site string, quality int) {
	url := fmt.Sprintf("https://%s/", site)

	var buf []byte
	var ext string = "png"
	if quality < 100 {
		ext = "jpg"
	}
	log.Printf("take a screenshot from site %s", url)

	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(1400, 200000))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)
	actx, cancelFunc := chromedp.NewExecAllocator(context.Background(), options...)
	defer cancelFunc()

	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	tasks := chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.Sleep(3 * time.Second),
		chromedp.CaptureScreenshot(&buf),
	}
	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}
	filename := fmt.Sprintf("%s-%d-%v.%s", strings.Replace(site, "/", "-", -1), quality, time.Now().Unix(), ext)
	if err := os.WriteFile(filename, buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("screenshot save at %s", filename)
}
