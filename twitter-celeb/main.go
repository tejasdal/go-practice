package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/peterhellberg/giphy"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

var (
	TWITTER_CREDENTIAL          = getEnv("TWITTER_CREDENTIAL")
	TWITTER_CREDENTIAL_SECRET   = getEnv("TWITTER_CREDENTIAL_SECRET")
	TWITTER_ACCESS_TOKEN        = getEnv("TWITTER_ACCESS_TOKEN")
	TWITTER_ACCESS_TOKEN_SECRET = getEnv("TWITTER_ACCESS_TOKEN_SECRET")
	HOLIDAY_API_KEY             = getEnv("HOLIDAY_API_KEY")     //https://www.abstractapi.com/holidays-api#docs
)

func getEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Missing environment variable: %s", key))
	}
	return val
}

func main() {

	anaconda.SetConsumerKey(TWITTER_CREDENTIAL)
	anaconda.SetConsumerSecret(TWITTER_CREDENTIAL_SECRET)
	twitter := anaconda.NewTwitterApi(TWITTER_ACCESS_TOKEN, TWITTER_ACCESS_TOKEN_SECRET)

	for _, event := range whatIsToday() {
		giphy := giphy.DefaultClient
		gifs, err := searchGifs(giphy, []string{event})
		if err != nil {
			log.Println(err)
			return
		}

		gifURL := selectRandomGif(gifs)

		base64Gif, size, err := parseGifToBase64(gifURL)
		if err != nil {
			log.Println(err)
			return
		}

		err = tweetAGif(twitter, "#justforfunc #dalu #gophers #giphy", base64Gif, size)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Successfully post tweet with GIF url: %s for event keyword: %s", gifURL, event)
	}
}

func whatIsToday() []string {

	var events = make([]string, 0)
	holiday := getHolidayTodayInCountry("IN")
	if holiday != "" {
		events = append(events, holiday)
	}
	time.Sleep(time.Second)

	holiday = getHolidayTodayInCountry("US")
	if holiday != "" {
		events = append(events, holiday)
	}
	time.Sleep(time.Second)

	holiday = getHolidayTodayInCountry("CA")
	if holiday != "" {
		events = append(events, holiday)
	}

	if len(events) != 0 {
		return events
	}

	now := time.Now()
	if now.Weekday() == 5 {
		events = append(events, "Weekend")
	} else {
		events = append(events, fmt.Sprintf("%s", now.Weekday()))
	}
	return events
}

func getHolidayTodayInCountry(country string) string {
	now := time.Now()

	url := fmt.Sprintf("https://holidays.abstractapi.com/v1/?api_key=%s&country=%s&year=%d&month=%d&day=%d",
		HOLIDAY_API_KEY, country, now.Year(), now.Month(), now.Day())

	res, err := http.Get(url)
	if err != nil {
		log.Printf("could not fetch holiday from Holiday API, URL: %v, error: %v", url, err)
		return ""
	}
	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("could not parse response to bytes from Holiday API, URL: %v, error: %v", url, err)
		return ""
	}

	var holiday = make([]map[string]interface{}, 0)
	if err := json.Unmarshal(bytes, &holiday); err != nil {
		log.Printf("could not parse response to bytes from Holiday API, URL: %v, error: %v", url, err)
		return ""
	}
	if len(holiday) == 0 {
		log.Printf("no holiday today: %v", now)
		return ""
	}

	fmt.Printf("%+v \n", holiday[0]["name"])
	return holiday[0]["name"].(string)
}

func tweetAGif(twitter *anaconda.TwitterApi, tweet string, base64Gif string, gifSize int) error {
	media, err := twitter.UploadVideoInit(gifSize, "image/gif")

	err = twitter.UploadVideoAppend(media.MediaIDString, 0, base64Gif)
	if err != nil {
		return fmt.Errorf("could not append gif data with media id: %s : %v", media.MediaIDString, err)
	}

	gifMedia, err := twitter.UploadVideoFinalize(media.MediaIDString)
	if err != nil {
		return fmt.Errorf("could not send finalize cmd for gif with media id: %s : %v", media.MediaIDString, err)
	}

	_, err = twitter.PostTweet(tweet, url.Values{"media_ids": []string{gifMedia.MediaIDString}})
	if err != nil {
		return fmt.Errorf("count not post tweet: %v\n", err)
	}
	return nil
}

func parseGifToBase64(gifURL string) (string, int, error) {
	res, err := http.Get(gifURL)
	defer res.Body.Close()
	if err != nil {
		return "", -1, fmt.Errorf("could not fetch gif from %s: %v\n", gifURL, err)
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", -1, fmt.Errorf("could not convert gif response to byte array: %v", err)
	}

	return base64.StdEncoding.EncodeToString(bytes), len(bytes), nil
}

func selectRandomGif(gifs []giphy.Data) string {
	s1 := rand.NewSource(time.Now().UnixNano())
	randomSeek := rand.New(s1)
	index := randomSeek.Intn(len(gifs) - 1)
	return gifs[index].Images.FixedHeight.URL
}

func searchGifs(client *giphy.Client, keywords []string) ([]giphy.Data, error) {
	result, err := client.Search(keywords)
	if err != nil {
		return nil, fmt.Errorf("could not search giphy: %v\n", err)
	}
	if len(result.Data) <= 0 {
		return nil, fmt.Errorf("no giphy found for given keywords: %+v", keywords)
	}
	return result.Data, nil
}
