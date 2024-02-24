package google

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	}

	newList := make([]*ListNode, 0)
	for i := 0; i < len(lists); i = i + 2 {
		if i+1 < len(lists) {
			newList = append(newList, merge2Lists(lists[i], lists[i+1]))
		} else {
			newList = append(newList, merge2Lists(lists[i], nil))
		}
	}
	return mergeKLists(newList)
}

func merge2Lists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	fakeHead := &ListNode{
		Next: nil,
	}
	curr := fakeHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
			curr = curr.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
			curr = curr.Next
		}
	}
	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}
	return fakeHead.Next
}
