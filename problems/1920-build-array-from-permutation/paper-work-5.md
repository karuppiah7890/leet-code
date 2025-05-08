index: [0,1,2,3,4,5]
nums:  [0,2,1,5,3,4]

ans:   [0,1,2,4,5,3]

i    x
|    ^
|   /|
|  / |
| /  |
|/   |
x    y

i < x or i > x or i = x

if nums[i] = x and nums[x] = i , that is, i = y

nums[nums[i]]
=> nums[nums[i]] = nums[x] . AS nums[i] = x
=> nums[x] = i . AS nums[x] = i

ans[i] = nums[nums[i]] = i , if nums[i] = x and nums[x] = i

nums[nums[x]]
=> nums[nums[x]] = nums[i] . AS nums[x] = i
=> nums[i] = x . AS nums[i] = x

ans[x] = nums[nums[x]] = x , if nums[i] = x and nums[x] = i

---

index: [0,1,2,3,4,5]
nums:  [0,2,1,5,3,4]
nums:  [0,1,2,4,5,3]

map of old values?

---

index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

i = 0

nums: [0]

i = 1

map of old values?

1 -> 11

nums: [0, 1]

i = 2

map of old values?

1 -> 11
2 -> 9

nums: [0, 1, 3]

i = 3

map of old values?

1 -> 11
2 -> 9
3 -> 4

nums: [0, 1, 3, 7]

i = 4

map of old values?

1 -> 11
2 -> 9
3 -> 4
4 -> 7

nums: [0, 1, 3, 7, 8]

i = 5

map of old values?

1 -> 11
2 -> 9
3 -> 4
4 -> 7
5 -> 10

nums: [0, 1, 3, 7, 8, 2]

i = 6

map of old values?

1 -> 11
2 -> 9
3 -> 4
4 -> 7
5 -> 10 x
6 -> 5

nums: [0, 1, 3, 7, 8, 2, 10]

i = 7

map of old values?
1 -> 11
2 -> 9
3 -> 4
4 -> 7
6 -> 5
7 -> 8

nums: [0, 1, 3, 7, 8, 2, 10, 6]

i = 8

map of old values?
1 -> 11
2 -> 9
3 -> 4
4 -> 7
6 -> 5 x
7 -> 8
8 -> 6

nums: [0, 1, 3, 7, 8, 2, 10, 6, 5]

i = 9

map of old values?
1 -> 11
2 -> 9
3 -> 4 x
4 -> 7
7 -> 8
8 -> 6
9 -> 3

nums: [0, 1, 3, 7, 8, 2, 10, 6, 5, 4]

i = 9

map of old values?
1 -> 11
2 -> 9
3 -> 4 x
4 -> 7
7 -> 8
8 -> 6
9 -> 30
nums: [0, 1, 3, 7, 8, 2, 10, 6, 5, 4]

i = 10

map of old values?
1 -> 11
2 -> 9 x
4 -> 7
7 -> 8
8 -> 6
9 -> 3
10 -> 2

nums: [0, 1, 3, 7, 8, 2, 10, 6, 5, 4, 9]


i = 10

map of old values?
1 -> 11 x
4 -> 7
7 -> 8
8 -> 6
9 -> 3
10 -> 2
11 -> 1

nums: [0, 1, 3, 7, 8, 2, 10, 6, 5, 4, 9, 11]
