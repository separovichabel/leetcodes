import { addItemShiftOthers } from "."

describe('addItemShiftOthers', () => {
    it('first case', () => expect(addItemShiftOthers([-4,-1,0,3,10], NaN, 0)).toEqual([NaN,-4,-1,0,3,10]))
    it('second', () => expect(addItemShiftOthers([-4,-1,0,3,10], NaN, 2)).toEqual([-4,-1,NaN,0,3,10]))

    it('Third', () => expect(addItemShiftOthers([-4,-1,0,3,10], NaN, 5)).toEqual([-4,-1,0,3,10,NaN,]))
})

