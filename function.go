package main
 
import "fmt"
 
func main() {
    fmt.Println(add(20, 30))
}
 
func add(x int, y int) int {
    total := 0
    total = x + y
    return total
}
