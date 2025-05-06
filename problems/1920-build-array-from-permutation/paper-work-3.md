Input:
nums: [0,2,1,5,3,4]


Expected Output:
ans: [0,1,2,4,5,3]

---

       0,1,2,3,4,5
nums: [0,2,1,5,3,4]

for i = 0

= nums[nums[i]]
= nums[nums[0]]
= nums[0]
= 0

old value is 0
new value is also 0

for i = 1

= nums[nums[i]]
= nums[nums[1]]
= nums[2]
= 1

old value is 2
new value is 1

if we replace the 2 with 1, if someone needs the value 2,
they will not be able to use it.

For example, for some k, where nums[k] = 1, we can't find
ans[k] = nums[nums[k]] = nums[1] because the initial / old value of nums[1] will be overwritten, and we do not want the overwritten value

---
