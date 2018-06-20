package main

import ("fmt" 
		"os"
		"log")
		
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Printf("Trouble has hit while reading the test.txt file" + err.Error())
		return 
	}
	//delay this call to the end of this block of code - main()
	defer file.Close()
	stat, err := file.Stat()
	if(err != nil){
		log.Printf("Touble getting stats " + err.Error())
		return
	}
	bytes := make([]byte, stat.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return
	}
	value := string(bytes)
	fmt.Println(value)
}