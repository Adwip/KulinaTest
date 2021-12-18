package main
import "fmt"
import "strconv"

var value int

func main() {
    fmt.Print("Input : ")
    _, err := fmt.Scanf("%d", &value)
    if err != nil{
        fmt.Println(err.Error())
        return
    }
    var result = make([]string, value)

    for i:=0; i<value; i++{
        if (i+1)%3==0 && (i+1)%5==0{
            result[i] = "Kulina x Food"
        }else if (i+1)%3==0 {
            result[i] = "Kulina"
        }else if (i+1)%5==0{
            result[i] = "Food"
        }else{
            result[i] = strconv.Itoa(i+1)
        }
    }
    
    fmt.Printf("Output : ")
    
        fmt.Println(result)
}