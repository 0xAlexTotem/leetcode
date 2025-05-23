package java;

public class FinalArrayState {
    class Solution {
        public int[] getFinalState(int[] nums, int k, int multiplier) {
            for (int i = 0; i < k; i++) {
                int min = 0;
                int index = 0;
    
                for (int j = 0; j < nums.length; j++) {
                    if (j == 0 || nums[j] < min) {
                        min = nums[j];
                        index = j;
                    }
                }
    
                nums[index] = min * multiplier;
            }
    
            return nums;
        }
    }
}