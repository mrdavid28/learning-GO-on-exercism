package electionday

import (
	"strconv"
)

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
/*
Create a function `NewVoteCounter` that accepts the number of initial votes for a candidate
and returns a pointer referring to an `int`, initialized with the given number of initial votes.

```go
var initialVotes int
initialVotes = 2

var counter *int
counter = NewVoteCounter(initialVotes)
*counter == initialVotes // true
```*/
func NewVoteCounter(initialVotes int) *int {
	var pointer_to_votes *int
	pointer_to_votes = &initialVotes
	return *&pointer_to_votes
}

// VoteCount extracts the number of votes from a counter.
/*
##2 New system will also need a way to get the number of votes from a counter.

Create a function `VoteCount` that will take a counter (`*int`) as an argument
and will return the number of votes in the counter.
If the counter is `nil`
you should assume the counter has no votes:

```go
var votes int
votes = 3

var voteCounter *int
voteCounter = &votes

VoteCount(voteCounter)
// => 3

var nilVoteCounter *int
VoteCount(nilVoteCounter)
// => 0*/
func VoteCount(counter *int) int {
	var vote_counter int = 0
	if counter == nil {
		return 0
	}
	vote_counter = *counter
	return vote_counter
}

// IncrementVoteCount increments the value in a vote counter.
/*
##3 Create a function `IncrementVoteCount` that will take a counter (`*int`)
as an argument and a number of votes, and will increment the counter by that number of votes.
You can assume the pointer passed will never be `nil`.

```go
var votes int
votes = 3

var voteCounter *int
voteCounter = &votes

IncrementVoteCount(voteCounter, 2)

votes == 5          // true
*voteCounter == 5   // true
*/
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
/*## 4. Create the election results
Create a function `NewElectionResult` that receives the name of a candidate
and their number of votes and
returns a new election result.

```go
var result *ElectionResult
result = NewElectionResult("Peter", 3)

result.Name == "Peter"  // true
result.Votes == 3       // true
```
The election result struct is already created for you and it's defined as:

```go
type ElectionResult struct {
    // Name of the candidate
    Name    string
    // Number of votes the candidate had
    Votes   int
}
```*/
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result ElectionResult = *&ElectionResult{candidateName, votes}
	return &result
}

// DisplayResult creates a message with the result to be displayed.
/* ## 5. Announce the results
The message should show the name of the new president and the votes they had, in the following format: `<candidate_name> (<votes>)`.
This is an example of such message: `"Peter (51)"`.

Create a function `DisplayResult` that will receive an `*ElectionResult` as an argument
and will return a string with the message to display.



}
```*/
func DisplayResult(result *ElectionResult) string {

	var vote_result string = result.Name + " " + "(" + strconv.Itoa(result.Votes) + ")"
	return vote_result
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
/*## 6. Vote recounting

To make sure the final results were accurate, the votes were recounted. In the recount, it was found that the number votes for some of the candidates
was off by one.

Create a function `DecrementVotesOfCandidate` that receives the final results and the name of a candidate for which you should decrement its vote count. The final results are given in the form of a `map[string]int`, where the keys are the names of the candidates and the values are its total votes.

```go
var finalResults = map[string]int{
    "Mary":  10,
    "John":  51,
}

DecrementVotesOfCandidate(finalResults, "Mary")

finalResults["Mary"]
// => 9*/
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	results[candidate] -= 1
}
