index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

iterating from i = 0, but NOT going one by one
(by increasing i by 1, that is i++), but instead
travelling to the next index which is nums[i]

i = 0

nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

i = 1

map of old values?
1 -> 11

i <- 11

nums: [0, 1, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

i = 11

map of old values?
1 -> 11 x

i <- 1 THIS IS WRONG OR TRICKY OR PROBLEMATIC

nums: [0, 1, 9, 4, 7, 10, 5, 8, 6, 3,  2,  11]

WE DON't WANT TO DO ans[ans[i]]. We only want to do ans[i] = nums[nums[i]]

map or array of overwritten values?
11 -> null
1 -> null

11 -> true
1 -> true
