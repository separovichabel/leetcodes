import { duplicateZeros } from "."

describe('duplicateZeros', () => {
    it('first case', () => {
        const arr = [1,0,2,3,0,4,5,0]
        duplicateZeros(arr)
        expect(arr).toEqual([1,0,0,2,3,0,0,4])
    })
    it('second', () => {
        const arr = [1,2,3]
        duplicateZeros(arr)
        expect(arr).toEqual([1,2,3])
    })

})

