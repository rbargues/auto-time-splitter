package main
import (
	"encoding/json"
	"os"
	"io/ioutil"
	"fmt"
)

func main() {
	// jsonFile, err := os.Open("worlds.json")
	// if err != nil {
	// 	panic(err)
	// }
	// defer jsonFile.Close()

	// byteVal, _ := ioutil.ReadAll(jsonFile)
	// var result []string
	// json.Unmarshal(byteVal, &result)
	// for i := 0; i < len(result); i++ {
	// 	fmt.Printf("%v\n",result[i])
	// }
	
	jsonFile, err := os.Open("levels.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteVal, _ := ioutil.ReadAll(jsonFile)
	var result map[string][]interface{}
	json.Unmarshal([]byte(byteVal), &result)
	// fmt.Printf("%v\n", result)
	for key, value := range result {
		fmt.Printf("%v: %v\n",key,len(value))
		for i := 0; i < len(value); i++ {
			fmt.Printf("%v\n", value[i])
		}
	}
}