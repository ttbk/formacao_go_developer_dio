package main

func main() {

	for horas := 0; horas <= 12; horas++ {
		println("Horas:", horas)

		for min := 0; min < 60; min++ {
			print(min, " ")
		}

		println()
	}

}
