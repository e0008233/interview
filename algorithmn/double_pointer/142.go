package double_pointer

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

func detectCycle(head *ListNode) *ListNode {
	var slow = head
	var quick = head
	if head == nil {
		return nil
	}
	for {
		if quick.Next == nil || quick.Next.Next == nil {
			return nil
		}
		quick = quick.Next.Next
		slow = slow.Next
		if slow == quick {
			quick = head

			for {
				if quick == slow {
					return slow
				}
				quick = quick.Next
				slow = slow.Next

			}

		}
	}
}
