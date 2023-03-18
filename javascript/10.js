/**
 * @param {string} s
 * @param {string} p
 * @return {boolean}
 */
var isMatch = function(s, p) {
  let j = 0;
  for (let i = 0; i < p.length; i++) {
      if (i !== p.length-1 && p[i+1] === '*') {
          if (p[i] === '.') {
            return true;
          }
          while (p[i] === s[j]) {
            console.log(j);
            j++;
          }
          i++;
      } else {
        if (p[i] === '.') {
          j++;
        } else if (p[i] === s[j]) {
          j++;
        } else {
          return false;
        }
      }
  }
  if (j === s.length-1) {
    return true;
  }
  return false;
};

console.log(isMatch("aa", "a*"))