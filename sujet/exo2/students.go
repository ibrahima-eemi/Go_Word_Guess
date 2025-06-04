package students

import (
	"errors"
	"fmt"
	"io"
	"sort"
)

// 1. Structure Student
type Student struct {
	Name  string
	Age   int
	Grade float64
}

// 2. Structure StudentList
type StudentList struct {
	students []Student
}

// 3. Fonction NewStudent
func NewStudent(name string, age int, grade float64) (*Student, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if age < 1 || age > 99 {
		return nil, errors.New("age must be between 1 and 99")
	}
	if grade < 0 || grade > 20 {
		return nil, errors.New("grade must be between 0 and 20")
	}
	return &Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}, nil
}

// 4. Méthodes de StudentList

// AddStudents ajoute un ou plusieurs étudiants
func (sl *StudentList) AddStudents(students ...Student) {
	sl.students = append(sl.students, students...)
}

// RemoveStudent retire les étudiants par nom
func (sl *StudentList) RemoveStudent(name string) {
	filtered := sl.students[:0]
	for _, s := range sl.students {
		if s.Name != name {
			filtered = append(filtered, s)
		}
	}
	sl.students = filtered
}

// Sort trie les étudiants par note décroissante
func (sl *StudentList) Sort() StudentList {
	sorted := make([]Student, len(sl.students))
	copy(sorted, sl.students)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Grade > sorted[j].Grade
	})

	return StudentList{students: sorted}
}

// Print écrit les étudiants dans le writer
func (sl *StudentList) Print(out io.Writer) {
	for _, s := range sl.students {
		fmt.Fprintf(out, "%s (%d): %.1f\n", s.Name, s.Age, s.Grade)

	}
}
