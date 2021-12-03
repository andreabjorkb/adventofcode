package main

import (
  "fmt"
  "os"
  "errors"
  "bufio"
  "strings"
  "strconv"
)	

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

    f, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Can't open file:", os.Args[1])
        panic(err)
    }

    r := bufio.NewReader(f)
    line, err := Readln(r)

    course := map[string]int{"forward": 0, "up": 0, "down": 0, "depth": 0}
    for err == nil {
        split := strings.Split(line, " ") 
        if len(split) != 2 {
          err = errors.New("Reached end of input")
          break
        }

        key := split[0]
        val, err := strconv.Atoi(split[1])
        if err != nil {
          fmt.Println("Error converting value to int: %+v", split[1])
          panic(err)
        }

        course[key] += val
        
        if key == "forward" {
          course["depth"] += val*(course["down"]-course["up"])
        }
        //fmt.Println("=================")
        //fmt.Println(split)
        //fmt.Println(course)
        //fmt.Println("=================")

        line, err = Readln(r)
    }
    

    fmt.Println(course)
    fmt.Printf("Depth: %d\n", course["depth"])
    fmt.Printf("Forward: %d\n", course["forward"])
    fmt.Printf("Multiple: %d\n", course["depth"]*course["forward"])
}

func Readln(r *bufio.Reader) (string, error) {
  var (isPrefix bool = true
       err error = nil
       line, ln []byte
      )

  for isPrefix && err == nil {
      line, isPrefix, err = r.ReadLine()
      ln = append(ln, line...)
  }
  return string(ln),err
}