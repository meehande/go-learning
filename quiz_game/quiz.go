package main
import ("encoding/csv"
	"os"
	"fmt"
	"flag"
	"bufio"
	"reflect"
	"strings"
	//"time"
	//"strconv"

)

type problem struct {
    q string
    a string
}

func parseProblems(lines [][]string) []problem{
    problems := make([]problem, len(lines))
    for i, qa := range lines{
        problems[i] = problem{
            q: qa[0],
            a: qa[1],
        }
    }
    return problems
}

func read_quiz(filepath string)[]problem{
	csvFile, e := os.Open(filepath)
	if e != nil {
		fmt.Print(e)
	}
	r := csv.NewReader(csvFile)

	records, e := r.ReadAll()
	if e != nil {
		fmt.Print(e)
	}
	return parseProblems(records)
}

func readInput() string{
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimRight(text, "\n")
}

func askQuestions(questions []problem){
	correct := 0

	for _, p := range questions {
		fmt.Println(p.q)
		inputAns := readInput()
		fmt.Println(inputAns)
		
		if inputAns == p.a {
			fmt.Println("correct")
			correct += 1
		} else {
			fmt.Println("incorrect, expected ", p.a)
			fmt.Println(reflect.TypeOf(p.a))
			fmt.Println(reflect.TypeOf(inputAns))
		}
	}
	
	fmt.Println("result: ", correct, "/", len(questions))
}

func main(){
	filePath := flag.String("file", "problems.csv", "file containing quiz questions in csv format")
	//timeLimit := flag.String("time", "", "optional time limit, seconds")
	flag.Parse()
	fmt.Println("starting quiz")
	problems := read_quiz(*filePath)

// 	if *timeLimit != "" {
//         seconds, err := strconv.Atoi(*timeLimit)
//         if err != nil {
//             // handle error
//             fmt.Println(err)
//             os.Exit(2)
//         }
//         durationSeconds := time.Duration(seconds)
//         fmt.Println("setting timer ", seconds)
//
//         quitter := make(chan bool)
// 	    timedQuestions(quitter, questions)
// 	    time.Sleep(durationSeconds * time.Second)
// 	    quitter <- true
// 	    fmt.Println("exiting")
// 	}
    //fmt.Println(problems)
	askQuestions(problems)

}

