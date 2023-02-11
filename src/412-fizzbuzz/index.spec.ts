import { fizzBuzz } from "."

describe('fizzBuzz', () => {

    it('first case', () => {
        expect(fizzBuzz(1)).toEqual(['1'])
    })

    it('second', () => {
        expect(fizzBuzz(3)).toEqual(['1', '2', 'Fizz'])
    })

    it('Third', () => {
        expect(fizzBuzz(5)).toEqual(['1', '2', 'Fizz', '4', 'Buzz'])
    })

    
    it('Fourth', () => {
        expect(fizzBuzz(15)).toEqual(["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"])
    })

})
