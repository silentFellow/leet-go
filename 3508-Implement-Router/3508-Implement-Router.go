package leetcode

import "sort"

type Route struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	memoryLimit    int
	routes         []Route
	routeMap       map[Route]struct{}
	destinationMap map[int][]int // only store destination - timestamp
}

func Constructor(memoryLimit int) Router {
	return Router{
		memoryLimit:    memoryLimit,
		routes:         make([]Route, 0),
		routeMap:       make(map[Route]struct{}),
		destinationMap: make(map[int][]int),
	}
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	route := Route{
		source,
		destination,
		timestamp,
	}

	if _, ok := this.routeMap[route]; ok {
		return false
	}

	if len(this.routes) >= this.memoryLimit {
		this.ForwardPacket()
	}

	this.routes = append(this.routes, route)
	this.routeMap[route] = struct{}{}
	this.destinationMap[destination] = append(this.destinationMap[destination], timestamp)

	return true
}

func (this *Router) ForwardPacket() []int {
	if len(this.routes) == 0 {
		return []int{}
	}

	first := this.routes[0]
	this.routes = this.routes[1:]
	delete(this.routeMap, first)
	this.destinationMap[first.destination] = this.destinationMap[first.destination][1:]

	return []int{first.source, first.destination, first.timestamp}
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	timestamps := this.destinationMap[destination]

	left := sort.Search(len(timestamps), func(i int) bool {
		return timestamps[i] >= startTime
	})

	right := sort.Search(len(timestamps), func(i int) bool {
		return timestamps[i] > endTime
	})

	return right - left
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
