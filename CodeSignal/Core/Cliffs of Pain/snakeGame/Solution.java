public class Solution {

  char[][] solution(char[][] gameBoard, String commands) {

    ArrayList<Integer[]> snake = new ArrayList<Integer[]>();
    int[] DIR = new int[2];
    int tailCnt = 0;
    //Parse Board
    for (int i = 0; i < gameBoard.length; i++) {
      for (int j = 0; j < gameBoard[0].length; j++) {
        switch (gameBoard[i][j]) {
          case '.':
            break;
          case '*':
            tailCnt++;
            break;
          case '^':
            snake.add(new Integer[]{i, j});
            DIR = UP;
            gameBoard[i][j] = '.';
            break;
          case 'v':
            snake.add(new Integer[]{i, j});
            DIR = DOWN;
            gameBoard[i][j] = '.';
            break;
          case '<':
            snake.add(new Integer[]{i, j});
            DIR = LEFT;
            gameBoard[i][j] = '.';
            break;
          case '>':
            snake.add(new Integer[]{i, j});
            DIR = RIGHT;
            gameBoard[i][j] = '.';
            break;
        }
      }
    }
    //Get Tails and place in snake
    int[] chkDIR = new int[]{DIR[0], DIR[1]};
    Integer[] spot = {snake.get(0)[0], snake.get(0)[1]};
    while (tailCnt > 0) {
      if (spot[0] - chkDIR[0] > -1 && spot[0] - chkDIR[0] < gameBoard.length
        && spot[1] - chkDIR[1] > -1 && spot[1] - chkDIR[1] < gameBoard[0].length) {
        if (gameBoard[spot[0] - chkDIR[0]][spot[1] - chkDIR[1]] == '*') {
          gameBoard[spot[0] - chkDIR[0]][spot[1] - chkDIR[1]] = '.';
          tailCnt--;
          spot[0] -= chkDIR[0];
          spot[1] -= chkDIR[1];
          snake.add(new Integer[]{spot[0], spot[1]});
          continue;
        }
      }
      if (spot[0] - chkDIR[1] > -1 && spot[0] - chkDIR[1] < gameBoard.length
        && spot[1] - chkDIR[0] > -1 && spot[1] - chkDIR[0] < gameBoard[0].length) {
        if (gameBoard[spot[0] - chkDIR[1]][spot[1] - chkDIR[0]] == '*') {
          gameBoard[spot[0] - chkDIR[1]][spot[1] - chkDIR[0]] = '.';
          tailCnt--;
          spot[0] -= chkDIR[1];
          spot[1] -= chkDIR[0];
          snake.add(new Integer[]{spot[0], spot[1]});
          swapRC(chkDIR);
          continue;
        }
      }
      if (spot[0] + chkDIR[1] > -1 && spot[0] + chkDIR[1] < gameBoard.length
        && spot[1] + chkDIR[0] > -1 && spot[1] + chkDIR[0] < gameBoard[0].length) {
        if (gameBoard[spot[0] + chkDIR[1]][spot[1] + chkDIR[0]] == '*') {
          gameBoard[spot[0] + chkDIR[1]][spot[1] + chkDIR[0]] = '.';
          tailCnt--;
          spot[0] += chkDIR[1];
          spot[1] += chkDIR[0];
          snake.add(new Integer[]{spot[0], spot[1]});
          swapRC(chkDIR);
          chkDIR[1] *= -1;
          chkDIR[0] *= -1;
          continue;
        }
      }
    }

    //Move snake
    boolean isDead = false;
    for (char in : commands.toCharArray()) {
      if (in == 'L') {
        DIR = left(DIR);
      } else if (in == 'R') {
        DIR = right(DIR);
      } else {
        spot = snake.get(0);

        if (spot[0] + DIR[0] < 0 ||
          spot[0] + DIR[0] >= gameBoard.length ||
          spot[1] + DIR[1] < 0 ||
          spot[1] + DIR[1] >= gameBoard[0].length) {
          isDead = true;
          break;
        }

        Integer[] chkSpot = {spot[0] + DIR[0], spot[1] + DIR[1]};
        for (int i = 1; i < snake.size(); i++) {
          if (Arrays.equals(snake.get(i), chkSpot)) {
            isDead = true;
            break;
          }
        }
        if (isDead) {
          break;
        }

        for (int i = snake.size() - 1; i > 0; i--) {
          snake.get(i)[0] = snake.get(i - 1)[0];
          snake.get(i)[1] = snake.get(i - 1)[1];
        }

        snake.get(0)[0] = chkSpot[0];
        snake.get(0)[1] = chkSpot[1];

      }
    }

    for (int i = 0; i < snake.size(); i++) {
      if (isDead == true) {
        gameBoard[snake.get(i)[0]][snake.get(i)[1]] = 'X';
      } else if (i > 0) {
        gameBoard[snake.get(i)[0]][snake.get(i)[1]] = '*';
      } else {
        if (DIR == LEFT) {
          gameBoard[snake.get(i)[0]][snake.get(i)[1]] = '<';
        }
        if (DIR == RIGHT) {
          gameBoard[snake.get(i)[0]][snake.get(i)[1]] = '>';
        }
        if (DIR == UP) {
          gameBoard[snake.get(i)[0]][snake.get(i)[1]] = '^';
        }
        if (DIR == DOWN) {
          gameBoard[snake.get(i)[0]][snake.get(i)[1]] = 'v';
        }
      }
    }

    return gameBoard;


  }

  static int[] LEFT = {0, -1}, RIGHT = {0, 1}, UP = {-1, 0}, DOWN = {1, 0};

  void swapRC(int[] rc) {
    int tmp = rc[0];
    rc[0] = rc[1];
    rc[1] = tmp;
  }

  int[] left(int[] dir) {
    if (dir == RIGHT) {
      return UP;
    } else if (dir == UP) {
      return LEFT;
    } else if (dir == LEFT) {
      return DOWN;
    } else {
      return RIGHT;
    }
  }

  int[] right(int[] dir) {
    if (dir == RIGHT) {
      return DOWN;
    } else if (dir == UP) {
      return RIGHT;
    } else if (dir == LEFT) {
      return UP;
    } else {
      return LEFT;
    }
  }

}
