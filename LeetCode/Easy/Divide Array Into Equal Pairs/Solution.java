class Solution {
  public boolean divideArray(int[] nums) {
    Arrays.sort(nums);
    for(int i=0;i<nums.length;i=i+2){
      if(nums[i]!=nums[i+1]){
        return false;
      }
    }
    return true;
  }

  public boolean divideArray1(int[] nums) {
    int[] answer = new int[501];
    for(int i=0;i<nums.length;i++){
      answer[nums[i]]++;
    }
    for(int i=1;i<=500;i++){
      if(answer[i]%2 !=0) return false;
    }
    return true;
  }
}
