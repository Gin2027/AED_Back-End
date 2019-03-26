package models

import (
	db "aed/database"
	"fmt"
	"log"
	"time"
)
type Student struct {
	Id int `json:"id"`
	Account string `json:"account"`
	Password string `json:"password"`
	Name string `json:"name"`
	Age int `json:"age"`
	Location string `json:"location"`
	School string `json:"school"`
	Phone string `json:"phone"`
	QQ string `json:"qq"`
	Tid int `json:"tid"`
}




func (s *Student) AddStudent() bool {
	rs, err := db.SqlDB.Exec("INSERT INTO student(account,password,name,age,location,school,phone,qq,tid) values (?,?,?,?,?,?,?,?,?)",s.Account, s.Password, s.Name, s.Age, s.Location, s.School, s.Phone, s.QQ,s.Tid)
	if err != nil {
		return false
	}
	id, err := rs.LastInsertId()
	if err!=nil {
		return false
	} else {
		log.Println("成功注册学生,ID为",id)
		return true
	}
}

func (s *Student) EditStudent() bool {
	rs, err :=db.SqlDB.Exec("UPDATE student set password=? where id=?",s.Password,s.Id)
	if err != nil {
		return false
	}
	id, err:=rs.RowsAffected()
	if err!=nil{
		return false
	} else {
		log.Println("学生修改密码，ID为",id)
		return true
	}
}

func GetAllStudents() (students []Student)  {
	students = make([]Student,0)
	rows, err :=db.SqlDB.Query("SELECT * FROM student")
	if err != nil{
		return
	}
	for rows.Next() {
		var student Student
		rows.Scan(&student.Id,&student.Account,&student.Password,&student.Name,&student.Age,&student.Location,&student.School,&student.Phone,&student.QQ,&student.Tid)
		students = append(students,student)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return students
}

func GetStudentByAccount(account string) (s *Student)  {
	var student Student
	err := db.SqlDB.QueryRow("SELECT * FROM student WHERE account=?",account).Scan(
		&student.Id,&student.Account,&student.Password,&student.Name,&student.Age,&student.Location,&student.School,&student.Phone,&student.QQ,&student.Tid)

	if err != nil {
		log.Println(err)
	}

	return &student
}

func GetStudentById(id int) (s *Student) {
	var student Student
	err := db.SqlDB.QueryRow("SELECT * FROM student WHERE id=?",id).Scan(
		&student.Id,&student.Account,&student.Password,&student.Name,&student.Age,&student.Location,&student.School,&student.Phone,&student.QQ,&student.Tid)

	if err != nil {
		log.Println(err)
	}

	return &student
}

func SetTeacher(sid int,tid int) bool {
	rs, err :=db.SqlDB.Exec("UPDATE student set tid=? where id=?",tid,sid)
	if err != nil {
		return false
	}
	id, err:=rs.RowsAffected()
	if err!=nil{
		return false
	} else {
		log.Println("学生",id,"配对教师",tid)
		return true
	}
}

func AddAdvice(sid int,tid int,message string) bool {

	t:=time.Now()
	nt := t.Format("2006-01-02 15:04:05")
	fmt.Println(nt)
	rs, err := db.SqlDB.Exec("INSERT INTO advice(sid,tid,message,reply,gtime,rtime) VALUES (?,?,?,?,?,?)", sid, tid, message, nil, nt,nil)
	if err != nil {
		return false
	}
	aid, err := rs.LastInsertId()
	if err != nil {
		return false
	} else {
		log.Println("成功留言,留言ID为", aid)
		return true
	}
}

func GetAllReplys(sid int) (advices []Advice) {
	advices = make([]Advice,0)
	rows, err :=db.SqlDB.Query("SELECT * FROM advice where sid=? and reply!=''",sid)
	if err!= nil{
		return
	}
	for rows.Next() {
		var advice Advice
		rows.Scan(&advice.Aid,&advice.Sid,&advice.Tid,&advice.Message,&advice.Reply,&advice.Gtime,&advice.Rtime)
		advices = append(advices,advice)
	}
	if err = rows.Err(); err != nil {
		return nil
	}
	return advices
}