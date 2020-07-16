/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */

// function ListNode(val, next) {
//     return {
//         val: val,
//         next: next
//     }
// }

/**
 * @param {ListNode[]} lists
 * @return {ListNode}
 */
var mergeKLists = function(lists) {
    let head= null;
    let tail = null;
    while (true) {
        let minVal = Number.MAX_VALUE;
        let minI = -1;
        let nextNode = null;
        for (let i = 0; i < lists.length; i++) {
            if (lists[i] !== null && lists[i].val < minVal) {
                minVal = lists[i].val;
                minI = i;
                nextNode = lists[i].next;
            }
        }
        // console.log('min val', minVal);
        if (minI === -1) {
            break;
        }

        if (head === null) {
            head = lists[minI];
            tail = lists[minI];
        } else {
            tail.next = lists[minI];
            tail = tail.next;
        }

        lists[minI] = nextNode;
    }

    return head;
};

// /**
//  * @param {ListNode} list
//  * @return
//  */
// var traverse = function(list) {
//     while(list !== null) {
//         console.log(list.val);
//         list = list.next;
//     }
// }

// l13 = ListNode(5, null);
// l12 = ListNode(4, l13);
// l11 = ListNode(1, l12);

// l23 = ListNode(4, null);
// l22 = ListNode(3, l23);
// l21 = ListNode(1, l22);

// l32 = ListNode(6, null);
// l31 = ListNode(2, l32);

// traverse(mergeKLists([l11, l21, l31]));
// traverse(l31)