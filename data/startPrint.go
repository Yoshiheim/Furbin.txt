package data

import "fmt"

func PrintAllConfigs() {
	fmt.Printf("\t[Host%s. Port:%s]\t\n", Configs.Host, Configs.Port)
	fmt.Printf("\t\tCreated By '%s'\t\t\n", Configs.Name)
	fmt.Println(string(Logo))
}
