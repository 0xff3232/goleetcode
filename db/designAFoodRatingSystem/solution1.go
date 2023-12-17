/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */

package designAFoodRatingSystem

import (
    "log"
    "sort"
)

// Food represents a single food item with its name, rating, and associated cuisine.
type Food struct {
    Name    string
    Rating  int
    Cuisine string
}

// FoodRatings holds the entire system's data, mapping food names to Food structs,
// grouping foods by cuisine, and tracking the maximum rating for each cuisine.
type FoodRatings struct {
    FoodMap    map[string]*Food   // Maps food names to Food structs.
    CuisineMap map[string][]*Food // Groups foods by their cuisine.
    MaxRating  map[string]int     // Tracks the maximum rating for each cuisine.
}

// Constructor initializes the FoodRatings system with the given foods, cuisines, and ratings.
func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
    fr := FoodRatings{
        FoodMap:    make(map[string]*Food),
        CuisineMap: make(map[string][]*Food),
        MaxRating:  make(map[string]int),
    }
    for i, food := range foods {
        f := &Food{Name: food, Rating: ratings[i], Cuisine: cuisines[i]}
        fr.FoodMap[food] = f
        fr.CuisineMap[cuisines[i]] = append(fr.CuisineMap[cuisines[i]], f)
        // Update the maximum rating for each cuisine.
        if ratings[i] > fr.MaxRating[cuisines[i]] {
            fr.MaxRating[cuisines[i]] = ratings[i]
        }
    }
    return fr
}

// ChangeRating updates the rating of a specific food and adjusts the maximum rating for the cuisine if necessary.
func (fr *FoodRatings) ChangeRating(food string, newRating int) {
    f, exists := fr.FoodMap[food]
    if !exists {
        log.Printf("food '%s' not found", food)
        return
    }
    f.Rating = newRating
    cuisine := f.Cuisine
    // Update the maximum rating for the cuisine if this food's new rating is higher.
    if newRating > fr.MaxRating[cuisine] {
        fr.MaxRating[cuisine] = newRating
    } else if newRating < fr.MaxRating[cuisine] && f.Rating == fr.MaxRating[cuisine] {
        // Find the new maximum rating for this cuisine if the top-rated food's rating is decreased.
        newMax := 0
        for _, food := range fr.CuisineMap[cuisine] {
            if food.Rating > newMax {
                newMax = food.Rating
            }
        }
        fr.MaxRating[cuisine] = newMax
    }
}

// HighestRated returns the name of the highest-rated food for a given cuisine.
// In case of a tie in ratings, it returns the lexicographically smallest name.
func (fr *FoodRatings) HighestRated(cuisine string) string {
    var highestRatedFood *Food
    for _, food := range fr.CuisineMap[cuisine] {
        if highestRatedFood == nil || food.Rating > highestRatedFood.Rating || 
           (food.Rating == highestRatedFood.Rating && food.Name < highestRatedFood.Name) {
            highestRatedFood = food
        }
    }
    if highestRatedFood != nil {
        return highestRatedFood.Name
    }
    return ""
}

// sortFoods sorts a slice of *Food based on rating and name.
func sortFoods(foods []*Food) {
    sort.Slice(foods, func(i, j int) bool {
        if foods[i].Rating == foods[j].Rating {
            return foods[i].Name < foods[j].Name
        }
        return foods[i].Rating > foods[j].Rating
    })
}