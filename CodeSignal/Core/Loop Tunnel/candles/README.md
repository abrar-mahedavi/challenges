`Medium`	`Codewriting` 	`300`

## Description

---

When a candle finishes burning it leaves a leftover. <code>makeNew</code> leftovers can be combined to make a new candle, which, when burning down, will in turn leave another leftover.

You have <code>candlesNumber</code> candles in your possession. What's the total number of candles you can burn, assuming that you create new candles as soon as you have enough leftovers?

**Example**

For <code>candlesNumber = 5</code> and <code>makeNew = 2</code>, the output should be
<code>candles(candlesNumber, makeNew) = 9</code>.

Here is what you can do to burn 9 candles:

- burn <code>5</code> candles, obtain <code>5</code> leftovers;
- create <code>2</code> more candles, using <code>4</code> leftovers (<code>1</code> leftover remains);
- burn <code>2</code> candles, end up with <code>3</code> leftovers;
- create another candle using <code>2</code> leftovers (<code>1</code> leftover remains);
- burn the created candle, which gives another leftover (<code>2</code> leftovers in total);
- create a candle from the remaining leftovers;
- burn the last candle.

Thus, you can burn <code>5 + 2 + 1 + 1 = 9</code> candles, which is the answer.

**Input/Output**

- **[execution time limit] 4 seconds (js)**

- **[input] integer candlesNumber**

  The number of candles you have in your possession.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>1 \leq candlesNumber \leq 15</code>.

- **[input] integer makeNew**

  The number of leftovers that you can use up to create a new candle.

  _Guaranteed constraints:_<br>
  <code type='math/tex'>2 \leq makeNew \leq 5</code>.

- **[output] integer**

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
