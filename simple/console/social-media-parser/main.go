package main

import (
  "fmt"
  "log"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)
func main(){

var scrap string;
  
for true{
  fmt.Println("Hangi Platformdan bilgi almak istiyorsunuz?");
  fmt.Println("1 | İnstagram");
  fmt.Println("2 | Github");
  fmt.Println("3 | Twitter");
  fmt.Scan(&scrap);
  switch(scrap){
    case "1":
  var kullanici string;
  fmt.Println("Kullanıcı adı giriniz");
  fmt.Scan(&kullanici);
 res, err := http.Get("https://www.instagram.com/"+kullanici)
  if err != nil {
    log.Fatal(err)
    return
  }
  doc, err := goquery.NewDocumentFromReader(res.Body)
  followers := doc.Find("div > div > div > div > div > div > div > div > div > div > section > main > div > header > section > ul > li:nth-child(2) > a > div > span").Text()
		fmt.Printf("Takipçi: %s\n",followers)
    break;
    case "2":
    var git string;
  fmt.Println("Kullanıcı adı giriniz");
  fmt.Scan(&git);
 res, err := http.Get("https://github.com/"+git)
  if err != nil {
    log.Fatal(err)
    return
  }
  doc, err := goquery.NewDocumentFromReader(res.Body)
  follower := doc.Find("div.flex-order-1.flex-md-order-none.mt-2.mt-md-0 > div > a:nth-child(1) > span").Text();
  following := doc.Find("div.flex-order-1.flex-md-order-none.mt-2.mt-md-0 > div > a:nth-child(2) > span").Text();
		fmt.Println("Takipçi: ",follower);
    fmt.Println("Takip edilen", following);
    break;
    case "3":
    var kullanici string;
    fmt.Println("Kullanıcı adı giriniz");
    fmt.Scan(&kullanici);
    res, err := http.Get("https://twitter.com/"+kullanici)
    if err != nil {
      log.Fatal(err)
      return
    }
    doc, err := goquery.NewDocumentFromReader(res.Body)
    follower := doc.Find("div > div > div > main > div > div > div > div > div > div:nth-child(3) > div > div > div > div > div > a > span > span").Text();
    fmt.Println("Takipçi : ",follower)
    break;
  }
  }
}
