package furniture

import "fmt"

type shelf struct {
	is_cloathing     bool
	is_thigable      bool
	colour           string
	material         string
	name             string
	location_by_code int
}

type bed struct {
	persons_to_sleep int
	colour           string
	material         string
	name             string
	location_by_code int
}

type sofa struct {
	is_sleepable     bool
	is_folding       bool
	colour           string
	material         string
	name             string
	location_by_code int
}

type chair struct {
	is_pc            bool
	colour           string
	material         string
	name             string
	location_by_code int
}

type table struct {
	is_pc            bool
	colour           string
	material         string
	name             string
	location_by_code int
}

type shelfs_s struct {
	shelfs []shelf
}

type beds_s struct {
	beds []bed
}

type chairs_s struct {
	chairs []chair
}

type sofs_s struct {
	sofs []sofa
}

type tables_s struct {
	tables []table
}

func Tables_creation() tables_s {
	tables_massive := []table{{is_pc: true, colour: "Древесный темный", material: "Дерево и металл", name: "Компьютерный стол в спальне", location_by_code: 0},
		{is_pc: true, colour: "Древесный темный", material: "Дерево и металл", name: "Компьютерный стол в зале", location_by_code: 1},
		{is_pc: false, colour: "Белый", material: "Дерево и стекло", name: "Кухонный стол", location_by_code: 4}}
	return tables_s{tables_massive}
}
func Sofs_creation() sofs_s {
	sofs_massive := []sofa{{is_sleepable: true, is_folding: true, colour: "Серый", material: "Велюр", name: "Диван в зале", location_by_code: 1},
		{is_sleepable: false, is_folding: false, colour: "Фиолетовый", material: "Велюр", name: "Кухонный диван", location_by_code: 4}}
	return sofs_s{sofs_massive}
}

func Chairs_creation() chairs_s {
	chairs_massive := []chair{{is_pc: false, colour: "Серый", material: "Велюр", name: "кухонный стул", location_by_code: 4},
		{is_pc: false, colour: "Серый", material: "Велюр", name: "кухонный стул", location_by_code: 4},
		{is_pc: false, colour: "Серый", material: "Велюр", name: "кухонный стул", location_by_code: 4},
		{is_pc: false, colour: "Серый", material: "Велюр", name: "кухонный стул", location_by_code: 4},
		{is_pc: true, colour: "Белый", material: "Кожзам", name: "Компьютерный стул в зале", location_by_code: 1},
		{is_pc: true, colour: "Футбольный блин))", material: "Кожзам", name: "Компьютерный стул в спальне", location_by_code: 0}}
	return chairs_s{chairs_massive}
}

func Shelfs_creation() shelfs_s {
	shelfs_massive := []shelf{{is_cloathing: false, is_thigable: true, colour: "Коричневый", material: "Дерево", name: "Шкаф для инструментов", location_by_code: 0},
		{is_cloathing: true, is_thigable: false, colour: "Белый", material: "Дерево", name: "Шкаф для вещей родителей", location_by_code: 0},
		{is_cloathing: true, is_thigable: true, colour: "Слоновая кость", material: "Дерево", name: "Шкаф в зале", location_by_code: 1},
		{is_cloathing: true, is_thigable: false, colour: "Белый", name: "Шкаф для верхней одежды", location_by_code: 2},
		{is_cloathing: true, is_thigable: false, colour: "Белый", name: "Шкаф для старых вещей", location_by_code: 2},
		{is_cloathing: false, is_thigable: true, colour: "Слоновая кость", name: "Шкаф в ванной", location_by_code: 3}}
	return shelfs_s{shelfs_massive}
}

func Beds_creation() beds_s {
	beds_massive := []bed{{persons_to_sleep: 2, colour: "Фиолетовый", material: "Дерево", name: "Кровать в спальне", location_by_code: 0}}
	return beds_s{beds_massive}
}

func Beds_info(b beds_s, id int) {
	for n := 0; n < len(b.beds); n++ {
		if id == b.beds[n].location_by_code {
			fmt.Println("--------------------------")
			fmt.Println("КРОВАТЬ")
			fmt.Printf("ЦВЕТ:\t%s\nМАТЕРИАЛ:\t%s\nМОЖЕТ СПАТЬ ЛЮДЕЙ:\t%d\nИМЯ:\t%s\n", b.beds[n].colour, b.beds[n].material, b.beds[n].persons_to_sleep, b.beds[n].name)
			fmt.Println("--------------------------")
		}
	}

}

func Chair_info(c chairs_s, id int) {
	for n := 0; n < len(c.chairs); n++ {
		if id == c.chairs[n].location_by_code {
			fmt.Println("--------------------------")
			fmt.Println("Стул")
			fmt.Printf("ЦВЕТ:\t%s\nМАТЕРИАЛ:\t%s\nИМЯ:\t%s\n", c.chairs[n].colour, c.chairs[n].material, c.chairs[n].name)
			fmt.Println("--------------------------")
		}
	}
}

func Table_info(t tables_s, id int) {
	for n := 0; n < len(t.tables); n++ {
		if id == t.tables[n].location_by_code {
			fmt.Println("--------------------------")
			fmt.Println("Стол")
			fmt.Printf("ЦВЕТ:\t%s\nМАТЕРИАЛ:\t%s\nИМЯ:\t%s\n", t.tables[n].colour, t.tables[n].material, t.tables[n].name)
			fmt.Println("--------------------------")
		}
	}
}

func Get_shelf_info(s shelfs_s, id int) {
	for n := 0; n < len(s.shelfs); n++ {
		if id == s.shelfs[n].location_by_code {
			if s.shelfs[n].is_thigable == true {
				fmt.Println("--------------------------")
				fmt.Println("ШКАФ ДЛЯ ВЕЩЕЙ")
				fmt.Printf("ЦВЕТ:\t%s\nМАТЕРИАЛ:\t%s\nИМЯ:\t%s\n", s.shelfs[n].colour, s.shelfs[n].material, s.shelfs[n].name)
				fmt.Println("--------------------------")
			} else {
				fmt.Println("--------------------------")
				fmt.Println("ШКАФ ДЛЯ ХЛАМА")
				fmt.Printf("ЦВЕТ:\t%s\nМАТЕРИАЛ:\t%s\nИМЯ:\t%s\n", s.shelfs[n].colour, s.shelfs[n].material, s.shelfs[n].name)
				fmt.Println("--------------------------")
			}
		}
	}

}
func Sofs_info(s sofs_s, id int) {
	for n := 0; n < len(s.sofs); n++ {
		if id == s.sofs[n].location_by_code {
			fmt.Println("--------------------------")
			fmt.Println("Диван")
			fmt.Printf("ЦВЕТ:\t%s\nМАТЕРИАЛ:\t%s\nИМЯ:\t%s\n", s.sofs[n].colour, s.sofs[n].material, s.sofs[n].name)
			fmt.Println("--------------------------")
		}
	}
}
