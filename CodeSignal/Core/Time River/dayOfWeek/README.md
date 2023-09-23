`Medium`	`Codewriting` 	`300`

## Description

---

Whenever you decide to celebrate your birthday you always do this your favorite café, which is quite popular and as such usually very crowded. This year you got lucky: when you and your friend enter the café you're surprised to see that it's almost empty. The waiter lets slip that there are always very few people on this day of the week.

You enjoyed having the café all to yourself, and are now curious about the next time you'll be this lucky. Given the current <code>birthdayDate</code>, determine the number of years until it will fall on the same day of the week.

For your convenience, here is the list of months lengths (from January to December, respectively):

- Months lengths: <code>31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31</code>.

Please, note that in leap years February has <code>29</code> days. If your birthday is on the <code>29<sup>th</sup></code> of February, you celebrate it once in four years. Otherwise you birthday is celebrated each year.

Note (leap years): Year is leap if its number is divisible by 4 and isn’t divisible by 100 or if its number is divisible by 400.

**Example**

For <code>birthdayDate = "02-01-2016"</code>, the output should be
<code>dayOfWeek(birthdayDate) = 5</code>.

February 1 in <code>2016</code> is a Monday. The next year in which this same date will be Monday too is <code>2021</code>. <code>2021 - 2016 = 5</code>, which is the answer.

**Input/Output**

- **[execution time limit] 4 seconds (js)**
- **[input] string birthdayDate**

  A string representing the correct date in the format <code>mm-dd-yyyy</code>, where <code>mm</code> is the number of month (1-based, i.e. <code>01</code> for January, <code>02</code> for February and so on), dd is the day, and <code>yyyy</code> is the year.

  _Guaranteed constraints:_<br>
  <code>1 ≤ int(mm) ≤ 12</code>,<br>
  <code>1 ≤ int(dd) ≤ 31</code>,<br>
  <code>1900 ≤ int(yyyy) ≤ 2100</code>.

* **[output] integer**

  - An integer describing the number of years until your birthday falls on the same day of the week.


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
