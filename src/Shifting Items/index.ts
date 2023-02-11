export function addItemShiftOthers<T>(arr: T[], item: T, index: number): T[] {
    let curIndex = arr.length
    for (; curIndex > index; curIndex--) {
        arr[curIndex] = arr[curIndex-1]
    }
    arr[index] = item
    return arr
};
