export function removeDuplicates(nums: number[]): number {
    let totalLength = nums.length
    let insertIndex = 1
    for (let readIndex = 1; readIndex < nums.length; readIndex++) {
        if (nums[readIndex - 1] != nums[readIndex]) {
            nums[insertIndex] = nums[readIndex]
            insertIndex++
            continue
        }

        totalLength--
    }
    return totalLength
};