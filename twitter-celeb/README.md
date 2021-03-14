### Twitter Celeb Bot
Twitter Celeb Bot is a go application which tweets a GIF for an event. It leverages [anaconda](https://pkg.go.dev/github.com/ChimeraCoder/anaconda)
library to communicate with Twitter API. Basically, it is supposed to run as a cron job to tweet daily from given Twitter account.

#### How to get event?
Twitter Celeb Bot uses [Holiday API](//https://www.abstractapi.com/holidays-api#docs) that takes two-letter country code 
and current year, month, and day numeric value. If it's holiday in given country, then it returns holiday name.
<br>The bot queries this API for three countries, viz. India (IN), USA (US), and Canada (CA).
If there are no holidays in these countries, then it checks for *Friday* to wish **Weekend**. Otherwise, it returns current weekday.

#### How to get GIF?
Twitter Celeb Bot uses [/peterhellberg/giphy](https://pkg.go.dev/github.com/peterhellberg/giphy) library to search GIFs from *Giphy*.
It takes keywords from *event* function, search GIFs, and randomly selects a GIF from first page of the response.


### Run it locally

#### Pre-requisite

- Git
- Go
- Twitter Credentials
- Holiday API Credential

#### Steps

- Clone this repo and change directory:
```
 git clone https://github.com/tejasdal/go-practice/ &&  cd go-practice/twitter-celeb/
```

- Set the following environment variable: 
```bigquery
export TWITTER_CREDENTIAL = '{{your_credentials}}'
export TWITTER_CREDENTIAL_SECRET = '{{your_credentials}}'
export TWITTER_ACCESS_TOKEN = '{{your_credentials}}'
export TWITTER_ACCESS_TOKEN_SECRET = '{{your_credentials}}'
export HOLIDAY_API_KEY = '{{your_credentials}}'
```

- Run main.go:
```bigquery
go run main.go
```