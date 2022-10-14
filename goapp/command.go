package goapp

import (
	"github.com/mattermost/mattermost-plugin-apps/apps"
)

type Command struct {
	Name        string
	Hint        string
	Description string
	Icon        string
	BaseSubmit  *apps.Call
	BaseForm    *apps.Form

	Handler func(CallRequest) apps.CallResponse
}

func (c Command) Path() string {
	return "/" + c.Name
}

func (c Command) Submit(creq CallRequest) *apps.Call {
	if c.BaseSubmit == nil {
		return nil
	}
	s := *c.BaseSubmit.PartialCopy()
	if s.Path == "" {
		s.Path = c.Path()
	}
	return &s
}

func (c Command) Form(creq CallRequest) *apps.Form {
	if c.BaseForm == nil {
		return nil
	}
	f := *c.BaseForm.PartialCopy()
	if f.Icon == "" {
		f.Icon = creq.App.Icon
	}
	if f.Submit == nil {
		f.Submit = c.Submit(creq)
	} else if f.Submit.Path == "" {
		f.Submit.Path = c.Path()
	}

	f.Fields = creq.AppendDebugJSON(f.Fields)
	return &f
}

func (c Command) Binding(creq CallRequest) apps.Binding {
	b := apps.Binding{
		Location:    apps.Location(c.Name),
		Icon:        c.Icon,
		Hint:        c.Hint,
		Description: c.Description,
		Submit:      c.Submit(creq),
		Form:        c.Form(creq),
	}
	if b.Icon == "" {
		b.Icon = creq.App.Icon
	}
	return b
}
