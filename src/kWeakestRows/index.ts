export function kWeakestRows(mat: number[][], k: number): number[] {
    return mat
        .map((list, index) => ({soldires: list.reduce((a, b) => a + b), index}))
        .sort((a, b) => a.soldires - b.soldires)
        .slice(0, k)
        .map(v => v.index)
};