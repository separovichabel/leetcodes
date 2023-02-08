import { maximumWealth, MaximumWealth2 } from "."

describe('maximumWealth', () => {
    it('first case', () => expect(maximumWealth([[1,2,3],[3,2,1]])).toEqual(6))
    it('second', () => expect(maximumWealth([[1,5],[7,3],[3,5]])).toEqual(10))
    it('Third', () => expect(maximumWealth([[2,8,7],[7,1,3],[1,9,5]])).toEqual(17))
})


describe('maximumWealth2', () => {
    it('first case', () => expect(MaximumWealth2([[1,2,3],[3,2,1]])).toEqual(6))
    it('second', () => expect(MaximumWealth2([[1,5],[7,3],[3,5]])).toEqual(10))
    it('Third', () => expect(MaximumWealth2([[2,8,7],[7,1,3],[1,9,5]])).toEqual(17))
})