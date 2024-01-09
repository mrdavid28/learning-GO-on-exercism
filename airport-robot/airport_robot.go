package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(visitors_name string) string
}

type Italian struct {
}

/*


## 2. Implement Italian

Now your job is to make the robot work for people that scan Italian passports.

For that, create a struct `Italian` and implement the two methods that are needed
for the struct to fulfill the `Greeter` interface you set up in task 1.
You can greet someone in Italian with `"Ciao {name}!"`.

*/

func (Italian_Language Italian) LanguageName() string {
	return " Italian"
}

func (Italian_Language Italian) Greet(name string) string {
	var greet_in_Italian string = " Ciao " + name + "!"
	return greet_in_Italian
}

type Portuguese struct {
}

func (Portuguese_Langueage Portuguese) LanguageName() string {
	return " Portuguese"
}

func (Portuguese_Language Portuguese) Greet(name string) string {
	var greet_in_Portuguese string = " Olá " + name + "!"
	return greet_in_Portuguese
}

func SayHello(visitors_name string, greet Greeter) string {
	var final_input string = "I can speak" + greet.LanguageName() + ":" + greet.Greet(visitors_name)
	return final_input
}
