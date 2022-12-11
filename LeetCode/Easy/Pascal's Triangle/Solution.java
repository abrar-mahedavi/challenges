class Solution {
  public List<List<Integer>> generate(int numRows) {
    List<List<Integer>> answer = new ArrayList<List<Integer>>();
    List<Integer> row, pre= null;
    for(int i =0;i<numRows;i++){
      row = new ArrayList<Integer>();
      for(int j=0;j<=i;j++){
        if(j==0 || j==i){
          row.add(1);
        }else{
          // System.out.println(Arrays.toString(pre.toArray()));
          row.add(pre.get(j-1)+pre.get(j));
        }
      }
      pre = row;
      answer.add(row);
    }
    return answer;
  }
}
