export class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}

export function isPalindrome(head: ListNode | null): boolean {
    let fastCursor = head;
    let slowCursor = head;

    while (!!fastCursor && fastCursor!.next) {
        fastCursor = fastCursor.next.next
        slowCursor = slowCursor!.next
    }

    let cur: ListNode | null = slowCursor
    let next: ListNode | null = null;
    let prev: ListNode | null = null;

    while (!!cur) {
        next = cur!.next

        cur!.next = prev
        prev = cur

        cur = next
    }

    let inverted = prev

    while (inverted) {
        if (inverted!.val !== head!.val) return false
        inverted = inverted!.next
        head = head!.next
    }

    return true
};