`Medium`	`Codewriting` 	`300`

## Description

---

You have been watching a video for some time. Knowing the total video duration find out what portion of the video you have already watched.

**Example**

For <code>part = "02:20:00"</code> and <code>total = "07:00:00"</code>, the output should be
<code>videoPart(part, total) = [1, 3]</code>.

You have watched <code>1 / 3</code> of the whole video.

**Input/Output**

- **[execution time limit] 4 seconds (js)**
- **[input] string part**

  A string of the following format <code>"hh:mm:ss"</code> representing the time you have been watching the video.

- **[input] string total**

  A string of the following format <code>"hh:mm:ss"</code> representing the total video duration.

* **[output] array.integer**

  - An array of the following format <code>[a, b]</code> (where <code>a / b</code> is a reduced fraction).

  Note (reduced fraction): We call a fraction **reduced** if its numerator and denominator are relatively prime.

  Note (relatively prime): Two integers are said to be **relatively prime** (or coprime) if the only positive integer that evenly divides both of them is 1.


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
