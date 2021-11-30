package main
import ("encoding/csv"
	"os"
	"fmt"
	"flag"
	"bufio"
	"strings"
	"time"
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

func askTimedQuestions(questions []problem, timeLimit int){
    timer := time.NewTimer(time.Duration(timeLimit)*time.Second)
	correct := 0

	for _, p := range questions {
        fmt.Println(p.q)

        answerChannel := make(chan string)
        go func(){
            var answer string
            fmt.Scanf("%s\n", answer)
            answerChannel <- answer
        }()


	    select{
        case <-timer.C:
            fmt.Printf("Score: %d/%d", correct, len(questions))
            return
        case answer := <- answerChannel:
            if answer == p.a {
                correct += 1
            }
	    }


	}

	fmt.Println("result: ", correct, "/", len(questions))
}

func main(){
	filePath := flag.String("file", "problems.csv", "file containing quiz questions in csv format")
	timeLimit := flag.Int("time", 30, "optional time limit, seconds")
	flag.Parse()
	fmt.Printf("Starting your timed quiz! you have %d seconds to complete the problems, good luck!\n", *timeLimit)
	problems := read_quiz(*filePath)
	askTimedQuestions(problems, *timeLimit)

}

