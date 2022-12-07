class Solution {
  public void setZeroes(int[][] matrix) {
    int row=matrix.length;
    int col=matrix[0].length;
    int cols=1;

    for(int i=0;i<row;i++){
      if(matrix[i][0]==0) cols=0;
      for(int j=1;j<col;j++){
        if(matrix[i][j]==0){
          matrix[i][0]=0;
          matrix[0][j]=0;
        }
      }
    }

    for(int i=row-1;i>=0;i--){
      for(int j=col-1;j>=1;j--){
        if(matrix[i][0]==0  || matrix[0][j]==0)
          matrix[i][j]=0;
      }
      if(cols==0){
        matrix[i][0]=0;
      }
    }
  }
}
