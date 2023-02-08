// solução mais rapida
export function isArrayPalindrome(list: Array<number>): boolean {
    const list2 = [...list].reverse()

    let cursor = 0
    while (true) {
        if (cursor >= list.length) {
            return true
        }

        if (list[cursor] !== list2[cursor]) {
            return false
        }
        cursor++
    }
};


 // Definition for singly-linked list.
export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
 

export function isPalindrome(head: ListNode | null): boolean {
    let fast = head;
    let slow = head;

    while (fast?.next) {
        fast = fast.next.next
        slow = slow!.next
    }

    let slowReversed = reverseLinkedList(slow)

    let firstHalf = slowReversed
    let secondHalf = head

    while (firstHalf) {
        if (firstHalf.val !== secondHalf!.val) {
            return false
        }
        firstHalf = firstHalf.next
        secondHalf = secondHalf!.next
    }

    return true
};

export function reverseLinkedList(head: ListNode | null): ListNode | null {
    let next = null;
    let current = head
    let prev = null

    while (current) {
        next = current.next
        current.next = prev 
        prev = current
        current = next
    }

    return prev
}
