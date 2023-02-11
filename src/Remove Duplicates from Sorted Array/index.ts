export function removeDuplicates(nums: number[]): number {
    let cursor = 0
    let cursorInsert = 0
    let totalLength = nums.length

    const seen: {[string: number]: boolean} = {}

    let curNum: number;
    while (cursor < nums.length) {
        curNum = nums[cursor]

        if (seen[curNum]) {
            cursor++
            totalLength--
            continue
        }

        seen[curNum] = true
        nums[cursorInsert] = curNum

        cursorInsert++
        cursor++
    }

    return totalLength
};