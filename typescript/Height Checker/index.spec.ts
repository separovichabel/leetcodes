import { heightChecker } from "."

describe('heightChecker', () => {

    it('first case', () => expect(heightChecker([1,1,4,2,1,3])).toEqual(3))

    it('second', () => expect(heightChecker([5,1,2,3,4])).toEqual(5))
    it('third', () => expect(heightChecker([1,2,3,4,5])).toEqual(0))
})
