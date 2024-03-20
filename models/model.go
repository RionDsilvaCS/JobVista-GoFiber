package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty" validate:"required"`
	Mail_ID  string             `json:"mail_id,omitempty" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Role     string             `json:"role,omitempty" validate:"required"`
}

type Job struct {
	Id            primitive.ObjectID `json:"id,omitempty"`
	Job_title     string             `json:"job_title,omitempty" validate:"required"`
	Job_des       string             `json:"job_des,omitempty" validate:"required"`
	Company       string             `json:"company,omitempty" validate:"required"`
	Location      string             `json:"location,omitempty" validate:"required"`
	Start_date    string             `json:"start_date,omitempty" validate:"required"`
	Last_date     string             `json:"last_date,omitempty" validate:"required"`
	Salary        int16              `json:"salary,omitempty" validate:"required"`
	Recruit_count int8               `json:"recruit_count,omitempty" validate:"required"`
	Type          string             `json:"type,omitempty" validate:"required"`
	Duration      string             `json:"duration,omitempty" validate:"required"`
	Applicants    string             `json:"applicants,omitempty" validate:"required"`
	Skills        map[string]string  `json:"skills,omitempty" validate:"required"`
	User_ID       primitive.ObjectID `json:"user_id,omitempty" validate:"required"`
}

type Testimonial struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Company string             `json:"company,omitempty" validate:"required"`
	Role    string             `json:"role,omitempty" validate:"required"`
	Title   string             `json:"title,omitempty" validate:"required"`
	Body    string             `json:"body,omitempty" validate:"required"`
	Credits map[string]string  `json:"credits,omitempty" validate:"required"`
	User_ID primitive.ObjectID `json:"user_id,omitempty" validate:"required"`
}

type Application struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty"`
	Mail_ID     string             `json:"mail_id,omitempty"`
	Mobile_no   string             `json:"mobile_id,omitempty"`
	Education   string             `json:"education,omitempty"`
	Work_exp    map[string]string  `json:"work_exp,omitempty"`
	Skills      map[string]string  `json:"skills,omitempty"`
	Work_link   map[string]string  `json:"work_link,omitempty"`
	Resume_link string             `json:"resume_link,omitempty"`
	User_ID     primitive.ObjectID `json:"user_id,omitempty" validate:"required"`
}
