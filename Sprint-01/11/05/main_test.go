package main

import "testing"

func TestWorkerSort(t *testing.T) {

	type Worker struct {
		Name            string
		Position        string
		Salary          uint
		ExperienceYears uint
	}

	tests := []struct {
		workers  []Worker
		expected []string
	}{
		{
			workers: []Worker{
				{Name: "Михаил", Position: "директор", Salary: 200, ExperienceYears: 5},
				{Name: "Игорь", Position: "зам. директора", Salary: 180, ExperienceYears: 3},
				{Name: "Николай", Position: "начальник цеха", Salary: 120, ExperienceYears: 2},
				{Name: "Андрей", Position: "мастер", Salary: 90, ExperienceYears: 10},
				{Name: "Виктор", Position: "рабочий", Salary: 80, ExperienceYears: 3},
			},
			expected: []string{
				"Михаил - 12000 - директор",
				"Андрей - 10800 - мастер",
				"Игорь - 6480 - зам. директора",
				"Николай - 2880 - начальник цеха",
				"Виктор - 2880 - рабочий",
			},
		},
	}

	for _, test := range tests {
		c := &Company{}
		for _, worker := range test.workers {
			err := c.AddWokerInfo(worker.Name, worker.Position, worker.Salary, worker.ExperienceYears)
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
		}
		result, err := c.SortWorkers()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("Expected: %v, got: %v", test.expected[i], v)
			}
		}
	}
}
