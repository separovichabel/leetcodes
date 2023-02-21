export function moveZeroes(nums: number[]): void {
    let insert = 0

    for (let read = 0; read < nums.length; read++) {
        if (nums[read] !== 0) {
            nums[insert] = nums[read]
            insert++
        }
    }

    while (insert < nums.length) {
        nums[insert] = 0
        insert++
    }
};