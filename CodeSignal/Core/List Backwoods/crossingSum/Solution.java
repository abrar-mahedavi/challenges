public class Solution {

  int solution(int[][] matrix, int a, int b) {
    int sum=0;
    for (int i=0;i<matrix.length;i++)
    {
      for (int j=0; j<matrix[0].length;j++)
      {
        if (i==a || j==b)
        {
          sum+=matrix[i][j];
        }
      }
    }
    return sum;
  }

}
