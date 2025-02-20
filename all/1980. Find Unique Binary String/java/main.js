/**
 * @param {string[]} nums
 * @return {string}
 */
var findDifferentBinaryString = function(nums) {
    let n = nums.length;
    let set = new Set(nums);

    // Generate a binary string that doesn't exist in the set
    for (let i = 0; i < (1 << n); i++) {
        let binaryStr = i.toString(2).padStart(n, '0');
        if (!set.has(binaryStr)) return binaryStr;
    }
};