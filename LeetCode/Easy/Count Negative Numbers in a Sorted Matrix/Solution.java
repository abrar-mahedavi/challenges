class Solution {
  public int countNegatives(int[][] grid) {
    int countNegative =0;
    int rows = grid.length;
    for(int i=0;i<rows;i++){
      int cols = grid[i].length-1;
      for(int j=cols;j>=0;j--){
        if(grid[i][j]<0) countNegative++;
      }
    }
    return countNegative;
  }
}
