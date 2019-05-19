package main

/**
 * @DATE        2019/5/19
 * @NAME        Joshua Conero
 * @DESCRIPIT   迷宫寻物
 **/

func main() {
	exampleDir := "D:/conero/phpapps/apps/vr360/"
	nm := NewMaze(exampleDir)
	nm.Search()
	nm.Gather()
}
