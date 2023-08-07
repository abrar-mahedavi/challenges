`Medium`	`Codewriting` 	`300`

## Description

---

Yesterday you found some shoes in the back of your closet. Each shoe is described by two values:

- _type_ indicates if it's a left or a right shoe;
- _size_ is the size of the shoe.

Your task is to check whether it is possible to pair the shoes you found in such a way that each pair consists of a right and a left shoe of an equal size.

**Example**

- For

  <code>
  shoes = [[0, 21],<br>
          [1, 23],<br>
          [1, 21],<br>
          [0, 23]]<br>
  </code>

  the output should be
  `pairOfShoes(shoes) = true`;

- For
  <code>
  shoes = [[0, 21],<br>
  [1, 23],<br>
  [1, 21],<br>
  [1, 23]]<br>
  </code>
  the output should be
  `pairOfShoes(shoes) = false`.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.array.integer shoes**

  Array of shoes. Each shoe is given in the format <code>[type, size]</code>, where type is either <code>0</code> or <code>1</code> for left and right respectively, and size is a positive integer.<br>

  _Guaranteed constraints:_<br>
  <code>1 ≤ shoes.length ≤ 200</code>,<br> <code>1 ≤ shoes[i][1] ≤ 100</code>.

- **[output] boolean**
  - <code>true</code> if it is possible to pair the shoes, <code>false</code> otherwise.

## [Java] Syntax Tips

``` java
// Prints help message to the console
// Returns a string
// 
// Globals declared here will cause a compilation error,
// declare variables inside the function instead!
String helloWorld(String name) {
    System.out.println("This prints to the console when you Run Tests");
    return "Hello, " + name;
}
```
