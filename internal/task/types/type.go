package types

import "reflect"

const (
	Runnable = iota
	Running
	Success
	Failed
)

type TaskArgs interface {
	Map() map[string]interface{}
	TaskName() string
}

type AudioArgs struct {
	TextInput string
	ToneInput string
	OutputDir string
	FileName  string
}

func (a AudioArgs) Map() map[string]interface{} {
	return struct2map(a)
}

func (a AudioArgs) TaskName() string {
	return "audio"
}

type DpArgs struct {
	AudioInput string
	ImageInput string
	OutputDir  string
	FileName   string
}

func (a DpArgs) Map() map[string]interface{} {
	return struct2map(a)
}

func (a DpArgs) TaskName() string {
	return "dp"
}

func struct2map(data interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	reflectValue := reflect.ValueOf(data)
	if reflectValue.Kind() != reflect.Struct {
		return nil
	}
	reflectType := reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fieldName := reflectType.Field(i).Name
		fieldValue := reflectValue.Field(i).Interface()
		result[fieldName] = fieldValue
	}

	return result
}

// TaskExcutor struct
// + Excutor(execTmpl) (status int, err error)
// + Error()
type TaskExcutorIntf interface {
	Execute()
	Error() error
	GetTaskId() string
}

// TaskManager intf
// + taskCh chan TaskExcutor
// + resCh chan TaskRes
//
// + Start()
// + UpdateTask(taskId string, status int)
// + RegisterTask(taskArgsSli TaskArgs...)

type TaskManagerIntf interface {
	Start()
	RegisterTask(taskId string, taskArgsSli ...TaskArgs) error
	UpdateTask(taskId string, data interface{})
	QueryTask(taskId string) (interface{}, bool)
}
