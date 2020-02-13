<?php
/*
 * @Author: 百里
 * @Date: 2020-02-11 08:55:45
 * @LastEditTime : 2020-02-11 09:15:46
 * @LastEditors  : 百里
 * @Description: 二进制1的个数
 * @FilePath: \leetcode\20.剑指offer\15.二进制1的个数.php
 * @https://github.com/yezihack
 */
//计算二进制里的1有多少个.但是遇到负数则死循环.
function numberOf1InBinaryV1(int $n) {
    $count = 0; //记录统计1的次数.
    while($n) { //整除到0为止
        if($n & 1) { //n和1进行与运行. 因为与的规划是: 1&1=1, 0&1=0 
            $count++;//累加1的次数
        }
        $n >>= 1; //向右移动1位,相当于除于2, 但是当负数的话, 左边补得是1
        //所以会出现死循环.那我们换一种思路.向左移动.
    }
    return $count;
}
//计算二进制里的1有多少个.
//我们先判断$n的最低位是不是1. 然后判断第二位是不是1.
//所以需要一个标识flag, 从1开始判断, 然后左移一位, 再判断, 以此类推.
function numberOf1InBinaryV2(int $n) {
    $count = 0; //记录统计1的次数.
    $flag = 1; //先判断最低位.
    while($flag) { //整除到0为止
        if($n & $flag) { //n和flag进行与运行. 因为与的规划是: 1&1=1, 0&1=0 
            $count++;//累加1的次数
        }
        // $n <<= 1; //向左移动1位,相当于乘以2
        $flag = $flag << 1;//写法与上面一样.
    }
    return $count;
}
//把一个整数减去1，再和原来的整数做相与运算，会把该整数二进制的最右边的1变成0。
function numberOf1InBinaryV3(int $n) {
    $count = 0;
    while($n) {
        $count ++;
        $n = $n & ($n - 1);
    }
    return $count;
}
$n = 9;
var_dump(numberOf1InBinaryV1($n));
var_dump(numberOf1InBinaryV2($n));
var_dump(numberOf1InBinaryV3($n));