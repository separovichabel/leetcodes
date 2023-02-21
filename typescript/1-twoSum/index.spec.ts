import { twoSum, twoSum2, twoSum3 } from "."

describe('twoSum', () => {
    it('first case', () => expect(twoSum([2,7,11,15], 9)).toEqual([0,1]))
    it('second', () => expect(twoSum([3,2,4], 6)).toEqual([1,2]))
    it('Third', () => {
        const result = twoSum([3,3], 6)
        expect(result.find(v => v === 0)).not.toBe(undefined)
        expect(result.find(v => v === 1)).not.toBe(undefined)
    })
})

describe('twoSum2', () => {
    it('first case', () => expect(twoSum2([2,7,11,15], 9)).toEqual([0,1]))
    it('second', () => expect(twoSum2([3,2,4], 6)).toEqual([1,2]))
    it('Third', () => {
        const result = twoSum2([3,3], 6)
        expect(result.find(v => v === 0)).not.toBe(undefined)
        expect(result.find(v => v === 1)).not.toBe(undefined)
    })
    it('fifth', () => {
        const result = twoSum2([2,5,5,11], 10)
        expect(result.length).toBe(2)
        expect(result.find(v => v === 1)).not.toBe(undefined)
        expect(result.find(v => v === 2)).not.toBe(undefined)
    })
})

describe('twoSum3', () => {
    it('first case', () => expect(twoSum3([2,7,11,15], 9)).toEqual([0,1]))
    it('second', () => expect(twoSum3([3,2,4], 6)).toEqual([1,2]))
    it('Third', () => {
        const result = twoSum3([3,3], 6)
        expect(result.find(v => v === 0)).not.toBe(undefined)
        expect(result.find(v => v === 1)).not.toBe(undefined)
    })
    it('fifth', () => {
        const result = twoSum3([2,5,5,11], 10)
        expect(result.length).toBe(2)
        expect(result.find(v => v === 1)).not.toBe(undefined)
        expect(result.find(v => v === 2)).not.toBe(undefined)
    })
})