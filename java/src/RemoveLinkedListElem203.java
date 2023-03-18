import java.util.List;

public class RemoveLinkedListElem203 {
     public static class ListNode {
         int val;
         ListNode next;
         ListNode() {}
         ListNode(int val) { this.val = val; }
         ListNode(int val, ListNode next) { this.val = val; this.next = next; }
     }

    static class Solution {
        public ListNode removeElements(ListNode head, int val) {
            ListNode prevNode = head;
            ListNode currentNode = head;
            while (currentNode != null) {
                if (currentNode.val == val) {
                    if (prevNode == head && currentNode == prevNode) {
                        prevNode = currentNode.next;
                        head = prevNode;
                    } else {
                        prevNode.next = currentNode.next;
                    }
                }
                else {
                    prevNode = currentNode;
                }
                currentNode = currentNode.next;
            }
            return head;
        }
    }
}
