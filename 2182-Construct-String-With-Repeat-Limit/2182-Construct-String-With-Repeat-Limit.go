package leetcode

// imports
import (
	"container/heap"
	"strings"
)

// custom type that contains rune with freq for heap
type freqRune struct {
	val  rune // character
	freq int // frequency
}

// heap
type repStr []freqRune

// sort interface
func (r repStr) Len() int { return len(r) }

func (r repStr) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func (r repStr) Less(i, j int) bool {
	return r[i].val > r[j].val
}

// heap interface
func (r *repStr) Push(val any) { *r = append(*r, val.(freqRune)) }

func (r *repStr) Pop() any {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[:n-1]
	return x
}

func repeatLimitedString(s string, repeatLimit int) string {
	hmap := make(map[rune]int) // character with freq

	// count frequency
	for _, val := range s {
		hmap[val]++
	}

	// heap
	rh := &repStr{}
	heap.Init(rh)
	for key, val := range hmap {
		neo := freqRune{
			val:  key,
			freq: val,
		}

		heap.Push(rh, neo)
	}

	// main logic
	var ans strings.Builder
	for rh.Len() > 0 {
		top := heap.Pop(rh).(freqRune) // take first element

		// you only want to repeat till repeatLimit
		// but also may be current top frequency is small, so you cannot repeat till repeatLimit
		// so find minimum of it
		minRep := min(top.freq, repeatLimit)

		ans.WriteString(strings.Repeat(string(top.val), minRep)) // concat
		top.freq -= minRep // reduce frequency

		// if only top still have more occurances
		// also only when another element in heap exists
		if top.freq > 0 && rh.Len() > 0 {
			// just put single time next character
			// and redo the current top
			// so lexicographically largest will be get
			// and no exceed of repeatLimit
			next := (*rh)[0]
			ans.WriteRune(next.val)
			(*rh)[0].freq--

			// there is a possiblity that next element frequency becomes zero after this operation
			// if thats the case pop it from the heap
			if (*rh)[0].freq == 0 {
				heap.Pop(rh)
			}

			heap.Push(rh, top)
		}
	}

	return ans.String()
}
