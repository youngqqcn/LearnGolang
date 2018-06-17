package main

import "fmt"

/**
*作者: yqq
*日期: 2018/6/17  16:01
*描述:
*/

/*
go语言中的defer就像c++中的析构函数，但是go语言中defer的对象是函数（或者对象的方法），defer能保证在函数结束最后执行该方法（函数），但是有例外：如果在定义的方法中defer定义的方法如果在panic后面，defer定义的方法就无法执行到。

panic 是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值作为参数。panic 的作用就像我们平常接触的异常。不过Go可没有try…catch，所以，panic一般会导致程序挂掉（除非recover）。所以，Go语言中的异 常，那真的是异常了。你可以试试，调用panic看看，程序立马挂掉，然后Go运行时会打印出调用栈。
但是，关键的一点是，即使函数执行的时候 panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时 候 defer 有点类似 try-catch-finally 中的 finally。
panic就是这么简单。抛出个真正意义上的异常。

panic的函数并不会立刻返回，而是先defer，再返回，如果有办法将panic捕获到，并阻止panic传递，就正常处理，如果没有没有捕获，程序直接异常终止（可以注释掉下面程序中的recover试一试）。

Go语言提供了recover内置函数，前面提到，一旦panic，逻辑就会走到defer那，那我们就在defer那等着，调用recover函 数将会捕获到当前的panic（如果有的话），被捕获到的panic就不会向上传递了，于是，世界恢复了和平。你可以干你想干的事情了。

不过要注意的是，recover之后，逻辑并不会恢复到panic那个点去，函数还是会在defer之后返回。

 */

func DoubleDefer1()  {

	//defer func() {
	//	fmt.Println(recover())
	//}()

	defer func() {
		panic("second defer's panic")
	}()

	panic("DoubleDefer's panic")

	/*
	panic: DoubleDefer's panic
		panic: second defer's panic
	 */
}


func DoubleDefer2()  {

	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("second defer's panic")
	}()

	panic("DoubleDefer's panic")

	/*
	second defer's panic

	疑惑:  既然panic 会中断程序, 那么, 为什么还会调用 延迟函数?

	解疑:
		发生panic之后, 程序并不是立即向上抛出, 而是到 defer, 然后才返回
		注意, 在panic后面的操作不会被执行, 即使recover了, 也会直接返回
	 */
}


func main() {

	//DoubleDefer1()
	DoubleDefer2()

}
