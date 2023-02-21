import { validMountainArray } from "."

describe('validMountainArray', () => {
    it('first case', () => expect(validMountainArray([2,1])).toEqual(false))
    it('second', () => expect(validMountainArray( [3,5,5])).toEqual(false))

    it('Third', () => expect(validMountainArray([0,3,2,1])).toEqual(true))
    it('fourth', () => expect(validMountainArray([9,8,7,6,5,4,3,2,1,0])).toEqual(false))
    it('fifth', () => expect(validMountainArray([1,2,3,4,5,6,7,8,9])).toEqual(false))
    
})

