package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"newgo/test/model"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

var cache redis.Conn

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

/*
handler for curd opertion of user profile
*/

func initCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {

	}
	cache = conn

}
func Login(w http.ResponseWriter, r *http.Request) {
	initCache()
	fmt.Print("in login")
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Print(creds.Username)
	fmt.Print(creds.Password)
	// Get the expected password from our in memory map
	expectedPassword := "password"

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create a new random session token
	sessionToken, err := uuid.NewV4()
	fmt.Print("session       ")
	fmt.Print(sessionToken)
	fmt.Print("      end")
	if err == nil {
		fmt.Print("     no error")
	}
	// Set the token in the cache, along with the user whom it represents
	// The token has an expiry time of 120 seconds
	_, err = cache.Do("SETEX", sessionToken, "120", creds.Username)
	if err != nil {
		// If there is an error in setting the cache, return an internal server error
		fmt.Print("       error")
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	// Finally, we set the client cookie for "session_token" as the session token we just generated
	// we also set an expiry time of 120 seconds, the same as the cache
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken.String(),
		Expires: time.Now().Add(1200 * time.Second),
	})
	w.Write([]byte(fmt.Sprintf("Welcome %s!", "logedd")))
}

// userprofile handler
func AddUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.User_profile
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.Adduserprofile(&u)
}

// update user profile handler
func UpdateUserProfileHandler(w http.ResponseWriter, r *http.Request) {

}

// delete user profile handler
func DeleteUserProfileHandler(w http.ResponseWriter, r *http.Request) {

}

// get user profile handler
func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	roll_no := params["roll_no"]
	u, err := model.GetuserProfile(roll_no)
	if err == nil {
		fmt.Print(u)
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)

}

// course update handler
func UpdateCourseHandler(w http.ResponseWriter, r *http.Request) {

}

// course delete handler
func DeleteCourseHandler(w http.ResponseWriter, r *http.Request) {

}
func isauthenticated(w http.ResponseWriter, r *http.Request) bool {

	// We can obtain the session token from the requests cookies, which come with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return false
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return false
	}
	sessionToken := c.Value

	// We then get the name of the user from our cache, where we set the session token
	response, err := cache.Do("GET", sessionToken)
	if err != nil {
		// If there is an error fetching from cache, return an internal server error status
		w.WriteHeader(http.StatusInternalServerError)
		return false
	}
	if response == nil {
		// If the session token is not present in cache, return an unauthorized error
		w.WriteHeader(http.StatusUnauthorized)
		return false
	}
	// Finally, return the welcome message to the user

	w.Write([]byte(fmt.Sprintf("Welcome %s!", "response")))
	return true
}

// course getting handler
func GetCourseHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	college_id := params["course_id"]
	u, err := model.GetCourse(college_id)
	if err != nil {
		print("error in getting task")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

// add course handler
func AddCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.Course
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddCourse(&u)
}

/* curd operation for college data
 */
// add college data handler
func AddCollegeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.College
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddCollege(&u)
}

// update college data handler
func UpdateCollegeHandler(w http.ResponseWriter, r *http.Request) {

}

// delete college data handler
func DeleteCollegeHandler(w http.ResponseWriter, r *http.Request) {

}

// get college handler
func GetCollegeHandler(w http.ResponseWriter, r *http.Request) {
	u, err := model.GetCollege("")
	if err == nil {
		fmt.Print(u)
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

/*
 handlers for curd operation of reading material
*/
// add reading material handler
func AddReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.ReadingMaterial
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddReadingMaterial(&u)
}

//update reading material handler
func UpdateReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {

}

// delete reading material handler
func DeleteReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {

}

// get reading material handler
func GetReadingMaterialHandler(w http.ResponseWriter, r *http.Request) {

}

/*
handler for curd operation of task
*/

//  add task handler
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.Task
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddTask(&u)
}

// update task handler
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
}

// delete task handler
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

}
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	task_id := params["task_id"]
	u, err := model.GetTask(task_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	task_ids := params["ids"]
	fmt.Print(task_ids)
	u, err := model.GetTasks(task_ids)

	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

// get task handler
func GetBatchHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	batch_id := params["batchid"]
	u, err := model.GetBatch(batch_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

func AddBatchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.Batch
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddBatch(&u)
}

// update task handler
func UpdateBatchHandler(w http.ResponseWriter, r *http.Request) {
}

// delete task handler
func DeleteBatchHandler(w http.ResponseWriter, r *http.Request) {

}

// get task handler

func GetTaskReleasedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	batch_id := params["batchid"]
	u, err := model.GetTaskReleased(batch_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

func GetTaskReleasedNoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	batch_id := params["batchid"]
	u, err := model.GetTaskReleasedNo(batch_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

func AddTaskReleasedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var u model.TaskReleased
	_ = json.NewDecoder(r.Body).Decode(&u)
	json.NewEncoder(w).Encode(u)

	model.AddTaskReleased(&u)
}

// update task handler
func UpdateTaskReleasedHandler(w http.ResponseWriter, r *http.Request) {
}

// delete task handler
func DeleteTaskReleasedHandler(w http.ResponseWriter, r *http.Request) {

}

func AddProblemHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var problem model.Problem
	_ = json.NewDecoder(r.Body).Decode(&problem)
	json.NewEncoder(w).Encode(problem)

	model.AddProblem(&problem)
}

func GetProblemsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problem_id := params["problem_ids"]
	u, err := model.GetProblems(problem_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}

func DeleteProblemHandler(w http.ResponseWriter, r *http.Request) {

}

func UpdateProblemHandler(w http.ResponseWriter, r *http.Request) {

}

func AddProblemSolvedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var problemsolved model.Problem_Solved
	_ = json.NewDecoder(r.Body).Decode(&problemsolved)
	json.NewEncoder(w).Encode(problemsolved)

	model.AddProblemSolved(&problemsolved)
}

func GetProblemSolvedHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problemsolved_id := params["problemsolved_ids"]
	u, err := model.GetProblemSolved(problemsolved_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)
}
func AddProblemNotificationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var problemnotification model.Problem_Notification
	_ = json.NewDecoder(r.Body).Decode(&problemnotification)
	json.NewEncoder(w).Encode(problemnotification)

	model.AddProblemNotification(&problemnotification)
}
func GetProblemNotificationHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	picker_id := params["pickerid"]
	u, err := model.GetProblemNotification(picker_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)

}
func UpdateProblemNotificationHandler(w http.ResponseWriter, r *http.Request) {}

func AddChatHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	var chat model.ChatHistory
	_ = json.NewDecoder(r.Body).Decode(&chat)
	json.NewEncoder(w).Encode(chat)

	model.AddChat(&chat)
}
func GetChatHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	picker_id := params["pickerid"]
	u, err := model.GetProblemNotification(picker_id)
	if err != nil {
		print("error in getting batch")
	}
	enableCors(&w)
	json.NewEncoder(w).Encode(u)

}
