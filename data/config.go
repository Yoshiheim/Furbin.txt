package data

import (
	"encoding/json"
	"log"
	"os"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"golang.org/x/time/rate"
)

// All configs for 'config.json' for avoid hardcoding :3
type Config struct {
	Port        string     `json:"port"`
	Host        string     `json:"host"`
	Name        string     `json:"name"`
	Description []string   `json:"Description"`
	CreatedBy   string     `json:"created_by"`
	Theme       uint       `json:"theme"`
	FaviconPath string     `json:"favicon_path"`
	DBFilename  string     `json:"db_filename"`
	Logo        LogoCfg    `json:"logo"`
	ClearTimer  ClearTimer `json:"clear_timer"`
	Limit       Limit      `json:"limit"`
	Topics      []Topic    `json:"topics"`
	Pastes      []Paste    `json:"pastes"`
}

type LogoCfg struct {
	Path    string `json:"logo_path"`
	Width   int    `json:"width"`
	Heigth  int    `json:"heigth"`
	CharMap string `json:"charmap"`
}

type Limit struct {
	LimitSec    rate.Limit `json:"limit_time"`
	LimitPerSec int        `json:"posts"`
}

// "There we go, it should do something now, wow it didn't, why?...
// weird. Let's do this instead, okay that worked... yep, it crashed."
// - Notch™, 2011.
// (for content: https://www.youtube.com/watch?v=BES9EKK4Aw4&t=153s )
type ClearTimer struct {
	ClearPinned bool   `json:"destroy_pinned"`
	Temp        string `json:"tick"`

	//Temp        time.Duration `json:"tick"`
}

// Wanna make DB topic without inconvenience? //
// Its should solve it.
// But i CANT make feature for update date when configs.json updated,
// so i you want change it, will delete you DB file :(
type Topic struct {
	Name        string `json:"name"`
	Description string `json:"descr"`
}

type Paste struct {
	Title      string `json:"title"`
	Content    string `json:"content"`
	TopicIndex uint   `json:"topic_index"`
	IsTitled   bool   `json:"is_titled"`
}

// var of configs from config.json file
var Configs Config

// The Logo of Pastebin's main page
var Logo []byte

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

	flags := aic_package.DefaultFlags()
	if Configs.Logo.Heigth == 0 || Configs.Logo.Width == 0 {
		flags.Full = true
	}
	flags.Dimensions = []int{Configs.Logo.Width, Configs.Logo.Heigth}
	//flags.Negative = true
	flags.Dither = true
	if Configs.Logo.CharMap == "" {
		flags.CustomMap = aic_package.DefaultFlags().CustomMap //" .:-~+=\"*%&)@"
	} else {
		flags.CustomMap = Configs.Logo.CharMap
	}

	asciiArt, err := aic_package.Convert(Configs.Logo.Path, flags)
	if err != nil {
		log.Fatalln(err)
	}
	Logo = []byte(asciiArt)

	/*Logo, err = os.ReadFile(Configs.Logo.LogoPath)
	if err != nil {
		log.Fatalln(err)
	}
	*/
}

func LoadLogo(c *Config) {
	// "<<", "furry_raptor.txt"

}

// TODO: REMOVE SHITCODE MOTHERFUCKER.
// idk why, its shitcode in GitHub ngl,
// so my OCD make me pretty bad when thinking about it.
// I WANT MAKE THE HOXT WEBSITE BETTER.
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
