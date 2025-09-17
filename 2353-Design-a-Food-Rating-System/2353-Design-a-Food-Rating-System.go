package leetcode

import "container/heap"

type FoodStruct struct {
	food   string
	cusine string
	rating int
}

type RatingHeap []*FoodStruct

func (this RatingHeap) Len() int {
	return len(this)
}

func (this RatingHeap) Less(i, j int) bool {
	if this[i].rating == this[j].rating {
		return this[i].food < this[j].food
	}
	return this[i].rating > this[j].rating
}

func (this RatingHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *RatingHeap) Push(val any) {
	*this = append(*this, val.(*FoodStruct))
}

func (this *RatingHeap) Pop() any {
	old := *this
	n := len(old)
	item := old[n-1]
	*this = old[:n-1]
	return item
}

type FoodRatings struct {
	foodMap    map[string]*FoodStruct
	cuisineMap map[string]*RatingHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	foodMap := make(map[string]*FoodStruct)
	cuisineMap := make(map[string]*RatingHeap)

	for i := range foods {
		item := &FoodStruct{
			food:   foods[i],
			cusine: cuisines[i],
			rating: ratings[i],
		}

		foodMap[foods[i]] = item
		if _, ok := cuisineMap[cuisines[i]]; !ok {
			cuisineMap[cuisines[i]] = &RatingHeap{}
		}
		heap.Push(cuisineMap[cuisines[i]], &FoodStruct{
			food:   foods[i],
			cusine: cuisines[i],
			rating: ratings[i],
		})
	}

	return FoodRatings{
		foodMap:    foodMap,
		cuisineMap: cuisineMap,
	}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	item := this.foodMap[food]
	item.rating = newRating
	heap.Push(this.cuisineMap[item.cusine], &FoodStruct{
		food:   item.food,
		cusine: item.cusine,
		rating: newRating,
	})
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	h := this.cuisineMap[cuisine]
	for {
		top := (*h)[0]
		if this.foodMap[top.food].rating != top.rating {
			heap.Pop(h)
			continue
		}
		return top.food
	}
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
