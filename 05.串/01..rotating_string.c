/*
 * @Author: 百里
 * @Date: 2020-02-13 07:46:32
 * @LastEditTime : 2020-02-13 07:56:32
 * @LastEditors  : 百里
 * @Description: 
 * @FilePath: \leetcode\05.串\01.旋转字符串.c
 * @https://github.com/yezihack
 */
#include <stido.h>
/*
题目描述
给定一个字符串，要求把字符串前面的若干个字符移动到字符串的尾部，
如把字符串“abcdef”前面的2个字符'a'和'b'移动到字符串的尾部，
使得原字符串变成字符串“cdefab”。请写一个函数完成此功能，
要求对长度为n的字符串操作的时间复杂度为 O(n)，空间复杂度为 O(1)。
*/
void RotatingString(char* s, int n) {
    char t = s[0];
    for(int i = 1; i < n; i ++) {
        s[i-1] = s[i];
    }
    s[n-1] = t;
}
void LeftRotatingString(char* s,int n, int m) {
    while (m--)
    {
        RotatingString(s, n);
    }
}
void main() {
    string s = "abcdef";
    int n = sizeof(s);
    LeftRotatingString(s, n, 2);
}