package main

import (
	"fmt"

	"github.com/0xff3232/goleetcode/db/designAFoodRatingSystem"
)

func main() {
    fmt.Println("Solution I'm currently working on..")


	t := designAFoodRatingSystem.Constructor([]string{"kimchi", "miso"},
											[]string{"korean", "japanese"},
											[]int{9, 12})
	


	for foodName, rating := range t.FoodMap{
		fmt.Printf("Name: %s, Rating: %d\n", foodName, rating.Rating)
	}
}
