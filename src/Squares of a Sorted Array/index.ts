export function sortedSquares(nums: number[]): number[] {
    const negativesSqr: number[] = []
    const sqrs: number[] = []
    let sqr

    for (let num of nums) {
        sqr = num ** 2

        if (num < 0) {
            negativesSqr.unshift(sqr)
            continue
        }
        
        while (negativesSqr.length && sqr >= negativesSqr[0]) {
            sqrs.push(negativesSqr.shift()||NaN)
        }
        sqrs.push(sqr)
    }

    if (negativesSqr.length) {
        sqrs.push(...negativesSqr)
    }

    return sqrs
};