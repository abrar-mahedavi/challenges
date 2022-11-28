# Max Val and Number of Occurences - O(n)

Given an array of integers, return the maximum value and its number of occurences.

## Example:

```bash
Input: nums = [2, 7, 11, 8, 11, 8, 3, 11]
Output: [11, 3]
Explanation: The maximum value is 11 and it appears 3 times
```

### Note:

Your algorithm should run in O(n) time and use O(1) space.

Follow up:
Could you do this in one pass?

### Hints:

<details>
  <summary>Hint #1</summary>
  
Let's start by thinking about an easier version of the problem. Suppose we only want to return the maximum value in the array and we don't care how many times it occurs. Have we solved a similar easier problem in the past? Can we adapt that solution here?
</details>

<details>
  <summary>Hint #2</summary>

That similar problem is "Minimum Value Of Three". There we tried out each of the three numbers and used a variable minVal to store the smallest one found so far throught the algorithm:

```bash
minVal = A
if B < minVal:
	minVal = B
if C < minVal:
	minVal = C
return minVal
```

How can we adapt this approach for an array of numbers instead of just three?

</details>

<details>
  <summary>Hint #3</summary>

The principle is the same. We will iterate through each value in the array nums and update the maximum accordingly. The variable is now maxVal and stores the greatest element found so far:

```bash
maxVal = nums[0]
for val in nums:
	if val > maxVal:
		maxVal= val
return maxVal
```

Now how can we get the number of occurences of this value in nums?

</details>

<details>
  <summary>Hint #4</summary>

One idea would be that after we found maxVal to traverse the array again and count the occurences in another variable count:

```bash
maxVal = nums[0]
for val in nums:
	if val > maxVal:
		maxVal = val

count = 0
for val in nums:
	if val == maxVal:
		count += 1

return [maxVal, count]
```

This is good, but the follow up asks us to solve the question in one pass. We can also compute count in the same for loop where we compute maxVal. Take a small example and try to figure out how to update both maxVal and count for each val when necessary.

</details>

<details>
  <summary>Hint #5</summary>

For each element val of nums, there are two scenarios that are important:
val > maxVal => We found a new maximum value, so we should update maxVal to val and reset count to 1 as this is the first occurence ever of this value.
val == maxVal => We found a new occurence of the current solution, so we should increment count.
This is how the code would look like:
maxVal, count = nums[0], 0

```bash
for val in nums:
	if val > maxVal:
		maxVal = val
		count = 1
	else if val == maxVal:
		count += 1

return [maxVal, count]
```

</details>
