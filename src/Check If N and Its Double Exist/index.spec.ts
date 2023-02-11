import { checkIfExist } from "."

describe('checkIfExist', () => {
    it('first case', () => expect(checkIfExist([10,2,5,3])).toEqual(true))
    it('second', () => expect(checkIfExist([3,1,7,11])).toEqual(false))
})

