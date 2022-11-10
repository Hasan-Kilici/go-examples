package main

import (
	"fmt"
  "math/rand"
  "strings"
)

func main() {
	var name string
  var i bool = true
  var command string
  
  var animals = []string{}

  var animalList = []string{
  "Örümcek",
  "Ahtapot",
  "Kalamar",
  "Koyun",
  "İnek",
  "Salyangoz",
  "Aslan",
  "Kaplan",
  "Çita",
  "Jaguar",
  "Solucan",
  "Köstebek",
  "Fil",
  "Deve",
  "Arap",
  "Eşek",
  "At",
  "Unicorn",
  "Domuz",
  "Kelebek",
  "Baykuş",
  "Kartal",
  "Yengeç",
  "Sincap",
  "Kedi",
  "Papağan",
  "Güvercin",
  "Serçe",
  "Örümcek",
  "Ahtapot",
  "Kalamar",
  "Koyun",
  "İnek",
  "Salyangoz",
  "Aslan",
  "Kaplan",
  "Çita",
  "Jaguar",
  "Solucan",
  "Köstebek",
  "Fil",
  "Deve",
  "Arap",
  "Eşek",
  "At",
  "Unicorn",
  "Domuz",
  "Kelebek",
  "Baykuş",
  "Kartal",
  "Yengeç",
  "Sincap",
  "Kedi",
  "Papağan",
  "Güvercin",
  "Serçe",
  "Örümcek",
  "Ahtapot",
  "Kalamar",
  "Koyun",
  "İnek",
  "Salyangoz",
  "Aslan",
  "Kaplan",
  "Çita",
  "Jaguar",
  "Solucan",
  "Köstebek",
  "Fil",
  "Deve",
  "Arap",
  "Eşek",
  "At",
  "Unicorn",
  "Domuz",
  "Kelebek",
  "Baykuş",
  "Kartal",
  "Yengeç",
  "Sincap",
  "Kedi",
  "Papağan",
  "Güvercin",
  "Serçe",
  "Dinazor",
  }
 
  var para int = 0; 
  var rastpar int = 100;
  var olta bool = false;
  var ag bool = false;
  var upgpara int = 400;
  
  fmt.Println("-------------------");
  fmt.Println("    İsim nedir?    ");
  fmt.Println("-------------------");
  fmt.Scan(&name);
  for(i){
   fmt.Println("-------------------");
   fmt.Println("   Komut giriniz   ");
   fmt.Println("-------------------");
   fmt.Println("      KOMUTLAR     ");
   fmt.Println("-------------------");
   fmt.Println("");
   fmt.Println("avlan");   
   fmt.Println("hayvanlarım");
   fmt.Println("param");
   fmt.Println("market");
   fmt.Println("satin-al-esyaadi");  
   fmt.Println("güçlendir");
   fmt.Println("");
   fmt.Println("-------------------");
   fmt.Scan(&command);
   switch(strings.ToLower(command)){
     case "avlan":
     var kazanilanpara int = rand.Intn(rastpar);
     hunted := rand.Intn(len(animalList));
     if(olta == false){
     animals = append(animals, animalList[hunted]);
     para = para + kazanilanpara;
     } 
     if(olta == true) {
     hunted2 := rand.Intn(len(animalList));
     kazanilanpara = rand.Intn(rastpar)*2;
     animals = append(animals, animalList[hunted], animalList[hunted2]);
     fmt.Println("");
     fmt.Println("Tebrikler",animalList[hunted2],"Yakaladın!");
     fmt.Println("");   
     }
     if(ag == true){
      kazanilanpara = rand.Intn(rastpar);
     }
     
     animals = append(animals, animalList[hunted]);
     para = para + kazanilanpara;
     fmt.Println("");
     fmt.Println("Tebrikler",animalList[hunted],"Yakaladın!");
     fmt.Println("Kazanılan para :", kazanilanpara);
     fmt.Println("");
     
     break;
     case "hayvanlarım":
     fmt.Println("");
     fmt.Println(animals);
     fmt.Println("");
     break;
     case "param":
     fmt.Println("");
     fmt.Println(para,"TL");
     fmt.Println("");
     break;
     case "market":
     fmt.Println("");
     fmt.Println("- Ağ | Daha Çok para kazanmanızı sağlar [2x] | 200 TL");
     fmt.Println("- Olta | Daha çok hayvan yakalamanızı sağlar [2x] | 500 TL");
     fmt.Println("");
     break;
     case "satin-al-olta":
     if(para > 499){
     fmt.Println("");
     fmt.Println("Olta Satın alındı");
     olta = true;
     fmt.Println("");
     } else {
     fmt.Println("");
     fmt.Println("Yetersiz bakiye!");
     fmt.Println("");
     }
     break;
     case "satin-al-ag":
     if(para > 199){
     fmt.Println("");
     fmt.Println("Ağ Satın alındı");
     ag = true;
     rastpar = 200;
     fmt.Println("");
     } else {
     fmt.Println("");
     fmt.Println("Yetersiz bakiye!");
     fmt.Println("");
     }
     break;
     case "güçlendir":
     fmt.Println("");
     fmt.Println("para-arttır | Kazanılan parayı arttır | ",upgpara, "TL");
     fmt.Println("");
     break;
     case "para-arttır":
     if(para >= upgpara){
       upgpara = upgpara + rand.Intn(upgpara*4);
       rastpar = rastpar + rand.Intn(rastpar*2);
       fmt.Println("");
       fmt.Println("Para Arrttırma seviye atladı!");
       fmt.Println("Artık max",rastpar," Para kazanacaksın!")
       fmt.Println("");
     } else {
      fmt.Println("");
      fmt.Println("Yetersiz bakiye!");
      fmt.Println("");
     }
     break;
   }
  }
}
