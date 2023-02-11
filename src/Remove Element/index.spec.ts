import { removeElement } from ".";


describe('removeElement', () => {
    it('first case', () => {
        const expectedResult = [2,2]
        const arr = [3,2,2,3]
        const resp = removeElement(arr, 3)
        
        expect(resp).toEqual(2)
        arr.sort((a,b) => a - b)
        expectedResult.sort((a,b) => a - b)

        for (let i = 0; i < resp; i++) {
            expect(arr[i]).toEqual(expectedResult[i])
        }
    })
    
});
