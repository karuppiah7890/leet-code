Input: [0,2,1,5,3,4]

Mutate Input

ans[i] = nums[nums[i]]

ans = nums

nums[i] = nums[nums[i]]

nums[i] - to read
nums[nums[i]] - to read

writing to nums[i], ans = nums

---

i = 0

nums[0] -> x

nums[nums[0]] -> y
nums[x] -> y

nums[0] = nums[nums[0]]

nums[0] -> y

nums[0] (x) is needed when nums[j] = 0, for example, for a k,
nums[k] = nums[nums[k]], where nums[k] = j, nums[k] = nums[j], then
nums[k] = 0 , then j and k are equal. j = k = 0

---

i

0 <= i <= len-1

0 <= nums[i] <= len-1

and every i is distinct and every nums[i] is distinct

read nums[i] -> x
read nums[nums[i]] -> nums[x] -> y
write y to nums[i], you lose the value x

who requires nums[i] old value, the value x

nums[k] = nums[nums[k]] will require the value x when nums[k] = i, and hence
nums[nums[k]] -> nums[i] -> x
write x to nums[k]

it means,

i < k

assuming nums[i] = x, nums[k] = i

assuming traversing from left to right

input (nums):
index  ....i......k....
values ....x......i....

output (ans):
index  ....i......k....
values ....y......x....

cases:
if nums[i] -> i

x = i

y -> nums[nums[i]] -> nums[i] -> i

y = i

if nums[k] = i => k = i. then, x = k = i

---

relationship between i and x

i = x

OR

i < x

OR

i > x

---

if i = x, then let's look at our scenario

fact: i < k

assuming nums[i] = x, nums[x] = y, nums[k] = i

then ans[i] = nums[nums[i]] = nums[x] = y

assuming traversing from left to right

input (nums):
index  ....i......k....
values ....x......i....

output (ans):
index  ....i......k....
values ....y......x....

if i = x, and i < k then

x < k, because i < k and i = x

if i = x, and nums[i] = x

if nums[k] = i, and i = x, then nums[k] = x

if nums[k] = x and nums[i] = x, then k = i, as i and k can't be different,
then that would mean there are two x values, which is not possible, as all
values in nums[] are distinct

if i = x, and i = k, then i = x = k, nums[nums[i]] = nums[x] = y and
nums[x] = nums[i], then nums[i] = nums[x] = y, then nums[i] = y, also
ans[i] = y, so, nums[i] = ans[i] = y. Also, nums[i] = x and nums[i] = ans[i],
ans[i] = x, then x = y

i = x = k = y

ans[i] = nums[nums[i]]
ans[i] = nums[x] as nums[i] = x
ans[i] = nums[i] as i = x
ans[i] = x as nums[i] = x
ans[i] = i as i = x

ans[x] = x as i = x

---

if i < x, then let's look at our scenario

fact: i < k

assuming nums[i] = x, nums[x] = y, nums[k] = i

then ans[i] = nums[nums[i]] = nums[x] = y

Don't want to lose the value x if it's needed later, for index k of ans[].
That is, ans[k] = nums[nums[k]]; ans[k] = nums[i] as nums[k] = i; ans[k] = x

assuming traversing from left to right

input (nums):
index  ....i......k....
values ....x......i....

output (ans):
index  ....i......k....
values ....y......x....

if i < x, and i < k the

---

Index: [0,1,2,3,4,5]
Input: [0,2,1,5,3,4]

Expected Output: [0,1,2,4,5,3]

ans[0] = nums[nums[0]] = nums[0] = 0

ans[1] = nums[nums[1]] = nums[2] = 1

who needs 2? who needs nums[1]?

they will need nums[1] if

ans[y] = nums[nums[y]], where nums[y] = 1, then
ans[y] = nums[1] = 2

then, if nums[y] = 1, then y = 2
ans[y] = ans[2] = nums[nums[2]] = nums[1] = 2

who needs 1? who needs nums[2]?
they will need nums[2] if

ans[z] = nums[nums[z]], where nums[z] = 2, then
ans[z] = nums[2] = 1

---

Expected Output: [0,1,2,4,5,3]
