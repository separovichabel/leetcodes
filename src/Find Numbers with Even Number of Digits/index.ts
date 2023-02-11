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

export function findNumbers2(nums: number[]): number {
    let evenCount = 0

    let digits = 1
    for (let num of nums) {

        for (;num >= 10; num = Math.floor(num / 10)) {
            digits++
        }
        
        let isEven = !(digits % 2)

        if (isEven) {
            evenCount++
        }
        
        digits = 1
    }
    return evenCount
};
