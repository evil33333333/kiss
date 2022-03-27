//just going to remove any irrelevant comments or anything i dont care to respond to

/// the main function along with the R limit var is placed at the wrong area -  sai
// the location of the function doesnt matter - E
package main

import (
	//"syscall"
	"fmt"
	"log"
	"os"
	"strings"

	/// doesnt even have the same modules as me nor does it use a proxy module. - sai
  
  // i just removed it since divines proxies were in it, it would be very fair to him - E
	"github.com/fatih/color"
	"github.com/valyala/fasthttp"
)

var rl int
var attempts int
var claimed bool = false
var severe_err bool = false

/// useless boolean variables i do not have in my code - sai

// I would argue otherwise. 
// If a session goes invalid in your build, Fuck it, right? -E

/// where is my fasthttp global var gone? - sai
// oh yeah lets have all the goroutines run through one global client, great choice - E

/// WHERE ARE MY COLOUR VARIABLES !!!! - SAI
/// WHERE IS THE AUTH FUNCTION!!! - SAI

// removed, why would have your shit pastebin auth in the code? 
// also colour variables dont really matter in this case, do they? -E 

/// I HAVE SEPERATE FUNCTIONS FOR SPINNERS AND RPS WHERE THEY AT? - SAI
// first, rps with your turbo is dogshit so no need to further embarrass you
// second, you dont need a seperate spinner function since the while loop in the main function does the exact same thing -E
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
  
  /*
    Consider the following assembly.
    
    mov     eax, 7
    lea     rbx, [rsp+188h+var_128]
    call    syscall_Getrlimit
    
    This call shows that something with the equivalent of 7 was called for syscall.Getrlimit().
    This was probably an enum but it doesn't fucking matter because it is commented out.
    
    -E
  */
	//fmt.Printf("%s%s%s Successfully authorised\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	//Apart of the original binary ^^
	fmt.Printf("%s%s%s Kiss Instagram Main Swapper | Sai\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	//This color is incorrect but does it really matter?
  
	/// i do not use %s for calling variables in my print statements another embarassing attempt -  sai
  // oh yeah because "+" is so much better and wayyyy different (makes 0 change in the slightest) -E
  
	fmt.Printf("%s%s%s Sessionid: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	fmt.Println("(this turbo is not coded by sai if you have purchased this you have been scammed.)")
	var sessionid string /// why is the sessionid and target var being called here it could of been made a global var - sai
  /*
  I dont know who taught you how to code but using global variables are a bad practice. 
  Read this: http://www0.cs.ucl.ac.uk/staff/M.Harman/jss4.pdf
  
  -E
  */
  
	fmt.Scanln(&sessionid)
	fmt.Printf("%s%s%s Target: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var target string
	fmt.Scanln(&target)
	checkswap(sessionid, target)
  /// i dont call these vars inside my checkswap func cos my vars are global. -  sai
  // Read the sessionid comment. -E
  
	fmt.Printf("%s%s%s Thread amount: ", color.WhiteString("["), color.HiMagentaString("?"), color.WhiteString("]"))
	var threadcount int 
  ///  i dont call my thread var here and i do not call my var threadcount - sai
  /*
    I know. I prefer threadcount over "threads" because its alot easier to read.
    Be honest, i < threadcount is alot more readable than i < threads.
    
    -E
  */
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
    // I know how to make spinners. I read the assembly and figured this is most likely how you write your spinners. (considering your code generally horribly managed) -E
	}
	// I literally cringed so fucking hard at this part.
  
	/// ur cringing at ur own code pal AKA ribbit turbo - sai
  // vous m'avez tous sous-estim -E
	fmt.Printf("\n%s%s%s Claimed @%s after %d attempts          %s%s%s Kiss v5 | vous m'avez tous sous-estim\n\n", color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"), color.GreenString(target), attempts, color.WhiteString("["), color.GreenString("+"), color.WhiteString("]"))
	os.Exit(0)
	/// i dont use %s again and why isnt there a /n between the kiss v5 and claimed msg - sai
  // BECAUSE THERE IS NO NEW FUCKING LINE YOU MORON -E
}

func checkswap(sessionid string, username string) {
	client := &fasthttp.Client{} /// wtfawk - sai
  /*
      lea     rax, github_com_valyala_fasthttp_requestPool
      xchg    ax, ax
      
      I figured this was a reference. But it does the same thing.
  */
	req := &fasthttp.Request{}
	resp := &fasthttp.Response{}
	req.SetRequestURI("https://i.instagram.com/api/v1/accounts/set_username/")
	req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
	req.Header.SetMethod("POST")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("user-agent", "Instagram 5.9.2 Android (29/10; 420dpi; 1080x1794; Google/google; Android SDK built for x86_64; generic_x86_64; ranchu; db 'en_US; 132081655)")
	/// not my user agent what the fuck is this old ass shit - sai
  
  /*
    Consider the following assembly.
    lea     rbx, aUserAgentS ; "User-Agent\r\n--%s"
    mov     ecx, 0Ah
    lea     rdi, aX509InvalidSig+0B65h ; "Instagram 5.9.2 Android (29/10; 420dpi;"...
    mov     esi, 89h
    nop     dword ptr [rax+00h]
    call    github_com_valyala_fasthttp__ptr_RequestHeader_Set
    
    You might have changed it and forgot but the assembly doesn't lie. You do though.
    
    -E
  
  */
	req.SetBody([]byte(fmt.Sprintf("username=%s", username)))
	err := client.Do(req, resp)
	if err != nil {
		log.Fatal(err) 
    ///  i do not do this in my code unless i HAVE a proxy error and this turbo doesnt even use proxies - sai
    // So, you do 0 error checking? AT ALL?? Not even going to bother bringing up the assembly. You are just a sloth. -E
	}
	fasthttp.ReleaseRequest(req)
	buffer := resp.Body()
	if string(buffer) == "This username isn't available." {
		fmt.Println("This user is NOT swappable\n")
	} else if string(buffer) == "This username isn't available. Please try again." {
		fmt.Println("This user is swappable.\n") 
    /// incorrect strings and theres no colour???? embarassing attempt to copy me again - sai
    // oopsie, it said username not user. Yup, now i have been caught. -E
	} else {
		fmt.Println(string(buffer) + "\n")
	}
	fasthttp.ReleaseResponse(resp)
}

func claimRequest(sessionid string, username string) { 
  /// i do not call my string vars in the claim req cos they are global. -  sai
  /* 
    Despite the fact that this is a bad practice, how does this at all affect the program?
    It literally does the same thing. The same. exact. thing....
    
    -E
  */
  
	for !claimed && !severe_err {
		client := &fasthttp.Client{} 
    ///  again only do this if you want to make the slowest turbo in the entire world - sai
    /* 
      psst... hey.... im trying to help you not look like a moron.... 
      you manage resources like an infant...
    */
		req := fasthttp.AcquireRequest()
		resp := fasthttp.AcquireResponse()
		req.SetRequestURI("https://g.instagram.com/api/v1/accounts/set_username/")
		req.Header.Add("cookie", fmt.Sprintf("sessionid=%s", sessionid))
		req.Header.SetMethod("POST")
		req.Header.Add("content-type", "application/x-www-form-urlencoded")
		req.Header.Add("user-agent", "Instagram 113.0.0.39.122 Android (24/5.0; 515dpi; 1440x2416; 'huawei/google; Nexus 6P; angler; angler; en_US)")
		/// at least u got one user agent right either way... still embarassing - sai
    // no no no, i got both of them right. -E
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
      /// i do not have a claimed boolean variable i just print it in the claim func and os exit. - sai
      
      /*
        Consider the following moron.
        
        Let's say you have 20,000 goroutines running (as usual) and you claim a username.
        You have at the very least 100 goroutines spamming stdout because EVERY SINGLE GOROUTINE HAS ACCESS TO PRINT IN THE CLAIM FUNC. 
        This prevents that from happening by always only printing one time.
        
        -E
      */
		} else if strings.Contains(string(buffer), "logged") {
			severe_err = true
			fmt.Println("err::session went invalid")
			// my turbo doesnt even print this - Sai the Sloth
      // i know, my string just looks way cooler lol -E
			os.Exit(1)
		} else {
			attempts++
		}
		fasthttp.ReleaseResponse(resp)
	}
	/// this is so poorly coded if they put like 500 threads its only gonna do 500 attempts theres no for statement to make it loop. shit turbo -  ssai
  /* 
      That's funny considering that this is your program. Also, I think you can't read but 
      there is a for loop which checks if there is an error with the session or if it claims.
      
      "for !claimed && !severe_err"
      
      Here is a link to a Go tutorial: https://www.youtube.com/watch?v=YS4e4q9oBaU
      And here is a link explaining how loops work in Go: https://www.programiz.com/golang/while-loop
      -E
      
  */
  
	/// imagine failing at making a swap turbo something thats been accomplished many times by multiple other people.
  /*
    Imagine having your life being based around one general functionality of a social media platform. Actually. Actually no.
    Imagine having your life being based around one general functionality of a social media platform AND STILL GETTING
    OVERTAKEN BY NEWGENS. FUCKING. NEWGENS. You spend so much money (your lifesavings) on accounts and endpoints just to be 
    overshadowed by newgens swappers with a Github turbo. Do better.
    
    -E
  */
}
