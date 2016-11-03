# Sorter
## 简介
这是一个排序程序，能从指定文件中读取数据并按照指定算法排序后，写入到指定的输出文件中去。


## 用法
```
// 基本用法
USAGE: sortwe -i <in> -o <out> -a <qsort|bubblesort>

// 实际的一个例子
> sorter -i in.dat -o out.dat -a qsort
The sorting process costs 10us to complete.

// 对于不同的错误将会返回不同的信息，此处不做例举
```

## 程序结构

├─bin                                                                                                                                                                                          
├─pkg                                                                                                                                                                                          
├─sorter                                                                                                                                                                                       
│  │  sorter.go                                                                                                                                                                                
│  │                                                                                                                                                                                           
│  └─algorithms                                                                                                                                                                                
│      ├─bubblesort                                                                                                                                                                            
│      │      bubblesort.go                                                                                                                                                                    
│      │      bubblesort_test.go                                                                                                                                                               
│      │                                                                                                                                                                                       
│      └─qsort                                                                                                                                                                                 
│              qsort.go                                                                                                                                                                        
│              qsort_test.go                                                                                                                                                                   
│                                                                                                                                                                                              
└─src

## 程序工作
* 获取并解析命令行参数
* 从对应文件中读取输入数据
* 调用对应的排序函数
* 将排序的结果输出到对应的文件中
* 打印排序所花费的时间