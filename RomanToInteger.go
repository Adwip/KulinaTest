package main
import "fmt"
import "strings"

var romanNumber = map[string]int{"I":1,"V":5,"X":10,"L":50,"C":100,"D":500,"M":1000}
var numberResult int
var romanSlice []string



func main() {
    romanValue := "LIV"
    fmt.Printf("Angka romannya adalah : %s", romanValue)
    fmt.Println()
    romanSlice = strings.Split(romanValue,"")
    for _, val := range romanSlice{
        numberResult += romanNumber[val]
    }
    fmt.Printf("Hasil angkanya adalah : %d", numberResult)
}