package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// =============================================
// 一、io.Reader / io.Writer —— Go 数据流的核心接口
// =============================================

// ReadWriterDemo 演示 io.Reader 和 io.Writer 的基本用法
// 核心 API：
//   - io.Reader 接口：Read(p []byte) (n int, err error)
//   - io.Writer 接口：Write(p []byte) (n int, err error)
//   - io.Copy(dst Writer, src Reader)：将 src 的数据拷贝到 dst
//   - io.ReadAll(r Reader)：读取 Reader 全部内容
func ReadWriterDemo() {
	fmt.Println("===== io.Reader / io.Writer 示例 =====")

	// 1. strings.NewReader —— 将字符串包装为 io.Reader
	reader := strings.NewReader("Hello, Go 数据流!")

	// 2. io.ReadAll —— 一次性读取 Reader 中的全部数据
	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ReadAll 读取结果:", string(data))

	// 3. io.Copy —— 将 Reader 中的数据流式拷贝到 Writer（此处拷贝到标准输出）
	reader.Reset("io.Copy 流式输出演示\n")
	written, err := io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("io.Copy 共写入 %d 字节\n", written)
}

// =============================================
// 二、bufio —— 带缓冲的读写，提升 IO 性能
// =============================================

// BufIODemo 演示 bufio 包的常用操作
// 核心 API：
//   - bufio.NewReader(rd io.Reader)：创建带缓冲的 Reader
//   - bufio.NewScanner(r io.Reader)：创建按行扫描器
//   - bufio.NewWriter(w io.Writer)：创建带缓冲的 Writer
func BufIODemo() {
	fmt.Println("\n===== bufio 缓冲读写示例 =====")

	// ---------- 按行读取 ----------
	input := "第一行\n第二行\n第三行\n"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// scanner.Scan() 每次读取一行，返回 true 表示还有数据
	lineNo := 1
	for scanner.Scan() {
		//time.Sleep(1000 * time.Millisecond) // 模拟处理每行数据的时间
		// scanner.Text() 返回当前行的字符串（不含换行符）
		fmt.Printf("  第 %d 行: %s\n", lineNo, scanner.Text())
		lineNo++
	}
	// 检查扫描过程中是否出错
	if err := scanner.Err(); err != nil {
		log.Fatal("扫描出错:", err)
	}

	// ---------- 按单词读取 ----------
	wordInput := "Go 语言 数据流 编程"
	wordScanner := bufio.NewScanner(strings.NewReader(wordInput))
	// 设置分割函数为按单词分割（默认是按行）
	wordScanner.Split(bufio.ScanWords)
	fmt.Print("  按单词扫描: ")
	for wordScanner.Scan() {
		fmt.Printf("[%s] ", wordScanner.Text())
	}
	fmt.Println()

	// ---------- 带缓冲的 Writer ----------
	var buf bytes.Buffer
	// bufio.NewWriter 包装后写入会先存入缓冲区，减少系统调用次数
	writer := bufio.NewWriter(&buf)
	writer.WriteString("缓冲写入的内容\n")
	writer.WriteString("再写一行\n")
	// Flush() 将缓冲区中的数据真正写入底层 Writer —— 非常重要，不调用则数据丢失
	writer.Flush()
	fmt.Println("  缓冲写入结果:", buf.String())
}

// =============================================
// 三、bytes.Buffer —— 内存中的读写缓冲区
// =============================================

// BytesBufferDemo 演示 bytes.Buffer 的流式读写
// 核心 API：
//   - bytes.NewBufferString(s string)：从字符串创建 Buffer
//   - buf.Write(p []byte) / buf.WriteString(s string)：写入数据
//   - buf.Read(p []byte)：读取数据
//   - buf.String()：获取 Buffer 中全部内容的字符串
func BytesBufferDemo() {
	fmt.Println("\n===== bytes.Buffer 示例 =====")

	// 创建一个 Buffer 并写入数据
	var buf bytes.Buffer
	buf.WriteString("Hello ")
	buf.Write([]byte("Buffer!"))
	fmt.Println("  Buffer 内容:", buf.String())

	// Buffer 同时实现了 io.Reader 和 io.Writer，可以在流式管道中使用
	// 从 Buffer 读取固定长度
	p := make([]byte, 5)
	n, _ := buf.Read(p)
	fmt.Printf("  读取 %d 字节: %s\n", n, string(p[:n]))
	fmt.Println("  剩余内容:", buf.String())
}

// =============================================
// 四、io.Pipe —— 同步的内存管道（生产者-消费者模式）
// =============================================

// PipeDemo 演示 io.Pipe 在协程间传递数据流
// 核心 API：
//   - io.Pipe() 返回 (*PipeReader, *PipeWriter)
//   - 写端写入的数据会阻塞直到读端读取，适合协程间流式传输
func PipeDemo() {
	fmt.Println("\n===== io.Pipe 管道示例 =====")

	// 创建管道：pr 是读端，pw 是写端
	pr, pw := io.Pipe()

	// 启动一个协程作为生产者，往管道写数据
	go func() {
		defer pw.Close() // 写完后关闭写端，读端会收到 io.EOF
		for i := 1; i <= 3; i++ {
			fmt.Fprintf(pw, "管道消息 %d\n", i)
		}
	}()

	// 主协程作为消费者，从管道读数据
	scanner := bufio.NewScanner(pr)
	for scanner.Scan() {
		fmt.Println("  收到:", scanner.Text())
	}
}

// =============================================
// 五、io.TeeReader —— 读取的同时复制数据流
// =============================================

// TeeReaderDemo 演示 io.TeeReader（类似 Unix tee 命令）
// 核心 API：
//   - io.TeeReader(r Reader, w Writer)：从 r 读取数据时，会同时写入 w
//   - 常用场景：读取数据的同时记录日志、计算校验和等
func TeeReaderDemo() {
	fmt.Println("\n===== io.TeeReader 示例 =====")

	original := strings.NewReader("这是需要处理的数据流内容")

	// 创建一个 Buffer 用于保存副本
	var backup bytes.Buffer

	// TeeReader：从 original 读取时，数据会同时写入 backup
	tee := io.TeeReader(original, &backup)

	// 读取 tee（数据会同时流向 backup）
	processed, _ := io.ReadAll(tee)
	fmt.Println("  处理的数据:", string(processed))
	fmt.Println("  备份的数据:", backup.String())
}

// =============================================
// 六、io.MultiReader / io.MultiWriter —— 多路合并
// =============================================

// MultiReaderWriterDemo 演示数据流的合并与分发
// 核心 API：
//   - io.MultiReader(readers ...Reader)：将多个 Reader 串联为一个
//   - io.MultiWriter(writers ...Writer)：写入时同时写入多个 Writer
func MultiReaderWriterDemo() {
	fmt.Println("\n===== io.MultiReader / io.MultiWriter 示例 =====")

	// ---------- MultiReader：串联多个数据源 ----------
	r1 := strings.NewReader("数据源1 -> ")
	r2 := strings.NewReader("数据源2 -> ")
	r3 := strings.NewReader("数据源3")
	// 读取 multiReader 时会依次读取 r1、r2、r3
	multiReader := io.MultiReader(r1, r2, r3)
	combined, _ := io.ReadAll(multiReader)
	fmt.Println("  合并读取:", string(combined))

	// ---------- MultiWriter：同时写入多个目标 ----------
	var buf1, buf2 bytes.Buffer
	// 写入 multiWriter 的数据会同时写入 buf1 和 buf2
	multiWriter := io.MultiWriter(&buf1, &buf2)
	multiWriter.Write([]byte("广播数据"))
	fmt.Println("  buf1 收到:", buf1.String())
	fmt.Println("  buf2 收到:", buf2.String())
}

// =============================================
// 七、JSON 流式编解码 —— 处理大量 JSON 数据
// =============================================

// JSONStreamDemo 演示 JSON 的流式编码和解码
// 核心 API：
//   - json.NewEncoder(w Writer)：创建 JSON 编码器，将对象编码后写入 Writer
//   - json.NewDecoder(r Reader)：创建 JSON 解码器，从 Reader 逐条解码
//   - 适用场景：HTTP 响应、日志处理、大文件逐条解析
func JSONStreamDemo() {
	fmt.Println("\n===== JSON 流式编解码示例 =====")

	// 定义数据结构
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	// ---------- 流式编码：逐条写入 JSON ----------
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf) // 编码器绑定到 Buffer

	users := []User{
		{Name: "张三", Age: 28},
		{Name: "李四", Age: 32},
		{Name: "王五", Age: 25},
	}
	for _, u := range users {
		// Encode 会将对象编码为 JSON 并写入底层 Writer（每条一行）
		encoder.Encode(u)
	}
	fmt.Println("  编码结果:")
	fmt.Print("  ", buf.String())

	// ---------- 流式解码：逐条读取 JSON ----------
	decoder := json.NewDecoder(&buf) // 解码器绑定到 Buffer
	fmt.Println("  解码结果:")
	for decoder.More() { // More() 检查是否还有更多 JSON 值
		var u User
		if err := decoder.Decode(&u); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("    姓名: %s, 年龄: %d\n", u.Name, u.Age)
	}
}

// =============================================
// 八、CSV 流式读写 —— 处理表格数据
// =============================================

// CSVStreamDemo 演示 CSV 的流式读写
// 核心 API：
//   - csv.NewWriter(w Writer)：创建 CSV 写入器
//   - csv.NewReader(r Reader)：创建 CSV 读取器
//   - reader.Read()：逐行读取
//   - reader.ReadAll()：一次读取全部
func CSVStreamDemo() {
	fmt.Println("\n===== CSV 流式读写示例 =====")

	// ---------- 写入 CSV ----------
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// 写入表头
	writer.Write([]string{"姓名", "城市", "职业"})
	// 写入数据行
	writer.Write([]string{"张三", "北京", "工程师"})
	writer.Write([]string{"李四", "上海", "设计师"})
	writer.Flush() // 刷新缓冲区

	fmt.Println("  CSV 内容:")
	fmt.Print("  ", buf.String())

	// ---------- 逐行读取 CSV ----------
	reader := csv.NewReader(&buf)
	fmt.Println("  逐行解析:")
	for {
		// Read() 每次读取一行，返回 []string
		record, err := reader.Read()
		if err == io.EOF {
			break // 读到末尾退出
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("    %v\n", record)
	}
}

// =============================================
// 九、文件流式处理 —— 大文件逐行读取
// =============================================

// FileStreamDemo 演示文件的流式读写（大文件友好）
// 核心 API：
//   - os.Create(name string)：创建文件，返回 *os.File（实现了 io.ReadWriteCloser）
//   - os.Open(name string)：只读打开文件
//   - bufio.NewScanner(r)：逐行扫描文件
//   - defer file.Close()：确保文件关闭，防止资源泄漏
func FileStreamDemo() {
	fmt.Println("\n===== 文件流式处理示例 =====")

	filename := "dataStream_test_output.txt"

	// ---------- 写入文件 ----------
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal("创建文件失败:", err)
	}

	// 使用带缓冲的 Writer 提升写入性能
	writer := bufio.NewWriter(file)
	for i := 1; i <= 5; i++ {
		fmt.Fprintf(writer, "这是第 %d 行数据\n", i)
	}
	writer.Flush() // 刷新缓冲写入磁盘
	file.Close()

	// ---------- 逐行读取文件 ----------
	file, err = os.Open(filename)
	if err != nil {
		log.Fatal("打开文件失败:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("  文件内容:")
	for scanner.Scan() {
		fmt.Println("   ", scanner.Text())
	}

	// 清理测试文件
	os.Remove(filename)
}

// =============================================
// 十、io.LimitReader —— 限制读取量（防止内存溢出）
// =============================================

// LimitReaderDemo 演示限制数据流的读取量
// 核心 API：
//   - io.LimitReader(r Reader, n int64)：最多从 r 读取 n 字节
//   - 常用场景：限制 HTTP 请求体大小、防止恶意大数据攻击
func LimitReaderDemo() {
	fmt.Println("\n===== io.LimitReader 示例 =====")

	// 原始数据较长
	original := strings.NewReader("这是一段很长的数据流内容，我们只需要前面一部分")

	// LimitReader 限制最多读取 30 字节（中文 UTF-8 每字符 3 字节）
	limited := io.LimitReader(original, 30)
	data, _ := io.ReadAll(limited)
	fmt.Printf("  限制读取 30 字节: %s\n", string(data))
}

// =============================================
// RunAll 运行所有示例
// =============================================

func RunAll() {
	ReadWriterDemo()
	BufIODemo()
	BytesBufferDemo()
	PipeDemo()
	TeeReaderDemo()
	MultiReaderWriterDemo()
	JSONStreamDemo()
	CSVStreamDemo()
	FileStreamDemo()
	LimitReaderDemo()
}

func main() {
	// ReadWriterDemo()
	BufIODemo()
}
