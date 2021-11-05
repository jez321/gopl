// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/jez321/gopl/ch04/4.10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	sort.Slice(result.Items, func(i, j int) bool {
		return result.Items[i].CreatedAt.After(result.Items[j].CreatedAt)
	})

	newIssues := []*github.Issue{}
	recentIssues := []*github.Issue{}
	oldIssues := []*github.Issue{}

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(time.Now().AddDate(0, -1, 0)):
			newIssues = append(newIssues, item)
		case item.CreatedAt.After(time.Now().AddDate(-1, 0, 0)):
			recentIssues = append(recentIssues, item)
		default:
			oldIssues = append(oldIssues, item)
		}
	}

	printCategory("New issues", newIssues)
	printCategory("Recent issues", recentIssues)
	printCategory("Old issues", oldIssues)
}

func printCategory(category string, issues []*github.Issue) {
	fmt.Printf("%s:", category)
	if len(issues) == 0 {
		fmt.Printf(" None\n")
	}
	fmt.Println()
	for _, item := range issues {
		fmt.Printf("#%-6d %10.10s %9.9s %.55s\n",
			item.Number, item.CreatedAt.Format("2006-01-02 03:04"), item.User.Login, item.Title)
	}
}
