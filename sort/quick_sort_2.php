<?php
/**
 * 快速排序
 * 思路:利用二分思想，找一个基准值，小于基准值放在左边，　大于基准值放在右边．以次递归实现最终排序
 * 提示：一个元素无需排序
 */
/**
 * @param $list
 * @param $low
 * @param $high
 * @return mixed
 */
function QuickSortPartition(&$list, $low , $high) {
    if(count($list) < 1) {
        return $list;
    }
    //find pivot value
    $pivot = $list[$low];
    while ($low < $high){
        //提示: 必须先从右边找起
        //如果最右边的标记high与pivot比较时, 大于的话则只需向左移动标记, 否则停止,等待交换
        while($low < $high && $list[$high] > $pivot) {
            $high --; //向右移动, 当然是--啦.就是下标移动
        }
        //程序走到这一步,证明上面的左标记大于pivot, 右标记小于pivot, 此时需要对low与higt的位置进行一次互换
        $list[$low] = $list[$high];
        //如果最左边标记low与pivot比较时, 小于的话则只需向右移动标记, 否则停止,等待交换
        while ($low < $high && $list[$low] < $pivot) {
            $low ++; //向右移动,需要++即可
        }
        $list[$high] = $list[$low];
    }
    //程序走到这一步, 证明low, high标记相遇了.需要将pivot放置到low的位置上
    $list[$low] = $pivot;
    var_dump($list, $low, $low, $pivot);
    return $low;
}
function QuickSort($list, $low, $high) {
   if($low < $high) {
       $pivot = QuickSortPartition($list, $low, $high);
       QuickSort($list, 0, $pivot - 1);
       QuickSort($list, $pivot+1, $high);
   }
}

function TestQuickSort() {
    $list = [3, 9, 7, 4, 5, 2, 9, 6, 1];
    QuickSort($list, 0, count($list) - 1);
    var_dump($list);
}
echo '<pre>';
TestQuickSort();