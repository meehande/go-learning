package main
import ("encoding/csv"
	"os"
	"fmt"
	"flag"
	"bufio"
	"reflect"
	"strings"
)

func read_quiz(filepath string)[][]string{
	/*
	dat, e := os.ReadFile(filepath)
	if e != nil {
		panic(e)
	}

	fmt.Print(string(dat))
	*/
	csvFile, e := os.Open(filepath)
	if e != nil {
		fmt.Print(e)
	}
	r := csv.NewReader(csvFile)

	records, e := r.ReadAll()
	if e != nil {
		fmt.Print(e)
	}
	//fmt.Print(records)
	return records
}

func readInput() string{
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\n")
}

func askQuestions(questions [][]string){
	correct := 0

	for _, qa := range questions {
		q := qa[0]
		//a := qa[1]
		fmt.Println(q)
		inputAns := readInput()
		fmt.Println(inputAns)
		
		if inputAns == qa[1] {
			fmt.Println("correct")
			correct += 1
		} else {
			fmt.Println("incorrect, expected ", qa[1])
			fmt.Println(reflect.TypeOf(qa[1]))
			fmt.Println(reflect.TypeOf(inputAns))
		}
	}
	
	fmt.Println("result: ", correct, "/", len(questions))
}

func main(){
	filePath := flag.String("file", "problems.csv", "file containing quiz questions. defaults to problems.csv")
	flag.Parse()
	fmt.Println("starting quiz")
	questions := read_quiz(*filePath)
	askQuestions(questions)
}
