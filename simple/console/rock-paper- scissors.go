package main

import (
	"fmt"
  "math/rand"
)


func main() {
	var game bool = true;
  
  var chose string;
  var botchose int = rand.Intn(3);
  
  for game {
    fmt.Println("Taş : 1");
    fmt.Println("Kağıt : 2");
    fmt.Println("Makas : 3");
    
    fmt.Scan(&chose);
    botchose = rand.Intn(3);
    //Berabere durumları
    if chose == "1" && botchose == 1{
      fmt.Println("Sen : Taş");
      fmt.Println("Rakip : Taş");
      fmt.Println("Berabere !!");
    } else if chose == "2" && botchose == 2 {
      fmt.Println("Sen : Kağıt");
      fmt.Println("Rakip : Kağıt");
      fmt.Println("Berabere !!");      
    } else if chose == "3" && botchose == 3 {
      fmt.Println("Sen : Makas");
      fmt.Println("Rakip : Makas");
      fmt.Println("Berabere !!");      
    } else if chose == "1" && botchose == 3 { //Kazanma durumları
      fmt.Println("Sen : Taş");
      fmt.Println("Rakip : Makas");
      fmt.Println("Kazandın !!");      
    } else if chose == "2" && botchose == 1 {
      fmt.Println("Sen : Kağıt");
      fmt.Println("Rakip : Taş");
      fmt.Println("Kazandın !!");      
    } else if chose == "3" && botchose == 2 {
      fmt.Println("Sen : Makas");
      fmt.Println("Rakip : Kağıt");
      fmt.Println("Kazandın !!");      
    } else if chose == "1" && botchose == 2 { //Kaybetme durumları
      fmt.Println("Sen : Taş");
      fmt.Println("Rakip : Kağıt");
      fmt.Println("Kaybettin !!");      
    } else if chose == "2" && botchose == 3 {
      fmt.Println("Sen : Kağıt");
      fmt.Println("Rakip : Makas");
      fmt.Println("Kaybettin !!");      
    } else if chose == "3" && botchose == 1 {
      fmt.Println("Sen : Makas");
      fmt.Println("Rakip : Taş");
      fmt.Println("Kaybettin !!");       
    }
  }
}
