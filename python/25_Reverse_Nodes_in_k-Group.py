# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
class Solution(object):
    def reverseKGroup(self, head, k):
        l = []
        while head != None:
            l.append(head.val)
            head = head.next
        i = 0
        n = []
        while i < len(l) - k:
            u = l[i: i+k]
            u.reverse()
            n.extend(u)
            i += k
        n.extend(l[i:])
        next = None
        for i in range(len(n)-1, -1, -1):
            cur = ListNode(n[i])
            cur.next = next
            next = cur

        return cur

    def setTail(self, head, t):
        if head == None:
            return ListNode(t)
        elif head.next == None:
            head.next = ListNode(t)
            return head
        else:
            ListNode(head.val, self.setTail(head.next, t))




    def reverse(self, head):
        if head == None:
            return None
        return ListNode()


    def concat(self, l1, l2):
        if l1 == None:
            return l2
        h = l1
        while l1 != None:
            if l1.next == None:
                l1.next = l2
                return h
            else:
                l1 = l1.next

    def print(self, head):
        while head != None:
            print(head.val)
            head = head.next


if __name__ == '__main__':
    solution = Solution()
    h5 = ListNode(5)
    h4 = ListNode(4, h5)
    h3 = ListNode(3, h4)
    h2 = ListNode(2, h3)
    h1 = ListNode(1, h2)
    solution.print(solution.reverseKGroup(h1, 2))
#    solution.print(solution.concat(h1, None))
