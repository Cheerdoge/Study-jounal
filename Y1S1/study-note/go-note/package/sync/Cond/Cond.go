/*任务描述：
实现一个简单的任务调度器，使用 sync.Cond 来协调任务的提交和执行。要求：
创建一个 TaskScheduler 结构体，包含:
个任务队列（字符串切片）
一个 sync.Cond
一个标志位表示调度器是否已关闭
实现以下方法：
NewTaskScheduler() - 创建并初始化调度器
Submit(task string) - 提交任务到队列
Execute() - 从队列中取出并执行任务（模拟执行）
Close() - 关闭调度器，优雅停止所有 worker
在主函数中：
创建调度器
启动 3 个 worker goroutine 执行任务
提交 10 个任务
关闭调度器并等待所有任务完成
要求：
使用 sync.Cond 进行任务通知
正确处理调度器关闭，避免 goroutine 泄漏
使用 for 循环检查条件*/