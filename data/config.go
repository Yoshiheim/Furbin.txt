package data

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/time/rate"
)

// All configs for 'config.json' for avoid hardcoding :3
type Config struct {
	// change view of pastebin here
	Port string `json:"port"` // port like 8080
	Host string `json:"host"` // its usually 127.0.0.1
	/*
		Name           string         `json:"name"`             // Name of Pastebinn
		Logo           LogoCfg        `json:"logo"`             // embed ASCII art as logo to website
		PasteDiv       CreatePasteDiv `json:"create_paste_div"` // Configs for div for created pastes on index page
		TextLogo       TextLogoCfg    `json:"text_logo"`        // embed ASCII text by figlet to the website
		Description    Descriptions   `json:"description"`      // Array of string as Description of pastebin
		CreatorsGithub string         `json:"creators_github"`  // Like to Your Github like "https://github.com/<CreatorsGithub>"
	*/
	Topics []Topic `json:"topics"` // All set topic of website
	Pastes []Paste `json:"pastes"` // All set paste of website

	// ect.
	PasteLens            LenOfPaste `json:"paste_lens"` // the limit of texts length in pastes
	CheckMemoryUsageTick string     `json:"check_memory_usage_tick"`
	Theme                uint       `json:"theme"`        // its should be for change color in pastebin but its doesn't work at all actually
	FaviconPath          string     `json:"favicon_path"` // Path of favicon of website, by default is "./data/favicon.ico"
	DBFilename           string     `json:"db_filename"`  // filename of sqlite database file, like "data.db"
	ClearTimer           ClearTimer `json:"clear_timer"`  // Timer for clear all pastes(but doesn't deletes pinned pastes, but you can change on '/HOXT/data/config/.json' by setting "delete-pinned" to true for delete pinned paste too)
	Limit                Limit      `json:"limit"`        // limit doesn't works, but idk, its already works, but you should hardcode it :(
}

type DynamicConfig struct {
	Name           string         `json:"name"`             // Name of Pastebinn
	Logo           LogoCfg        `json:"logo"`             // embed ASCII art as logo to website
	PasteDiv       CreatePasteDiv `json:"create_paste_div"` // Configs for div for created pastes on index page
	Description    Descriptions   `json:"description"`      // Array of string as Description of pastebin
	CreatorsGithub string         `json:"creators_github"`  // Like to Your Github like "https://github.com/<CreatorsGithub>"
}

type LenOfPaste struct {
	TitleLen   int `json:"title_len"`   // the len of paste's Title
	ContentLen int `json:"content_len"` // like and of paste's Content
	AuthorLen  int `json:"author_len"`  // and of name of Author too
}

type CreatePasteDiv struct {
	Hide          bool `json:"hide"`
	ForTopicIndex int  `json:"for_topic_index"`
}

type Descriptions struct {
	Hide bool     `json:"hide"` // hide descripting on the pastebin
	Text []string `json:"text"` // content
}

type LogoCfg struct {
	Hide    bool   `json:"hide"` // hide ASCII art logo on the pastebin
	Path    string `json:"path"` // and path to image, in the pastebin its "./data/cute_furry_raptor.png"
	Color   RGB    `json:"color"`
	Width   int    `json:"width"`  // width of ASCII art
	Heigth  int    `json:"heigth"` // and heigth too
	Size    int    `json:"size"`
	CharMap string `json:"charmap"` // charmap when image converted to ASCII art
}

type RGB struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

// "There we go, it should do something now, wow it didn't, why?...
// weird. Let's do this instead, okay that worked... yep, it crashed."
// - Notch™, 2011.
// (for content: https://www.youtube.com/watch?v=BES9EKK4Aw4&t=153s )
type Limit struct {
	LimitSec    rate.Limit `json:"limit_time"`
	LimitPerSec int        `json:"posts"`
}

// Clear All Paste every `Temp` var time
type ClearTimer struct {
	ClearPinned bool   `json:"destroy_pinned"`
	Temp        string `json:"tick"` // (its string because its for parse).

	//Temp        time.Duration `json:"tick"`
}

// Wanna make DB topic without inconvenience??!?/!/!??!
// Its should solve it.
// But i CANT make feature for update date when configs.json updated,
// so i you want change it, will delete you DB file :(
type Topic struct {
	Name        string `json:"name"`
	Description string `json:"descr"`
}

// SAME SHIT LIKE I SAID ON Topic!!11!!!!
type Paste struct {
	Title      string `json:"title"`   // paste's title
	Content    string `json:"content"` // and you know it
	TopicIndex uint   `json:"topic_index"`
	IsTitled   bool   `json:"is_titled"`
}

// var of configs from config.json file
var Configs Config

// The Logo of Pastebin's main page
var Logo []byte

// I hope you use figlet in linux or WSL in windows
var TextLogo []byte

// Embed configs values to "data.Configs" var.
func InitConfig(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(file, &Configs)
	if err != nil {
		log.Fatalln(err)
	}
	if Configs.Port == "" {
		Configs.Port = "127.0.0.1"
	}
	if Configs.Host == "" {
		Configs.Host = "8080"
	}
	fmt.Printf("[GET config.json - OK]\n")
	/*
		if Configs.Logo.Hide == false && Configs.Logo.Path != "" {

			flags := aic_package.DefaultFlags()
			if Configs.Logo.Heigth == 0 || Configs.Logo.Width == 0 {
				flags.Full = true
			}
			flags.Dimensions = []int{Configs.Logo.Width, Configs.Logo.Heigth}
			flags.Dither = true

			if Configs.Logo.CharMap == "" {
				flags.CustomMap = " .:-~+=*%&)@" //aic_package.DefaultFlags().CustomMap //" .:-~+=\"*%&)@"
			} else {
				flags.CustomMap = Configs.Logo.CharMap //Configs.Logo.CharMap
			}

			asciiArt, err := aic_package.Convert(Configs.Logo.Path, flags)
			if err != nil {
				Logo = []byte("")
				fmt.Println(err)
			} else {

				Logo = []byte(asciiArt)
			}

		} else {
			Logo = []byte("")
		}
	*/
	/*
		if Configs.TextLogo.Hide == false {
			f, err := os.ReadFile(Configs.TextLogo.File)
			if err != nil {
				TextLogo = []byte("")
			} else {
				TextLogo = f
			}
		}
	*/

	/*Logo, err = os.ReadFile(Configs.Logo.LogoPath)
	if err != nil {
		log.Fatalln(err)
	}
	*/
}

// TODO: REMOVE SHITCODE MOTHERFUCKER.
// idk why, its shitcode in GitHub ngl,
// so my OCD make me pretty bad when thinking about it.
// I WANT MAKE THE HOXT WEBSITE BETTER.

// node: oh, i need this :3
func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func LoadDynamicConfig(path string) (*DynamicConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config DynamicConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GetDConfig(w http.ResponseWriter) *DynamicConfig {
	cfg, err := LoadDynamicConfig("./data/textconf.json")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error with some file...", http.StatusInternalServerError)
		return nil
	}
	if cfg == nil {
		log.Println("CONFIG IS NULL")
		http.Error(w, "Error with some file...", http.StatusInternalServerError)
		return nil
	}
	return cfg
}
