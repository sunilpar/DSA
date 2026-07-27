const str = "A man, a plan, a canal: Panama";
const str1 = "race a car";
const str2 = " ";
const str3 = "0P";
function isPalindrome(s) {
    const str = s.toLowerCase().replace(/[^a-z0-9]/g, "");
    console.log("trimmed str", str);
    let i = 0;
    let j = str.length - 1;
    while (i < j) {
        if (str[i] !== str[j]) {
            return false;
        }
        i++;
        j--;
    }
    return true;
}
console.log(isPalindrome(str3));
