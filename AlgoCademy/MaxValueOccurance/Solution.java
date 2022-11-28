package com.abu.algocademy.maxvalueoccurance;

public class Solution {
  public int[] maxValNumOfOccurrences(int[] nums) {
    int maxVal = nums[0], count = 0;
    for (int val : nums) {
      if (val > maxVal) {
        maxVal = val;
        count = 1;
      } else if (val == maxVal) {
        count += 1;
      }
    }
    return new int[]{maxVal, count};
  }
}
