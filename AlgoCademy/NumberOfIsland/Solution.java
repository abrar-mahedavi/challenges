package com.abu.algocademy.numberofisland;

public class Solution {

  boolean isValidCell(int[][] grid, int row, int col) {
    try {
      return grid[row][col] == 1;
    } catch (Exception exception) {
      return false;
    }
  }

  public void dfs(int[][] grid, int row, int col) {
    if (!isValidCell(grid, row, col)) {
      return;
    }
    grid[row][col] = 0;
    int[][] directions = new int[][]{{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
    for (int[] d : directions) {
      dfs(grid, row + d[0], col + d[1]);
    }

  }

  public int numIslands(int[][] grid) {
    int result = 0;
    for (int i = 0; i < grid.length; i++) {
      for (int j = 0; j < grid[0].length; j++) {
        if (grid[i][j] == 1) {
          result++;
          dfs(grid, i, j);
        }
      }
    }
    return result;
  }

}
