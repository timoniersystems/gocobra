# gocobra
CLI example based on Cobra

## Set up project

`go mod init github.com/timoniersystems/gocobra`

## Install cobra-cli

`go install github.com/spf13/cobra-cli@latest`

## Initialize Cobra CLI and viper

`cobra-cli init --viper`

## Test run the app

`go run main.go`

## Add commands and subcommands

```
cobra-cli add joke
cobra-cli add fact
```

## Public APIs

Jokes:
```
curl  -H "Accept: application/json" https://api.chucknorris.io/jokes/random
curl -H "Accept: application/json" https://icanhazdadjoke.com/
curl  -H "Accept: application/json" https://itsthisforthat.com/api.php\?json
```

Facts:
```
curl  -H "Accept: application/json" https://meowfacts.herokuapp.com/  
curl  -H "Accept: application/json"  http://numbersapi.com/42/trivia\?json
```

News
```
https://inshorts.deta.dev/news?category=science

Categories:
business
sports
world
politics
technology
startup
entertainment
miscellaneous
hatke
science
automobile
```