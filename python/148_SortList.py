# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def sortList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if head == None or head.next == None:
            return head
        prev = None
        slow = head
        fast = head
        while fast != None and fast.next != None:
            prev = slow
            slow = slow.next
            fast = fast.next.next
        prev.next = None
        l1 = self.sortList(head)
        l2 = self.sortList(slow)

        return self.mergeList(l1, l2)

    def mergeList(self, node1, node2):
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
    node3 = ListNode(3, None)
    node2 = ListNode(1, node3)
    node1 = ListNode(2, node2)
    root = ListNode(4, node1)

    solution = Solution()
    answer = solution.sortList(root)
    while answer != None:
        print(answer.val)
        answer = answer.next
