index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  1,  2]

expected ans: [0, 2, 3, 7, 8, 1, 10, 6, 5, 4, 11, 9]

i = 0

nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  1,  2]

---

i = 1

nums[i] = nums[1] = 11
ans[i] = nums[nums[1]] = nums[11] = 2

is there any other i, where nums[i] = 11, that needs the value 2? 2, which is
basically ans[i] = nums[nums[i]] = nums[11] = 2

