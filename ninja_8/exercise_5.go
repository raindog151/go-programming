package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (ba ByAge) Len() int { return len(ba)}
func (ba ByAge) Swap(a, b int) { ba[a], ba[b] = ba[b], ba[a]}
func (ba ByAge) Less(a, b int) bool { return ba[a].Age < ba[b].Age }

type ByLast []user

func (bl ByLast) Len() int { return len(bl)}
func (bl ByLast) Swap(a, b int) { bl[a], bl[b] = bl[b], bl[a]}
func (bl ByLast) Less(a, b int) bool { return bl[a].Last < bl[b].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
	sort.Sort(ByAge(users))
	fmt.Println(users)
	sort.Sort(ByLast(users))
	fmt.Println(users)

	for i, _ := range users {
		fmt.Printf("%s, %s - Age: %d\n", users[i].Last, users[i].First, users[i].Age)
		sort.Strings(users[i].Sayings)
		for s, _ := range users[i].Sayings {
			fmt.Println(users[i].Sayings[s])
		}
	}


}
