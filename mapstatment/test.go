package mapstatment

import "fmt"

func TestMap() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap["china"] = "beijing"
	fmt.Println(countryCapitalMap)
}
