package main
import (
	"fmt"
	"reflect"
	"log"
)

func checkTimeUp(quitChannel chan bool) int{

        select {
            case <- quitChannel:
                return -1
            default:
                return 0
            }
}


func timedQuestions(ch chan bool, questions [][]string){

	go func(){
        correct := 0

        for _, qa := range questions {
            log.Println("timer: asking question")
            select {
                case <-ch:
                    fmt.Println("result: ", correct, "/", len(questions))
                    return
                default:
                    fmt.Println("time not up")
            }

//             timeUp := checkTimeUp(ch)
//             if timeUp < 0 {
//                 fmt.Println("result: ", correct, "/", len(questions))
//                 return
//             }

            q := qa[0]
            //a := qa[1]
            fmt.Println(q)
            inputAns := readInput()

            select {
                case <-ch:
                    fmt.Println("result: ", correct, "/", len(questions))
                    return
               default:
                fmt.Println("time not up")
            }

//             timeUp := checkTimeUp(ch)
//             if timeUp < 0 {
//                 fmt.Println("result: ", correct, "/", len(questions))
//                 return
//             }
            fmt.Println(inputAns)

            if inputAns == qa[1] {
                //fmt.Println("correct")
                correct += 1
            } else {
                //fmt.Println("incorrect, expected ", qa[1])
                fmt.Println(reflect.TypeOf(qa[1]))
                fmt.Println(reflect.TypeOf(inputAns))
            }
        }
        fmt.Println("timer: result: ", correct, "/", len(questions))

    }()

}