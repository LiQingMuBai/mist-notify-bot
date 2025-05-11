package switcher

import "sync"

// 任务状态类型
type TaskStatus string

const (
	StatusBefore TaskStatus = "before" // 任务前状态
	StatusMiddle TaskStatus = "middle" // 任务中状态
	StatusAfter  TaskStatus = "after"  // 任务后状态
)

// 用户任务结构
type UserTask struct {
	UserID string     // 用户ID
	TaskID string     // 任务ID
	Status TaskStatus // 任务状态
}

// 任务流程管理器
type TaskFlowManager struct {
	mu    sync.RWMutex                    // 读写锁，保证并发安全
	tasks map[int64]map[string]TaskStatus // 外层map key是用户ID，内层map key是任务ID
}

// 创建新的任务流程管理器
func NewTaskFlowManager() *TaskFlowManager {
	return &TaskFlowManager{
		tasks: make(map[int64]map[string]TaskStatus),
	}
}

// 设置用户任务状态
func (m *TaskFlowManager) SetTaskStatus(userID int64, taskID string, status TaskStatus) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.tasks[userID]; !ok {
		m.tasks[userID] = make(map[string]TaskStatus)
	}
	m.tasks[userID][taskID] = status
}

// 获取用户任务状态
func (m *TaskFlowManager) GetTaskStatus(userID int64, taskID string) (TaskStatus, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	userTasks, ok := m.tasks[userID]
	if !ok {
		return "", false
	}

	status, ok := userTasks[taskID]
	return status, ok
}

// 获取用户所有任务状态
func (m *TaskFlowManager) GetUserTasks(userID int64) map[string]TaskStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.tasks[userID]
}

// 推进任务到下一个状态
func (m *TaskFlowManager) AdvanceTaskStatus(userID int64, taskID string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	userTasks, ok := m.tasks[userID]
	if !ok {
		return false
	}

	currentStatus, ok := userTasks[taskID]
	if !ok {
		return false
	}

	switch currentStatus {
	case StatusBefore:
		userTasks[taskID] = StatusMiddle
	case StatusMiddle:
		userTasks[taskID] = StatusAfter
	case StatusAfter:
		// 已经是最终状态，不再推进
		return false
	}

	return true
}
