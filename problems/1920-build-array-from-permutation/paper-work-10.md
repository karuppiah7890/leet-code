index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

iterating from i = 0, but NOT going one by one
(by increasing i by 1, that is i++), but instead
travelling to the next index which is nums[i]

if i and nums[i] are equal, we don't do travelling

to know if we have overwritten a value to a new value,
put new value x, by doing this:

- x - 1 OR -x-1 OR -(x+1)

y = f(x) = -(x+1)

x = - y - 1 = - f(x) - 1 = - (y + 1) = - ( f(x) + 1 )

for example ->  11 becomes -12, then -12 becomes 11 when reversing. How?

x  -> - x - 1 -> y
11 -> - 11 - 1 -> - 12

             y  -> - (  y ) - 1 OR - (   y + 1)           
reversing: - 12 -> - (- 12) - 1 OR - (- 12 + 1)

We also ensure we visit every index ONLY ONCE and overwrite it ONLY
ONCE

WE DON't WANT TO DO ans[ans[i]]. We only want to do ans[i] = nums[nums[i]]

WE DON'T WANT TO VISIT AN INDEX MULTIPLE TIMES, OVER OVERWRITE A VALUE
AT AN INDEX MULTIPLE TIMES, EVEN IF THE VALUE IS CORRECT OR RIGHT, AS
THIS IS A WASTE OF TIME.

---

index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

i = 0

index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [-1, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

while { if nums[i] is negative, do i++ } ?? or do a search from i = 0 and
do the same thing, and find first positive value nums[i]

i = 1

map of old values?
1 -> 11

i <- 11

index:[ 0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [-1, -2, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

i = 11

map of old values?
1 -> 11 x

index:[ 0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10,  11]
nums: [-1, -2, 9, 4, 7, 10, 5, 8, 6, 3,  2, -12]

i <- 1


---

if there is a k, where

index: ......i......k.....j....
nums:  ......x......y.....z....


---

index:[0,  1, 2, 3, 4,  5, 6, 7, 8, 9, 10, 11]
nums: [0, 11, 9, 4, 7, 10, 5, 8, 6, 3,  2,  1]

values of i
1 -> 11
11 -> 1
since 1 is already travelled to / overwritten
search for first positive value from i = 0
1 -> 0
0 -> 1 (i++ , since 0 is already travelled to / overwritten)
1 -> 2 (i++ , since 1 is already travelled to / overwritten)
search ends here
2 -> 9
9 -> 3
3 -> 4
4 -> 7
7 -> 8
8 -> 6
6 -> 5
5 -> 10
10 -> 2
since 1 is already travelled to / overwritten
search for first positive value (not travelled to) from i = 0
2 -> 0
0 -> 1 (i++ , since 0 is already travelled to / overwritten)
1 -> 2 (i++ , since 1 is already travelled to / overwritten)
2 -> 3 (i++ , since 2 is already travelled to / overwritten)
3 -> 4 (i++ , since 3 is already travelled to / overwritten)
4 -> 5 (i++ , since 4 is already travelled to / overwritten)
5 -> 6 (i++ , since 5 is already travelled to / overwritten)
6 -> 7 (i++ , since 6 is already travelled to / overwritten)
7 -> 8 (i++ , since 7 is already travelled to / overwritten)
8 -> 9 (i++ , since 8 is already travelled to / overwritten)
9 -> 10 (i++ , since 9 is already travelled to / overwritten)
10 -> 11 (i++ , since 10 is already travelled to / overwritten)
11 -> 12 (i++ , since 11 is already travelled to / overwritten)
search ends here, as i >= length

when i becomes length, then we stop search and say that a form of ans[] is
computed and just remove the signs (negative symbols) and subtract 1,
to get ans[]
