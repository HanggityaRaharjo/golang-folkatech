package main

import "fmt"

type Transaction struct {
    Buy  int
    Sell int
}

func MaxProfit(prices []int, k int) (int, []Transaction) {
    profit := 0
    transactions := []Transaction{}

    buyPrice := 0

    for i := 0; i < len(prices)-1; i++ {
        if buyPrice == 0 && prices[i] < prices[i+1] {
            buyPrice = prices[i]
        } else if buyPrice != 0 && prices[i] > prices[i+1] {
            sellPrice := prices[i]
            profit += sellPrice - buyPrice
            transactions = append(transactions, Transaction{Buy: buyPrice, Sell: sellPrice})
            buyPrice = 0
        }
    }

    if buyPrice != 0 {
        sellPrice := prices[len(prices)-1]
        profit += sellPrice - buyPrice
        transactions = append(transactions, Transaction{Buy: buyPrice, Sell: sellPrice})
    }

    if k < len(transactions) {
        transactions = transactions[:k]
    }

    return profit, transactions
}

func main() {
    prices := []int{4, 11, 2, 20, 59, 80}
    k := 2
    totalProfit, transactions := MaxProfit(prices, k)

    fmt.Printf("%d (", totalProfit)
    for i, trans := range transactions {
        fmt.Printf("buy:%d, sell:%d", trans.Buy, trans.Sell)
        if i < len(transactions)-1 {
            fmt.Print(", ")
        }
    }
    fmt.Println(")")
}
