<?php
/**
 * 冒泡排序
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/9
 * Time: 17:31
 */
/**
 * @param $list
 * @return mixed
 */
function BubblingSort($list) {
    $s = microtime(true);
    if(count($list) == 1) {
        return $list;
    }
    $count = $swapCount = 0;
    $len = count($list);
    for ($i = 0; $i < $len; $i ++ ) {
        for ($k = $i + 1; $k < $len; $k ++) {
            if ($list[$i] > $list[$k]) {
                $temp = $list[$i];
                $list[$i] = $list[$k];
                $list[$k] = $temp;
                $swapCount++;
            }
            $count ++;
        }
    }
    $diff = round(microtime(true) - $s, 4);
    echo sprintf('<br/>无flag版本的排序后:%s<br/>for执行%s次, 换交次数:%s, 耗时:%ss<br/>',  implode(',', $list),$count, $swapCount, $diff);
}
function BubblingSort2($list) {
    $s = microtime(true);
    if(count($list) < 2){
        return $list;
    }
    $count = $swapCount = 0;
    $len = count($list);
    for ($i = 0; $i < $len; $i ++ ) {
        $flag = false;
        for($j = 0; $j < $len - $i - 1; $j++ ) {
            if ($list[$j] > $list[$j + 1]) {
                $tmp = $list[$j];
                $list[$j] = $list[$j+1];
                $list[$j+1] = $tmp;
                $swapCount ++;
                $flag = true;
            }
            $count++;
        }
        if(!$flag) {
            break;
        }
    }
    $diff = round(microtime(true) - $s, 4);
    echo sprintf('<br/>有flag版本的排序后:%s<br/>for执行%s次, 换交次数:%s, 耗时:%ss<br/>',  implode(',', $list),$count, $swapCount, $diff);
}
function TestBubblingSort($list) {
    BubblingSort($list);
    BubblingSort2($list);
}
function createArrary($n = 10) {
    $list = [];
    for ($i = 0; $i < $n; $i ++) {
        $list[] = mt_rand(0, 999);
    }
    shuffle($list);
    return $list;
}

$cc = 10;
$i = 0;
for ($c = 1; $c <= 1000; $c *= $cc ){
    $i ++;
    $list = createArrary(10 * $c);
    echo sprintf('第%s次测试: %s条数据<br/>', $i, count($list));
    echo sprintf('排序前:%s, <br/>', implode(',', $list));
    TestBubblingSort($list);
    echo '<hr/>';
}