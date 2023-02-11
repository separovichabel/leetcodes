import { sortedSquares } from "."

describe('sortedSquares', () => {
    it('first case', () => expect(sortedSquares([-4,-1,0,3,10])).toEqual([0,1,9,16,100]))
    it('second', () => expect(sortedSquares([-7,-3,2,3,11])).toEqual([4,9,9,49,121]))

    it('Third', () => expect(sortedSquares([-1])).toEqual([1]))
})

