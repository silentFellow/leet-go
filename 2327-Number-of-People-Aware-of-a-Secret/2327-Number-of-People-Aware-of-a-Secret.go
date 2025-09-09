package leetcode

const MOD int = 1_000_000_007

func peopleAwareOfSecret(n int, delay int, forget int) int {
	ageBuckets := make([]int, forget)
	ageBuckets[0] = 1

	for day := 2; day <= n; day++ {
		newAgeBucket := make([]int, forget)

		for age := 0; age < forget-1; age++ {
			cnt := ageBuckets[age]
			if cnt == 0 {
				continue
			}

			newAgeBucket[age+1] = cnt
			if age+1 >= delay {
				newAgeBucket[0] = (newAgeBucket[0] + cnt) % MOD
			}
		}

		ageBuckets = newAgeBucket
	}

	ans := 0
	for _, v := range ageBuckets {
		ans = (ans + v) % MOD
	}

	return ans
}
