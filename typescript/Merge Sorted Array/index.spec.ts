import { merge, merge2 } from "."

describe('merge', () => {
    it('first case', () => {
        const arr = [1,2,3,0,0,0]
        merge(arr, 3, [2,5,6], 3)
        expect(arr).toEqual([1,2,2,3,5,6])
    })
    it('second', () => {
        const arr = [1]
        merge(arr, 1, [], 0)
        expect(arr).toEqual([1])
    })
    it('Third', () => {
        const arr: number[] = []
        merge(arr, 0, [1], 1)
        expect(arr).toEqual([1])
    })
});

describe('merge2', () => {
    it('first case', () => {
        const arr = [1,2,3,0,0,0]
        merge2(arr, 3, [2,5,6], 3)
        expect(arr).toEqual([1,2,2,3,5,6])
    })
    it('second', () => {
        const arr = [1]
        merge2(arr, 1, [], 0)
        expect(arr).toEqual([1])
    })
    it('Third', () => {
        const arr: number[] = [0]
        merge2(arr, 0, [1], 1)
        expect(arr).toEqual([1])
    })
})
