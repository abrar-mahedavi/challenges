`Medium`	`Codewriting` 	`300`

## Description

---

We define the _weakness_ of number <code>x</code> as the number of positive integers smaller than <code>x</code> that have more divisors than <code>x</code>.

It follows that the _weaker_ the number, the greater overall weakness it has. For the given integer <code>n</code>, you need to answer two questions:
- what is the _weakness_ of the weakest numbers in the range <code>[1, n]</code>?
- how many numbers in the range <code>[1, n]</code> have this _weakness_?

Return the answer as an array of two elements, where the first element is the answer to the first question, and the second element is the answer to the second question.

**Example**

For <code>n = 9</code>, the output should be
<code>weakNumbers(n) = [2, 2]</code>.

Here are the number of divisors and the specific weakness of each number in range <code>[1, 9]</code>:

- <code>1</code>: d(1) = 1, weakness(1) = 0;
- <code>2</code>: d(2) = 2, weakness(2) = 0;
- <code>3</code>: d(3) = 2, weakness(3) = 0;
- <code>4</code>: d(4) = 3, weakness(4) = 0;
- <code>5</code>: d(5) = 2, weakness(5) = 1;
- <code>6</code>: d(6) = 4, weakness(6) = 0;
- <code>7</code>: d(7) = 2, weakness(7) = 2;
- <code>8</code>: d(8) = 4, weakness(8) = 0;
- <code>9</code>: d(9) = 3, weakness(9) = 2.

As you can see, the maximal _weakness_ is <code>2</code>, and there are <code>2</code> numbers with that _weakness_ level.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer n**

  _Guaranteed constraints:_<br>
  <code>1 ≤ n ≤ 1000</code>.

- **[output] array.integer**
  - Array of two elements: the _weakness_ of the _weakest_ number, and the number of integers in range <code>[1, n]</code> with this _weakness_.

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
