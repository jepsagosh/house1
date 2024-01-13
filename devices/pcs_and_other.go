package devices

import (
	"fmt"
)

type pc struct {
	is_laptop        bool
	brand            string
	gpu              string
	cpu              string
	location_by_code int
	name             string
}

type console struct {
	brand            string
	model            string
	location_by_code int
	name             string
}
type tv struct {
	brand            string
	diagonal         int
	location_by_code int
	name             string
}
type pcs struct {
	pces []pc
}
type consoles struct {
	consols []console
}
type tves struct {
	tvs []tv
}

func Pces_massive() pcs {
	pces_massive := []pc{{is_laptop: false, brand: "очумелые ручки", gpu: "2080", cpu: "I9 9900k", location_by_code: 1, name: "Игровой компьютер"},
		{is_laptop: true, brand: "Xiaomi", gpu: "Отсутствует", cpu: "i5 12400", location_by_code: 0, name: "Рабочий ноутбук"}}
	return pcs{pces_massive}
}

func Consoles_massive() consoles {
	consoles_massive := []console{{brand: "Sony", model: "PLAYSTATION 5", location_by_code: 1, name: "Плойка"},
		{brand: "Nintendo", model: "Switch", location_by_code: 1, name: "Свитч"}}
	return consoles{consoles_massive}
}
func Tves_massive() tves {
	tves_massive := []tv{{brand: "Sony", diagonal: 67, location_by_code: 1, name: "Телевизор в зале"},
		{brand: "LG", diagonal: 38, location_by_code: 0, name: "Телевизор в спальне"},
		{brand: "Samsung", diagonal: 45, location_by_code: 4, name: "Телевизор на кухне"}}
	return tves{tves_massive}
}

func Get_tv_info(t tves, id int) {
	fmt.Println("--------------------------")
	for n := 0; n < len(t.tvs); n++ {
		if id == t.tvs[n].location_by_code {
			fmt.Printf("БРЕНД:\t%s\nДИАГОНАЛЬ:\t%d\nИМЯ: \t%s\n", t.tvs[n].brand, t.tvs[n].diagonal, t.tvs[n].name)
		}
	}
}

func Get_console_info(con consoles, id int) {
	fmt.Println("--------------------------")
	for n := 0; n < len(con.consols); n++ {
		if id == con.consols[n].location_by_code {
			fmt.Printf("БРЕНД:\t%s\nМОДЕЛЬ:\t%s\n", con.consols[n].brand, con.consols[n].model)
		}
	}
}

func Get_pcs_info(comp pcs, id int) {
	fmt.Println("--------------------------")
	for n := 0; n < len(comp.pces); n++ {
		if id == comp.pces[n].location_by_code {
			fmt.Printf("БРЕНД:\t%s\nВИДЕОКАРТА:\t%s\nПРОЦЕССОР:\t%s\nИМЯ:\t%s\n", comp.pces[n].brand, comp.pces[n].gpu, comp.pces[n].cpu, comp.pces[n].name)
		} else {
			break
		}
	}

}
