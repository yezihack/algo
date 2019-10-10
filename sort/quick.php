<?php
/**
 * quick sort
 * User: Administrator
 * Date: 2019/10/10
 * Time: 15:43
 */

function quick(&$list, $low, $high) {
    if($low < $high) {
        $pivot = $list[($low+ $high) / 2];
        $start = $low;
        $end = $high;
        while($low <= $high) {
            while ($list[$low] < $pivot) {
                $low ++;
            }
            while ($list[$high] > $pivot) {
                $high --;
            }
            if($low <= $high) {
                $tmp = $list[$low];
                $list[$low] = $list[$high];
                $list[$high] = $tmp;
                $low ++;
                $high--;
            }
        }
        if($start < $high ) {
            quick($list, $start, $high);
        }
        if($end > $low ) {
            quick($list, $low, $end);
        }
    }
}

function quickSort(&$list) {
    if(count($list) <= 1 ){
        return;
    }
    $pivot = $list[0];
    $low = 0;
    $high = count($list) - 1;
    while ($low < $high) {
        while ($low < $high && $list[$high] > $pivot) {
            $high --;
        }
        $list[$low] = $list[$high];
        while ($low < $high &&  $list[$low] < $pivot) {
            $low ++;
        }
        $list[$high] = $list[$low];
    }
    $list[$low] = $pivot;

    echo sprintf('pivot: %s,low:%s, 左边:%s, 右边:%s, list: %s' . PHP_EOL,$pivot, $low,
        implode(',', array_slice($list, 0, $low )),
        implode(',', array_slice($list, $low + 1)),
        implode(',', $list)
        );
    $left = array_slice($list, 0, $low);
    $right = array_slice($list, $low + 1);
    if(count($left) > 0) {
        quickSort(array_slice($list, 0, $low));
    }
    if(count($right) > 0){
        quickSort( array_slice($list, $low + 1));
    }
}
$list = [2, 3, 1, 4, 5, 8, 6, 7];
for($i = 0; $i < 0; $i ++ ){
    array_push($list, mt_rand(0, 99999));
}
//shuffle($list);
echo sprintf("排序后:%s\n", implode(',', $list));
//quick($list, 0, count($list) -1);
quickSort($list);
echo sprintf("排序后:%s\n", implode(',', $list));