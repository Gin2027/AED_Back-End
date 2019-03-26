package apis

import (
	"aed/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TRegister(c *gin.Context)  {
	var status string

	account := c.PostForm("account")
	password := c.PostForm("password")
	name := c.PostForm("name")
	introduce := c.PostForm("introduce")
	phone := c.PostForm("phone")
	cer := 0

	te := models.GetTeacherByAccount(account)
	if te.Id!=0 {
		status = "该用户名已经被占用"
	} else {
		password = models.GetMD5(password)
		teacher := models.Teacher{Account:account,Password:password,Name:name,Introduce:introduce,Phone:phone,Certification:cer,Notice:"暂无公告"}
		flag := teacher.AddTeacher()
		if flag {
			status = "注册成功"
		} else {
			status = "注册失败"
		}
	}
	c.String(http.StatusOK,status)
}

func TLogin(c *gin.Context)  {
	var flag bool
	var course string
	account := c.PostForm("account")
	password := c.PostForm("password")
	te := models.GetTeacherByAccount(account)
	if te.Id!=0 {
		if models.GetMD5(password) == te.Password {
			flag=true
			course = "登录成功"
		} else {
			flag=false
			course = "密码错误"
		}
	} else {
		flag=false
		course = "没有此用户名"
	}
	if flag!=true {te=nil}
	c.JSON(http.StatusOK,gin.H{
		"status" : flag,
		"message" :te,
		"course" : course,
	})
}

func ModifyTeacher(c *gin.Context)  {
	id := c.PostForm("id")
	tid,_ :=strconv.Atoi(id)
	te := models.GetTeacherById(tid)
	te.Introduce = c.PostForm("introduce")
	te.Name = c.PostForm("name")
	te.Phone = c.PostForm("phone")
	te.Notice = c.PostForm("notice")

	flag := te.EditTeacher()
	if flag!=false {
		c.JSON(http.StatusOK,gin.H{
			"status" : flag,
			"message" : te,
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"status" : flag,
			"message" : models.GetTeacherById(tid),
		})
	}
}

func GetTeacher(c *gin.Context) {
	tid :=c.PostForm("tid")
	var ntid int
	ntid,err :=strconv.Atoi(tid)
	if err !=nil {
	}
	teacher := models.GetTeacherById(ntid)
	c.JSON(http.StatusOK,gin.H{
		"id" :teacher.Id,
		"name" : teacher.Name,
		"introduce" : teacher.Introduce,
		"phone" :teacher.Phone,
		"notice" : teacher.Notice,
	})

}

func TeachersList(c *gin.Context)  {
	list := models.GetAllTeachers()
	c.JSON(http.StatusOK,gin.H{
		"teacherList" : list,
		"count" : len(list),
	})
}

func AdviceList(c *gin.Context)  {
	tid :=c.PostForm("tid")
	var ntid int
	ntid,err :=strconv.Atoi(tid)
	if err != nil {
	}

	list := models.GetAdviceList(ntid)
	c.JSON(http.StatusOK,gin.H{
		"advicelist" : list,
		"count" : len(list),
	})
}

func Reply(c *gin.Context) {
	aid :=c.PostForm("aid")
	reply := c.PostForm("reply")
	var naid int
	naid,err :=strconv.Atoi(aid)
	if err != nil {
	}

	models.ReplyAdvice(naid,reply)
	c.String(http.StatusOK,"回复成功")
}