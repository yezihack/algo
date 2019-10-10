<?php
$swapCount = 0;
$forCount = 0;
function QuickSort(&$list, $low, $high){
    global $swapCount, $forCount;
    if($low < $high) {
        $left = $low;
        $right = $high;
        $mid = $list[($low + $high) / 2];
        while($left <= $right) {
            while ($list[$left] < $mid) {
                $left ++;
            }
            while ($list[$right] > $mid) {
                $right --;
            }
            if($left <= $right) {
                $tmp = $list[$left];
                $list[$left] = $list[$right];
                $list[$right] = $tmp;
                $left ++;
                $right--;
                $swapCount++;
            }
            $forCount++;
        }
        if($low < $right) {
            QuickSort($list, $low, $right);
        }
        if($high > $left) {
            QuickSort($list, $left, $high);
        }
    }
}
$swapCount2 = 0;
$forCount2 = 0;
function bubblingSort(&$list) {
    global $swapCount2, $forCount2;
    for ($i = 0; $i < count($list); $i++) {
        $flag = false;
        for ($j = 0; $j < count($list) - 1 - $i; $j ++) {
            if($list[$j] > $list[$j+1]) {
                $tmp = $list[$j + 1];
                $list[$j + 1] = $list[$j];
                $list[$j] = $tmp;
                $flag = true;
                $swapCount2 ++;
            }
            $forCount2++;
        }
        if(!$flag) break;
    }
}

$list = [3, 5, 2, 8, 7, 4, 9];
for($i = 0; $i <100; $i++) {
    array_push($list, mt_rand(0, 10000));
}
shuffle($list);
$list2 = $list;
QuickSort($list, 0, count($list));
bubblingSort($list2);
echo sprintf('快排:共:%s数据排序后:%s, forCount:%s, swapCount:%s<br/>',count($list), implode(',', $list), $forCount, $swapCount);
echo sprintf('冒泡:共:%s数据排序后:%s, forCount:%s, swapCount:%s<br/>',count($list2), implode(',', $list2), $forCount2, $swapCount2);