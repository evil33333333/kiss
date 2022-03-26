package main

import (
	//"syscall"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

var rl int
var attempts int
var claimed bool = false
var severe_err bool = false

func main() {
	/*
		var rlimit syscall.Rlimit
		variable, err := syscall.Getrlimit(7, &rlimit)
		if err != nil {
			fmt.Println("couldn't get rlimit")
		}
		ok, err := syscall.Setrlimit(7, &rlimit)
		if err != nil {
			fmt.Println("err setting rlimit")
		}

		How do you adjust rlimit and still be slow LOL

		For linux users
	*/

	//fmt.Printf("%s%s%s Successfully authorised\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	//Apart of the original binary ^^
	fmt.Printf("%s%s%s Kiss Instagram Main Swapper | Sai\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	//This color is incorrect but does it really matter?
	fmt.Printf("%s%s%s Sessionid: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var sessionid string
	fmt.Scanln(&sessionid)
	fmt.Printf("%s%s%s Target: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var target string
	fmt.Scanln(&target)
	checkswap(sessionid, target)
	fmt.Printf("%s%s%s Thread amount: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var threadcount int
	fmt.Scanln(&threadcount)
	fmt.Printf("\n%s%s%s Press enter to start\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	fmt.Scanln()
	for i := 0; i < threadcount; i++ {
		go claimRequest(sessionid, target)
	}
	for !claimed {
		fmt.Printf("\r%s%s%s Requests: %d | RL: %d\r", color.WhiteString("["), color.GreenString("/"), color.WhiteString("]"), attempts, rl)
		fmt.Printf("\r%s%s%s Requests: %d | RL: %d\r", color.WhiteString("["), color.BlueString("-"), color.WhiteString("]"), attempts, rl)
		fmt.Printf("\r%s%s%s Requests: %d | RL: %d\r", color.WhiteString("["), color.RedString("\\"), color.WhiteString("]"), attempts, rl)
		fmt.Printf("\r%s%s%s Requests: %d | RL: %d\r", color.WhiteString("["), color.YellowString("|"), color.WhiteString("]"), attempts, rl)
	}
	// I literally cringed so fucking hard at this part.
	fmt.Printf("\n%s%s%s Claimed @%s after %d attempts          %s%s%s Kiss v5 | vous m'avez tous sous-estim\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"), color.GreenString(target), attempts, color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	os.Exit(0)
}

func checkswap(sessionid string, username string) {
	client := &fasthttp.Client{}
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/set_username/")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 5.9.2 Android (29/10; 420dpi; 1080x1794; Google/google; Android SDK built for x86_64; generic_x86_64; ranchu; db 'en_US; 132081655)")
	req.SetBody([]byte(fmt.Sprintf("username=%s", username)))
	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err)
	}
	fasthttp.ReleaseRequest(req)
	buffer := resp.Body()
	if string(buffer) == "This username isn't available." {
		fmt.Println("This user is NOT swappable\n")
	} else if string(buffer) == "This username isn't available. Please try again." {
		fmt.Println("This user is swappable.\n")
	} else {
		fmt.Println(string(buffer) + "\n")
	}
	fasthttp.ReleaseResponse(resp)
}

func claimRequest(sessionid string, username string) {
	for !claimed && !severe_err {
		client := &fasthttp.Client{}
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		req.SetRequestURI("https://i.instagram.com/api/v1/accounts/set_username/")
		req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
		req.Header.SetMethod("POST")
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("user-agent", "Instagram 113.0.0.39.122 Android (24/5.0; 515dpi; 1440x2416; 'huawei/google; Nexus 6P; angler; angler; en_US)")
		req.SetBody([]byte(fmt.Sprintf("username=%s", username)))
		err := client.Do(req, resp)
		if err != nil {
			log.Fatal(err)
		}
		fasthttp.ReleaseRequest(req)
		buffer := resp.Body()
		if strings.Contains(string(buffer), "few minutes") {
			rl++
		} else if strings.Contains(string(buffer), "email") {
			claimed = true
		} else if strings.Contains(string(buffer), "logged") {
			severe_err = true
			fmt.Println("err::session went invalid")
			os.Exit(1)
		} else {
			attempts++
		}
		fasthttp.ReleaseResponse(resp)
	}

}
