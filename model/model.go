package model

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Quries struct {
	gorm.Model
	ID                 string `json:"id"`
	Text               string `json:"text"`
	Posted_by          string `json:"posted_by"`
	Status             string `json:"status"`
	Resolved_by        string `json:"resolved_by"`
	Question_id        string `json:"question_id"`
	Reading_materialId string `json:"reading_materialid"`
	Task_Id            string `json:"task_id"`
	ResolutionStarted  string `json:"resolutionsstarted"`
}
type ChatHistory struct {
	gorm.Model
	ID      string `json:"id"`
	FromId  string `json:"fromid"`
	ToId    string `json:"toid"`
	QueryId string `json:"queryid"`
	Message string `json:"message"`
	status  string `json:"status"`
}

type Users struct {
	ID              string `json:"id"`
	User_Profile_Id string `json:"user_profileid"`
	Email_Id        string `json:"emailid"`
	Password        string `json:"password"`
	Auth_Id         string `json:"authid"`
	Created_at      string `json:"careatedat"`
	Update_at       string `json:"updateat"`
	Active_status   string `json:"activatestatus"`
	Points          string `json:"points"`
}
type Problem struct {
	gorm.Model
	Problem_ID             string `form:"problem_id" query:"problem_id" json:"problem_id"`
	Problem_statement_html string `form:"problem_statement_html" query:"problem_statement_html" json:"problem_statement_html"`
	Test_cases_ids         string `form:"test_cases_ids" query:"test_cases_ids" json:"test_cases_ids"`
}
type Problem_Notification struct {
	gorm.Model
	Problem_ID        string `form:"problem_id" query:"problem_id" json:"problem_id"`
	User_ID           string `form:"userid" query:"userid" json:"userid"`
	Picked            bool   `form:"picked" query:"picked" json:"picked"`
	Problem_Picker_id string `form:"problem_picker_id" query:"problem_picker_id" json:"problem_picker_id"`
}
type User_profile struct {
	gorm.Model
	Roll_no   string `form:"roll_no" query:"roll_no" json:"roll_no",not null`
	FirstName string `form:"first_name" query:"first_name" json:"first_name"`

	LastName string `form:"last_name" query:"last_name" json:"last_name"`

	Dob          string `form:"dob" query:"dob" json:"dob,omitempty",not null`
	College_id   string `form:"college_id" query:"college_id" json:"college_id",not null`
	Dept         string `form:"dept" query:"dept" json:"dept,omitempty",not null`
	Phone_no     string `form:"phone_no" query:"phone_no" json:"phone_no",not null`
	Alt_Phone_no string `form:"alt_phone_no" query:"alt_phone_no" json:"alt_phone_no"`
	Batch_id     string `form:"batch_id" query:"batch_id" json:"batch_id"`
	State        string `form:"state" query:"state" json:"state"`
	City         string `form:"city" query:"city" json:"city"`
	Srn_usn      string `form:"srn_usn" query:"srn_usn" json:"srn_usn"`
	Semester     string `form:"semester" query:"semester" json:"semester"`
}

type Batch struct {
	gorm.Model
	Current_course_id string `form:"current_course_id" query:"current_course_id" json:"current_course_id"`
	Batch_id          string `form:"batch_id" query:"batch_id" json:"batch_id"`
}
type Problem_Solved struct {
	gorm.Model
	Problem_ID string `form:"problem_id" query:"problem_id" json:"problem_id"`
	User_ID    string `form:"user_id" query:"user_id" json:"user_id"`
}
type TaskReleased struct {
	gorm.Model
	TaskId      string `form:"taskid" query:"taskid" json:"taskid"`
	Batch_id    string `form:"batch_id" query:"batch_id" json:"batch_id"`
	Dead_line_1 string `form:"dead_line_1" query:"dead_line_1" json:"dead_line_1"`
	Dead_line_2 string `form:"dead_line_2" query:"dead_line_2" json:"dead_line_2"`
}

type Course struct {
	gorm.Model
	CourseID          string `form:"courseid" query:"courseid" json:"courseid,omitempty"`
	CourseName        string `form:"coursename" query:"coursename" json:"coursename,omitempty"`
	CourseDescription string `form:"coursedescription" query:"coursedescription" json:"coursedescription,omitempty"`
	CourseSylabbus    string `form:"coursesylabbus" query:"coursesylabbus" json:"coursesylabbus,omitempty"`
}
type College struct {
	gorm.Model
	Name       string `form:"name" query:"name" json:"name"`
	City       string `form:"city" query:"city" json:"city"`
	State      string `form:"state" query:"state" json:"state"`
	University string `form:"university" query:"university" json:"university"`
	CollegeID  string `form:"collegeid" query:"collegeid" json:"collegeid"`
}
type ReadingMaterial struct {
	gorm.Model
	MaterialID string `form:"materialid" query:"materialid" json:"materialid"`
	Path       string `form:"path" query:"path" json:"path"`
	UpdatePath string `form:"updatepath" query:"updatepath" json:"updatepath"`
}
type Task struct {
	gorm.Model
	CourseID          string `form:"courseid" query:"courseid" json:"courseid"`
	TaskID            string `form:"taskid" query:"taskid" json:"taskid"`
	TaskName          string `form:"taskname" query:"taskname" json:"taskname"`
	ReadingMaterialID string `form:"readingmaterialid" query:"updareadingmaterialidtepath" json:"readingmaterialid"`
	VideoLink         string `form:"videolink" query:"videolink" json:"videolink"`
	Marks             string `form:"marks" query:"marks" json:"marks"`
	Description       string `form:"description" query:"description" json:"description"`
	QuestionIDs       string `form:"questionids" query:"questionids" json:"questionids"`
	McqID             string `form:"mcqid" query:"mcqid" json:"mcqid"`
}

func Adduser(u *Users) (*Users, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return u, err

}
func AddCourse(c *Course) (*Course, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err

}
func DeleteCourse(id int) error {
	var tempcollege College
	tempcollege.CollegeID = ""
	if db := db.Find(&tempcollege); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempcollege).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err

}
func UpdateCourse(course *Course) (*Course, error) {
	var tempCourse Course
	tempCourse.ID = course.ID
	if db := db.Find(&tempCourse); db.Error != nil {
		return nil, db.Error
	}
	if course.CourseName != "" {
		tempCourse.CourseName = course.CourseName
	}
	if course.CourseDescription != "" {
		tempCourse.CourseDescription = course.CourseDescription
	}
	if course.CourseSylabbus != "" {
		tempCourse.CourseSylabbus = course.CourseSylabbus
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempCourse).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempCourse, err

}
func GetCourse(college_id string) ([]Course, error) {
	var course []Course
	db.Where("course_id=?", college_id).Find(&course)
	return course, nil

}

func AddCollege(c *College) (*College, error) {

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func DeleteCollege(id int) error {
	var tempcollege College
	tempcollege.CollegeID = ""
	if db := db.Find(&tempcollege); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempcollege).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}
func UpdateCollege(college *College) (*College, error) {
	var tempCollege College
	tempCollege.ID = college.ID
	if db := db.Find(&tempCollege); db.Error != nil {
		return nil, db.Error
	}
	if college.Name != "" {
		tempCollege.Name = college.Name
	}
	if college.City != "" {
		tempCollege.City = college.City
	}
	if college.State != "" {
		tempCollege.State = college.State
	}
	if college.University != "" {
		tempCollege.University = college.University
	}
	if college.CollegeID != "" {
		tempCollege.CollegeID = college.CollegeID
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempCollege).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempCollege, err

}
func GetCollege(id string) ([]College, error) {
	var college []College

	db.Find(&college)

	return college, nil
}

// curd operation for reading material
func AddReadingMaterial(c *ReadingMaterial) (*ReadingMaterial, error) {

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err

}
func DeleteReadingMaterial(id int) error {
	var tempmaterial ReadingMaterial
	tempmaterial.MaterialID = ""
	if db := db.Find(&tempmaterial); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempmaterial).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}
func UpdateReadingMaterial(material *ReadingMaterial) (*ReadingMaterial, error) {
	var tempmaterial ReadingMaterial
	tempmaterial.MaterialID = material.MaterialID
	if db := db.Find(&tempmaterial); db.Error != nil {
		return nil, db.Error
	}

	if material.Path != "" {
		tempmaterial.Path = material.Path
	}
	if material.UpdatePath != "" {
		tempmaterial.UpdatePath = material.UpdatePath
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempmaterial).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempmaterial, err
}
func GetReadingMaterial(id string) (*ReadingMaterial, error) {
	var material ReadingMaterial
	material.MaterialID = id
	if db := db.Find(&material); db.Error != nil {
		return nil, db.Error
	}
	return &material, nil
}

// curd operation for Task

func AddTask(c *Task) (*Task, error) {

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err

}
func DeleteTask(id int) error {
	var temptask Task
	temptask.TaskName = ""
	if db := db.Find(&temptask); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&temptask).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err
}
func UpdateTask(task *Task) (*Task, error) {
	var temptask Task
	temptask.TaskName = task.TaskName
	if db := db.Find(&temptask); db.Error != nil {
		return nil, db.Error
	}

	if task.ReadingMaterialID != "" {
		temptask.ReadingMaterialID = task.ReadingMaterialID
	}
	if task.VideoLink != "" {
		temptask.VideoLink = task.VideoLink
	}
	if task.Description != "" {
		temptask.Description = task.Description
	}
	if task.QuestionIDs != "" {
		temptask.QuestionIDs = task.QuestionIDs
	}
	if task.McqID != "" {
		temptask.McqID = task.McqID
	}
	if task.Marks != "" {
		temptask.Marks = task.Marks
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&temptask).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &temptask, err
}
func GetTask(task_id string) (*Task, error) {
	var tasks Task
	db.Where("task_id=?", task_id).Find(&tasks)
	return &tasks, nil
}
func GetTasks(task_id string) ([]Task, error) {
	var tasks []Task
	//db.Where("task_id=?", "2001").Find(&tasks)
	var r []string = strings.Split(task_id, ",")
	db.Where("task_id IN (?)", r).Find(&tasks)
	return tasks, nil
}
func Init() {
	fmt.Print("in init")
	var err error
	db, err = gorm.Open("mysql", "root:securepassword@tcp(localhost:3306)/ptech?charset=utf8&parseTime=True")
	if err != nil {
		fmt.Print(err)
	}
	if err == nil {
		fmt.Print("no error")
	}
	db.AutoMigrate(&Users{})

	db.AutoMigrate(&User_profile{})
	db.AutoMigrate(&Course{})
	db.AutoMigrate(&College{})
	db.AutoMigrate(&ReadingMaterial{})
	db.AutoMigrate(&Task{})
	db.AutoMigrate(&Batch{})
	db.AutoMigrate(&TaskReleased{})
	db.AutoMigrate(&Problem{})
	db.AutoMigrate(&Problem_Solved{})
	db.AutoMigrate(&Problem_Notification{})
	//	db.AutoMigrate(&ChatHistory{})

}
func Adduserprofile(c *User_profile) (*User_profile, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func DeleteuserProfile(id int) error {
	var tempprofile User_profile
	tempprofile.Roll_no = ""
	if db := db.Find(&tempprofile); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempprofile).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err

}
func UpdateuserProfile(userprofile *User_profile) (*User_profile, error) {

	var tempprofile User_profile
	tempprofile.Roll_no = userprofile.Roll_no
	if db := db.Find(&tempprofile); db.Error != nil {
		return nil, db.Error
	}

	if userprofile.FirstName != "" {
		tempprofile.FirstName = userprofile.FirstName
	}
	if userprofile.LastName != "" {
		tempprofile.LastName = userprofile.LastName
	}
	if userprofile.Dob != "" {
		tempprofile.Dob = userprofile.Dob
	}
	if userprofile.College_id != "" {
		tempprofile.College_id = userprofile.College_id
	}
	if userprofile.Dept != "" {
		tempprofile.Dept = userprofile.Dept
	}
	if userprofile.Phone_no != "" {
		tempprofile.Phone_no = userprofile.Phone_no
	}
	if userprofile.Alt_Phone_no != "" {
		tempprofile.Alt_Phone_no = userprofile.Alt_Phone_no
	}
	if userprofile.Batch_id != "" {
		tempprofile.Batch_id = userprofile.Batch_id
	}
	if userprofile.State != "" {
		tempprofile.State = userprofile.State
	}
	if userprofile.Semester != "" {
		tempprofile.Semester = userprofile.Semester
	}
	if userprofile.Srn_usn != "" {
		tempprofile.Srn_usn = userprofile.Srn_usn
	}

	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Save(&tempprofile).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return &tempprofile, err
}
func GetuserProfile(roll_no string) (*User_profile, error) {
	var profile User_profile

	db.Where("roll_no=?", roll_no).Find(&profile)

	return &profile, nil
}
func AddBatch(c *Batch) (*Batch, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func DeleteBatch(id int) error {
	var tempBatch Batch
	tempBatch.Batch_id = ""
	if db := db.Find(&tempBatch); db.Error != nil {
		return db.Error
	}
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := tx.Unscoped().Delete(&tempBatch).Error; err != nil {
		tx.Rollback()
		return err
	}
	err := tx.Commit().Error
	return err

}
func UpdatBatch(userprofile *User_profile) (*User_profile, error) {
	return nil, nil
}
func GetBatch(batchid string) (*Batch, error) {
	var batch Batch

	db.Where("batch_id=?", batchid).Find(&batch)

	return &batch, nil
}

// gygfhgjh
type Account struct {
	gorm.Model
	AccountID string `gorm:"not null;unique;index"`
	Password  uint32
}

func AddTaskReleased(c *TaskReleased) (*TaskReleased, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func DeleteTaskReleased(id int) error {
	// var tempBatch Batch
	// tempBatch.Batch_id = ""
	// if db := db.Find(&tempBatch); db.Error != nil {
	// 	return db.Error
	// }
	// tx := db.Begin()
	// if tx.Error != nil {
	// 	return tx.Error
	// }
	// if err := tx.Unscoped().Delete(&tempBatch).Error; err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
	// err := tx.Commit().Error
	return nil

}
func UpdateTaskReleased(userprofile *User_profile) (*User_profile, error) {
	return nil, nil
}
func GetTaskReleased(batchid string) ([]TaskReleased, error) {
	var batch []TaskReleased

	db.Where("batch_id=?", batchid).Find(&batch)

	return batch, nil
}
func GetTaskReleasedNo(batchid string) (*int, error) {
	var taskReleased TaskReleased
	var taskno int

	db.Where("batch_id = ?", batchid).Find(&taskReleased).Count(&taskno)

	return &taskno, nil
}

func DeleteProblem(id int) error {

	return nil

}
func UpdateProblem(userprofile *User_profile) (*User_profile, error) {
	return nil, nil
}
func GetProblems(problem_ids string) ([]Problem, error) {
	var problem []Problem

	db.Where("problem_id=?", problem_ids).Find(&problem)

	return problem, nil
}
func AddProblem(c *Problem) (*Problem, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func GetProblemSolved(problem_ids string) ([]Problem_Solved, error) {
	var problemsolved []Problem_Solved
	db.Find(&problemsolved)
	//db.Where("problem_id=?", problem_ids).Find(&problem)

	return problemsolved, nil
}
func AddProblemSolved(c *Problem_Solved) (*Problem_Solved, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}
func AddProblemNotification(c *Problem_Notification) (*Problem_Notification, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}

func GetProblemNotification(problem_ids string) ([]Problem_Notification, error) {
	type Result struct {
		Problem_ID string `form:"problem_id" query:"problem_id" json:"problem_id"`
	}
	var s []string
	var result []Result
	db.Raw("SELECT problem_id FROM problem_solveds").Scan(&result)
	for _, num := range result {
		s = append(s, num.Problem_ID)
	}
	var problemsolved []Problem_Notification
	//db.Raw("select user_id from problem_solveds").Scan(&problem)
	fmt.Print(s)
	db.Where("problem_id IN (?)", s).Find(&problemsolved)
	//db.Find(&problemsolved)
	//db.Where("problem_id=?", problem_ids).Find(&problem)

	return problemsolved, nil
}
func AddChat(c *ChatHistory) (*ChatHistory, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	if err := tx.Create(c).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	err := tx.Commit().Error
	return c, err
}

func GetChat(queryid string) ([]ChatHistory, error) {

	var chats []ChatHistory
	//db.Raw("select user_id from problem_solveds").Scan(&problem)

	db.Where("queryid IN (?)", queryid).Find(&chats)
	//db.Find(&problemsolved)
	//db.Where("problem_id=?", problem_ids).Find(&problem)

	return chats, nil
}
