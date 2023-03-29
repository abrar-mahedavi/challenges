`Medium`	`Codewriting` 	`300`

Implement the missing code, denoted by ellipses. You may not modify the pre-existing code.
You are playing a game on a rectangular checkerboard with a specific set of rules.
Each game piece has a unique <code>id</code> and each cell on the board is defined by its <code>x</code> and <code>y</code> coordinates.

The current positions of your pieces on the board are stored in the **positions** table with the following structure:

- **id**: unique piece id;
- **x**: x coordinate of the cell that the piece with id <code>id</code> occupies;
- **y**: y coordinate of the cell that the piece with id <code>id</code> occupies.

In this game each piece on the board is said to defend its nearest neighbor. The distance between two pieces is calculated simply as the distance between the points <code>(x1, y1)</code> and <code>(x2, y2)</code>, where <code>(x1, y1)</code> and <code>(x2, y2)</code> are the coordinates of the cells occupied by the first and the second piece, respectively.
You thought it might be good idea to find what piece each of the game pieces defends.

Given the **positions** table, compose the resulting table with two columns: <code>id1</code> and <code>id2</code>, such that on each row the piece with id <code>id1</code> _defends_ the piece with <code>id2</code> (i.e. <code>id2</code> is the closest to <code>id1</code> piece).
The table should be sorted by the values of <code>id1</code> in _ascending_ order.
It's guaranteed that for each piece there is only one other piece closest to it.

## Example

- For the following tables **positions**

| id  | x   | y   |
| --- | --- | --- |
| 1   | 0   | 0   |
| 2   | 2   | 0   |
| 3   | 4   | 3   |
| 4   | 2   | 1   |

the output should be

| id1 | id2 |
| --- | --- |
| 1   | 2   |
| 2   | 4   |
| 3   | 4   |
| 4   | 2   |

![](img.png)

# Output
- [execution time limit] 10 seconds (mysql)

