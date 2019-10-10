<?php
/**
 * Created by PhpStorm.
 * User: Administrator
 * Date: 2019/10/9
 * Time: 23:12
 */

function bubbling($list) {
    echo sprintf('排序前:%s<br/>', implode(',', $list));
    for($i = 0; $i < count($list); $i ++) {
        for ($j = $i + 1; $j <count($list); $j++) {
            if($list[$i] > $list[$j]) {
                $tmp = $list[$i];
                $list[$i] = $list[$j];
                $list[$j] = $tmp;
                echo sprintf('外:%s, 内:%s, 排序中: %s<br/>', $i +1, $j +1, implode(',', $list));
            }

        }

    }
    echo sprintf('排序后:%s<br/>', implode(',', $list));
}
$list = [1, 5, 3, 6, 2];
bubbling($list);