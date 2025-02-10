var clearDigits = function (s) {
    let ans = [];
    for (let i = 0; i < s.length; i++) {
      if (s[i] >= '0' && s[i] <= '9') {
        if (ans.length > 0) {
          ans.pop();
        }
      } else {
        ans.push(s[i]);
      }
    }
    return ans.join('');
  };