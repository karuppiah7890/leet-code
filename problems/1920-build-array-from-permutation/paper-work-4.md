index:0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11
nums: 0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1

0 -> nums[nums[0]] -> nums[0] = 0

1 -> nums[nums[1]] -> nums[11] = 1

2 -> nums[nums[2]] -> nums[9] = 3

3 -> nums[nums[3]] -> nums[4] = 7

4 -> nums[nums[4]] -> nums[7] = 8

5 -> nums[nums[5]] -> nums[10] = 2

6 -> nums[nums[6]] -> nums[5] = 10

7 -> nums[nums[7]] -> nums[8] = 6

8 -> nums[nums[8]] -> nums[6] = 5

9 -> nums[nums[9]] -> nums[3] = 4

10 -> nums[nums[10]] -> nums[2] = 9

11 -> nums[nums[11]] -> nums[1] = 11

ans:
0, 1, 3, 7, 8, 2, 10, 6, 5, 4, 9, 11

---


index:0,  1, 2, 3, 4,  5,  6, 7, 8, 9, 10, 11
nums: 0, 11, 9, 4, 7, 10,  5, 8, 6, 3,  2,  1
ans:  0,  1, 3, 7, 8,  2, 10, 6, 5, 4,  9, 11
