package routes

import (
	"newgo/test/service/handler"

	"github.com/gorilla/mux"
)

// all routes
func Addroutes(r *mux.Router) {

	routesCollege(r)
	routesCourse(r)
	routesReadingMaterial(r)
	routesTask(r)
	routesUserProfile(r)
	taskreleased(r)
	batch(r)
	routesProblem(r)
	routesProblemSolved(r)
	routesProblemNotification(r)
	routesChat(r)
	r.HandleFunc("/login", handler.Login).Methods("POST")
}

// reading material curd operation routes
func routesChat(r *mux.Router) {
	r.HandleFunc("/addchat/{queryid}", handler.AddChatHandler).Methods("POST")
	r.HandleFunc("/getChat/{queryid}", handler.GetChatHandler).Methods("GET")
}
func routesProblemNotification(r *mux.Router) {
	r.HandleFunc("/addproblemnotification", handler.AddProblemNotificationHandler).Methods("POST")
	r.HandleFunc("/getproblemnotification/{pickerid}", handler.GetProblemNotificationHandler).Methods("GET")
	r.HandleFunc("/updateproblemnotification/{pickerid}", handler.UpdateProblemNotificationHandler).Methods("PUT")
}
func routesProblemSolved(r *mux.Router) {
	r.HandleFunc("/addproblemsolved", handler.AddProblemSolvedHandler).Methods("POST")
	r.HandleFunc("/getproblemssolved/{problemsolved_ids}", handler.GetProblemSolvedHandler).Methods("GET")

}
func routesProblem(r *mux.Router) {
	r.HandleFunc("/addproblem", handler.AddProblemHandler).Methods("POST")
	r.HandleFunc("/getproblems/{problem_ids}", handler.GetProblemsHandler).Methods("GET")
	r.HandleFunc("/deleteproblem/{id}", handler.DeleteProblemHandler).Methods("DELETE")
	r.HandleFunc("/updateproblem/{id}", handler.UpdateProblemHandler).Methods("PUT")

}
func routesReadingMaterial(r *mux.Router) {
	r.HandleFunc("/addreadingmaterial", handler.AddReadingMaterialHandler).Methods("POST")
	r.HandleFunc("/getreadingmaterial", handler.GetReadingMaterialHandler).Methods("GET")
	r.HandleFunc("/deletereadingmaterial/{id}", handler.DeleteReadingMaterialHandler).Methods("DELETE")
	r.HandleFunc("/updatereadingmaterial/{id}", handler.UpdateReadingMaterialHandler).Methods("PUT")

}

// colleges curd operation routes
func routesCollege(r *mux.Router) {
	r.HandleFunc("/addcollege", handler.AddCollegeHandler).Methods("POST")
	r.HandleFunc("/getcollege", handler.GetCollegeHandler).Methods("GET")
	r.HandleFunc("/deletecollege/{id}", handler.DeleteCollegeHandler).Methods("DELETE")
	r.HandleFunc("/updatecollege/{id}", handler.UpdateCollegeHandler).Methods("PUT")

}

// course curd operation routes
func routesCourse(r *mux.Router) {
	r.HandleFunc("/addcourse", handler.AddCourseHandler).Methods("POST")
	r.HandleFunc("/getcourse/{course_id}", handler.GetCourseHandler).Methods("GET")
	r.HandleFunc("/deletecourse/{id}", handler.DeleteCourseHandler).Methods("DELETE")
	r.HandleFunc("/updatecourse/{id}", handler.UpdateCourseHandler).Methods("PUT")

}

// task curd operation routes
func routesTask(r *mux.Router) {
	r.HandleFunc("/addtask", handler.AddTaskHandler).Methods("POST")
	r.HandleFunc("/gettask/{id}", handler.GetTaskHandler).Methods("GET")
	r.HandleFunc("/gettasks/{ids}", handler.GetTasksHandler).Methods("GET")
	r.HandleFunc("/deletetask/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/updatetask/{id}", handler.UpdateTaskHandler).Methods("PUT")

}

// userprofile curd operation routes
func routesUserProfile(r *mux.Router) {
	r.HandleFunc("/addprofile", handler.AddUserProfileHandler).Methods("POST")
	r.HandleFunc("/updateprofile/{roll_no}", handler.UpdateUserProfileHandler).Methods("PUT")
	r.HandleFunc("/deleteprofile/{roll_no}", handler.DeleteUserProfileHandler).Methods("DELETE")
	r.HandleFunc("/getprofile/{roll_no}", handler.GetUserProfileHandler).Methods("GET")

}
func batch(r *mux.Router) {
	r.HandleFunc("/addbatch", handler.AddBatchHandler).Methods("POST")
	r.HandleFunc("/updatebatch/{batchid}", handler.UpdateBatchHandler).Methods("PUT")
	r.HandleFunc("/deletebatch/{batchid}", handler.DeleteBatchHandler).Methods("DELETE")
	r.HandleFunc("/getbatch/{batchid}", handler.GetBatchHandler).Methods("GET")

}
func taskreleased(r *mux.Router) {
	r.HandleFunc("/addtaskreleased", handler.AddTaskReleasedHandler).Methods("POST")
	r.HandleFunc("/updatetaskreleased/{batchid}", handler.UpdateTaskReleasedHandler).Methods("PUT")
	r.HandleFunc("/deletetaskreleased/{batchid}", handler.DeleteTaskReleasedHandler).Methods("DELETE")
	r.HandleFunc("/gettaskreleased/{batchid}", handler.GetTaskReleasedHandler).Methods("GET")
	r.HandleFunc("/gettaskreleasedNo/{batchid}", handler.GetTaskReleasedNoHandler).Methods("GET")

}
