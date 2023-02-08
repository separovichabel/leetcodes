import { isArrayPalindrome, isPalindrome, ListNode, reverseLinkedList } from "."

describe('isArrayPalindrome', () => {

    it('should return false if A listNode is not a palindrome', () => {

        expect(isArrayPalindrome([1,2,3,2])).toEqual(false)
    })

    it('should return false if A listNode is empty', () => {
        expect(isArrayPalindrome([1,2,1])).toEqual(true)
    })

    it('should return true if A listNode is a palindrome', () => {
        expect(isArrayPalindrome([1,3,3,1])).toEqual(true)
    })
})

describe('isPalindrome', () => {
    it('should blame palindrome', () => {
        const obj = new ListNode(1, new ListNode(2, new ListNode(1, null)))
        expect(isPalindrome(obj)).toBe(true)
    })

    it('should blame palindrome', () => {
        const obj = new ListNode(1, new ListNode(2, null))
        expect(isPalindrome(obj)).toBe(false)
    })

    it('should not accuse palindrome', () => {
        const obj = new ListNode(1, new ListNode(2, new ListNode(3, null)))
        expect(isPalindrome(obj)).toBe(false)
    })
})

describe('reverseLinkedList', () => {
    it('should invert a linkedList', () => {
        const original = new ListNode(1, new ListNode(2, null))

        const resp = reverseLinkedList(original)
        expect(resp?.val).toBe(2)
        expect(resp!.next!.val).toBe(1)
    })

    it('should invert a linkedList', () => {
        const original = new ListNode(1, new ListNode(2, new ListNode(3, null)))

        const resp = reverseLinkedList(original)
        expect(resp?.val).toBe(3)
        expect(resp!.next!.val).toBe(2)
    })
})