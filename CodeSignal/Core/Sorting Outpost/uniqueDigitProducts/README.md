`Medium`	`Codewriting` 	`300`

## Description

---

Let's call <code>product(x)</code> the product of <code>x</code>'s digits. Given an array of integers <code>a</code>, calculate <code>product(x)</code> for each <code>x</code> in <code>a</code>, and return the number of distinct results you get.

**Example**

For <code>a = [2, 8, 121, 42, 222, 23]</code>, the output should be
<code>uniqueDigitProducts(a) = 3</code>.

Here are the products of the array's elements:

- <code>2</code>: <code>product(2) = 2</code>;
- <code>8</code>: <code>product(8) = 8</code>;
- <code>121</code>: <code>product(121) = 1 \* 2 \* 1</code> = 2;
- <code>42</code>: <code>product(42) = 4 \* 2 = 8</code>;
- <code>222</code>: <code>product(222) = 2 \* 2 \* 2</code> = 8;
- <code>23</code>: <code>product(23) = 2 \* 3 = 6</code>.

As you can see, there are only <code>3</code> different products: <code>2</code>, <code>6</code> and <code>8</code>.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] array.integer a**

  _Guaranteed constraints:_<br>
  <code>1 ≤ a.length ≤ 10<sup>5</sup>,<br>
  <code>1 ≤ a[i] ≤ 10<sup>9</sup></code>.

* **[output] integer**
  - The number of different digit products in <code>a</code>.

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
