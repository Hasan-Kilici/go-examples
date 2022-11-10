package main

import (
	"fmt"
)

func main() {
 var chose string
 var num1 int
 var num2 int
 var result int
  
 fmt.Println("İşlem seçiniz");
 fmt.Println("- Toplama : +");
 fmt.Println("- Çıkarma : -");
 fmt.Println("- Çarpma : *");
 fmt.Println("- Bölme : /");
 fmt.Scan(&chose);

 switch(chose){
   case "+":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 + num2;
   fmt.Println(result);
   break;
   case "-":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 - num2;
   fmt.Println(result);
   break;
   case "*":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 * num2;
   fmt.Println(result);
   break;
   case "/":
   fmt.Println("  1.Sayıyı giriniz");
   fmt.Scan(&num1);
   fmt.Println("  2.Sayıyı giriniz");
   fmt.Scan(&num2);
   result = num1 / num2;
   fmt.Println(result);
   break;
 }
}
