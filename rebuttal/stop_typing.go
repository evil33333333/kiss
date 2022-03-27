/// this is nikos instagram swapper he uses to lose swaps hes just changed the print names to Kiss v5
/// if you have purchased this code and was told its mine yoiu have been scmmed - sai
/// annotating ribbit turbo disguised as my swapper - sai
/// the main function along with the R limit var is placed at the wrong area -  sai
package main

import (
	//"syscall"
	"fmt"
	"log"
	"os"
	"strings"

	/// doesnt even have the same modules as me nor does it use a proxy module. - sai
	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

var rl int
var attempts int
var claimed bool = false
var severe_err bool = false

/// useless boolean variables i do not have in my code - sai
/// where is my fasthttp global var gone? - sai
/// WHERE ARE MY COLOUR VARIABLES !!!! - SAI
/// WHERE IS THE AUTH FUNCTION!!! - SAI
/// I HAVE SEPERATE FUNCTIONS FOR SPINNERS AND RPS WHERE THEY AT? - SAI
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
	/// the main function along with the R limit var is placed at the wrong area -  sai
	/// also gonna ask why the rlimit is at 7 are you guys fucking retarded? LOL - sai
	//fmt.Printf("%s%s%s Successfully authorised\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	//Apart of the original binary ^^
	fmt.Printf("%s%s%s Kiss Instagram Main Swapper | Sai\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	//This color is incorrect but does it really matter?
	/// i do not use %s for calling variables in my print statements another embarassing attempt -  sai
	//
	fmt.Printf("%s%s%s Sessionid: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	fmt.Println("(this turbo is not coded by sai if you have purchased this you have been scammed.)")
	var sessionid string /// why is the sessionid and target var being called here it could of been made a global var - sai
	fmt.Scanln(&sessionid)
	fmt.Printf("%s%s%s Target: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var target string
	fmt.Scanln(&target)
	checkswap(sessionid, target) /// i dont call these vars inside my checkswap func cos my vars are global. -  sai
	fmt.Printf("%s%s%s Thread amount: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var threadcount int ///  i dont call my thread var here and i do not call my var threadcount - sai
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
		/// what the fuck is this LMFAO do these kids not know how to make a spinner :rofl: - sai
	}
	// I literally cringed so fucking hard at this part.
	/// ur cringing at ur own code pal AKA ribbit turbo - sai
	fmt.Printf("\n%s%s%s Claimed @%s after %d attempts          %s%s%s Kiss v5 | vous m'avez tous sous-estim\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"), color.GreenString(target), attempts, color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	os.Exit(0)
	/// i dont use %s again and why isnt there a /n between the kiss v5 and claimed msg - sai
}

func checkswap(sessionid string, username string) {
	client := &fasthttp.Client{} /// wtfawk - sai
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/set_username/")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 5.9.2 Android (29/10; 420dpi; 1080x1794; Google/google; Android SDK built for x86_64; generic_x86_64; ranchu; db 'en_US; 132081655)")
	/// not my user agent what the fuck is this old ass shit - sai

	req.SetBody([]byte(fmt.Sprintf("username=%s", username)))
	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err) ///  i do not do this in my code unless i HAVE a proxy error and this turbo doesnt even use proxies - sai
	}
	fasthttp.ReleaseRequest(req)
	buffer := resp.Body()
	if string(buffer) == "This username isn't available." {
		fmt.Println("This user is NOT swappable\n")
	} else if string(buffer) == "This username isn't available. Please try again." {
		fmt.Println("This user is swappable.\n") /// incorrect strings and theres no colour???? embarassing attempt to copy me again - sai
	} else {
		fmt.Println(string(buffer) + "\n")
	}
	fasthttp.ReleaseResponse(resp)
}

func claimRequest(sessionid string, username string) { /// i do not call my string vars in the claim req cos they are global. -  sai
	for !claimed && !severe_err {
		client := &fasthttp.Client{} ///  again only do this if you want to make the slowest turbo in the entire world - sai
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		req.SetRequestURI("https://g.instagram.com/api/v1/accounts/set_username/")
		req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
		req.Header.SetMethod("POST")
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("user-agent", "Instagram 113.0.0.39.122 Android (24/5.0; 515dpi; 1440x2416; 'huawei/google; Nexus 6P; angler; angler; en_US)")
		/// at least u got one user agent right either way... still embarassing - sai
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
			claimed = true /// i do not have a claimed boolean variable i just print it in the claim func and os exit. - sai
		} else if strings.Contains(string(buffer), "logged") {
			severe_err = true
			fmt.Println("err::session went invalid")
			// my turbo doesnt even print this
			os.Exit(1)
		} else {
			attempts++
		}
		fasthttp.ReleaseResponse(resp)
	}
	/// this is so poorly coded if they put like 500 threads its only gonna do 500 attempts theres no for statement to make it loop. shit turbo -  ssai
	/// imagine failing at making a swap turbo something thats been accomplished many times by multiple other people.
}
