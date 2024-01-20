package handler

import "fmt"

type CreateTaskRequest struct {
	Description string `json:description`
	IsDone      *bool  `json:isdone`
}

func (c *CreateTaskRequest) validade() error {
	if c.Description == "" {
		return checksRequiredParameter("Description")
	}
	if c.IsDone == nil {
		return checksRequiredParameter("isdone")
	}
	return nil
}

func checksRequiredParameter(parameter string) error {
	return fmt.Errorf("Parameter %s is required", parameter)
}
