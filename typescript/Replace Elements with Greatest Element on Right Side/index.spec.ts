import { replaceElements } from "."

describe('replaceElements', () => {
    it('first case', () => expect(replaceElements([17,18,5,4,6,1])).toEqual([18,6,6,6,1,-1]))
    it('second', () => expect(replaceElements([400])).toEqual([-1]))
})

