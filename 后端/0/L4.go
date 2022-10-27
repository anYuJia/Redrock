package main

import "fmt"

func main() {
	users := []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution := make(map[string]int, len(users))
	gold_total := 0
	for _, k := range users {
		gold := 0
		for _, m := range k {
			switch m {
			case 'e':
				gold++
			case 'E':
				gold++
			case 'i':
				gold += 2
			case 'I':
				gold += 2
			case 'o':
				gold += 3
			case 'O':
				gold += 3
			case 'u':
				gold += 4
			case 'U':
				gold += 4
			default:
				gold += 0
			}
		}
		distribution[k] = gold
		gold_total += gold
	}
	for u, n := range distribution {
		fmt.Println(u, "分配到了", n, "个金币")
	}
	fmt.Printf("一共分配了%d个金币!\n", gold_total)
}
