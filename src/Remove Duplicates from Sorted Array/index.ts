export function removeDuplicates(nums: number[]): number {
    let cursor = 0
    let cursorInsert = 0
    let totalLength = nums.length

    let lastElement = -101 // less than the minimum value possible

    let curNum: number;
    while (cursor < nums.length) {
        curNum = nums[cursor]

        if (lastElement >= curNum) {
            cursor++
            totalLength--
            continue
        }

        lastElement = curNum
        nums[cursorInsert] = curNum

        cursorInsert++
        cursor++
    }

    return totalLength
};