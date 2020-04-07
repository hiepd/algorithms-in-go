package main

/*
Heap's algorithm generates all possible permutations of n items
https://en.wikipedia.org/wiki/Heap%27s_algorithm
*/

func Permute(nums []int) [][]int {
	res := make([][]int, 0)
	genPermutations(nums, len(nums), &res)
	return res
}

func genPermutations(nums []int, k int, res *[][]int) {
	// base case, add array to results
	if k == 1 {
		*res = append(*res, append([]int{}, nums...))
		return
	}

	// iterate k times here
	for i := 0; i < k; i++ {
		// generate permutations
		genPermutations(nums, k-1, res)
		// if k is even, swap current item with last item
		temp := nums[k-1]
		if k%2 == 0 {
			nums[k-1] = nums[i]
			nums[i] = temp
		} else { // if k is odd, swap first item with last item
			nums[k-1] = nums[0]
			nums[0] = temp
		}
	}

	return
}
