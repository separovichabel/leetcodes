import { ListNode, middleNode } from "."

describe('middleNode', () => {

    it('1', () => {
        expect(middleNode(new ListNode(1, new ListNode(2, new ListNode(3, ))))!.val).toEqual(2)
    })

    it('2', () => {
        expect(
            middleNode(new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5, new ListNode(6, null)))))))!.val
        ).toEqual(4)
    })

    // it('should return true if A listNode is a palindrome', () => {
    //     expect(isArrayPalindrome([1,3,3,1])).toEqual(true)
    // })
})
