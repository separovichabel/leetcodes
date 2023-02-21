import { kWeakestRows } from "."

describe('kWeakestRows', () => {

    it('first case', () => {
        const mat = [[1,1,0,0,0],
        [1,1,1,1,0],
        [1,0,0,0,0],
        [1,1,0,0,0],
        [1,1,1,1,1]]

        const k = 3
        expect(kWeakestRows(mat, k)).toEqual([2,0,3])
    })

    it('second', () => {
        const mat = [[1,0,0,0],
        [1,1,1,1],
        [1,0,0,0],
        [1,0,0,0]]

        const k = 2
        expect(kWeakestRows(mat, k)).toEqual([0,2])
    })
})
