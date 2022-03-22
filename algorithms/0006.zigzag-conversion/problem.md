# Zigzag Conversion

0006.zigzag-conversion

Medium

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

```
P   A   H   N
A P L S I I G
Y   I   R
```
And then read line by line: `"PAHNAPLSIIGYIR"`

Write the code that will take a string and make this conversion given a number of rows:

`string convert(string s, int numRows);`


Example 1:

```
Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Explanation:
P   A   H   N
A P L S I I G
Y   I   R   
```

Example 2:

```
Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:
P     I    N
A   L S  I G
Y A   H R
P     I
```

Example 3:

```
Input: s = "A", numRows = 1
Output: "A"
```

Added 4:

```
Input: s = "NOWISTHETIMEFORALLGOODMENTOCOMETOTHEAIDOFTHEIRCOUNTRY", numRows = 5
Output:    "NTLNOFUOEIALETTTOTONWHMRGMOEHDHCTITEOODCMEIERRSFOOAIY"
Explanation:

N       T       L       N       O       F       U
O     E I     A L     E T     T T     O T     O N
W   H   M   R   G   M   O   E   H   D   H   C   T
I T     E O     O D     C M     E I     E R     R
S       F       O       O       A       I       Y
```
```
length: 53
(0,0) (1,8) (2, 16) (3,24) (4,32) (5,48) (6,64)

```

Constraints:

```
1 <= s.length <= 1000
s consists of English letters (lower-case and upper-case), ',' and '.'.
1 <= numRows <= 1000
```

Results:

```
Success
Runtime: 20 ms, faster than 36.72% of Go online submissions for Zigzag Conversion.
Memory Usage: 6.9 MB, less than 50.80% of Go online submissions for Zigzag Conversion.
```
