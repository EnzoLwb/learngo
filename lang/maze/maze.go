package main

import (
	"fmt"
	"os"
)

//[][]int  第一个[]表示切片类型 后面[]int表示切片数值的类型 又是一个int类型的slice
func readFile(filename string) [][]int {
	//打开文件
	os, err := os.Open(filename)
	defer os.Close()
	if err != nil {
		panic(err)
	}

	//阅读第一行内容
	var row, col int
	fmt.Fscanf(os, "%d %d", &row, &col)

	//将其他内容放入maze
	maze := make([][]int, row)
	fmt.Println(maze)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(os, "%d", &maze[i][j])

		}
	}
	return maze
}

//上下左右
var dirs = [4]point{
	{-1, 0}, //上
	{0, -1}, //左
	{1, 0},  //下
	{0, 1},  //右
}

//点结构
type point struct {
	i, j int // i 是列 j是行标!!! 不是 xy
}

//两个点相加的方法 也就是求他下一个的位置
func (p point) add(r point) point {
	return point{
		p.i + r.i,
		p.j + r.j,
	}
}

//点的值 传入一组点 返回是否越界
func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		fmt.Println("上下越界了")
		return 0, false //上下越界了
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		fmt.Println("左右越界了")
		return 0, false //左右越界了
	}
	return grid[p.i][p.j], true
}

/*start end 是点*/
func walk(maze [][]int, start point, end point) [][]int {
	//走了多少步才走到我这一格 ；所有的点
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	fmt.Println(steps)
	Q := []point{
		start,
	} //放置需要探索的序列 里面是某个点 初始化就是起点 start

	//退出条件：1.走到终点 2.队列没有点了 证明死路
	//开始走
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break //发现终点了
		}
		//开始探索
		for _, dir := range dirs {
			next := cur.add(dir)
			//什么情况下才能探索下去呢
			//1.下个点 是0 && 2.step next 也是0 && 3 next 不等于start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}
			if next == start {
				continue
			}
			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1
			Q = append(Q, next)

		}

	}
	return steps
}

func main() {
	maze := readFile("lang/maze/maze.in")
	fmt.Println(maze)
	fmt.Println(len(maze))

	//走迷宫
	steps := walk(maze,
		point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val) //三位对齐
		}
		fmt.Println()
	}
}
