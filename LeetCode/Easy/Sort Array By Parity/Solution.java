class Solution {
  public int[] sortArrayByParity(int[] nums) {
    int left = 0;
    int right = nums.length-1;
    int[] answers = new int[nums.length];
    // int counter = nums.length -1;
    for(int i=0;i<nums.length;i++){
      if(nums[i]%2!=0){
        answers[right]= nums[i];
        right--;
      } else {
        answers[left]= nums[i];
        left++;
      }
    }
    return answers;
  }

  public int[] sortArrayByParity1(int[] nums) {
    int left = 0;
    for(int i=0;i<nums.length;i++){
      if(nums[i]%2==0){
        int temp = nums[i];
        nums[i] = nums[left];
        nums[left] = temp;
        left++;
      }
    }
    return nums;
  }
}
