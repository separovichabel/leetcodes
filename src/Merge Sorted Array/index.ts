/**
 Do not return anything, modify nums1 in-place instead.
*/
export function merge(nums1: number[], n1Length: number, nums2: number[], n2Length: number): void {
    let n1Index = 0
    let n2Index = 0
    let finalIndex = 0

    const nums1Original = [...nums1]
    const totalLength = nums1.length

    while (n1Index < n1Length || n2Index < n2Length) {
        if (n1Index < n1Length && n2Index < n2Length) {
            if (nums1Original[n1Index] < nums2[n2Index]) {
                nums1[finalIndex] = nums1Original[n1Index]
                finalIndex++, n1Index++
            } else if (nums1Original[n1Index] > nums2[n2Index]) {
                nums1[finalIndex] = nums2[n2Index]
                finalIndex++, n2Index++
            } else if (nums1Original[n1Index] === nums2[n2Index]){
                nums1[finalIndex] = nums2[n2Index]
                nums1[finalIndex + 1] = nums2[n2Index]
                n2Index++, n1Index++
                finalIndex += 2
            }
        } else if (n1Index < n1Length) {
            nums1[finalIndex] = nums1Original[n1Index]
            finalIndex++, n1Index++
        } else {
            nums1[finalIndex] = nums2[n2Index]
            finalIndex++, n2Index++
        }
    }

};

export function merge2(nums1: number[], n1Length: number, nums2: number[], n2Length: number): void {
    let index1 = n1Length - 1
    let index2 = n2Length - 1
    let cursor = nums1.length - 1

    while (index2 >= 0) {
        if (index1 >= 0 && nums1[index1] >= nums2[index2]) {
            nums1[cursor] = nums1[index1]
            cursor--, index1--
            continue
        }
        nums1[cursor] = nums2[index2]
        cursor--, index2--
    }

};