package main
 
import (
    "fmt"
    "os"
    "bufio"
    "time"
    "encoding/csv"
)
 
const (
    PATH = "slog.csv"
    LAYOUT = "2006 01/02 15:04"
)
 
func recordLog(log string) {
    record := []string{}
    record = append(record, time.Now().Format(LAYOUT))
    record = append(record, log)
    f, err := os.OpenFile(PATH, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    w := csv.NewWriter(f)
    w.Write(record)
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