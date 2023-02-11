export function findNumbers(nums: number[]): number {
    let evenCount = 0

    let strN
    for (let num of nums) {
        strN = num.toString()
        if (strN.length % 2) continue
        evenCount++
    }
    return evenCount
};