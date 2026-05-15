package main    
import "fmt"

type Wallet struct{
    Address string
    Balance string
}

func add(a,b int) int{
    return a+b
}

func main(){
    fmt.Println("================day 1==============")
    myWallet := Wallet{
        Address:"0x12672980183",
        Balance:"0.1ETH",
    }
    fmt.Println("my first wallet")
    fmt.Println("address:",myWallet.Address)
    fmt.Println("balance:",myWallet.Balance)

    sum:=add(10,20)
    fmt.Println("10+20=",sum)
}
