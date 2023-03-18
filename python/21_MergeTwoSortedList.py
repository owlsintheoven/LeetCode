# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def mergeTwoLists(self, node1, node2):
        head = None
        prev = None
        while node1 != None and node2 != None:
            if node1.val <= node2.val:
                curr = ListNode(node1.val, None)
                node1 = node1.next
            else:
                curr = ListNode(node2.val, None)
                node2 = node2.next
            if prev == None:
                prev = curr
                head = curr
            else:
                prev.next = curr
                prev = curr

        if node1 == None:
            curr = node2
        else:
            curr = node1
        if prev == None:
            head = curr
        else:
            prev.next = curr
        return head


if __name__ == '__main__':
    solution = Solution()
    l13 = ListNode(4, None)
    l12 = ListNode(2, l13)
    l11 = ListNode(1, l12)
    l23 = ListNode(4, None)
    l22 = ListNode(3, l23)
    l21 = ListNode(1, l22)
    #l11 = None
    #l21 = ListNode(0, None)
    l11 = None
    l21 = None
    answer = solution.mergeTwoLists(l11, l21)
    while answer != None:
        print(answer.val)
        answer = answer.next