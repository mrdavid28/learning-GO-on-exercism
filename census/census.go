// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{name, age, address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
/* 2. Validate required information

Residents may be reluctant to provide personal data to census workers.
In those cases it's necessary to determine if the resident provided
the required information to be counted in the census.

In order to be counted, a resident must provide a non-zero value for their name
and an address that contains a non-zero value for the `street` key.
All other information, such as the resident's age, is optional.
 Implement the `HasRequiredInfo` method that returns a boolean indicating if the resident
 has provided the required information.

```go
name := "Matthew Sanabria"
age := 0
address := make(map[string]string)

resident := NewResident(name, age, address)

resident.HasRequiredInfo()
// => false
```*/
func (r *Resident) HasRequiredInfo() bool {
	if r.Name == "" || r.Address["street"] == "" { //age is optional
		return false
	}
	return true
}

// Delete deletes a resident's information.
/*3. Delete resident information

Life moves fast and mistakes happen. A resident can move out of the city.
 A census worker can make mistakes when collecting data.
 In those cases,
  it's necessary to have the ability to delete a resident's data so they will not be counted.

Implement the `Delete` method that sets all of the fields the resident to their zero value.

```go
name := "Matthew Sanabria"
age := 29
address := map[string]string{"street": "Main St."}

resident := NewResident(name, age, address)

fmt.Println(resident)
// Output: &{Matthew Sanabria 29 map[street:Main St.]}

resident.Delete()

fmt.Println(resident)
// Output: &{ 0 map[]}*/
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	var residents_to_count int = 0
	for _, resident := range residents {
		if resident.Name != "" || resident.Age != 0 || resident.Address["street"] != "" {
			residents_to_count += 1
		}

	}
	return residents_to_count
}
