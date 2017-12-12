package repo

import (
	"github.com/urfave/cli"
	"fmt"
	"text/template"
	"os"
	"github.com/gboddin/ghcli/lib"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var ListCmd = cli.Command{
	Name:  "org-list",
	ShortName: "ols",
	Usage: "List repositories by organisation",
	Action: List,
	Flags: []cli.Flag{
		cli.StringFlag{Name: "type", Usage:"Repo type to list"},
	},
}

func List(c *cli.Context) error {
	client := lib.GetGithubClient(c)

	org := c.Args().First()

	if org == "" {
		fmt.Println("Org required")
		cli.HelpPrinter(os.Stdout, "",nil)
		os.Exit(1);
	}

	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 50},
	}

	tmpl, err := template.New("_").Parse(tmplRepoList)
	if(err != nil) {
		fmt.Printf("%+v\n", err)
	}
	for {
		Repos, resp, err := client.Repositories.ListByOrg(oauth2.NoContext, org, opt)
		if(err != nil) {
			fmt.Printf("%+v\n", err)
		}
		for _, Repo := range Repos {
			tmpl.Execute(os.Stdout, Repo)
		}
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	return nil
}

var tmplRepoList = "Name: {{ .FullName }}\n"