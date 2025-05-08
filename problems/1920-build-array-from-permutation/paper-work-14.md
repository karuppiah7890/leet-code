
initially, index = 0

loop over nums[] {

    if value at index is negative, then,
    index++
    continue

    ---

    if index == value at index, then:

    NEWnums[index] = nums[nums[index]]
    => NEWnums[index] = nums[value] . AS nums[index] = value
    => NEWnums[index] = nums[index] . AS index == value
    => NEWnums[index] = value . AS nums[index] = value
    => NEWnums[index] = value = index . AS index == value

    => NEWnums[index] = - NEWnums[index] - 1 
        ^^^^ TO MARK THIS AS TRAVELLED TO OR VISITED AND UPDATED / GOTTEN NEW VALUE

    index ++

    ---

    if index != value at index, then:

    initially, stopPointIndex = index
    initially, stopPointIndexOldValue = value at index

    loop {
        nextIndex = value
        if value == stopPointIndex {
            => NEWnums[index] = nums[value] (but original) . AS nums[index] = value
            For original value, we do this:
            => NEWnums[index] = stopPointIndexOldValue . AS nums[index] = value
            ^^^^ AS stopPointIndexOldValue is the old value or original
                value at stopPointIndex and we need the original value
                of nums[value]
        } else {
            => NEWnums[index] = nums[value] . AS nums[index] = value
        }

    => NEWnums[index] = - NEWnums[index] - 1 
        ^^^^ TO MARK THIS AS TRAVELLED TO OR VISITED AND UPDATED / GOTTEN NEW VALUE

        index <- nextIndex

        if index == stopPointIndex {
            break; // break out of loop
        }
    }
}

---

initially, index = 0

loop over nums[] {
    NEWnums[index] = - NEWnums[index] - 1
    index++
}

---

And above algorithm will work when NEWnums[] and nums[] point to the same
array
