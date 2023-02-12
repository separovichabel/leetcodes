// export function thirdMax(nums: number[]): number {
//     nums.sort((a,b) => b - a)

//     let counter = 0
//     for (let i = 1; i < nums.length; i++) {
//         if (counter == 3) {
//             return nums[i-1]
//         }

//         if (nums[i-1] < nums[i] ) {
//             counter++
//             continue
//         }
//     }

//     return nums[0]
// };

export function thirdMax(nums: number[]): number {
    let firstMax = {exists: false, value: -1}
    let secondMax = {exists: false, value: -1}
    let thirdMax = {exists: false, value: -1}

    for (let num of nums) {

        if (
            firstMax.exists && firstMax.value == num ||
            secondMax.exists && secondMax.value == num ||
            thirdMax.exists && thirdMax.value == num
        ) {
            continue
        }

        if (!firstMax.exists || firstMax.value < num) {
            thirdMax = secondMax
            secondMax = firstMax
            firstMax = {exists: true, value: num}
            continue
        }

        if (!secondMax.exists || secondMax.value < num) {
            thirdMax = secondMax
            secondMax = {exists: true, value: num}
            continue
        }

        if (!thirdMax.exists || thirdMax.value < num) {
            thirdMax = {exists: true, value: num}
        }

    }

    if (!thirdMax.exists) return firstMax.value

    return thirdMax.value
};

