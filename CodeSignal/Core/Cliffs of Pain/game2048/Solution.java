public class Solution {

  int[][] solution(int[][] grid, String path) {
    ArrayDeque<Integer> queue = new ArrayDeque<>();
    for (char c : path.toCharArray()) {
      for (int i = 0; i < 4; i++) {
        for (int j = 0; j < 4; j++) {
          if (c == 'U' && grid[j][i] != 0) {
            queue.addLast(grid[j][i]);
          } else if (c == 'D' && grid[j][i] != 0) {
            queue.addFirst(grid[j][i]);
          } else if (c == 'L' && grid[i][j] != 0) {
            queue.addLast(grid[i][j]);
          } else if (c == 'R' && grid[i][j] != 0) {
            queue.addFirst(grid[i][j]);
          }
        }
        queue = mergeNumbers(queue);
        for (int j = 0; j < 4; j++) {
          if (c == 'U') {
            grid[j][i] = queue.removeFirst();
          } else if (c == 'D') {
            grid[j][i] = queue.removeLast();
          } else if (c == 'L') {
            grid[i][j] = queue.removeFirst();
          } else if (c == 'R') {
            grid[i][j] = queue.removeLast();
          }
        }
      }
    }
    return grid;
  }

  private ArrayDeque<Integer> mergeNumbers(ArrayDeque<Integer> numbers) {
    int lastNumber = 0;
    int count = numbers.size();
    for (int i = 0; i < count; i++) {
      int newNumber = numbers.removeFirst();
      if (lastNumber == 0) {
        lastNumber = newNumber;
      } else if (lastNumber == newNumber) {
        lastNumber = 0;
        numbers.addLast(newNumber * 2);
      } else {
        numbers.addLast(lastNumber);
        lastNumber = newNumber;
      }
    }
    numbers.addLast(lastNumber);
    while (numbers.size() < 4) {
      numbers.addLast(0);
    }
    return numbers;
  }
}
