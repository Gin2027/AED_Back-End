package apis

import (
	"aed/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func MD5(c *gin.Context){
	str := c.Param("str")
	nstr := models.GetMD5(str)
	c.String(http.StatusOK,nstr)
}

func GetAll(c *gin.Context)  {
	datalist := models.GetAllStudents()
	c.JSON(http.StatusOK,gin.H{
		"datalist" : datalist,
		"count" : len(datalist),
	})

}
/*func GetOne(c *gin.Context) {
	var flag bool
	account := c.Query("account")
	stu := models.GetStudentByAccount(account)
	if stu.Id!=0 {
		flag=true
	} else {
		flag=false
	}
	c.JSON(http.StatusOK,gin.H{
		"success": flag,
		"message": stu,
	})
}*/

func Register(c *gin.Context)  {
	var status string

	account := c.PostForm("account")
	password := c.PostForm("password")
	name := c.PostForm("name")
	a := c.PostForm("age")
	var age int
	age,err := strconv.Atoi(a)
	if err != nil {
	}
	location := c.PostForm("location")
	school := c.PostForm("school")
	phone := c.PostForm("phone")
	qq := c.PostForm("qq")
	var tid int
	tid = 0

	stu := models.GetStudentByAccount(account)
	if stu.Id!=0 {
		status = "该用户名已经被占用"
	} else {
		password = models.GetMD5(password)
		student := models.Student{Account:account,Password:password,Name:name,Age:age,Location:location,School:school,Phone:phone,QQ:qq,Tid:tid}
		fmt.Println(student)
		flag := student.AddStudent()
		if flag {
			status = "注册成功"
		} else {
			status = "注册失败"
		}
	}
	c.String(http.StatusOK,status)
}

func Login(c *gin.Context)  {
	var flag bool
	var course string
	account := c.PostForm("account")
	password := c.PostForm("password")
	stu := models.GetStudentByAccount(account)
	if stu.Id!=0 {
		if models.GetMD5(password) == stu.Password {
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
	if flag!=true {stu=nil}
	c.JSON(http.StatusOK,gin.H{
		"status" : flag,
		"student" :stu,
		"course" : course,
	})
}

func ModifyStudent(c *gin.Context)  {
	id := c.PostForm("id")
	var nid int
	nid,err :=strconv.Atoi(id)
	if err != nil {}
	stu := models.GetStudentById(nid)
	oldpw := c.PostForm("old")
	newpw := c.PostForm("new")

	if (models.GetMD5(oldpw)!=stu.Password) {
		c.String(http.StatusOK, "旧密码不正确")
	} else {
		stu.Password = models.GetMD5(newpw)
		stu.EditStudent()
		c.String(http.StatusOK,"修改成功")
	}
}
func ChooseTeacher(c *gin.Context)  {

	//解除教师关系将tid设置为0即可

	tid := c.PostForm("tid")
	sid := c.PostForm("sid")

	var nsid int
	nsid,err :=strconv.Atoi(sid)
	if err != nil {
	}

	var ntid int
	ntid,err1 :=strconv.Atoi(tid)
	if err1 !=nil {
	}

	var flag bool
	stu := models.GetStudentById(nsid)
	if stu.Tid==0 {
		flag = models.SetTeacher(nsid,ntid)
	} else
	{
		flag = false
	}

	c.JSON(http.StatusOK,flag)
}
func GiveAdvice(c *gin.Context)  {
	sid := c.PostForm("sid")
	message := c.PostForm("message")
	var nsid int
	nsid,err :=strconv.Atoi(sid)
	if err != nil {
	}

	stu :=models.GetStudentById(nsid)
	models.AddAdvice(nsid,stu.Tid,message)

	c.JSON(http.StatusOK,true)
}

func GetReplies(c *gin.Context)  {
	sid := c.PostForm("sid")
	var nsid int
	nsid,_ =strconv.Atoi(sid)
	adviceList := models.GetAllReplys(nsid)
	c.JSON(http.StatusOK,gin.H{
		"adviceList" : adviceList,
		"count" : len(adviceList),
	})

}
