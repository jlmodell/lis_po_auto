package main

import (
	"time"

	"github.com/gonutz/auto"
)

var PURCH_ORDERS = []string{
	"E30928",
	"E30929",
	"E30930",
	"E30931",
	"E30932",
	"E30933",
	"E30934",
	"E30935",
	"E30936",
	"E30937",
	"E30938",
	"E30939",
	"E30940",
	"E30941",
	"E30942",
	"E30943",
	"E30944",
	"E30945",
	"E30946",
	"E30947",
	"E30948",
	"E30949",
	"E30950",
	"E30951",
	"E30952",
}

func main() {
	// x, y, err := auto.MousePosition()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(x, y)

	for _, po := range PURCH_ORDERS {
		auto.ClickLeftMouseAt(500, 12)
		auto.Type("GET.VALUE.OF.LIS.PO")
		auto.PressKey(auto.KeyEnter)
		auto.Type(po)
		auto.PressKey(auto.KeyEnter)
		auto.PressKey(auto.KeyF3)
		time.Sleep(2 * time.Second)
		auto.ClickLeftMouseAt(1133, 179)
		auto.ClickLeftMouseAt(1133, 179)
		time.Sleep(3 * time.Second)
	}

}
