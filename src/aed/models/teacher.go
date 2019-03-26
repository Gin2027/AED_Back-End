package models

import (
	db "aed/database"
	"log"
	"time"
)

type Teacher struct {
	Id int `json:"id"`
	Account string `json:"-"`
	Password string `json:"-"`
	Name string `json:"name"`
	Introduce string `json:"introduce"`
	Phone string `json:"phone"`
	Certification int `json:"-"`
	Notice string `json:"notice"`
}

type Advice struct {
	Aid int `json:"aid"`
	Sid int `json:"sid"`
	Tid int `json:"tid"`
	Message string `json:"message"`
	Reply string `json:"reply"`
	Gtime string`json:"gtime"`
	Rtime string `json:"rtime"`
}

func (t *Teacher) AddTeacher() bool {
	rs, err := db.SqlDB.Exec("INSERT INTO teacher(account,password,name,introduce,phone,certification,notice) values (?,?,?,?,?,?,?)",t.Account,t.Password,t.Name,t.Introduce,t.Phone,t.Certification,t.Notice)
	if err != nil {
		return false
	}
	id, err := rs.LastInsertId()
	if err!=nil {
		return false
	} else {
		log.Println("成功注册教师,ID为",id)
		return true
	}
}

func (t *Teacher) EditTeacher() bool {
	rs, err :=db.SqlDB.Exec("UPDATE teacher set introduce=?,name=?,phone=?,notice=? where id=?",t.Introduce,t.Name,t.Phone,t.Notice,t.Id)
	if err != nil {
		return false
	}
	id, err:=rs.RowsAffected()
	if err!=nil{
		return false
	} else {
		log.Println("教师修改信息，ID为",id)
		return true
	}
}

func GetTeacherByAccount(account string) (t *Teacher)  {
	var teacher Teacher
	err := db.SqlDB.QueryRow("SELECT * FROM teacher WHERE account=?",account).Scan(
		&teacher.Id,&teacher.Account,&teacher.Password,&teacher.Name,&teacher.Introduce,&teacher.Phone,&teacher.Certification,&teacher.Notice)

	if err != nil {
		log.Println(err)
	}

	return &teacher
}

func GetTeacherById(id int) (t *Teacher)  {
	var teacher Teacher
	err := db.SqlDB.QueryRow("SELECT * FROM teacher WHERE id=?",id).Scan(
		&teacher.Id,&teacher.Account,&teacher.Password,&teacher.Name,&teacher.Introduce,&teacher.Phone,&teacher.Certification,&teacher.Notice)

	if err != nil {
		log.Println(err)
	}

	return &teacher
}

func GetAllTeachers() (teachers []Teacher)  {
	teachers = make([]Teacher,0)
	rows, err :=db.SqlDB.Query("SELECT * FROM teacher")
	if err != nil{
		return
	}
	for rows.Next() {
		var teacher Teacher
		rows.Scan(&teacher.Id,&teacher.Account,&teacher.Password,&teacher.Name,&teacher.Introduce,&teacher.Phone,&teacher.Certification,&teacher.Notice)
		teacher.Account = "0"
		teacher.Password = "0"
		teachers = append(teachers,teacher)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return teachers
}

func GetAdviceList(tid int) (advices []Advice) {
	advices = make([]Advice, 0)
	rows, err := db.SqlDB.Query("SELECT * FROM advice where tid = ?", tid)
	if err != nil {
		return
	}
	for rows.Next() {
		var advice Advice
		rows.Scan(&advice.Aid, &advice.Sid, &advice.Tid, &advice.Message, &advice.Reply, &advice.Gtime, &advice.Rtime)
		advices = append(advices, advice)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return advices
}

func ReplyAdvice(aid int,reply string) bool {
	t:=time.Now()
	nt := t.Format("2006-01-02 15:04:05")
	rs,err:=db.SqlDB.Exec("UPDATE advice set reply=?,rtime=? where aid=?",reply,nt,aid)
	if err != nil {
		return false
	}
	_,err =rs.RowsAffected()
	if err != nil {
		return false
	} else {
		return true
	}
}