package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	"github.com/bartholdbos/golegram"
	"github.com/gin-gonic/gin"
)

//Webhook represents push information from the webhook
type Webhook struct {
	Before, After, Ref, User_name string
	User_id, Project_id           int
	Repository                    RepositoryInterface
	Commits                       []Commit
	Total_commits_count           int
	Pusher                        Author
}

//Config struct represents the config filé
type Config struct {
	Bottoken     string
	Address      string
	Port         int64
	Repositories []ConfigRepository
}

//ConfigRepository represents a repository from the config file
type ConfigRepository struct {
	Clone_url      string
	Telegramtarget string
	Commands       []string
}

func PanicIf(err error, what ...string) {
	if err != nil {
		if len(what) == 0 {
			panic(err)
		}

		panic(errors.New(err.Error() + what[0]))
	}
}

var config Config
var configFile string

//TODO: toml instead of json
func loadConfig(configFile string) Config {
	var file, err = os.Open(configFile)
	PanicIf(err)

	defer func() {
		err := file.Close()
		PanicIf(err)
	}()

	buffer := make([]byte, 1024)
	count := 0

	count, err = file.Read(buffer)
	PanicIf(err)

	err = json.Unmarshal(buffer[:count], &config)
	PanicIf(err)

	return config
}

func processWebhook(wh Webhook) {

	bot, _ := golegram.NewBot(config.Bottoken)
	//our beautiful git push sticker
	// nope
	//bot.SendSticker(userID, "BQADBAADjgEAAlIvPQAB7_h8b5RYj3sC")
	gitSystem := "gogs"

	if len(wh.Repository.Full_name) > 0 {
		gitSystem = "github"
	}

	var fullRepo string

	if gitSystem == "github" {
		fullRepo = wh.Repository.Full_name
	} else {
		fullRepo = wh.Repository.Owner.Username + "/" + wh.Repository.Name
	}

	for _, repo := range config.Repositories {

		if repo.Clone_url != wh.Repository.Clone_url {
			continue
		}
		branch := strings.Replace(wh.Ref, "refs/heads/", "", -1) //doe iets wegstrippen later

		announceMsg := "⬆️ " + wh.Pusher.Name + " pushed to _" + branch + "_ of _" + fullRepo + "_"

		_, _ = bot.SendMessage(repo.Telegramtarget, announceMsg)

		for _, commit := range wh.Commits { //for each commit in our received json, show a seperate message
			msg := "➕ [" + commit.Message + "](" + commit.Url + ")\n"
			msg += "👤 " + commit.Author.Name
			msg += ""
			_, _ = bot.SendMessage(repo.Telegramtarget, msg)
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP)
	go func() {
		<-sigc
		config = loadConfig(configFile)
		log.Println("config reloaded")
	}()

	config := loadConfig("config.json")

	r.POST("/", func(c *gin.Context) {
		var json Webhook
		if c.BindJSON(&json) == nil {
			processWebhook(json)
			c.String(http.StatusOK, "Ok.")
		} else {
			c.String(401, "Nope.")
		}

	})
	address := config.Address + ":" + strconv.FormatInt(config.Port, 10)
	r.Run(address)
}
