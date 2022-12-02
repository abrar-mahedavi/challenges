# Binary Strings with at most K Consecutive One - O(n * k)

Given two non-negative integers `N` and `K`, return the number of binary strings of length `N` with at most `K` consecutive ones

## Example

```bash
Input: N = 4, K = 2
Output: 13
Explanation: [
    "0000",
    "0001",
    "0010",
    "0011",
    "0100",
    "0101",
    "0110",
    "1000",
    "1001",
    "1010",
    "1011",
    "1100",
    "1101",
]
"0111", "1110" and "1111" are not valid because they have more than 2 consecutive ones

```

Note:
Notes: K <= N <= 30


## Hints:

<details>
  <summary>Hint #1</summary>

Let's come up with a brute force approach. Brute force tries each and every candidate to our solution. For this problem, the candidates are all the subarrays of `nums`. How can we try out each subarray?

</details>

<details>
  <summary>Hint #2</summary>

A subarray is defined by its left endpoint, let's name it index `i` and its right point, let's name it index `j`. We should go through each pair of indices `(i, j)` such that `0 <= i <= j < n`. How can we achieve this?

</details>

<details>
  <summary>Hint #3</summary>

Using two nested for loops! The first loop will try out each possible `i` between `0` and n - 1. The second loop will be written inside the first one and will try out each possible `j` between `i` and n - 1.

```bash
for each i : 0 -> n - 1:
		for each j : i -> n - 1:
```

We will use a variable maxSum to store the greatest sum found so far and for each pair (i, j), update maxSum with `sum(nums[i...j])` if it's the case:

```bash
maxSum = nums[0]
for each i : 0 -> n - 1:
	for each j : i -> n - 1:
		maxSum = max(maxSum, sum(nums[i...j]))
return maxSum
```

All we have to do now is find a way of getting `sum(nums[i...j])` in O(1) time. How can we do this?

</details>

<details>
  <summary>Hint #4</summary>

One way is using Partial Sums. But that implies using an extra array which gives extra space of O(n). There is a better way. We can use a variable sum to store the `sum` of the current subarray `sum(nums[i...j])`.

```bash
maxSum = nums[0]
for each i : 0 -> n - 1:
	for each j : i -> n - 1:
		maxSum = max(maxSum, sum)
return maxSum
```

How should we update this variable throught the algorithm to make sure that it always stores the correct sum?

</details>

<details>
  <summary>Hint #5</summary>
  
Each time we increment `j` in the second loop, we should add `nums[j]` to `sum`. Also, each time we increment `i` in the first loop, we start with a new empty subarray so we should reset `sum` to 0:
```bash
maxSum = nums[0]
for each i : 0 -> n - 1:
	sum = 0
	for each j : i -> n - 1:
		sum += nums[j]
		maxSum = max(maxSum, sum)
return maxSum
```
