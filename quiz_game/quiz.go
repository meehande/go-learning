package main
import ("encoding/csv"
	"os"
	"fmt"
	"flag"
	"bufio"
	"reflect"
	"strings"
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
	flag.Parse()
	fmt.Println("starting quiz")
	problems := read_quiz(*filePath)
	askQuestions(problems)

}

