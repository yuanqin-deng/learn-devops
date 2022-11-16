package two_sum

func twoSum(nums []int, target int) []int {
	//哈希表：使用哈希表，以数组值为key,索引下标为value，记录在哈希表中，
	//若target减去value得到的key在哈希表中存在，即找到这两个数字，访问哈希表返回下标即可
	numToIndex := make(map[int]int, len(nums)/2) //使用map，一定要make
	for index, value := range nums {
		o, ok := numToIndex[target-value]

	}
}
