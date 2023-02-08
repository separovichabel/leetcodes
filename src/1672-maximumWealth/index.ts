export function maximumWealth(accounts: number[][]): number {
    const a = accounts.map(v => v.reduce((a ,b) => a + b))
        .sort((a, b) => a - b)
    
    return a[a.length - 1]
};

export function MaximumWealth2(accounts: number[][]): number {
    let richestWealth = 0

    for (let personAccounts of accounts) {
        let wealth = 0

        for (let account of personAccounts) {
            wealth += account
        }

        if (wealth > richestWealth) richestWealth = wealth
    }
    return richestWealth
};