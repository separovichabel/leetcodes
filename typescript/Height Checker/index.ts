// export function heightChecker(heights: number[]): number {
//     const expected: number[] = [...heights].sort((a,b) => a - b)
    
//     let outOfIndex = 0
//     for (let i in heights) {
//         if (heights[i] != expected[i]) {
//             outOfIndex++
//         }
//     }
//     return outOfIndex

// };

export function heightChecker(heights: number[]): number {
    const heightCounters: {[number: string]: number} = {}

    for (let height of heights) {
        if (heightCounters[height]) {
            heightCounters[height]++
            continue
        }
        heightCounters[height] = 1
    }

    let height = 1
    let counter = 0

    for (let i = 0; i < heights.length; i++) {
        while (!heightCounters[height]) {
            height++
        }

        if (heights[i] != height) {
            counter++
        }

        heightCounters[height]--
    }
    return counter
}