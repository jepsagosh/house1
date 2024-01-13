package animate

import "fmt"

type animate struct {
	is_human bool
	kind     string
	name     string
	age      int
	gender   string
}

type animate_massive struct {
	animate_massive []animate
}

func Animate_creations() animate_massive {
	animate_creation := []animate{{is_human: false, kind: "Собака", name: "Умка", age: 7, gender: "Мужской"},
		{is_human: true, name: "Олег", age: 20, gender: "Мужской"}, {is_human: true, name: "Отец", age: 43, gender: "Мужской"},
		{is_human: true, name: "Мать", age: 46, gender: "Женский"}}
	return animate_massive{animate_massive: animate_creation}
}

func Get_animate_creations_info(alive animate_massive) {
	fmt.Println("------------ЖИВУТ ДОМА-----------")
	for n := 0; n < len(alive.animate_massive); n++ {
		if alive.animate_massive[n].is_human != true {
			fmt.Println("--------------------------")
			fmt.Println("Животное")
			fmt.Printf("Вид:\t%s\nИмя:\t%s\nВозраст:\t%d\nПол:\t%s\n", alive.animate_massive[n].kind, alive.animate_massive[n].name, alive.animate_massive[n].age, alive.animate_massive[n].gender)
			fmt.Println("--------------------------")
		} else {
			fmt.Println("--------------------------")
			fmt.Println("\nЧеловек")
			fmt.Printf("Имя:\t%s\nВозраст:\t%d\nПол:\t%s\n", alive.animate_massive[n].name, alive.animate_massive[n].age, alive.animate_massive[n].gender)
			fmt.Println("--------------------------")
		}
	}
	fmt.Println("--------------------------")
}
