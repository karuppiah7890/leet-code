
index: .....i......x.....z....y.....a.....c....b...
nums:  .....x......y.....a....z.....b.....d....c...

we start at i and let's say we end up in x and we keep
looping back to x and never end up in i, ever. is this possible?

Let's say it's possible

It's possible, if, for some d in nums, d = x

BUT THEN, that's ONLY Possible, if i = c and d = x, because there
can't be two x in nums

if i = c and d = x, then from b, we go to c, which is i, so, we
go back to i, and again start looping at i and no where between

---


index: ...k...i......x.....z....y.....a.....c....b.....j....d....
nums:  ...l...x......y.....a....z.....b.....d....c.....k....j....

started at i, ended up at l and have passed through all the numbers in
index and nums. is it possible?

if there's an l in nums, there's an l in index, if that's not i, that is
l != i , let's say it's some m (l = m), and if all the numbers have been passed
through, including m (l = m), then, we are saying that, we ended up in m twice
while starting from i, which is not possible, unless, m = i, which means,
l = m = i

so, you do end up in i, if you start from i

