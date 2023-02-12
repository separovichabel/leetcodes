import { moveZeroes } from ".";


describe('moveZeroes', () => {
    it('first moveZeroes', () => {
        const expectedResult = [1,3,12,0,0]
        const arr = [0,1,0,3,12]
        const resp = moveZeroes(arr)
        
        expect(resp).toEqual(expectedResult)
    })

    it('second moveZeroes', () => {
        const expectedResult = [0]
        const arr = [0]
        const resp = moveZeroes(arr)
        
        expect(resp).toEqual(expectedResult)
    })
    
});
