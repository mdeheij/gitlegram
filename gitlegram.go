package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bartholdbos/golegram"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

//Repository represents repository information from the webhook
type Repository struct {
	Name        string
	Url         string
	Description string
	Home        string
}

//Commit represents commit information from the webhook
type Commit struct {
	Id        string
	Message   string
	Timestamp string
	Url       string
	Author    Author
}

//Author represents author information from the webhook
type Author struct {
	Name  string
	Email string
}

//Webhook represents push information from the webhook
type Webhook struct {
	Before, After, Ref, User_name string
	User_id, Project_id           int
	Repository                    Repository
	Commits                       []Commit
	Total_commits_count           int
}

//Config struct represents the config filé
type Config struct {
	Logfile      string
	Bot          string
	Address      string
	Port         int64
	Repositories []ConfigRepository
}

//ConfigRepository represents a repository from the config file
type ConfigRepository struct {
	Name           string
	Telegramtarget int32
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

func main() {
	args := os.Args

	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP)

	go func() {
		<-sigc
		config = loadConfig(configFile)
		log.Println("config reloaded")
	}()

	//if we have a "real" argument we take this as conf path to the config file
	if len(args) > 1 {
		configFile = args[1]
	} else {
		configFile = "config.json"
	}

	config := loadConfig(configFile)
	writer, err := os.OpenFile(config.Logfile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	PanicIf(err)

	//close logfile on exit
	defer func() {
		writer.Close()
	}()

	log.SetOutput(writer)
	http.HandleFunc("/", hookHandler)

	address := config.Address + ":" + strconv.FormatInt(config.Port, 10)
	log.Println("Listening on " + address)

	//starting server
	err = http.ListenAndServe(address, nil)
	if err != nil {
		log.Println(err)
	}
}

func hookHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	var hook Webhook
	//read request body
	var data, err = ioutil.ReadAll(r.Body)
	PanicIf(err, "while reading request")

	//fmt.Println(string(data))
	//unmarshal request body

	err = json.Unmarshal(data, &hook)
	PanicIf(err, "while unmarshaling request")

	fmt.Println(hook)

	//find matching config for repository name
	for _, repo := range config.Repositories {

		if repo.Name != hook.Repository.Name {
			continue
		}
		bot, _ := Golegram.NewBot(config.Bot)
		//our beautiful git push sticker
		bot.SendSticker(repo.Telegramtarget, "BQADBAADjgEAAlIvPQAB7_h8b5RYj3sC")

		//find out which branch by replacing with nothing
		branch := strings.Replace(hook.Ref, "refs/heads/", "", -1) //doe iets wegstrippen later
		announceMsg := "➡️ " + hook.User_name + " pushed to " + branch + " of IS203-4"
		_, _ = bot.SendMessage(repo.Telegramtarget, announceMsg)

		for _, commit := range hook.Commits { //for each commit in our received json, show a seperate message
			msg := "➕ " + commit.Message
			msg += " (" + commit.Author.Name
			msg += ")"
			fmt.Println(msg)
			_, _ = bot.SendMessage(repo.Telegramtarget, msg)
		}
	}

}
