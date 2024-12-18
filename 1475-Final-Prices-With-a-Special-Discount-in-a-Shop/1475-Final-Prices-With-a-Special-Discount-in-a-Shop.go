package leetcode

type stack []int

// some utility functions for stack
func (s *stack) Push(v int) { *s = append(*s, v) }

func (s *stack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func (s *stack) Len() int { return len(*s) }

func (s *stack) Peek() int { return (*s)[len(*s)-1] }

// main logic
func finalPrices(prices []int) []int {
	s := stack{}

	// loog prices array
	for i := len(prices) - 1; i >= 0; i-- {
		// check stack for potential discounts
		// or until stack becomes empty
		for s.Len() > 0 && s.Peek() > prices[i] {
			s.Pop()
		}

		// if non-empty, change prices[i]
		// before changing, push the value to stack
		// since it may be discount for previous elements
		if s.Len() != 0 {
			top := s.Peek()
			s.Push(prices[i])
			prices[i] -= top
			continue
		}

		// if empty push current
		s.Push(prices[i])
	}

	// return modified prices
	return prices
}
