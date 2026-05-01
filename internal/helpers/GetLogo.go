package helpers

import (
	"fmt"
	"hoxt/data"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

func UpdateLogo() {
	c, err := data.LoadDynamicConfig("./data/textconf.json")
	if err != nil {
		fmt.Println(err)
	}
	flags := aic_package.DefaultFlags()

	flags.Dimensions = []int{c.Logo.Heigth, c.Logo.Width}
	if c.Logo.CharMap == "" {
		flags.CustomMap = " .:-~+=*%&)@"
	} else {
		flags.CustomMap = c.Logo.CharMap
	}
	flags.Full = false
	flags.Dither = true
	asciiArt, err := aic_package.Convert(c.Logo.Path, flags)
	if err != nil {
		fmt.Printf("-- %s --\n", err)
		data.Logo = []byte("")
	} else {
		fmt.Printf("[CONVERT IMAGE TO ASCIII - OK]\n")
		data.Logo = []byte(asciiArt)
	}
}
