学习笔记

回溯模板
func dfs (step int) {
    if 判断边界条件 {
        相应逻辑
        return
    }

    for 尝试每种可能 {
        满足check条件
        标记
        下探到下一层 dfs(step + 1)
        恢复初始状态（回溯的时候要用到）
    }
}