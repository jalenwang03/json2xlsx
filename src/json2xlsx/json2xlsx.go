package json2xlsx

import (
	"bufio"
	"encoding/json"
	"fmt"
	xlsx "github.com/tealeg/xlsx"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	_ "reflect"
	"strconv"
	"strings"
	"time"
)

var (
	file        *xlsx.File
	sheet       *xlsx.Sheet
	row         *xlsx.Row
	cell        *xlsx.Cell
	err         error
	curr_path   string
	users       []user_interface
	usersMap    map[string]user_interface
	projects    []project_interface
	ProjectsMap map[int]project_interface
	drives      []drivers_interface
	DrivesMap   map[string]drivers_interface
	tasksMap    map[string][]task_interface
	filelist    []string
	inputReader *bufio.Reader
)

func init() {
	usersMap = make(map[string]user_interface)
	ProjectsMap = make(map[int]project_interface)
	tasksMap = make(map[string][]task_interface)
	DrivesMap = make(map[string]drivers_interface)
}

func Do() {
	curr_path = getcurrentdir()
	// fmt.Println(FileReader(curr_path + "\\users.json"))
	log.Printf("载入用户列表。。。\n")
	loadUsers(curr_path + "\\users.json")
	log.Printf("用户列表载入完成\n")
	// fmt.Println()
	loadProjects(curr_path + "\\projects.json")
	loadDrivers(curr_path + "\\drives.json")
	log.Println("载入Drivers...")
	log.Println("Drivers载入完成，共" + strconv.Itoa(WriteDrivers()) + "个Driver")
	log.Println("载入项目...")
	log.Println("项目载入完成，共" + strconv.Itoa(WritePrijects()) + "个项目")
	log.Printf("载入任务。。。\n")
	getTasks(curr_path + "\\task")
	log.Printf("任务载入完成\n")
	log.Printf("写入用户。。。\n")
	a := WriteUser()
	log.Printf("用户写入完成，共%d个用户\n", a)
	log.Printf("写入任务。。。\n")
	a = WriteTasks()
	log.Printf("任务写入完成，共%d个项目，共%d个任务\n", len(tasksMap), a)
	fmt.Println("回车退出")
	inputReader = bufio.NewReader(os.Stdin)
	_, _ = inputReader.ReadString('\n')
}
func loadUsers(filename string) {
	data := (FileReader(filename))

	json.Unmarshal([]byte(data), &users)
	for i := 0; i < len(users); i++ {
		// fmt.Println(users[i])
		usersMap[users[i].Uid] = users[i]
	}
}
func getUser(uid string) string {
	u := usersMap[uid]
	return u.showDisplay_name()
}

func WriteUser() int {
	a := 0
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Users")
	line := []string{"Name", "Display_name", "Desc", "Email", "Created_at", "Updated_at", "Role", "status"}
	WriteLine(sheet, line)
	if err != nil {
		log.Printf(err.Error())
	}
	for _, user := range usersMap {
		a++
		// fmt.Println(reflect.TypeOf(user))
		line = []string{user.Name, user.Display_name, user.Desc, user.Email, parseDate(int64(user.Created_at)), parseDate(int64(user.Updated_at)), user.Role, user.Status}
		WriteLine(sheet, line)
	}
	// fmt.Println("saving files")
	err = file.Save("users.xlsx")
	if err != nil {
		log.Println(err.Error())
	}
	return a
}
func (u *user_interface) showDisplay_name() string {
	return u.Display_name
}

func loadProjects(filename string) {
	data := (FileReader(filename))
	json.Unmarshal([]byte(data), &projects)
	for i := 0; i < len(projects); i++ {
		// fmt.Println(users[i])
		ProjectsMap[i] = projects[i]
	}
}

// func getProject(team string) string {
// 	return ProjectsMap[team].description
// }
func WritePrijects() int {
	b := 0
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Projects")
	line := []string{"name", "team", "description", "visibility", "permission", "created_at", "created_by", "updated_at", "updated_by", "status_history", "entries", "color", "members", "is_deleted", "is_archived", "archived_at"}
	WriteLine(sheet, line)
	for _, project := range ProjectsMap {
		str := ""
		for a := 0; a < len(project.Mambers); a++ {
			str = str + "  " + getUser(project.Mambers[a].Uid) + "  " + strconv.Itoa(project.Mambers[a].Permission)
		}
		b++
		line = []string{project.Name, project.Team, project.Description, strconv.Itoa(project.Visibility), strconv.Itoa(project.Permission), parseDate(int64(project.Created_at)), getUser(project.Created_by), parseDate(int64(project.Updated_at)), getUser(project.Updated_by), project.Status_history, fmt.Sprint(project.Entries), project.Color, str, strconv.Itoa(project.Is_deleted), strconv.Itoa(project.Is_archived), parseDate(int64(project.Archived_at))}
		WriteLine(sheet, line)
	}
	err = file.Save("Projects.xlsx")
	if err != nil {
		log.Println(err.Error())
	}
	return b
}
func getGroups() {

}
func loadDrivers(filename string) {
	data := (FileReader(filename))

	json.Unmarshal([]byte(data), &drives)
	for i := 0; i < len(drives); i++ {
		// fmt.Println(users[i])
		DrivesMap[drives[i].Id] = drives[i]
	}
}
func WriteDrivers() int {
	a := 0
	file = xlsx.NewFile()
	sheet, _ = file.AddSheet("Drivers")
	line := []string{"id", "type", "title", "created_at", "created_by", "updated_at", "updated_by", "scope", "tags", "addition"}
	WriteLine(sheet, line)
	for _, driver := range DrivesMap {
		scopemember := ""
		for s := 0; s < len(driver.Scope.Members); s++ {
			scopemember = scopemember + "  " + driver.Scope.Members[s].ID + "  " + getUser(driver.Scope.Members[s].Uid) + "  " + strconv.Itoa(driver.Scope.Members[s].Permission)
		}
		line = []string{driver.Id, driver.Drivers_type, driver.Title, parseDate(int64(driver.Created_at)), getUser(driver.Created_by), parseDate(int64(driver.Updated_at)), getUser(driver.Updated_by), scopemember, fmt.Sprint(driver.Tags), driver.Addition.Ext + "  " + fmt.Sprint(driver.Addition.Size) + "   " + driver.Addition.Path}
		WriteLine(sheet, line)
		a++
	}
	err = file.Save("Drivers.xlsx")
	if err != nil {
		log.Println(err.Error())
	}
	return a
}
func getTasks(path string) {
	i := 0
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		path = strings.Replace(path, "./", "", -1)
		if strings.HasSuffix(path, "json") {
			// fmt.Println(path)
			filelist = append(filelist, path)
			// fmt.Println(i)
			i++
		}
		return nil
	})
	if err != nil {
		log.Printf("文件查找失败失败 %v\n", err)
	}
	for i := 0; i < len(filelist); i++ {
		var tasks []task_interface
		data := (FileReader(filelist[i]))
		json.Unmarshal([]byte(data), &tasks)
		fl := strings.Split(filelist[i], "\\")
		filename := fl[len(fl)-1]
		tasksMap[filename[0:len(filename)-5]] = tasks
	}

}

func WriteTasks() int {
	a := 0
	file = xlsx.NewFile()
	for filename, tasklist := range tasksMap {
		b := 0
		sheet, err = file.AddSheet(filename)
		line := []string{"id", "identifier", "title", "description", "assignment", "watchers", "due_date", "attachments", "completion", "visibility", "extensions", "is_archived", "priority", "is_deleted", "created_at", "created_by", "updated_at", "updated_by", "is_cascading_deleted", "is_cascading_archived", "comments"}
		WriteLine(sheet, line)
		for j := 0; j < len(tasklist); j++ {
			a++
			b++
			line = []string{}
			task := tasklist[j]
			line = append(line, task.Id)
			line = append(line, strconv.Itoa(task.Identifier))
			line = append(line, task.Title)
			line = append(line, task.Description)
			line = append(line, task.Assignment)
			line = append(line, task.Watchers)
			line = append(line, parseDate(int64(task.Due_date.Due_date)))

			var str string
			for x := 0; x < len(task.Attachments); x++ {
				str = str + " " + task.Attachments[x].Url
			}
			line = append(line, str)
			line = append(line, strconv.Itoa(task.Completion.Is_completed))
			line = append(line, task.Visibility)
			line = append(line, fmt.Sprintln(task.Extensions))
			line = append(line, strconv.Itoa(task.Is_archived))
			line = append(line, task.Priority)
			line = append(line, strconv.Itoa(task.Is_deleted))
			line = append(line, parseDate(int64(task.Created_at)))
			line = append(line, getUser(task.Created_by))
			line = append(line, parseDate(int64(task.Updated_at)))
			line = append(line, getUser(task.Updated_by))
			line = append(line, strconv.Itoa(task.Is_cascading_deleted))
			line = append(line, strconv.Itoa(task.Is_cascading_archived))
			line = append(line, fmt.Sprintln(task.Comments))
			WriteLine(sheet, line)
		}
	}
	err = file.Save("Tasks.xlsx")
	if err != nil {
		log.Println(err.Error())
	}
	return a
}
func getMessages() {

}
func getcurrentdir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	return dir
}
func FileReader(filename string) string {
	// var returnline string
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	returnline, _ := ioutil.ReadAll(f)
	return string(returnline)
}
func parseDate(a int64) string {
	// var b int64
	// b = 1457943446
	if a > 1000000000 {
		baseTime := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
		date := baseTime.Add(time.Duration(a) * time.Second)
		// fmt.Println(date)
		return date.String()
	} else {
		return strconv.Itoa(int(a))
	}
}
func WriteLine(sheet *xlsx.Sheet, line []string) {
	row = sheet.AddRow()
	for i := 0; i < len(line); i++ {
		cell = row.AddCell()
		cell.Value = line[i]
	}
}
