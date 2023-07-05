package main


func main(){

	nums := []int {4,1,98,9}
	println(twoSum(nums , 5))
}

// Time: O(n*n) = O(n^2)
// Space: O(1)
func twoSum(nums []int, target int) []int {
    // O(n)
    for i := 0; i < len(nums)-1; i++ {
        // O(n)
        for j := i+1; j < len(nums); j++ {
			// Time: O(1)
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}