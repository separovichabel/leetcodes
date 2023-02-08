export function canConstruct(ransomNote: string, magazine: string): boolean {
    const map: {[string: string]: number} = {}

    magazine.split('').forEach(char => {
        if (map[char]) {
            map[char]++
        } else {
            map[char] = 1
        }
    })

    for (let char of ransomNote.split('')) {
        if (map[char] && map[char] > 0) {
            map[char]--
        } else {
            return false
        }
    }
    
    return true
};

canConstruct('aa', 'aab')