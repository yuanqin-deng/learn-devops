import java.util.HashMap;

/**
 * @Author: Yuanqin DENG
 * @Date: 2022/11/17
 */
public class two_sum {
    public static int[] twoSum(int []nums, int target) {
        if (nums == null || nums.length < 2) {
            return new int[]{ -1, -1 };
        }
        int[] result = new int[]{ -1, -1 };
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            if (map.containsKey(target - nums[i])) {
                result[0] = map.get(target - nums[i]);
                result[1] = i;
                break;
            }
            map.put(nums[i], i);
        }
        return result;
    }

    public static void main(String[] args) {
        int[] nums;
        nums = new int[]{2, 11, 3, 7, 15};
        int target = 9;
        System.out.println("[" + twoSum(nums, target)[0] + "," + twoSum(nums, target)[1] + "]");
    }
}
