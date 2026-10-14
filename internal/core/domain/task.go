package domain

import (
	"fmt"
	"time"

	core_errors "github.com/MaximKachkov/ToDooo/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserId int
}

func NewTask(id int, version int, title string, description *string, completed bool, createdAt time.Time, completedAt *time.Time, authorID int) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
		AuthorUserId: authorID,
	}
}

func NewTaskUninitialized(title string, description *string, authorID int) Task {
	return NewTask(UninitializedID, UninitializedVersion, title, description, false, time.Now().UTC(), nil, authorID)

}

func (t Task) CompletionDuration() *time.Duration {
	if !t.Completed {
		return nil
	}
	if t.CompletedAt == nil {
		return nil
	}

	duration := t.CompletedAt.Sub(t.CreatedAt)

	return &duration
}

func (t Task) Validate() error {
	lenTitle := len([]rune(t.Title))
	lenDescription := len([]rune(*t.Description))

	if lenTitle < 1 || lenTitle > 100 {
		return fmt.Errorf("title len need to be between 1 and 100 : %w", core_errors.ErrInvalidArgument)
	}
	if lenDescription > 1000 {
		return fmt.Errorf("description len must be between 1 and 1000 symbols : %w", core_errors.ErrInvalidArgument)
	}
	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("error completed is true but completed at is nil : %w", core_errors.ErrInvalidArgument)
		}
		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("completed at cannot be before started at :%w", core_errors.ErrInvalidArgument)
		}

	} else {
		if t.CompletedAt != nil {
			fmt.Errorf("completed at exists while completed is false:%w", core_errors.ErrInvalidArgument)
		}
	}
	if t.AuthorUserId < 0 {
		return fmt.Errorf("Author id cannot be negative:%w", core_errors.ErrInvalidArgument)
	}
	return nil
}

type TaskPatch struct {
	Title       Nullable[string]
	Description Nullable[string]
	Completed   Nullable[bool]
}

func (t *TaskPatch) Validate() error {
	if t.Title.Set && t.Title.Value == nil {
		return fmt.Errorf("title cannot be nil:%w", core_errors.ErrInvalidArgument)
	}
	if t.Description.Set && t.Description.Value == nil {
		empty := ""
		t.Description.Value = &empty
	}
	if t.Completed.Set && t.Completed.Value == nil {
		return fmt.Errorf("complteted cannot be nil:%w", core_errors.ErrInvalidArgument)
	}

	return nil
}

func NewTaskPatch(title Nullable[string], desc Nullable[string], completed Nullable[bool]) TaskPatch {
	var taskPatch TaskPatch

	taskPatch.Title = title
	taskPatch.Description = desc
	taskPatch.Completed = completed

	return taskPatch
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	err := patch.Validate()
	if err != nil {
		return fmt.Errorf("error apply patch:%w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		if patch.Description.Value == nil {
			empty := ""
			tmp.Description = &empty
		} else {
			tmp.Description = patch.Description.Value
		}
	}
	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value

		if tmp.Completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt

		} else {
			tmp.CompletedAt = nil
		}
	}

	err = tmp.Validate()
	if err != nil {
		return fmt.Errorf("error validating tmp:%w", err)
	}

	*t = tmp
	return nil
}
