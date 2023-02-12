function sortArrayByParity(nums: number[]): number[] {
    const ordered: number[] = new Array(nums.length)
    let even = 0
    let odd = nums.length - 1

    for (let i = 0; i < nums.length; i++) {
        if (nums[i] % 2 == 0) {
            ordered[even] = nums[i]
            even++
            continue
        }
        ordered[odd] = nums[i]
        odd--
    }

    return ordered
};