package main

import (
	"flag"
	"fmt"
	"os"
	"bufio"
	"io"
	"strconv"
	"time"
    "./algorithms/qsort"
    "./algorithms/bubblesort"
)

var (
    infile = flag.String("i","infile","File contains values for sorting")
    outfile = flag.String("o","outfile","File to resive sorted values")
    algorithm = flag.String("a","qsort","Sort algorithm")
)

func readValues(infile string) (values []int, err error) {
    // 打开文件
    file, err := os.Open(infile)
    if err != nil{
        fmt.Println("Failed to open the input file ", infile)
        return
    }
    defer file.Close()

    // 创建读取流
    br := bufio.NewReader(file)
    
    // 创建一个 int 类型的切片
    values = make([]int, 0)

    // 循环读取数据
    for{
        // 读取一行的数据
        // 如果当前行的数据超过 buffer 则 isPrefix 的值为 true
        line, isPrefix, err1 := br.ReadLine()
        if err1 != nil{
            if err1 != io.EOF{
                err = err1
            }
            break
        }
        // 超过 buffer 的长度
        if isPrefix {
            fmt.Println("A too long line, seems unexpected.");
            return
        }

        // 将字符数组转为字符串
        str := string(line)

        // 将 string 类型转换为 int 类型
        value, err1 := strconv.Atoi(str)
        if err1 != nil{
            err = err1
            return
        }

        // 将数据加入到切片中
        values = append(values, value)
    }

    return
}

func writeValues(values []int, outfile string) error {
    // 创建文件并打开
    file, err := os.Create(outfile)
    if err != nil{
        fmt.Println("Failed to create the output file ", err)
        return err
    }
    defer file.Close()

    // 循环读取切片中的数据并将其写入到文件中
    for _, value := range values{
        // 将 int 类型转换为 string 类型
        str := strconv.Itoa(value)
        // 写入到文件中
        file.WriteString(str + "\n")
    }

    return nil
}

func main()  {
    flag.Parse()

    // 打印参数
    // if infile != nil{
    //     fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ",
    //         *algorithm)
    // }

    // 读取文件
    values, err := readValues(*infile)

    if err == nil{
        t1 := time.Now()
        
        switch *algorithm {
            
            case "qsort":
                qsort.QuikSort(values)
            
            case "bubblesort":
                bubblesort.BubbleSort(values)
            
            default:
                fmt.Println("Sorting algrithm", *algorithm, 
                    "is either unknown or unsupported.")
        }
        
        t2 := time.Now()

        fmt.Println("The sorting process costs ", t2.Sub(t1), " to complete.")

        writeValues(values, *outfile)

    } else {
        
        fmt.Println(err)
        
    }   
}
