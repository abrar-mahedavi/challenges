class Solution {
  public boolean containsDuplicate(int[] nums) {
    Map<Integer, Boolean> numsCounter = new HashMap<Integer, Boolean>();
    for (int i = 0; i < nums.length; i++) {
      if (numsCounter.containsKey(nums[i])) {
        return true;
      } else {
        numsCounter.put(nums[i], true);
      }
    }
    return false;
  }
}
