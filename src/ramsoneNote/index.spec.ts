import { canConstruct } from "."

describe('canConstruct', () => {

    it('first case', () => {
        expect(canConstruct('a', 'b')).toEqual(false)
    })

    it('second', () => {
        expect(canConstruct('aa', 'ab')).toEqual(false)
    })

    it('third case ', () => {
        expect(canConstruct('ab', 'ab')).toEqual(true)
    })

    it('fourth case ', () => {
        expect(canConstruct('aa', 'aab')).toEqual(true)
    })
})
