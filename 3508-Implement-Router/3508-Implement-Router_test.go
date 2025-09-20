package leetcode

import (
	"reflect"
	"testing"
)

func TestRouter_EndToEnd(t *testing.T) {
	router := Constructor(3)

	// Case 1: Add unique packets
	if !router.AddPacket(1, 4, 90) {
		t.Errorf("expected AddPacket(1,4,90) = true")
	}
	if !router.AddPacket(2, 5, 90) {
		t.Errorf("expected AddPacket(2,5,90) = true")
	}

	// Case 2: Add duplicate
	if router.AddPacket(1, 4, 90) {
		t.Errorf("expected AddPacket(1,4,90) = false (duplicate)")
	}

	// Case 3: Eviction when memory full
	if !router.AddPacket(3, 5, 95) {
		t.Errorf("expected AddPacket(3,5,95) = true")
	}
	if !router.AddPacket(4, 5, 105) {
		t.Errorf("expected AddPacket(4,5,105) = true, should evict oldest (1,4,90)")
	}

	// Case 4: Forward oldest
	got := router.ForwardPacket()
	want := []int{2, 5, 90}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ForwardPacket() = %v, want %v", got, want)
	}

	// Case 5: Add new packet after eviction
	if !router.AddPacket(5, 2, 110) {
		t.Errorf("expected AddPacket(5,2,110) = true")
	}

	// Case 6: GetCount in range
	count := router.GetCount(5, 100, 110)
	if count != 1 {
		t.Errorf("GetCount(5,100,110) = %d, want %d", count, 1)
	}
}

func TestRouter_ForwardEmpty(t *testing.T) {
	router := Constructor(2)

	// Forward with no packets
	got := router.ForwardPacket()
	want := []int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ForwardPacket() = %v, want %v", got, want)
	}
}

func TestRouter_GetCountEmpty(t *testing.T) {
	router := Constructor(2)

	// No packets stored
	if c := router.GetCount(5, 1, 100); c != 0 {
		t.Errorf("GetCount(empty) = %d, want %d", c, 0)
	}
}

func TestRouter_MultipleDestinations(t *testing.T) {
	router := Constructor(5)

	router.AddPacket(1, 10, 10)
	router.AddPacket(2, 20, 20)
	router.AddPacket(3, 10, 30)
	router.AddPacket(4, 20, 40)
	router.AddPacket(5, 10, 50)

	// Count packets with destination 10 in range [15, 50]
	if c := router.GetCount(10, 15, 50); c != 2 {
		t.Errorf("GetCount(10,15,50) = %d, want %d", c, 2)
	}

	// Count packets with destination 20 in range [1, 100]
	if c := router.GetCount(20, 1, 100); c != 2 {
		t.Errorf("GetCount(20,1,100) = %d, want %d", c, 2)
	}
}

func TestRouter_FullEviction(t *testing.T) {
	router := Constructor(2)

	router.AddPacket(1, 1, 10)
	router.AddPacket(2, 2, 20)
	router.AddPacket(3, 3, 30) // should evict (1,1,10)

	// Forward first → should return (2,2,20)
	got := router.ForwardPacket()
	want := []int{2, 2, 20}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ForwardPacket() = %v, want %v", got, want)
	}

	// Forward second → should return (3,3,30)
	got = router.ForwardPacket()
	want = []int{3, 3, 30}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ForwardPacket() = %v, want %v", got, want)
	}

	// Forward empty → []
	got = router.ForwardPacket()
	want = []int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("ForwardPacket() = %v, want %v", got, want)
	}
}
