package main
import(
  "fmt"
)

func main(){
  var x float32 = 123.78
  var y float32 = 3.4e+38
  fmt.Printf("Veri tipi: %T, değeri: %v\n", x, x)
  fmt.Printf("Veri tipi: %T, değeri: %v", y, y)
}
