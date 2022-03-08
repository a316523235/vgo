package models

type Switch struct {
	TaskSwitch bool
	PrintKey bool
}

func (s *Switch) CloseTask() {
	s.TaskSwitch = false
}

func (s *Switch) OpenTask()  {
	s.TaskSwitch = true
}

func (s *Switch) IsTaskOpen() bool {
	return s.TaskSwitch == true
}

func (s *Switch) IsPrintKey() bool {
	return s.PrintKey == true
}