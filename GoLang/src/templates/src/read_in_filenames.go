package main

import(
	"fmt"
	"io/ioutil"
	"log"
)
// Define constant for filepath. path should be in form "C:/way_to_my_path"
const adr string = "C:/absolute/path/to/src/fails"

func get_file_names(par_address string)[]string  { // define function that takes in string and gives out strind type slice

	var sl_filenames []string // define empty string type slice

	// func ReadDir(dirname string) ([]os.FileInfo, error)
	files, err := ioutil.ReadDir(par_address)
	if err != nil {
		log.Fatal(err)
	}

	// add filenames to slice
	for _, file := range files{
		sl_filenames = append(sl_filenames, file.Name())
	}
	// return final appended slice
	return sl_filenames
}

func main()  {
	fmt.Println(get_file_names(adr))

}
