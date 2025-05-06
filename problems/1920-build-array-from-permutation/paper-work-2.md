Input:
nums: [0,2,1,5,3,4]


Expected Output:
ans: [0,1,2,4,5,3]

---

NOT use another full array to create the result. Instead, mutate the
input array maybe. Are there any other ways? Not sure, I don't know

---

nums: [0,2,1,5,3,4]

i = 0
nums[i] = nums[0] = 0
nums[nums[i]] = nums[0] = 0

---

Input:
index: .....i......k.....
nums:  .....x......i.....

Output:
index: .....i......k.....
ans:   .....y......x.....

Assuming:
- i < k
- 0 <= i < length -1
- 0 <= k < length -1
- nums[i] = x
- 0 <= x < length -1
- nums[k] = i
- nums[x] = y
- 0 <= y < length -1


Based on the above assumptions, we can say that -

ans[i] = nums[nums[i]]
=> ans[i] = nums[x] . AS nums[i] = x
=> ans[i] = y . AS nums[x] = y

Based on the above assumptions and current inferences from those
assumptions, we can say that -

ans[k] = nums[nums[k]]
=> ans[k] = nums[i] . AS nums[k] = i
=> ans[k] = x . AS nums[i] = x

NOW, let's assume, i = x

Assuming:
- i = x
- i < k
- 0 <= i < length -1
- 0 <= k < length -1
- nums[i] = x
- 0 <= x < length -1
- nums[k] = i
- nums[x] = y
- 0 <= y < length -1

Inferences from some of the assumptions
- ans[i] = y
- ans[k] = x

Based on the above assumptions and current inferences from those
assumptions, we can say that -

ans[i] = y
=> ans[i] = ans[x] = y . AS i = x
=> ans[x] = y

ans[k] = x
=> ans[k] = i . AS x = i

ans[x] = y
=> ans[ans[k]] = y . AS ans[k] = x

if

i = x

AND

nums[i] = x = i

AND

nums[k] = i = x

then, nums[i] = nums[k], which means it's only possible if i = k,
because, there can't two numbers in nums[] that are same with
different indices / indexes, i and k. So, i must be equal to k, OR
i = k

if

i = k 

AND

i = x

THEN

i = k = x

then

ans[i] = nums[nums[i]]
=> ans[i] = nums[x] . AS nums[i] = x
=> ans[i] = nums[i] . AS x = i

nums[x] = y
=> nums[i] = y . AS x = i
=> x = y . AS nums[i] = x

SO, if i = k = x

AND x = y

THEN i = k = x = y

