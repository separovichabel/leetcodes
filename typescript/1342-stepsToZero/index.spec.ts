import { numberOfSteps, recursionNumberOfSteps } from "."

describe('numberOfSteps', () => {
    it('first case', () => expect(numberOfSteps(14)).toEqual(6))
    it('second', () => expect(numberOfSteps(8)).toEqual(4))
    it('Third', () => expect(numberOfSteps(123)).toEqual(12))
})

describe('recursionNumberOfSteps', () => {
    it('first case', () => expect(recursionNumberOfSteps(14)).toEqual(6))
    it('second', () => expect(recursionNumberOfSteps(8)).toEqual(4))
    it('Third', () => expect(recursionNumberOfSteps(123)).toEqual(12))
})