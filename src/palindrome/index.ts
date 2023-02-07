
 // Definition for singly-linked list.
class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
 

function isPalindrome(head: ListNode | null): boolean {

    let curHead = head;
    let midHead = head;

    while (true) {
        if (curHead === null) {
            break;
        }
    }

    while (true) {
        if (curHead === null && midHead === null) {
            return true
        }
    }
};

console.log(isPalindrome(new ListNode(1, new ListNode(2, new ListNode(2, new ListNode(1))))))