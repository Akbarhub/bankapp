package card

import(
	"git/pkg/bank/types"
	"fmt"
)


func ExampleTotal (){
	cards := []types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
		{
			Balance: 2_000_00,
			Active: true,
		},
		{
			Balance: 2_000_00,
			Active: false,
		},
	}
	fmt.Println(Total(cards))
	// Output: 300000
}