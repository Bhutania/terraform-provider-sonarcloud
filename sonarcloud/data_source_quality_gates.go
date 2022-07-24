package sonarcloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/reinoudk/go-sonarcloud/sonarcloud/qualitygates"
)

type dataSourceQualityGateType struct{}

func (d dataSourceQualityGateType) GetSchema(__ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "This data source retrieves a Quality Gate for the configured organization.",
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:          types.StringType,
				Attributes:    nil,
				Description:   "The index of the Quality Gate",
				Optional:      false,
				Computed:      true,
				Sensitive:     false,
				Validators:    []tfsdk.AttributeValidator{},
				PlanModifiers: []tfsdk.AttributePlanModifier{},
			},
			"qualityGates": {
				Description: "A quality gate",
				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
					"id": {
						Type:          types.Float64Type,
						Attributes:    nil,
						Description:   "Id created by SonarCloud",
						Required:      false,
						Optional:      false,
						Computed:      true,
						Sensitive:     false,
						Validators:    []tfsdk.AttributeValidator{},
						PlanModifiers: []tfsdk.AttributePlanModifier{},
					},
					"name": {
						Type:          types.StringType,
						Attributes:    nil,
						Description:   "Name of the Quality Gate",
						Required:      true,
						Optional:      false,
						Computed:      false,
						Sensitive:     false,
						Validators:    []tfsdk.AttributeValidator{},
						PlanModifiers: []tfsdk.AttributePlanModifier{},
					},
					"isDefault": {
						Type:          types.BoolType,
						Attributes:    nil,
						Description:   "Is this the default Quality gate for this project?",
						Required:      false,
						Optional:      true,
						Computed:      false,
						Sensitive:     false,
						Validators:    []tfsdk.AttributeValidator{},
						PlanModifiers: []tfsdk.AttributePlanModifier{},
					},
					"isBuiltIn": {
						Type:          types.BoolType,
						Attributes:    nil,
						Description:   "Is this Quality gate built in?",
						Required:      false,
						Optional:      true,
						Computed:      false,
						Sensitive:     false,
						Validators:    []tfsdk.AttributeValidator{},
						PlanModifiers: []tfsdk.AttributePlanModifier{},
					},
					"actions": {
						Description:   "What actions can be performed on this Quality Gate",
						Required:      false,
						Optional:      true,
						Computed:      true,
						Sensitive:     false,
						Validators:    []tfsdk.AttributeValidator{},
						PlanModifiers: []tfsdk.AttributePlanModifier{},
						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{
							"rename": {
								Type:          types.BoolType,
								Description:   "Whether this object can be renamed",
								Required:      false,
								Optional:      true,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"setAsDefault": {
								Type:          types.BoolType,
								Description:   "Whether this object can be set as Default",
								Required:      false,
								Optional:      true,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"copy": {
								Type:          types.BoolType,
								Description:   "Whether this object can be copied",
								Required:      false,
								Optional:      true,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"associateProjects": {
								Type:          types.BoolType,
								Description:   "Whether this object can be associated with Projects",
								Required:      false,
								Optional:      true,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"delete": {
								Type:          types.BoolType,
								Description:   "Whether this object can be deleted",
								Required:      false,
								Optional:      true,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"manageConditions": {
								Type:          types.BoolType,
								Description:   "Whether this object can be managed",
								Required:      false,
								Optional:      true,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
						}),
					},
					"conditions": {
						Computed:    true,
						Optional:    true,
						Description: "The conditions of this quality gate.",
						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{
							"id": {
								Type:          types.Float64Type,
								Description:   "ID of the Condition.",
								Required:      false,
								Optional:      false,
								Computed:      true,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"metric": {
								Type:        types.StringType,
								Attributes:  nil,
								Description: "The metric on which the condition is based. Must be one of: https://docs.sonarqube.org/latest/user-guide/metric-definitions/",
								Required:    true,
								Optional:    false,
								Computed:    false,
								Sensitive:   false,
								Validators: []tfsdk.AttributeValidator{
									allowedOptions("security_rating", "ncloc_language_distribution", "test_execution_time", "statements", "lines_to_cover", "quality_gate_details", "new_reliabillity_remediation_effort", "tests", "security_review_rating", "new_xxx_violations", "conditions_by_line", "new_violations", "ncloc", "duplicated_lines", "test_failures", "test_errors", "reopened_issues", "new_vulnerabilities", "duplicated_lines_density", "test_success_density", "sqale_debt_ratio", "security_hotspots_reviewed", "security_remediation_effort", "covered_conditions_by_line", "classes", "sqale_rating", "xxx_violations", "true_positive_issues", "violations", "new_security_review_rating", "new_security_remediation_effort", "vulnerabillities", "new_uncovered_conditions", "files", "branch_coverage_hits_data", "uncovered_lines", "comment_lines_density", "new_uncovered_lines", "complexty", "cognitive_complexity", "uncovered_conditions", "functions", "new_technical_debt", "new_coverage", "coverage", "new_branch_coverage", "confirmed_issues", "reliabillity_remediation_effort", "projects", "coverage_line_hits_data", "code_smells", "directories", "lines", "bugs", "line_coverage", "new_line_coverage", "reliability_rating", "duplicated_blocks", "branch_coverage", "new_code_smells", "new_sqale_debt_ratio", "open_issues", "sqale_index", "new_lines_to_cover", "comment_lines", "skipped_tests"),
								},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"op": {
								Type:        types.StringType,
								Attributes:  nil,
								Description: "Operation on which the metric is evaluated must be either: LT, GT",
								Required:    false,
								Optional:    true,
								Computed:    false,
								Sensitive:   false,
								Validators: []tfsdk.AttributeValidator{
									allowedOptions("LT", "GT"),
								},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
							"error": {
								Type:          types.Float64Type,
								Attributes:    nil,
								Description:   "The value on which the condition errors.",
								Required:      true,
								Optional:      false,
								Computed:      false,
								Sensitive:     false,
								Validators:    []tfsdk.AttributeValidator{},
								PlanModifiers: []tfsdk.AttributePlanModifier{},
							},
						}),
					},
				}),
			},
		},
	}, nil
}

func (d dataSourceQualityGateType) NewDataSource(_ context.Context, p tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return dataSourceQualityGate{
		p: *(p.(*provider)),
	}, nil
}

type dataSourceQualityGate struct {
	p provider
}

func (d dataSourceQualityGate) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, resp *tfsdk.ReadDataSourceResponse) {
	var diags diag.Diagnostics

	request := qualitygates.ListRequest{}

	response, err := d.p.client.Qualitygates.List(request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Could not read the Quality Gate",
			fmt.Sprintf("The SearchAll request returned an error: %+v", err),
		)
		return
	}

	result := QualityGates{}
	allQualityGates := make([]QualityGate, len(response.Qualitygates))
	for i, qualityGate := range response.Qualitygates {
		allConditions := make([]Condition, len(qualityGate.Conditions))
		for j, condition := range qualityGate.Conditions {
			allConditions[j] = Condition{
				Error:  types.Float64{Value: condition.Error},
				ID:     types.Float64{Value: condition.Id},
				Metric: types.String{Value: condition.Metric},
				Op:     types.String{Value: condition.Op},
			}
		}
		allQualityGates[i] = QualityGate{
			ID:        types.Float64{Value: qualityGate.Id},
			IsBuiltIn: types.Bool{Value: qualityGate.IsBuiltIn},
			IsDefault: types.Bool{Value: qualityGate.IsDefault},
			Name:      types.String{Value: qualityGate.Name},
		}
		allQualityGates[i].Actions = Action{
			Copy:             types.Bool{Value: qualityGate.Actions.Copy},
			Delete:           types.Bool{Value: qualityGate.Actions.Delete},
			ManageConditions: types.Bool{Value: qualityGate.Actions.ManageConditions},
			Rename:           types.Bool{Value: qualityGate.Actions.Rename},
			SetAsDefault:     types.Bool{Value: qualityGate.Actions.SetAsDefault},
		}
		allQualityGates[i].Conditions = allConditions
	}
	result.QualityGates = allQualityGates
	result.ID = types.String{Value: d.p.organization}

	diags = resp.State.Set(ctx, result)

	resp.Diagnostics.Append(diags...)
}