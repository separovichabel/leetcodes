import { findMaxConsecutiveOnes } from "."

describe('findMaxConsecutiveOnes', () => {
    it('first case', () => expect(findMaxConsecutiveOnes([1,1,0,1,1,1])).toEqual(3))
    it('second', () => expect(findMaxConsecutiveOnes([1,0,1,1,0,1])).toEqual(2))
    // it('Third', () => expect(minDeletions("ceabaacb")).toEqual(2))
})
