package demo

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Three() {
	defer func() {
		fmt.Println("1111111111111111")
	}()
	defer func() {
		fmt.Println("2222222222222222")

	}()
	defer func() {
		fmt.Println("3333333333333333")

	}()

	// 获取当前工作目录
	appDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	// 打开文件，这里使用你的文件路径，记得修改成实际的路径
	file, err := os.Open(appDir + "/demo/ac.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个 Scanner 以逐行读取文件内容
	scanner := bufio.NewScanner(file)

	res := make(map[string]int)
	// 遍历每一行
	for scanner.Scan() {
		line := scanner.Text()
		// 调用函数解析每一行的数组
		for _, k := range parseArray(line) {
			fmt.Println("kkkk", k)
			res[k]++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// 将 map 转换为切片
	var keyValuePairs []KeyValue
	for key, value := range res {
		keyValuePairs = append(keyValuePairs, KeyValue{Key: key, Value: value})
	}

	// 使用自定义排序
	sort.Slice(keyValuePairs, func(i, j int) bool {
		return keyValuePairs[i].Value > keyValuePairs[j].Value
	})

	// 列出前 10 个键
	topKeys := make([]string, 0, 50)
	for i, kv := range keyValuePairs {
		if i < 50 {
			topKeys = append(topKeys, kv.Key)
		} else {
			break
		}
	}

	// 打印结果
	fmt.Println("Top 10 keys by value:")
	for _, key := range topKeys {
		fmt.Printf("%s: %d\n", key, res[key])
	}
}

// KeyValue 结构体用于存储键值对
type KeyValue struct {
	Key   string
	Value int
}

// 解析每行的数组
func parseArray(line string) []string {
	// 使用空格分割字符串得到数组元素
	elements := strings.Fields(line)

	// 检查数组长度是否符合要求
	if len(elements) != 7 {
		fmt.Println("Invalid array length in line:", line)
		return nil
	}

	// 将字符串数组转换为整数数组
	var intArray []int
	for _, element := range elements[:6] {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Println("Error converting to integer:", element)
			return nil
		}
		intArray = append(intArray, num)
	}

	return generateTriplets(intArray)
}

// 生成数组的所有三元组
func generateTriplets(arr []int) []string {
	res := make([]string, 0)

	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				triplet := []int{arr[i], arr[j], arr[k]}
				sort.Ints(triplet) // 为了考虑无序性，对三元组进行排序
				str := strconv.Itoa(triplet[0]) + "," + strconv.Itoa(triplet[1]) + "," + strconv.Itoa(triplet[2])

				res = append(res, str)
			}
		}
	}

	return res
}
