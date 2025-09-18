package leetcode

import "testing"

func TestTaskManager_BasicFlow(t *testing.T) {
	tm := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	tm.Add(4, 104, 5)
	tm.Edit(102, 8)
	if got := tm.ExecTop(); got != 3 {
		t.Errorf("ExecTop() = %v, want %v", got, 3)
	}
	tm.Rmv(101)
	tm.Add(5, 105, 15)
	if got := tm.ExecTop(); got != 5 {
		t.Errorf("ExecTop() = %v, want %v", got, 5)
	}
}

func TestTaskManager_EmptyExecTop(t *testing.T) {
	tm := Constructor([][]int{})
	if got := tm.ExecTop(); got != -1 {
		t.Errorf("ExecTop() = %v, want %v", got, -1)
	}
}

func TestTaskManager_EditAndRemove(t *testing.T) {
	tm := Constructor([][]int{{1, 101, 10}})
	tm.Edit(101, 20)
	if got := tm.ExecTop(); got != 1 {
		t.Errorf("ExecTop() = %v, want %v", got, 1)
	}
	tm.Add(2, 102, 15)
	tm.Rmv(102)
	if got := tm.ExecTop(); got != -1 {
		t.Errorf("ExecTop() = %v, want %v", got, -1)
	}
}

func TestTaskManager_TieBreaker(t *testing.T) {
	tm := Constructor([][]int{{1, 101, 10}, {2, 102, 10}})
	if got := tm.ExecTop(); got != 2 { // taskId 102 > 101
		t.Errorf("ExecTop() = %v, want %v", got, 2)
	}
	if got := tm.ExecTop(); got != 1 {
		t.Errorf("ExecTop() = %v, want %v", got, 1)
	}
}

func TestTaskManager_ReAddSameTaskId(t *testing.T) {
	tm := Constructor([][]int{{1, 101, 8}, {2, 102, 20}, {3, 103, 5}})
	tm.Rmv(101)
	tm.Add(50, 101, 8)
	if got := tm.ExecTop(); got != 2 {
		t.Errorf("ExecTop() = %v, want %v", got, 2)
	}
	tm.Edit(101, 9)
	if got := tm.ExecTop(); got != 50 {
		t.Errorf("ExecTop() = %v, want %v", got, 50)
	}
}
