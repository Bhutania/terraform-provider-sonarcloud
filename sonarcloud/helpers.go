package sonarcloud

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/projects"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/qualitygates"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/user_groups"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/user_tokens"
)

// changedAttrs returns a map where the keys are the names of all the attributes that were changed
// Note that the name is not the full path, but only the AttributeName of the last path step.
func changedAttrs(req tfsdk.UpdateResourceRequest, diags diag.Diagnostics) map[string]struct{} {
	diffs, err := req.Plan.Raw.Diff(req.State.Raw)
	if err != nil {
		diags.AddError(
			"Could not diff plan with state",
			"This should not happen and is an error in the provider.",
		)
	}

	changes := make(map[string]struct{})
	for _, diff := range diffs {
		steps := diff.Path.Steps()
		index := len(steps) - 1

		if !diff.Value1.Equal(*diff.Value2) {
			attr := steps[index].(tftypes.AttributeName)
			changes[string(attr)] = struct{}{}
		}
	}
	return changes
}

// findGroup returns the group with the given name if it exists in the response
func findGroup(response *user_groups.SearchResponseAll, name string) (Group, bool) {
	var result Group
	ok := false
	for _, g := range response.Groups {
		if g.Name == name {
			result = Group{
				ID:           types.String{Value: big.NewFloat(g.Id).String()},
				Default:      types.Bool{Value: g.Default},
				Description:  types.String{Value: g.Description},
				MembersCount: types.Number{Value: big.NewFloat(g.MembersCount)},
				Name:         types.String{Value: g.Name},
			}
			ok = true
			break
		}
	}
	return result, ok
}

// findGroup returns the group member with the given login if it exists in the response
func findGroupMember(response *user_groups.UsersResponseAll, group string, login string) (GroupMember, bool) {
	var result GroupMember
	ok := false
	for _, u := range response.Users {
		if u.Login == login {
			result = GroupMember{
				Group: types.String{Value: group},
				Login: types.String{Value: login},
			}
			ok = true
			break
		}
	}
	return result, ok
}

// tokenExists returns whether a token with the given name exists in the response
func tokenExists(response *user_tokens.SearchResponse, name string) bool {
	for _, t := range response.UserTokens {
		if t.Name == name {
			return true
		}
	}
	return false
}

// findProject returns the project with the given key if it exists in the response
func findProject(response *projects.SearchResponseAll, key string) (Project, bool) {
	var result Project
	ok := false
	for _, p := range response.Components {
		if p.Key == key {
			result = Project{
				ID:         types.String{Value: p.Key},
				Name:       types.String{Value: p.Name},
				Key:        types.String{Value: p.Key},
				Visibility: types.String{Value: p.Visibility},
			}
			ok = true
			break
		}
	}
	return result, ok
}

// findQualityGate returns the quality gate with the given name if it exists in a response
func findQualityGate(response *qualitygates.ListResponse, name string) (QualityGate, bool) {
	var result QualityGate
	ok := false
	for _, q := range response.Qualitygates {
		if q.Name == name {
			result = QualityGate{
				ID:        types.String{Value: fmt.Sprintf("%d", int(q.Id))},
				GateId:    types.Float64{Value: q.Id},
				Name:      types.String{Value: q.Name},
				IsBuiltIn: types.Bool{Value: q.IsBuiltIn},
				IsDefault: types.Bool{Value: q.IsDefault},
			}
			for _, c := range q.Conditions {
				result.Conditions = append(result.Conditions, Condition{
					Error:  types.String{Value: c.Error},
					ID:     types.Float64{Value: c.Id},
					Metric: types.String{Value: c.Metric},
					Op:     types.String{Value: c.Op},
				})
			}
			ok = true
			break
		}
	}
	return result, ok
}

// findSelection returns a Selection{} struct with the given project keys if they exist in a response
// this can be sped up using hashmaps, but I didn't feel like introducing a new dependency/taking code from somewhere.
// Ex library: https://pkg.go.dev/github.com/juliangruber/go-intersect/v2
func findSelection(response *qualitygates.SearchResponse, keys []attr.Value) (Selection, bool) {
	projectKeys := make([]attr.Value, 0)
	ok := true
	for _, k := range keys {
		ok = false
		for _, s := range response.Results {
			if k.Equal(types.String{Value: s.Key}) {
				projectKeys = append(projectKeys, types.String{Value: strings.Trim(s.Key, "\"")})
				ok = true
				break
			}
		}
		if !ok {
		  break
		}
	}
	return Selection{
		ProjectKeys: types.Set{ElemType: types.StringType, Elems: projectKeys},
	}, ok
}
