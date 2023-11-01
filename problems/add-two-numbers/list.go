package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(digits []int) *ListNode {
	var firstNode *ListNode = nil
	var lastNode *ListNode = nil

	for index := 0; index < len(digits); index++ {
		newLastNode := &ListNode{
			Val:  digits[index],
			Next: nil,
		}
		if lastNode != nil {
			lastNode.Next = newLastNode
		}
		lastNode = newLastNode

		if index == 0 {
			firstNode = lastNode
		}
	}

	return firstNode
}

func (list *ListNode) Add(anotherList *ListNode) *ListNode {
	carryOver := 0

	var firstNode *ListNode = nil
	var lastNode *ListNode = nil
	digit := list
	anotherDigit := anotherList

	for digit != nil && anotherDigit != nil {
		sum := digit.Val + anotherDigit.Val
		carryOver = sum / 10
		sumDigit := sum % 10

		newLastNode := &ListNode{
			Val:  sumDigit,
			Next: nil,
		}
		if lastNode != nil {
			lastNode.Next = newLastNode
		}
		lastNode = newLastNode

		if firstNode == nil {
			firstNode = lastNode
		}

		digit = digit.Next
		anotherDigit = anotherDigit.Next
	}

	return firstNode
}

func (list *ListNode) Equals(anotherList *ListNode) bool {
	digit := list
	anotherDigit := anotherList
	for digit != nil && anotherDigit != nil {
		if digit.Val != anotherDigit.Val {
			return false
		}

		digit = digit.Next
		anotherDigit = anotherDigit.Next
	}

	if digit != nil || anotherDigit != nil {
		return false
	}

	return true
}

func (list *ListNode) Digits() []int {
	node := list
	digits := []int{}

	for node != nil {
		digits = append(digits, node.Val)
		node = node.Next
	}

	return digits
}
