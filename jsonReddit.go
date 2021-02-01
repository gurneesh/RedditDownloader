package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	// "reflect"
)

const (
	PROXY_ADDRESS = "127.0.0.1:9050"
	BASE_URL      = "https://www.reddit.com/r/"
	SUBREDDIT     = "wallpapers"
)

type RedditJson struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
				Subreddit                  string        `json:"subreddit"`
				Selftext                   string        `json:"selftext"`
				AuthorFullname             string        `json:"author_fullname"`
				Saved                      bool          `json:"saved"`
				ModReasonTitle             interface{}   `json:"mod_reason_title"`
				Gilded                     int           `json:"gilded"`
				Clicked                    bool          `json:"clicked"`
				Title                      string        `json:"title"`
				LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
				Hidden                     bool          `json:"hidden"`
				Pwls                       int           `json:"pwls"`
				LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
				Downs                      int           `json:"downs"`
				TopAwardedType             interface{}   `json:"top_awarded_type"`
				HideScore                  bool          `json:"hide_score"`
				Name                       string        `json:"name"`
				Quarantine                 bool          `json:"quarantine"`
				LinkFlairTextColor         string        `json:"link_flair_text_color"`
				UpvoteRatio                float64       `json:"upvote_ratio"`
				AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
				SubredditType              string        `json:"subreddit_type"`
				Ups                        int           `json:"ups"`
				TotalAwardsReceived        int           `json:"total_awards_received"`
				MediaEmbed                 struct {
					Content   string `json:"content"`
					Width     int    `json:"width"`
					Scrolling bool   `json:"scrolling"`
					Height    int    `json:"height"`
				} `json:"media_embed"`
				AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           struct {
					Type   string `json:"type"`
					Oembed struct {
						ProviderURL     string `json:"provider_url"`
						Version         string `json:"version"`
						Title           string `json:"title"`
						Type            string `json:"type"`
						ThumbnailWidth  int    `json:"thumbnail_width"`
						Height          int    `json:"height"`
						Width           int    `json:"width"`
						HTML            string `json:"html"`
						AuthorName      string `json:"author_name"`
						ProviderName    string `json:"provider_name"`
						ThumbnailURL    string `json:"thumbnail_url"`
						ThumbnailHeight int    `json:"thumbnail_height"`
						AuthorURL       string `json:"author_url"`
					} `json:"oembed"`
				} `json:"secure_media"`
				IsRedditMediaDomain bool        `json:"is_reddit_media_domain"`
				IsMeta              bool        `json:"is_meta"`
				Category            interface{} `json:"category"`
				SecureMediaEmbed    struct {
					Content        string `json:"content"`
					Width          int    `json:"width"`
					Scrolling      bool   `json:"scrolling"`
					MediaDomainURL string `json:"media_domain_url"`
					Height         int    `json:"height"`
				} `json:"secure_media_embed"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				AuthorPremium       bool          `json:"author_premium"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              interface{}   `json:"edited"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
				} `json:"gildings"`
				ContentCategories        interface{}   `json:"content_categories"`
				IsSelf                   bool          `json:"is_self"`
				ModNote                  interface{}   `json:"mod_note"`
				Created                  float64       `json:"created"`
				LinkFlairType            string        `json:"link_flair_type"`
				Wls                      int           `json:"wls"`
				RemovedByCategory        interface{}   `json:"removed_by_category"`
				BannedBy                 interface{}   `json:"banned_by"`
				AuthorFlairType          string        `json:"author_flair_type"`
				Domain                   string        `json:"domain"`
				AllowLiveComments        bool          `json:"allow_live_comments"`
				SelftextHTML             interface{}   `json:"selftext_html"`
				Likes                    interface{}   `json:"likes"`
				SuggestedSort            interface{}   `json:"suggested_sort"`
				BannedAtUtc              interface{}   `json:"banned_at_utc"`
				URLOverriddenByDest      string        `json:"url_overridden_by_dest"`
				ViewCount                interface{}   `json:"view_count"`
				Archived                 bool          `json:"archived"`
				NoFollow                 bool          `json:"no_follow"`
				IsCrosspostable          bool          `json:"is_crosspostable"`
				Pinned                   bool          `json:"pinned"`
				Over18                   bool          `json:"over_18"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				TreatmentTags            []interface{} `json:"treatment_tags"`
				Visited                  bool          `json:"visited"`
				RemovedBy                interface{}   `json:"removed_by"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          string        `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    string        `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    struct {
					Type   string `json:"type"`
					Oembed struct {
						ProviderURL     string `json:"provider_url"`
						Version         string `json:"version"`
						Title           string `json:"title"`
						Type            string `json:"type"`
						ThumbnailWidth  int    `json:"thumbnail_width"`
						Height          int    `json:"height"`
						Width           int    `json:"width"`
						HTML            string `json:"html"`
						AuthorName      string `json:"author_name"`
						ProviderName    string `json:"provider_name"`
						ThumbnailURL    string `json:"thumbnail_url"`
						ThumbnailHeight int    `json:"thumbnail_height"`
						AuthorURL       string `json:"author_url"`
					} `json:"oembed"`
				} `json:"media"`
				IsVideo bool `json:"is_video"`
			} `json:"data,omitempty"`
		}
	}
}

func request(url string) {
	fmt.Println("Client without socks5")
	// resp, err := http.Get(url)

	// if err != nil {
	// 	fmt.Println(err)
	// }
}

func download_save(url string) {
	fmt.Println("Downloading ", url)
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDRESS, nil, proxy.Direct)
	if err != nil {
		fmt.Println("Couldn't connect to proxy")
		os.Exit(1)
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	slicedURL := strings.Split(url, "/")
	filehandle, err := os.Create(slicedURL[len(slicedURL)-1])
	defer filehandle.Close()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)s

	_, err = io.Copy(filehandle, resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Downloaded ", slicedURL[len(slicedURL)-1])
}

func socks5_request() []uint8 {
	dialer, err := proxy.SOCKS5("tcp", PROXY_ADDRESS, nil, proxy.Direct)
	if err != nil {
		fmt.Println("Couldn't connect to proxy")
		os.Exit(1)
	}

	httpTransport := &http.Transport{}
	httpClient := &http.Client{Transport: httpTransport}
	httpTransport.Dial = dialer.Dial

	req, err := http.NewRequest("GET", BASE_URL+SUBREDDIT+".json", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36")
	resp, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return body
}

func download(urls []string) {
	for _, url := range urls {
		if strings.HasSuffix(url, ".jpg") || strings.HasSuffix(url, ".jpeg") || strings.HasSuffix(url, ".png") || strings.HasSuffix(url, ".gif") || strings.HasSuffix(url, ".gifv") {
			fmt.Println("Image", url)
			go download_save(url)
		}
	}
}

func main() {
	var urls []string
	body := socks5_request()
	// fmt.Println(reflect.TypeOf(body))

	red := RedditJson{}
	if err := json.Unmarshal(body, &red); err != nil {
		fmt.Println(err)
	}

	for _, i := range red.Data.Children {
		// fmt.Println(i.Data.URL)
		urls = append(urls, i.Data.URL)
	}

	fmt.Println("URLS")

	for _, i := range urls {
		fmt.Println(i)
	}

	download(urls)

	fmt.Scanln()
}
