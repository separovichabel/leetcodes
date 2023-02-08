export function numberOfSteps(num: number): number {
    let counter = 0

    while (num !== 0) {
        if (num % 2) num--
        else num = num / 2
        counter++
    }

    return counter
};


export function recursionNumberOfSteps(num: number): number {
    return recurStep(num, 0)
};

function recurStep(num: number, counter: number): number {
    if (num === 0) return counter
    if (num % 2) return recurStep(num - 1, counter + 1)
    return recurStep(num / 2, counter + 1)
}