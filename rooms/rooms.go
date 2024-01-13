package rooms

import (
	"awesomeProject1/animate"
	"awesomeProject1/devices"
	"awesomeProject1/furniture"
	"fmt"
)

type room struct {
	name    string
	square  int
	room_id int
}

type rom_data struct {
	room_data []room
}

func Room_creations() rom_data {
	room_creation := []room{{name: "Спальная", square: 13, room_id: 0},
		{name: "Зал", square: 15, room_id: 1},
		{name: "Коридор", square: 17, room_id: 2},
		{name: "Ванная", square: 9, room_id: 3},
		{name: "Кухня", square: 16, room_id: 4}}
	return rom_data{room_creation}

}

func Room_built(r rom_data) {
	animate.Get_animate_creations_info(animate.Animate_creations())
	fmt.Println("-------CАМА КВАРТИРА---------")
	for n := 0; n < len(r.room_data); n++ {
		println()
		fmt.Printf("НАЗВАНИЕ:\t%s\nПЛОЩАДЬ:\t%d м^2\n", r.room_data[n].name, r.room_data[n].square)
		devices.Get_console_info(devices.Consoles_massive(), n)
		devices.Get_tv_info(devices.Tves_massive(), n)
		devices.Get_pcs_info(devices.Pces_massive(), n)
		furniture.Sofs_info(furniture.Sofs_creation(), n)
		furniture.Beds_info(furniture.Beds_creation(), n)
		furniture.Chair_info(furniture.Chairs_creation(), n)
		furniture.Table_info(furniture.Tables_creation(), n)
		furniture.Get_shelf_info(furniture.Shelfs_creation(), n)
		println()
	}

}
