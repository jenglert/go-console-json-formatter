package main

import "os"
import "fmt"
import "io/ioutil"
import "encoding/json"

func printIndent(indentLevel int ) {
	// Print the indent
	for i := 0; i < indentLevel; i++ {
	  fmt.Print(" ")
	}
}

func printWithIndent(indentLevel int, str string) {
	printIndent(indentLevel)
	fmt.Print(str)
}

func printlnWithIndent(indentLevel int, str string) {
	printIndent(indentLevel)
	fmt.Println(str)
}

func colorPunc(str string) string {
  return "\033[33m" + str + "\033[0m"
}



func printValue(value interface{}, indentLevel int) {
	switch vTyped := value.(type) {
		case string:
			fmt.Println("\"" + vTyped + "\"")
		case float64:
			fmt.Println(vTyped)
		case map[string]interface{}:
			printlnWithIndent(indentLevel, colorPunc("{"))
			printMap(vTyped, indentLevel + 2)
			printlnWithIndent(indentLevel, colorPunc("}"))
		case nil:
			fmt.Println("null")
		case []interface{}:
			fmt.Println(colorPunc("["))
			for _, av := range vTyped {
				printIndent(indentLevel + 2)
				printValue(av, indentLevel + 2)
			}
			printlnWithIndent(indentLevel, colorPunc("]"))
		default:
			fmt.Println(fmt.Sprintf("default %T", value))

	}

}

func printMap(m map[string]interface{}, indentLevel int) {

	for key, value := range m {

		printWithIndent(indentLevel, "\"\033[35m" + key + "\033[0m\" : ")

		printValue(value, indentLevel)
	}
}

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)

	var f interface{}
	jsonErr := json.Unmarshal(bytes, &f)

	if jsonErr != nil {
		fmt.Println("Unable to parse json: ", jsonErr)
	}

	fmt.Println("{")
	printMap(f.(map[string]interface{}), 2)
	fmt.Println("}")

}