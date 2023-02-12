import { thirdMax } from "."

describe('thirdMax', () => {
    it('first case', () => expect(thirdMax([3,2,1])).toEqual(1))
    it('second', () => expect(thirdMax([1,2])).toEqual(2))
    it('third', () => expect(thirdMax([2,2,3,1])).toEqual(1))

})

