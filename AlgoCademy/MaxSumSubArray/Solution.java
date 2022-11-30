package com.abu.algocademy.maxsubarray;

public class Solution {
  public int maxSumSubarray(int[] nums) {
    int maxSum = -99999;
    for (int i = 0; i < nums.length; i++) {
      int sum = 0;
      for (int j = i; j < nums.length; j++) {
        sum = sum + nums[j];
        if (sum > maxSum) {
          maxSum = sum;
        }
      }
    }
    return maxSum;
  }
}
