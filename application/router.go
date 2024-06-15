package application

import (
	"fmt"
	"yo/application/handlers"
	"yo/pkg/logic"
	data "yo/pkg/repository/data"
)

func StartReading() {
	parkingRepo := data.NewParkingRepo()
	parkingLotLogic := logic.NewParkingLot(parkingRepo)
	handler := handlers.NewApiHandler(parkingLotLogic)
	var commandMap = map[string]func(w2, w3, w4 string){
		"create":       handler.Create,
		"park":         handler.Insert,
		"leave":        handler.Remove,
		"status":       handler.Status,
		"color":        handler.GetByColor,
		"registration": handler.GetByRegistrationNo,
	}

	fmt.Println("input text:")
	for {
		var w1, w2, w3, w4 string
		_, err := fmt.Scanln(&w1, &w2, &w3, &w4)
		if err != nil && err.Error() != "unexpected newline" {
			fmt.Println(err.Error())
			continue
		}
		if val, ok := commandMap[w1]; ok {
			val(w2, w3, w4)
		} else {
			fmt.Println("Incorrect Command")
		}
	}
	//line, err := buffer.ReadString('\n')
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(line)
}
