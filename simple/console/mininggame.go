package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var game bool = true

	var madenler = []string{"Emrald", "Diamond", "Diamond", "Gold", "Gold", "Gold", "Iron", "Iron", "Iron", "Iron", "Bronz", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Silver", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Coal", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone", "Stone"}

	var name string
	var command string
	var buy string
	var upgrade string

	var para int = 0
	var kazanc int = 100

	var kazilacaklar int = 1

	var mineupg int = 400
	var cashupg int = 200
	var minecountupg int = 200

	type picaxes struct {
		wooden  bool
		stone   bool
		iron    bool
		silver  bool
		gold    bool
		diamond bool
		emrald  bool
	}

	pickaxe := picaxes{
		wooden:  false,
		stone:   false,
		iron:    false,
		silver:  false,
		diamond: false,
		emrald:  false,
	}

	fmt.Println("What is your name gamer?");
	fmt.Scan(&name);
	for game {

		fmt.Println("_-_ Enter Command _-_");
		fmt.Scan(&command);
		switch command {
		case "mine":
			i := 1
			for i <= kazilacaklar {
				kazilan := madenler[rand.Intn(len(madenler))];
        switch(kazilan){
        case "Stone":
        kazancin := rand.Intn(kazanc);
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Coal":
        kazancin := rand.Intn(kazanc-40)*2;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Iron":
        kazancin := rand.Intn(kazanc-40)*3;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Silver":
        kazancin := rand.Intn(kazanc-40)*4;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Gold":
        kazancin := rand.Intn(kazanc-40)*5;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Diamond":
        kazancin := rand.Intn(kazanc-40)*6;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        case "Emrald":
        kazancin := rand.Intn(kazanc-40)*10;
				fmt.Println("WoW u find", kazilan, "!!");
				fmt.Println(kazancin, "money deposited");
				para = para + kazancin;
				i = i + 1;
        break;
        }
			}
			break
		case "market":
			fmt.Println("")
			fmt.Println("_-_  Pickaxes _-_")
			fmt.Println("1 >   Stone Pickaxes | 200 Coin");
			fmt.Println("2 >   Iron Pickaxe | 2000 Coin");
			fmt.Println("3 >   Silver Picaxe | 8000 Coin");
			fmt.Println("4 >   Gold Picaxe | 10000 Coin");
			fmt.Println("5 >   Diamond Picaxe | 17500 Coin");
			fmt.Println("6 >   Emrald Picaxe | 58000 Coin");
			fmt.Println("_-_ Upgrade List _-_");
			fmt.Println("7 >   Mining |", mineupg, "Coin");
			fmt.Println("8 >   Cash |", cashupg, "Coin");
			fmt.Println("9 >   Mining Count", minecountupg, "Coin");
			fmt.Println("")
			break
		case "upgrade":
			fmt.Println("_-_ Upgrade List _-_");
			fmt.Println("1 >   Mining |", mineupg, "Coin");
			fmt.Println("2 >   Cash |", cashupg, "Coin");
			fmt.Println("3 >   Mining Count", minecountupg, "Coin");
			fmt.Scan(&upgrade)
      switch(upgrade){
        case "1":
        if(para >= mineupg){
          para = para - mineupg;
          kazanc = kazanc + rand.Intn(kazanc);
          kazilacaklar = kazilacaklar + 1; 
          mineupg = mineupg + rand.Intn(mineupg*4);
          fmt.Println("Upgraded Cash, now you earn",kazanc,"per mine");
          fmt.Println("Mining Count upgraded, Now you mining",kazilacaklar,"Times!!");
        } else {
					fmt.Println("You have not money");          
        }
        break;
        case "2":
        if(para >= cashupg){
          para = para - cashupg;
          kazanc = kazanc + rand.Intn(kazanc);
          cashupg = cashupg + rand.Intn(cashupg);
          fmt.Println("Upgraded Cash, now you earn",kazanc,"per mine");
        } else {
 					fmt.Println("You have not money");         
        }
        break;
        case "3":
        if(para >= minecountupg){
          para = para - minecountupg;
          kazilacaklar = kazilacaklar + 1;
          minecountupg = minecountupg + rand.Intn(minecountupg);
          fmt.Println("Mining Count upgraded, Now you mining",kazilacaklar,"Times!!");
        } else {
					fmt.Println("You have not money");          
        }
        break;
      }
			break
		case "buy":
			fmt.Println("Select a Item");
			fmt.Println("If you don't know what to buy try the market command");
			fmt.Println("1 >   Stone Pickaxes | 200 Coin");
			fmt.Println("2 >   Iron Pickaxe | 2000 Coin");
			fmt.Println("3 >   Silver Picaxe | 8000 Coin");
			fmt.Println("4 >   Gold Picaxe | 10000 Coin");
			fmt.Println("5 >   Diamond Picaxe | 17500 Coin");
			fmt.Println("6 >   Emrald Picaxe | 58000 Coin");
			fmt.Scan(&buy)
			switch buy {
			case "1":
				if para >= 200 && pickaxe.stone == false {
					pickaxe.stone = true;
					para = para - 200;
          kazilacaklar = kazilacaklar + 1;
					fmt.Println("You Got a Stone Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "2":
				if para >= 2000 && pickaxe.iron == false {
					pickaxe.iron = true;
          para = para - 2000;
          kazilacaklar = kazilacaklar + 2;
					fmt.Println("You Got a Iron Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "3":
				if para >= 8000 && pickaxe.silver == false {
					pickaxe.silver = true;
          para = para - 8000;
          kazilacaklar = kazilacaklar + 3;
					fmt.Println("You Got a Sılver Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "4":
				if para >= 10000 && pickaxe.gold == false {
          pickaxe.gold = true;
          para = para - 10000;
          kazilacaklar = kazilacaklar + 4;
					fmt.Println("You Got a Gold Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "5":
				if para >= 17500 && pickaxe.diamond == false {
          pickaxe.diamond = true;
          para = para - 175000;
          kazilacaklar = kazilacaklar + 6;
					fmt.Println("You Got a Diamond Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			case "6":
				if para >= 58000 && pickaxe.emrald == false {
          pickaxe.emrald = true;
          para = para - 58000;
          kazilacaklar = kazilacaklar + 10;
					fmt.Println("You Got a Emrald Pickaxe");
				} else {
					fmt.Println("You have not money");
				}
				break
			}
			break
		case "cash":
			fmt.Println("You have a", para, "Coin");
			break
		}
	}
}
