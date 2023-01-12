package main

import(
	"git/pkg/bank/card"
	"git/pkg/bank/types"
	"fmt"
)

func main() {
	cs := []types.Card{
		{
			Balance: 1000,
			Active:  true,
		},
		{
			Balance: 1000,
			Active:  true,
		},
		{
			Balance: 1000,
			Active:  true,
		},
		{
			Balance: -1000,
			Active:  true,
		},
		{
			Balance: 1000,
			Active:  false,
		},
	}
	sum := card.Total(cs)
	fmt.Println(sum)
}