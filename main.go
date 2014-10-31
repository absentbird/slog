package main
 
import (
    "fmt"
    "os"
    "bufio"
    "time"
    "encoding/csv"
)
 
const (
    PATH = "slog.log"
    LAYOUT = "2006 01/02 15:04"
)

type record struct {
    message string
    time time.Time
}

func (r record) Line() []string {
    line := []string{}
    line = append(line, r.time.Format(LAYOUT))
    line = append(line, r.message)
    return line
}

func recordLog(log string) {
    r := new(record)
    r.time = time.Now()
    r.message = log
    f, err := os.OpenFile(PATH, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    w := csv.NewWriter(f)
    w.Write(r.Line())
    w.Flush()
}
 
func main() {
    f, err := os.Stdin.Stat()
    if err != nil {
        panic(err)
    }
    if len(os.Args) > 1 {
        recordLog(string(os.Args[1]))
        return
    }
    reader := bufio.NewReader(os.Stdin)
    if f.Mode() & os.ModeNamedPipe == 0 {
        fmt.Print("Input record: ")
    }
    input, _, err := reader.ReadLine()
    if err != nil {
        panic(err)
    }
    recordLog(string(input))
    return
}
