# Reverse Integer

0007.reverse-integer

Medium

Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).


Example 1:

```
Input: x = 123
Output: 321
```

Example 2:

```
Input: x = -123
Output: -321
```

Example 3:

```
Input: x = 120
Output: 21
```

Constraints:

```
-231 <= x <= 231 - 1
```

Results:

```
Runtime: 9 ms, faster than 6.23% of Go online submissions for Reverse Integer.
Memory Usage: 2.1 MB, less than 34.70% of Go online submissions for Reverse Integer.
```

Notes:

Constraints were incorrect. Problem statement of avoiding Int32 overflow was what the judge was looking for.