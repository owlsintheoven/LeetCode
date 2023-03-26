import scala.util.control.Breaks.breakable
import scala.util.control.Breaks.break
//Definition for singly-linked list.
class ListNode(_x: Int = 0, _next: ListNode = null) {
	var next: ListNode = _next
	var x: Int = _x
}

object Solution {
	def setTail(t: Int, h: ListNode): ListNode = {
		if (h == null) ListNode(t)
		else if (h.next == null) ListNode(h.x, ListNode(t))
		else ListNode(h.x, setTail(t, h.next))
	}

	def insertAtK(h: ListNode, n: Int, k: Int): ListNode = {
		if (k == 0) ListNode(n, h)
		else if (h == null)
			if (k > 0) ListNode(n)
			else null
		else ListNode(h.x, insertAtK(h.next, n, k - 1))
	}

	def splitAtK(h: ListNode, k: Int): ListNode = {
		if (k == 0 || h == null) null
		else ListNode(h.x, splitAtK(h.next, k - 1))
	}

	def foldRight(h: ListNode)(f: (Int, ListNode) => ListNode): ListNode = {
		if (h == null) null
		else f(h.x, foldRight(h.next)(f))
	}

	def length(h: ListNode): Int =
		if (h == null) 0
		else 1 + length(h.next)

	def fold(h: ListNode)(f: (ListNode, Int, Int) => ListNode): ListNode = {
		if (h == null) null
		else f(fold(h.next)(f), h.x, length(h) + 1)
	}

//	def foldK(h: ListNode, k: Int, c: Int)(f: (ListNode, Int, Int) => ListNode): ListNode = {
//		if (h == null) null
//		else if (c == k) concat(c)
//		else if (c == 0)  f(foldK(h.next, k, k - 1)(f), h.x, k)
//		else f(foldK(h.next, k, c - 1)(f), h.x, c)
//	}

	def concat(h1: ListNode, h2: ListNode): ListNode = {
		if (h1 == null) h2
		else if (h1.next == null) {
			h1.next = h2
			h1
		}
		else
			ListNode(h1.x, concat(h1.next, h2))
	}

	def reverse(h: ListNode) =
		foldRight(h)(setTail)

	def reverse2(h: ListNode) =
		fold(h)(insertAtK)

	def traverseKSteps(h: ListNode, k: Int): ListNode = {
		if (k == 0) h
		else traverseKSteps(h.next, k - 1)
	}

	def reverseKGroup(h: ListNode, k: Int): ListNode = {
		if (length(h) < k) h
		else concat(reverse(splitAtK(h, k)), reverseKGroup(traverseKSteps(h, k), k))
	}

	def printListNode(h: ListNode): Unit = {
		if (h != null) {
			println(h.x)
			printListNode(h.next)
		}
	}

	def main(args: Array[String]): Unit = {
		var n = ListNode(1, ListNode(2, ListNode(3, ListNode(4, ListNode(5)))))
//		var n = ListNode(1, ListNode(2, ListNode(3, ListNode(4))))
//		var n = ListNode(1, ListNode(2, ListNode(3)))
//		var n2 = ListNode(10, ListNode(20))
//		printListNode(splitAtK(n, 2))
//		println(length(n))
//		printListNode(reverse(n))
//		printListNode(reverse2(n))
		printListNode(reverseKGroup(n, 2))
//		printListNode(concat(n, n2))
//		printListNode(insertAtK(n, 100, 5))
//		printListNode(setTail(100, n))
//		printListNode(reverseKGroup(n, 2))
	}
}
