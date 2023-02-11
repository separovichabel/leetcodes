import { findNumbers } from "."

describe('findNumbers', () => {
    it('first case', () => expect(findNumbers([12,345,2,6,7896])).toEqual(2))
    it('second', () => expect(findNumbers([555,901,482,1771])).toEqual(1))
    // it('Third', () => expect(minDeletions("ceabaacb")).toEqual(2))
})
