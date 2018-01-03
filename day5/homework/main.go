package main

import (
	"fmt"
	"os"
)

//定义书籍类
type Book struct{
	Name string          `json:"name"`
	Count int   		 `json:"count"`
	Author string		 `json:"author"`
	Publish_Date string  `json:"pub_date"`
}

//定义学生类
type Student struct{
	Name string          `json:"name"`
	Class string         `json:"class"`
	ID string            `json:"id"`
	Gender string        `json:"gender"`
	Books []*Book        `json:"books"`
}

//全局变量，图书馆，存放书籍的切片
var library []*Book
//全局变量，学校，存放学生的切片 
var school []*Student

//书籍的构造方法
func NewBook(name,author,pub_date string,count int) *Book{
	return &Book{
		Name:name,
		Author:author,
		Publish_Date:pub_date,
		Count:count,
	}
} 

//学生的构造方法
func NewStudent(id,name,class,gender string,books []*Book) *Student{
	return &Student{
		ID:id,
		Name:name,
		Gender:gender,
		Books:books,
	}
}

//录入书籍的函数
func BookInput(){
	var (
		name string
		author string
		count int
		pub_date string
	)
	// library = make([]*Book,1000)  
	fmt.Println("输入书名：")
	fmt.Scanln(&name)
	fmt.Println("输入作者：")
	fmt.Scanln(&author)
	fmt.Println("输入出版时间：")
	fmt.Scanln(&pub_date)
	fmt.Println("输入数量：")
	fmt.Scanln(&count)
	// fmt.Println(name,author,pub_date,count)
	library = append(library,NewBook(name,author,pub_date,count))
}

// 书名查找
func SelectBookName() (int,*Book){
	var book_name string
	fmt.Println("请输入书名：")
	fmt.Scanln(&book_name)
	for index,book := range library{
		if book_name == book.Name{
			//找到返回下标和书籍对象
			return index,book
		}
	}
	return -1,nil //没有找到返回-1,nil
}

// 查找所有书
func SelectBookAll(){
	for _,book := range library{
		fmt.Printf("书名:%s 作者:%s 出版时间:%s 数量:%d\n",book.Name,book.Author,book.Publish_Date,book.Count)
	}
}

// 书籍查询
func SelectBook(){
	// 输出所有书籍信息
	if library != nil{
		var sel int
		fmt.Println("选择1显示所有书籍信息，选择2条件查询")
		fmt.Scanln(&sel)
		switch sel{
			case 1:
				SelectBookAll()
			case 2:
				_,book := SelectBookName()
				if book != nil{
					fmt.Printf("书名:%s 作者:%s 出版时间:%s 数量:%d\n",book.Name,book.Author,book.Publish_Date,book.Count)
				}else{
					fmt.Println("您查询的书籍不存在！")
				}
			default:
				fmt.Println("您的输入有误！")	
		}
	}else{
		fmt.Println("图书馆里没有书籍，请录入！")
	}
}

// 学生信息录入
func StudentInput(){
	var (
	name string
	class string
	id string
	gender string
	books []*Book
	)
	// library = make([]*Book,1000)  
	fmt.Println("输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("输入性别：")
	fmt.Scanln(&gender)
	fmt.Println("输入班级：")
	fmt.Scanln(&class)
	fmt.Println("输入身份证：")
	fmt.Scanln(&id)
	// fmt.Println(name,author,pub_date,count)
	school = append(school,NewStudent(id,name,class,gender,books))
}

//按身份证查找学生
func SelectStudentId() *Student{
	var id string
	fmt.Println("请输入学生ID：")
	fmt.Scanln(&id)
	for _,stu := range school{
		if id == stu.ID {
			// fmt.Println(stu.ID,stu.Name,stu.Class,stu.Gender,stu.Books)
			return stu
		}
	}
	return nil
}

//借书管理
func BookManage(){
	stu := SelectStudentId()
	if stu != nil{
		index,book := SelectBookName()
		if book != nil && index != -1 {
			stu.Books = append(stu.Books,book)
			fmt.Println(stu)
		}else{
			fmt.Println("书籍不存在！")
		}
	}
}

// 书籍状态
func BookStatus(){
	index,book := SelectBookName()
	fmt.Printf("%d,%v",index,book)
}

// 我的借阅
func MyBorrow(){
	stu := SelectStudentId()
	for _,book := range stu.Books{
		fmt.Printf("%s",book.Name)
	}
}

func menu(){
	for {
		fmt.Println(`
欢迎使用图书管理系统
	1.书籍录入
	2.书籍查询
	3.学生管理
	4.借书管理
	5.书籍状态
	6.我的借阅
	7.退出系统
请选择输入编号：
			`)
		var num int
		fmt.Scanln(&num)
		switch num{
		case 1:
			fmt.Println("书籍录入")
			BookInput()
		case 2:
			fmt.Println("书籍查询")
			SelectBook()
		case 3:
			fmt.Println("学生管理")
			StudentInput()
		case 4:
			fmt.Println("借书管理")
			BookManage()
		case 5:
			fmt.Println("书籍状态")
			BookStatus()
		case 6:
			fmt.Println("我的借阅")
			MyBorrow()
		case 7:
			fmt.Println(library)
			fmt.Println(school)
			os.Exit(0)
		default:
			fmt.Println("输入无效！")
		}
	}
}

func main(){
	menu()
	// fmt.Println(library)
 }