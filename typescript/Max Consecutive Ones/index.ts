export function findMaxConsecutiveOnes(nums: number[]): number {
    let maxOccorrencies = 0
    let currentOccorrencies = 0

    for (let num of nums) {
        if (num === 1) {
            if (maxOccorrencies === currentOccorrencies) {
                maxOccorrencies++
            }
            currentOccorrencies++
        } else {
            currentOccorrencies = 0
        }
    }
    return maxOccorrencies
};