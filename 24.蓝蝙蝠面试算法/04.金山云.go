package _4_蓝蝙蝠面试算法

import (
	"io"
	"os"
	"strconv"
)
type FileOrderer interface {
	ReadFile(path string) (arr []int, err error) //读取文件操作
	WriteFile(arr []int) (err error) // 写文件操作
	FastOrder(arr []int)  //对数组进行排序
	BackMonitor()         //监听缓冲大小.
	MinHead(n int) int        //小顶堆
}
type FileOrder struct {
	size int   //每次读取大小 1024*2
	cacheSize int //缓冲大小
	capacity int//上限设置
	cache []int //缓冲数组.
	index int //记录每次读取位置
	fileIndex int //文件号
}
func NewFileOrder(size, cacheSize, capacity int) *FileOrderer {
	fo := &FileOrder{}
	fo.size = size
	fo.cacheSize = cacheSize
	fo.capacity = capacity
	//fo.arr = make([]int, size)
	fo.cache = make([]int, cacheSize)
	return fo
}

func Options() {
	// command
	fo := NewFileOrder()
	arr, err := fo.ReadFile(path)

}
func (f *FileOrder) ReadFile(path string) (arr []int, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*2)
	for {
		_, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		f.index ++
		//
	}
	arr = make([]int, 0)
	//byte -> string -> int
	return
}
func (f *FileOrder) WriteFile(arr []int) (err error) {
	//todo
	filename := "file-" + strconv.Itoa(f.fileIndex)
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	//todo
	return nil
}
func (f *FileOrder) FastOrder(arr []int) {
	f.FastSort(arr, 0, len(arr))
}
func (f *FileOrder) FastSort(arr []int, left, right int) {
	if left < right {
		index := f.Partition(arr, left, right)
		f.FastSort(arr, left, index-1)
		f.FastSort(arr, index+1, right)
	}
}
func (f *FileOrder) Partition(arr []int, left, right int) int {
	povit := arr[left]
	i, j := left, right
	for i < j {
		for i < j && arr[j] >= povit {
			j --
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= povit {
			i ++
		}
		arr[j] = arr[i]
	}
	arr[i] = povit
	return i
}
func (f *FileOrder) BackMonitor() {
	//todo
}
func (f *FileOrder) MinHead(n int) int {
	//todo
}
