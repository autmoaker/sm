package main


import
("github.com/dghubble/go-twitter/twitter"
"github.com/dghubble/oauth1"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
	"github.com/fatih/color"
	"time"
	"strconv"
)

func main(){

	config := oauth1.NewConfig("FnBTZChB9TzrFk31ujsIlhY6M", "OdRRbEBRVVv5DHEzLAdT49lH4NS30ezfSDqzfhK10BjKNa5ZwB")
	token := oauth1.NewToken("856494136018444288-VEKxmXAuaJeYPgj2fUCUkmDD8oBqhDr", "8uXEVjXjzg7mcJW0GueDRFItzKwm3bt9WDf1mbXWURom1")
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	fmt.Println(tweets,resp,err)
	// Send a Tweet
	arr,best:=chogadiya()
	fmt.Println(arr)
	d := color.New(color.FgGreen, color.Bold)
	t:=time.Now()

	y,m,da:=t.Date()
	daa:=strconv.Itoa(da)+"/"+m.String()+"/"+strconv.Itoa(y)
	s:=d.Sprintln("Aaj ka ("+daa+") shubh muhurat "+ arr + "\n" + " amrit "+best+"#muhurt#muhurat#aajkashubhmuhurat")
	tweet, resp, err := client.Statuses.Update("Aaj ka ("+daa+") \n shubh muhurat-"+ arr + "\n" + " amrit-"+best+"\n #muhurt #muhurat #aajkashubhmuhurat", nil)
	fmt.Println(tweet,s)
	// Status Show

	fmt.Println(s)
	str:=color.BlueString("Prints %s in blue.", arr)
	fmt.Println(str)

	// More default foreground colors..
	color.Red("We have red")
	color.Yellow("Yellow color too!")
	color.Magenta("And many others ..")

	// Hi-intensity colors
	color.HiGreen("Bright green color.")
	color.HiBlack("Bright black means gray..")
	color.HiWhite("Shiny white color!")
	//tweet, resp, err = client.Statuses.Show(585613041028431872, nil)
	//fmt.Println(tweet,resp,err)
	// Search Tweets
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query: "gopher",
	})
	fmt.Println(search)
	// User Show
	user, resp, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: "dghubble",
	})
	fmt.Println(user)
	// Followers
	followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})
	fmt.Println(followers)
}

func chogadiya() (string,string){


	link:= "https://www.drikpanchang.com/muhurat/choghadiya.html"
	doc, err := goquery.NewDocument(link)
	if err != nil {

		log.Println(err)

	}
	//fmt.Println(doc.Text())
	var name,time []string
	shubh:= ""
	best:=""
	doc.Find(".dpMuhurtaName").Each(func(i int, s *goquery.Selection) {
		name  =append(name , s.Text())

	})

	doc.Find(".dpMuhurtaTime").Each(func(i int, s *goquery.Selection) {
		time  =append(time , s.Text())

	})
	for i,v:=range time{

		fmt.Println(name[i],v)
		if strings.Contains(name[i],"Good"){
			shubh=shubh  + v+ " | "
		}

		if strings.Contains(name[i],"Best"){
			best=best+v+" | "
		}
	}

	return shubh,best

}