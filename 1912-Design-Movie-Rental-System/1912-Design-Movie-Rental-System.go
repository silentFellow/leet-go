package leetcode

import "container/heap"

type Movie struct {
	shop  int
	movie int
	price int
	index int
}

type MovieHeap []Movie

func (h MovieHeap) Len() int { return len(h) }

func (h MovieHeap) Less(i, j int) bool {
	if h[i].price != h[j].price {
		return h[i].price < h[j].price
	}
	if h[i].shop != h[j].shop {
		return h[i].shop < h[j].shop
	}
	return h[i].movie < h[j].movie
}

func (h MovieHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *MovieHeap) Push(v any) {
	val := v.(Movie)
	val.index = len(*h)
	*h = append(*h, val)
}

func (h *MovieHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	x.index = -1 // for safety
	*h = old[:n-1]
	return x
}

type MovieRentingSystem struct {
	movieHeap       map[int]*MovieHeap
	rentedMovies    *MovieHeap
	rentedMovieMap  map[[2]int]struct{}
	shopMovieIdxMap map[[2]int]int
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	movieHeap := make(map[int]*MovieHeap)
	shopMovieIdxMap := make(map[[2]int]int)

	for i, entry := range entries {
		shop, movie, price := entry[0], entry[1], entry[2]
		if _, ok := movieHeap[movie]; !ok {
			movieHeap[movie] = &MovieHeap{}
			heap.Init(movieHeap[movie])
		}
		heap.Push(movieHeap[movie], Movie{shop: shop, movie: movie, price: price, index: i})
		shopMovieIdxMap[[2]int{shop, movie}] = i
	}

	rentedMovies := &MovieHeap{}
	heap.Init(rentedMovies)

	return MovieRentingSystem{
		movieHeap:       movieHeap,
		rentedMovies:    rentedMovies,
		rentedMovieMap:  make(map[[2]int]struct{}),
		shopMovieIdxMap: shopMovieIdxMap,
	}
}

func (this *MovieRentingSystem) Search(movie int) []int {
	ans := []int{}
	if heapPtr, ok := this.movieHeap[movie]; ok {
		temp := []Movie{}
		for heapPtr.Len() > 0 && len(ans) < 5 {
			top := (*heapPtr)[0]
			ans = append(ans, top.shop)
			// Store temporarily because we cannot remove it
			temp = append(temp, top)
			heap.Remove(heapPtr, top.index)
		}
		// Push them back
		for _, m := range temp {
			heap.Push(heapPtr, m)
		}
	}
	return ans
}

func (this *MovieRentingSystem) Rent(shop int, movie int) {
	h := this.movieHeap[movie]
	for i := 0; i < h.Len(); i++ {
		if (*h)[i].shop == shop {
			// Remove from unrented heap
			m := heap.Remove(h, i).(Movie)
			// Add to rented heap
			heap.Push(this.rentedMovies, m)
			this.rentedMovieMap[[2]int{shop, movie}] = struct{}{}
			break
		}
	}
}

func (this *MovieRentingSystem) Drop(shop int, movie int) {
	// Remove from rented heap
	h := this.rentedMovies
	for i := 0; i < h.Len(); i++ {
		if (*h)[i].shop == shop && (*h)[i].movie == movie {
			m := heap.Remove(h, i).(Movie)
			heap.Push(this.movieHeap[movie], m)
			delete(this.rentedMovieMap, [2]int{shop, movie})
			break
		}
	}
}

func (this *MovieRentingSystem) Report() [][]int {
	ans := [][]int{}
	temp := []Movie{}
	h := this.rentedMovies
	for h.Len() > 0 && len(ans) < 5 {
		top := (*h)[0]
		ans = append(ans, []int{top.shop, top.movie})
		temp = append(temp, top)
		heap.Remove(h, top.index)
	}
	// Push them back
	for _, m := range temp {
		heap.Push(h, m)
	}
	return ans
}

/**
 * Your MovieRentingSystem object will be instantiated and called as such:
 * obj := Constructor(n, entries);
 * param_1 := obj.Search(movie);
 * obj.Rent(shop,movie);
 * obj.Drop(shop,movie);
 * param_4 := obj.Report();
 */
