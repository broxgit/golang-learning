package shopping

import (
	"fmt"
	"littlegobook/shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	if item == nil {
		return 0, false
	}
	return item.Price, true
}

func testVisibility() {
	fmt.Println("This is visible!")
}
