/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    if (strs.length === 0) {
        return "";
    }

    let minLength = Number.MAX_SAFE_INTEGER;
    for (let i = 0; i < strs.length; i++) {
        if (strs[i].length < minLength) {
            minLength = strs[i].length;
        }
    }

    let prefix = ""

    for (let j = 0; j < minLength; j++) {
        let current = strs[0][j];
        for (let i = 1; i < strs.length; i++) {
            if (current !== strs[i][j]) {
                return prefix;
            }
        }
        prefix += current;
    }

    return prefix;
};

console.log(longestCommonPrefix(["dog","racecar","car"]));

// check if current pos is less than the str length
// if all str[pos] are same append to prefix
// once diffent stop and return