package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid ref")
	}

	// parce que nous volons <nom du cours> + suffix course en maj
	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. expected='Golang'")
	}
}


func TestCourseEmptyValue(t *testing.T){
	_, err :=New("", "Bob", "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestNameEmptyValue(t *testing.T){
	_, err :=New("Golang", "", "2018-05-31")
	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}

func TestCourseTooLong(t *testing.T) {
	course :="azertyuiiiiiiiiiiiiiiiiiiiiiiiiiiioppppprrpefffnjenfenff"
	_, err := New(course, "Bob", "2018-05-31")
	if err==nil {
		t.Errorf("Error should be returned on a too long course name" )
	}
}

func TestNameTooLong(t *testing.T) {
	name:="azertyuiiiiiiiiiiiiiiiiiiiiiiiiiiioppppprrpefffnjenfenff"
	_, err := New("name", name, "2018-05-31")
	if err==nil {
		t.Errorf("Error should be returned on a too long course name" )
	}
}