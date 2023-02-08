export function fizzBuzz(n: number): string[] {
    const FBList = []
    let counter = 1

    let fizz = ''
    let buzz = ''

    while (counter <= n) {
        
        fizz = !(counter % 3) ? 'Fizz' : ''
        buzz = !(counter % 5) ? 'Buzz' : ''

        FBList.push(!!fizz || !!buzz? `${fizz}${buzz}` : `${counter}`)
        counter++
    }
    return FBList
};