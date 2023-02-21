export function validMountainArray(arr: number[]): boolean {
    if (arr.length < 3) return false

    let accending = true;

    for (let i = 0; i < arr.length - 1; i++) {
        if (accending && arr[i] < arr[i + 1]) continue
        if (accending && arr[i] > arr[i + 1] && i != 0) {
            accending = false
            continue
        }
        if (!accending && arr[i] > arr[i + 1]) continue
        return false
    }

    return !accending
};