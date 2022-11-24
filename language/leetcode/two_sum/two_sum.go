package two_sum

/**
题目
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
*/

func twoSum(nums []int, target int) []int {
	//哈希表：使用哈希表，以数组值为key,索引下标为value，记录在哈希表中，
	//若target减去value得到的key在哈希表中存在，即找到这两个数字，访问哈希表返回下标即可

	/**
	只遍历一次数组，在遍历的同时将数据使用HashMap存储起来，翻转key和value，建立数字和其坐标位置之间的映射。HashMap 是常数级的查找效率，这样在遍历数组的时候，用 target 减去遍历到的数字，就是另一个需要的数字了。直接在 HashMap 中查找其是否存在即可。
	由于当前数字在HashMap中还不存在，所以不会出现同一个数字使用两次的情况。比如 target 是4，遍历到了一个2，另外一个2还未进入HashMap中，肯定就是要找的值。
	整个实现步骤为：遍历一遍数组，查找target 减去当前数字是否存在于HashMap中；如果不存在，则建立 HashMap 映射。代码如下：
	*/
	numToIndex := make(map[int]int, len(nums)/2) //使用map，一定要make
	for index, value := range nums {
		if _, ok := numToIndex[target-value]; ok {
			return []int{numToIndex[target-value], index}
		}
		numToIndex[value] = index
	}
	return []int{}
}

//func main() {
//	arr := []int{2, 11, 3, 7, 15}
//	fmt.Println(twoSum(arr, 9))
//}
