package main
import ("encoding/csv"
	"os"
	"fmt"
	"flag"
	"bufio"
	"reflect"
	"strings"
	"time"
	"strconv"

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
	timeLimit := flag.String("time", "", "optional time limit, seconds")
	flag.Parse()
	fmt.Println("starting quiz")
	questions := read_quiz(*filePath)

	if *timeLimit != "" {
        seconds, err := strconv.Atoi(*timeLimit)
        if err != nil {
            // handle error
            fmt.Println(err)
            os.Exit(2)
        }
        durationSeconds := time.Duration(seconds)
        fmt.Println("setting timer ", seconds)

        quitter := make(chan bool)
	    timedQuestions(quitter, questions)
	    time.Sleep(durationSeconds * time.Second)
	    quitter <- true
	    fmt.Println("exiting")
	}

	//askQuestions(questions)

}

