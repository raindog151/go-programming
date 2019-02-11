package main

import "fmt"

func main() {

	x := map[string][]string {
		"bond_james": []string{"Shaken not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["thatcher_margaret"] = []string{"conservative", "mp", "shrimp and white wine"}

	for i, v := range x {
		fmt.Println(i, v)
	}
}
