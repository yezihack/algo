<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/9
 * Time: 12:47
 */

echo '<pre>';
$count = 0;
function swap(&$list, $left, $right) {
    global $count;
    if(count($list) < 2) {
        return $list;
    }
    $start = $left;
    $pivot = $list[$left];
    while ($left < $right) {
        while ($left < $right && $list[$right] >= $pivot) {
            $right --;
        }
        while($left < $right && $list[$left] <= $pivot) {
            $left ++;
        }
        if($left < $right) {//没有相遇
            $tmp = $list[$right];
            $list[$right] = $list[$left];
            $list[$left] = $tmp;
        }
    }
    $list[$start] = $list[$left];
    $list[$left] = $pivot;
    var_dump($list);
    $count ++;
    return $left;
}

function QuickSortTop(&$list, $left, $right) {
    if($left < $right) {
        $pivot = swap($list, $left, $right);
        QuickSortTop($list, 0, $pivot - 1);
        QuickSortTop($list, $pivot +1, $right);
    }
}

function TestSwap() {
    $list = [6, 1, 2, 7, 9, 3, 4, 5, 10, 8];
    swap($list, 0, count($list) - 1);
//    var_dump($list);
}
TestSwap();










/**
 * @param $arr
 * @param $left
 * @param $right
 */
function QuickSortSwap($arr, $left, $right) {
    if (count($arr) < 2) {
        return;
    }
    //找一个做基准值
    $pivot = $arr[$left];
    //左指针小于右指针
    while ($left < $right) {
        while ($arr[$right] > $pivot) { //判断右指针是否大于基准值, 如果大于则无需交替位置,只需移动指针即可
            $right --;
        }
        //否则是交互位置
        $arr[$left] = $arr[$right];
        while ($arr[$left] < $pivot) {
            $left ++;
        }
        //否则交互位置
        $arr[$right] = $left;
    }
    $arr[$left] = $pivot;
    return $left;
}

function quickSort2($arr) {
    if(count($arr) < 2) {
        return $arr;
    }
    $pivot = $arr[0];
    $low = $high = [];
    foreach ($arr as $value) {
        if($value == $pivot) {//不能将基准值参与计算
            continue;
        }
        if($pivot > $value) {
            $low[] = $value;
        } else {
            $high[] = $value;
        }
    }

    $low= quickSort2($low);
    $high = quickSort2($high);
    return array_merge($low, [$pivot], $high);
}
function TestQuickSort() {
    $arr = [8, 1, 6, 3, 9, 2];
//    QuickSort($arr, 0, count($arr));
    $ret = quickSort2($arr);
    var_dump($ret);
}
//TestQuickSort();

function quickSort3($array)
{
    if(!isset($array[1]))
        return $array;
    $mid = $array[0]; //获取一个用于分割的关键字，一般是首个元素
    $leftArray = array();
    $rightArray = array();

    foreach($array as $v)
    {
        if($v > $mid)
            $rightArray[] = $v;  //把比$mid大的数放到一个数组里
        if($v < $mid)
            $leftArray[] = $v;   //把比$mid小的数放到另一个数组里
    }

    $leftArray = quickSort3($leftArray); //把比较小的数组再一次进行分割
    $leftArray[] = $mid;        //把分割的元素加到小的数组后面，不能忘了它哦

    $rightArray = quickSort3($rightArray);  //把比较大的数组再一次进行分割
    return array_merge($leftArray,$rightArray);  //组合两个结果
}