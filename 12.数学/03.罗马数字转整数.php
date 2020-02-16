<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2020/2/15
 * Time: 15:37
 */

function RomanToInt(string $s) {
    if(empty($s)) {
        return 0;
    }
    //建立映射关系
    $hash = [
        "I" => 1,
        "V" => 5,
        "X" => 10,
        "L" => 50,
        "C" => 100,
        "D" => 500,
        "M" => 1000,
    ];
    var_dump($hash);

    $sum = 0;
    for($i = 0; $i < strlen($s)-1; $i ++) {
        if($hash[$s[$i]] < $hash[$s[$i+1]]) {
            $sum -= $hash[$s[$i]];
        } else {
            $sum += $hash[$s[$i]];
        }
    }
    $sum += $hash[$s[strlen($s)-1]];
    return $sum;
}
if(RomanToInt("III") != 3) {
    echo "err";
}
if(RomanToInt("IV") != 4) {
    echo "err";
}
if(RomanToInt("IX") != 9) {
    echo "err";
}
if(RomanToInt("LVIII") != 58) {
    echo "err";
}
if(RomanToInt("MCMXCIV") != 1994) {
    echo "err";
}

