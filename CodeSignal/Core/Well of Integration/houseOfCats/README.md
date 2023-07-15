`Medium`	`Codewriting` 	`300`

## Description

---

There are some people and cats in a house. You are given the number of legs they have all together. Your task is to return an array containing every possible number of people that could be in the house _sorted in ascending order_. It's guaranteed that each person has <code>2</code> legs and each cat has <code>4</code> legs.

**Example**

- For <code>legs = 6</code>, the output should be
  <code>houseOfCats(legs) = [1, 3]</code>.

  There could be either <code>1</code> cat and <code>1</code> person <code>(4 + 2 = 6)</code> or <code>3</code> people <code>(2 * 3 = 6)</code>.

- For <code>legs = 2</code>, the output should be
  <code>houseOfCats(legs) = [1]</code>.

  There can be only <code>1</code> person.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer legs**

  The total number of legs in the house. It's guaranteed,that this number is even.<br>

  _Guaranteed constraints:_<br>
  <code>0 ≤ legs < 50</code>.

- **[output] array.integer**
  - Every possible number of people that can be in the house.
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
