package domain

import "testing"

func TestTaskWithoutDescription(t *testing.T) {
	task := NewTaskUninitialized("Задача без описания", nil, 1)
	if err := task.Validate(); err != nil {
		t.Fatalf("validate task without description: %v", err)
	}
	completed := true
	if err := task.ApplyPatch(NewTaskPatch(Nullable[string]{}, Nullable[string]{}, Nullable[bool]{Set: true, Value: &completed})); err != nil {
		t.Fatalf("complete task without description: %v", err)
	}
	if !task.Completed || task.CompletedAt == nil {
		t.Fatal("completion must set status and timestamp")
	}
}
