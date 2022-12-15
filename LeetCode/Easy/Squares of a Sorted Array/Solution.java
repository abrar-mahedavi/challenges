class Solution {
  public int[] sortedSquares(int[] nums) {
    int[] answer = new int[nums.length];
    for(int i=0;i<nums.length;i++){
      answer[i] = nums[i]*nums[i];
    }
    Arrays.sort(answer);
    return answer;
  }

  public int[] sortedSquares1(int[] nums) {
    int left= 0;
    int right = (nums.length-1);
    int[] answer = new int[nums.length];
    int counter=nums.length-1;
    while(right>=left){
      if(Math.abs(nums[right])>=Math.abs(nums[left])){
        answer[counter]=nums[right]*nums[right];
        right--;
      }else{
        answer[counter]=nums[left]*nums[left];
        left++;
      }
      counter--;
    }
    return answer;
  }
}
