/**
 Do not return anything, modify arr in-place instead.
 */

 export function duplicateZeros(arr: number[]): void {
    const original = [...arr]
    let originalIndex = 0
    for (let index = 0; index < original.length; index++, originalIndex++) {
        arr[index] = original[originalIndex]
        
        if (arr[index] !== 0) {
            continue
        }

        index++
        if (index >= original.length) break
        arr[index] = 0
    }
};

