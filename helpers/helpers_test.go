package helpers_test

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"

	"github.com/followme1987/wcwhen/fakes"
	"github.com/followme1987/wcwhen/helpers"
	"github.com/followme1987/wcwhen/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/urfave/cli"
)

var _ = Describe("helpers", func() {
	Describe("Get Matches from API server", func() {
		It("should return matches from API server", func() {
			fake := fakes.FakeScraper{
				GetDataFunc: func() ([]byte, error) {
					testBytes := &model.Matches{
						MatchID: 12345,
					}
					data, err := json.Marshal(testBytes)
					Expect(err).ToNot(HaveOccurred())
					return data, nil
				},
				UnmarshalDataFunc: func(data []byte) []model.Matches {
					Expect(data).ToNot(BeEmpty())
					Expect(string(data)).To(ContainSubstring("\"matchId\":12345"))
					return []model.Matches{{
						MatchID: 12345,
					}}
				},
			}
			matches, err := helpers.RetrieveMatches(fake)
			Expect(err).ToNot(HaveOccurred())
			Expect(matches).ToNot(BeEmpty())
			Expect(matches[0].MatchID).To(Equal(12345))
		})

		It("should return message if failed to retrieve matches from API server", func() {
			fake := fakes.FakeScraper{
				GetDataFunc: func() ([]byte, error) {
					return nil, errors.New("something is wrong")
				},
				UnmarshalDataFunc: func(data []byte) []model.Matches {
					Expect(data).To(BeNil())
					return nil
				},
			}
			matches, err := helpers.RetrieveMatches(fake)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("something is wrong"))
			Expect(matches).To(BeNil())
		})
	})

	Describe("Get fixtures by team name", func() {

		It("should return fixtures for the given valid team", func() {
			set := flag.NewFlagSet("team", 0)
			set.String("name", "ireland", "doc")
			c := cli.NewContext(nil, set, nil)
			actual := helpers.RetrieveFixturesByTeam(c)
			Expect(actual).To(ContainSubstring("Scotland"))
			Expect(actual).To(ContainSubstring("Japan"))
			Expect(actual).To(ContainSubstring("Russia"))
			Expect(actual).To(ContainSubstring("Samoa"))
		})

		It("should return message if team name is not provided", func() {
			set := flag.NewFlagSet("team", 0)
			set.String("name", "", "doc")
			c := cli.NewContext(nil, set, nil)
			actual := helpers.RetrieveFixturesByTeam(c)
			Expect(actual).To(ContainSubstring("Please provide team name"))
		})

		It("should return message if team does not exist in any group", func() {
			set := flag.NewFlagSet("team", 0)
			team := "does not exist team"
			set.String("name", team, "doc")
			c := cli.NewContext(nil, set, nil)
			actual := helpers.RetrieveFixturesByTeam(c)
			expect := fmt.Sprintf("Sorry, there is no game for %s", team)
			Expect(actual).To(ContainSubstring(expect))
		})
	})

	Describe("Get teams by group Name", func() {

		It("should return teams with short name for the given group name", func() {
			set := flag.NewFlagSet("group", 0)
			set.String("name", "A", "doc")
			c := cli.NewContext(nil, set, nil)
			actual := helpers.RetrieveTeamsByGroupName(c)
			Expect(actual).To(ContainSubstring("Ireland"))
			Expect(actual).To(ContainSubstring("Scotland"))
			Expect(actual).To(ContainSubstring("Japan"))
			Expect(actual).To(ContainSubstring("Russia"))
			Expect(actual).To(ContainSubstring("Samoa"))
		})

		It("should return message if group name is not provided", func() {
			set := flag.NewFlagSet("group", 0)
			set.String("name", "", "doc")
			c := cli.NewContext(nil, set, nil)
			actual := helpers.RetrieveTeamsByGroupName(c)
			Expect(actual).To(Equal("Please provide group name"))
		})

		It("should return message if group does not exist in any match", func() {
			set := flag.NewFlagSet("group", 0)
			grp := "does not exist group"
			set.String("name", grp, "doc")
			c := cli.NewContext(nil, set, nil)
			actual := helpers.RetrieveTeamsByGroupName(c)
			Expect(actual).To(ContainSubstring(fmt.Sprintf("Sorry, group name %s is invalid", grp)))
		})
	})
})
