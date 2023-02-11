export function twoSum(nums: number[], target: number): number[] {
    const indexes: {[number: number]: number[]} = {}

    const orderdPossibles = nums.filter((v, i) => {
            if (indexes[v]) indexes[v].push(i)
            else indexes[v] = [i]
            return true
        })
        .sort((a, b) => a - b);
    
    
    let lowestIndex = 0
    let greatestIndex = orderdPossibles.length - 1


    while (orderdPossibles[lowestIndex] + orderdPossibles[greatestIndex] !== target) {
        if (orderdPossibles[lowestIndex] + orderdPossibles[greatestIndex] < target) lowestIndex++
        else greatestIndex--
    }

    const lowestValue = orderdPossibles[lowestIndex]
    const greatestValue = orderdPossibles[greatestIndex]

    return [indexes[lowestValue].pop() || 0, indexes[greatestValue].pop() || 0]
};

export function twoSum2(nums: number[], target: number): number[] {
    const numsIndexes: {[number: number]: number[]} = {}

    nums.forEach((v, i) => numsIndexes[v] ? numsIndexes[v].push(i) : numsIndexes[v] = [i])    
    
    let lowestIndex = 0
    
    let firstNum: number, secondNum: number;

    while (true) {
        firstNum = nums[lowestIndex]
        secondNum = target - firstNum

        if (!!numsIndexes[firstNum] &&
            !!numsIndexes[secondNum] &&
            firstNum !== secondNum || numsIndexes[secondNum]?.length > 1
        ) {
            return [numsIndexes[firstNum].pop() || 0, numsIndexes[secondNum].pop() || 0]
        } else {
            delete numsIndexes[firstNum]
            delete numsIndexes[secondNum]
        }
        lowestIndex++
    }
};

export function twoSum3(nums: number[], target: number): number[] {
    const seenNumbers: {[number: number]: number} = {}

    let curNumber: number, curComplement: number

    for (let index = 0; index < nums.length; index++) {
        curNumber = nums[index]
        curComplement = target - curNumber
        // console.log(`curNumber: ${curNumber} curComplement: ${curComplement} - seenNumber: ${seenNumbers[curComplement]} `)
        if (seenNumbers[curComplement] !== undefined)
            return [seenNumbers[curComplement], index]
        seenNumbers[curNumber] = index
    }  
    return [];
};
