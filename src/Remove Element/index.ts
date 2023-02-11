export function removeElement(nums: number[], val: number): number {
    let cursor = 0
    let nextPosition = 0
    let totalLength = nums.length

    while (cursor < nums.length) {
        if (nums[cursor] === val) {
            cursor++
            totalLength--
            continue
        }
        nums[nextPosition] = nums[cursor]
        nextPosition++
        cursor++
    }
    return totalLength
};