package main
import "fmt"
import "strings"

var string1, string2 string
var string1Slice, string2Slice, result []string
var sum int
func main() {

    string1 = "axbcde"
    string1Slice = strings.Split(string1, "")
    string2 = "acde"
    string2Slice = strings.Split(string2, "")
    
    fmt.Printf("String ke-1 : %s", string1)
    fmt.Println()
    fmt.Printf("String ke-2 : %s", string2)
    fmt.Println()
    
    for _, s1 := range string1Slice{
        sum = 0
        for _, s2 := range string2Slice{
            if s1==s2{
               sum++ 
            }
        }
        if sum == 0{
            result = append(result, s1)
        }
    }
    
    fmt.Println("Result : ")
    fmt.Print(result)
}