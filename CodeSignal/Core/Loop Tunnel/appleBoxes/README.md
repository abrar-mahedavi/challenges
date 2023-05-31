`Medium`	`Codewriting` 	`300`

## Description

---

You have <code>k</code> apple boxes full of apples. Each square box of size <code>m</code> contains <code>m × m</code> apples. You just noticed two interesting properties about the boxes:

1. The smallest box is size <code>1</code>, the next one is size <code>2</code>,..., all the way up to size <code>k</code>.
1. Boxes that have an odd size contain only yellow apples. Boxes that have an even size contain only red apples.

Your task is to calculate the difference between the number of red apples and the number of yellow apples.

**Example**

For <code>k = 5</code>, the output should be
<code>appleBoxes(k) = -15</code>.

There are <code>1 + 3 \* 3 + 5 \* 5 = 35</code> yellow apples and <code>2 \* 2 + 4 \* 4 = 20</code> red apples, making the answer <code>20 - 35 = -15</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer k**

  A positive integer.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>1 \leq k \leq 40</code>.

- **[output] integer**
  - The difference between the two types of apples.

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
