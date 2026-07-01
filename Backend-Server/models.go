package main

import (
	"database/sql"
	"time"

	"github.com/sanrose-bhets/Job-Grinder/internal/database"
)

type Account struct {
	ID     int32  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"mail"`
	Role   string `json:"role"`
	Apikey string `json:"api"`
}
type Company struct {
	Companyid         int32          `json:"id"`
	Companyname       string         `json:"name"`
	Companylocation   string         `json:"location"`
	Companywebsite    string         `json:"site"`
	Companyemail      string         `json:"mail"`
	Hiringmanagername sql.NullString `json:"hiringmanager"`
}
type Application struct {
	ID          int32         `json:"id"`
	Department  string        `json:"department"`
	Status      string        `json:"status"`
	Company     int32         `json:"company"`
	Position    int32         `json:"position"`
	Interviewid sql.NullInt32 `json:"interviewId"`
	DateApplied time.Time     `json:"dateApplied"`
}
type Interview struct {
	ID            int32     `json:"id"`
	Status        string    `json:"status"`
	Interviewdate time.Time `json:"interviewDate"`
}

type Job struct {
	Postionid            int32  `json:"id"`
	Positionname         string `json:"name"`
	Companyid            int32  `json:"companyid"`
	Experiencerequired   string `json:"experienceRequired"`
	Expectedcompensation string `json:"expectedCompensation"`
	Employmenttype       string `json:"employmentType"`
	Joblocationtype      string `json:"jobLocationType"`
}

func databaseAcctoAcc(dbUser database.Account) Account {
	return Account{
		ID:     dbUser.ID,
		Name:   dbUser.Name,
		Email:  dbUser.Email,
		Role:   dbUser.Role,
		Apikey: dbUser.Apikey,
	}
}

func databaseComptoComp(dbUser database.Company) Company {
	return Company{
		Companyid:         dbUser.Companyid,
		Companyname:       dbUser.Companyname,
		Companylocation:   dbUser.Companylocation,
		Companywebsite:    dbUser.Companywebsite,
		Companyemail:      dbUser.Companyemail,
		Hiringmanagername: dbUser.Hiringmanagername,
	}
}

func databaseIntvwtoIntvw(dbUser database.Interview) Interview {
	return Interview{
		ID:            dbUser.ID,
		Status:        dbUser.Status,
		Interviewdate: dbUser.Interviewdate,
	}
}

func databaseAppToApp(dbUser database.Application) Application {
	return Application{
		ID:          dbUser.ID,
		Department:  dbUser.Department,
		Status:      dbUser.Status,
		Company:     dbUser.Company,
		Position:    dbUser.Position,
		Interviewid: dbUser.Interviewid,
		DateApplied: dbUser.DateApplied,
	}
}

func databaseJobtoJob(dbUser database.Job) Job {
	return Job{
		Postionid:            dbUser.Postionid,
		Positionname:         dbUser.Positionname,
		Companyid:            dbUser.Companyid,
		Experiencerequired:   dbUser.Experiencerequired,
		Expectedcompensation: dbUser.Expectedcompensation,
		Employmenttype:       dbUser.Employmenttype,
		Joblocationtype:      dbUser.Joblocationtype,
	}
}
