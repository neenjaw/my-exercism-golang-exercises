package tournament

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Sprintf format for printing the results once totaled
const tableFormat = "%-30v | %2v | %2v | %2v | %2v | %2v\n"

// The header for the table
var tableHeader = fmt.Sprintf(tableFormat, "Team", "MP", "W", "D", "L", "P")

type TeamName string
type GameResult int

const (
	Win GameResult = iota
	Loss
	Draw
)

// create a GameResult from a string
func makeResult(resultStr string) (GameResult, error) {
	switch resultStr {
	case "win":
		return Win, nil
	case "loss":
		return Loss, nil
	case "draw":
		return Draw, nil
	default:
		return 0, errors.New("result format error")
	}
}

func (r GameResult) invert() GameResult {
	switch r {
	case Win:
		return Loss
	case Loss:
		return Win
	default:
		return Draw
	}
}

type TeamRecord struct {
	name   TeamName
	played int
	won    int
	lost   int
	drawn  int
	points int
}

func (team *TeamRecord) annotateResult(result GameResult) {
	team.played += 1

	switch result {
	case Win:
		team.won += 1
		team.points += 3
	case Draw:
		team.drawn += 1
		team.points += 1
	case Loss:
		team.lost += 1
	}
}

func (t TeamRecord) String() string {
	return fmt.Sprintf(tableFormat, t.name, t.played, t.won, t.drawn, t.lost, t.points)
}

// for sorting the slice of TeamRecord
type ByPointsAndName []TeamRecord

func (a ByPointsAndName) Len() int      { return len(a) }
func (a ByPointsAndName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPointsAndName) Less(i, j int) bool {
	// order by points
	if a[i].points > a[j].points {
		return true
	}
	if a[i].points < a[j].points {
		return false
	}
	// if points are equal, then order by name
	return a[i].name < a[j].name
}

func Tally(reader io.Reader, writer io.Writer) error {
	teamMap := make(map[TeamName]*TeamRecord)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		homeTeam, awayTeam, result, err := splitLine(line)
		if err != nil {
			return err
		}

		// annotate the home team
		team := getTeamRecord(teamMap, homeTeam)
		team.annotateResult(result)

		// annotate the away team
		team = getTeamRecord(teamMap, awayTeam)
		team.annotateResult(result.invert())
	}

	// Sort the team records
	teams := make([]TeamRecord, 0, len(teamMap))
	for _, team := range teamMap {
		teams = append(teams, *team)
	}
	sort.Sort(ByPointsAndName(teams))

	// output the results to a buffer
	var buffer bytes.Buffer
	buffer.WriteString(tableHeader)
	for _, team := range teams {
		buffer.WriteString(team.String())
	}

	writer.Write(buffer.Bytes())
	return nil
}

func splitLine(line string) (TeamName, TeamName, GameResult, error) {
	split := strings.Split(line, ";")
	if len(split) != 3 {
		return "", "", 0, errors.New("improper format")
	}
	if split[0] == split[1] {
		return "", "", 0, errors.New("team names can't match")
	}
	result, err := makeResult(split[2])
	if err != nil {
		return "", "", 0, err
	}
	return TeamName(split[0]), TeamName(split[1]), result, nil
}

func getTeamRecord(teamRecordMap map[TeamName]*TeamRecord, name TeamName) *TeamRecord {
	team, ok := teamRecordMap[name]
	if !ok {
		team = &TeamRecord{name: name}
		teamRecordMap[name] = team
	}
	return team
}
