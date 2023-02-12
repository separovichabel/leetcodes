export function replaceElements(arr: number[]): number[] {
    let greatest = -1
    let curItem = -1
    for (let index = arr.length-1; index >= 0; index--) {
        curItem = arr[index]

        arr[index] = greatest

        if (curItem > greatest) {
            greatest = curItem
        }
    }

    return arr
};