package main

import "encoding/xml"

type GetAllProjects struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllProjects"`
}

type GetAllProjectsResponse struct {
	XMLName              xml.Name                  `xml:"http://replicon.com/ GetAllProjectsResponse"`
	GetAllProjectsResult *ArrayOfProjectReference1 `xml:"http://replicon.com/ GetAllProjectsResult"`
}

type GetProjectDetails struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetProjectDetails"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetProjectDetailsResponse struct {
	XMLName                 xml.Name         `xml:"http://replicon.com/ GetProjectDetailsResponse"`
	GetProjectDetailsResult *ProjectDetails1 `xml:"http://replicon.com/ GetProjectDetailsResult"`
}

type BulkGetProjectDetails struct {
	XMLName     xml.Name `xml:"http://replicon.com/ BulkGetProjectDetails"`
	ProjectUris *string  `xml:"http://replicon.com/ projectUris"`
}

type BulkGetProjectDetailsResponse struct {
	XMLName                     xml.Name                `xml:"http://replicon.com/ BulkGetProjectDetailsResponse"`
	BulkGetProjectDetailsResult *ArrayOfProjectDetails1 `xml:"http://replicon.com/ BulkGetProjectDetailsResult"`
}

type CreateNewDraft struct {
	XMLName xml.Name `xml:"http://replicon.com/ CreateNewDraft"`
}

type CreateNewDraftResponse struct {
	XMLName              xml.Name `xml:"http://replicon.com/ CreateNewDraftResponse"`
	CreateNewDraftResult *string  `xml:"http://replicon.com/ CreateNewDraftResult"`
}

type CreateEditDraft struct {
	XMLName    xml.Name `xml:"http://replicon.com/ CreateEditDraft"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type CreateEditDraftResponse struct {
	XMLName               xml.Name `xml:"http://replicon.com/ CreateEditDraftResponse"`
	CreateEditDraftResult *string  `xml:"http://replicon.com/ CreateEditDraftResult"`
}

type UpdateName struct {
	XMLName    xml.Name `xml:"http://replicon.com/ UpdateName"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	Name       *string  `xml:"http://replicon.com/ name"`
}

type UpdateNameResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateNameResponse"`
}

type UpdateCode struct {
	XMLName    xml.Name `xml:"http://replicon.com/ UpdateCode"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	Code       *string  `xml:"http://replicon.com/ code"`
}

type UpdateCodeResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateCodeResponse"`
}

type UpdateTimeEntryDateRange struct {
	XMLName    xml.Name             `xml:"http://replicon.com/ UpdateTimeEntryDateRange"`
	ProjectUri *string              `xml:"http://replicon.com/ projectUri"`
	DateRange  *DateRangeParameter1 `xml:"http://replicon.com/ dateRange"`
}

type UpdateTimeEntryDateRangeResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateTimeEntryDateRangeResponse"`
}

type ApplyNewClient struct {
	XMLName    xml.Name `xml:"http://replicon.com/ ApplyNewClient"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	ClientUri  *string  `xml:"http://replicon.com/ clientUri"`
}

type ApplyNewClientResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ ApplyNewClientResponse"`
}

type UpdateProgram struct {
	XMLName    xml.Name `xml:"http://replicon.com/ UpdateProgram"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	ProgramUri *string  `xml:"http://replicon.com/ programUri"`
}

type UpdateProgramResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProgramResponse"`
}

type UpdateProjectLeader struct {
	XMLName    xml.Name `xml:"http://replicon.com/ UpdateProjectLeader"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	UserUri    *string  `xml:"http://replicon.com/ userUri"`
}

type UpdateProjectLeaderResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectLeaderResponse"`
}

type UpdateProjectLeaderApprovalIsRequired struct {
	XMLName    xml.Name `xml:"http://replicon.com/ UpdateProjectLeaderApprovalIsRequired"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	IsRequired *string  `xml:"http://replicon.com/ isRequired"`
}

type UpdateProjectLeaderApprovalIsRequiredResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectLeaderApprovalIsRequiredResponse"`
}

type UpdatePercentComplete struct {
	XMLName         xml.Name `xml:"http://replicon.com/ UpdatePercentComplete"`
	ProjectUri      *string  `xml:"http://replicon.com/ projectUri"`
	PercentComplete *int32   `xml:"http://replicon.com/ percentComplete"`
}

type UpdatePercentCompleteResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdatePercentCompleteResponse"`
}

type UpdateStatus struct {
	XMLName          xml.Name `xml:"http://replicon.com/ UpdateStatus"`
	ProjectUri       *string  `xml:"http://replicon.com/ projectUri"`
	ProjectStatusUri *string  `xml:"http://replicon.com/ projectStatusUri"`
}

type UpdateStatusResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateStatusResponse"`
}

type UpdateDescription struct {
	XMLName     xml.Name `xml:"http://replicon.com/ UpdateDescription"`
	ProjectUri  *string  `xml:"http://replicon.com/ projectUri"`
	Description *string  `xml:"http://replicon.com/ description"`
}

type UpdateDescriptionResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateDescriptionResponse"`
}

type UpdateAllowTimeEntryAgainstTasksOnly struct {
	XMLName                        xml.Name `xml:"http://replicon.com/ UpdateAllowTimeEntryAgainstTasksOnly"`
	ProjectUri                     *string  `xml:"http://replicon.com/ projectUri"`
	AllowTimeEntryAgainstTasksOnly *string  `xml:"http://replicon.com/ allowTimeEntryAgainstTasksOnly"`
}

type UpdateAllowTimeEntryAgainstTasksOnlyResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateAllowTimeEntryAgainstTasksOnlyResponse"`
}

type UpdateCostType struct {
	XMLName     xml.Name `xml:"http://replicon.com/ UpdateCostType"`
	ProjectUri  *string  `xml:"http://replicon.com/ projectUri"`
	CostTypeUri *string  `xml:"http://replicon.com/ costTypeUri"`
}

type UpdateCostTypeResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateCostTypeResponse"`
}

type UpdateEstimatedCost struct {
	XMLName       xml.Name         `xml:"http://replicon.com/ UpdateEstimatedCost"`
	ProjectUri    *string          `xml:"http://replicon.com/ projectUri"`
	EstimatedCost *MoneyParameter1 `xml:"http://replicon.com/ estimatedCost"`
}

type UpdateEstimatedCostResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateEstimatedCostResponse"`
}

type UpdateEstimatedHours struct {
	XMLName        xml.Name       `xml:"http://replicon.com/ UpdateEstimatedHours"`
	ProjectUri     *string        `xml:"http://replicon.com/ projectUri"`
	EstimatedHours *TaskDuration1 `xml:"http://replicon.com/ estimatedHours"`
}

type UpdateEstimatedHoursResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateEstimatedHoursResponse"`
}

type Validate struct {
	XMLName    xml.Name `xml:"http://replicon.com/ Validate"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type ValidateResponse struct {
	XMLName        xml.Name                   `xml:"http://replicon.com/ ValidateResponse"`
	ValidateResult *ValidationResultsSummary1 `xml:"http://replicon.com/ ValidateResult"`
}

type PublishDraft struct {
	XMLName  xml.Name `xml:"http://replicon.com/ PublishDraft"`
	DraftUri *string  `xml:"http://replicon.com/ draftUri"`
}

type PublishDraftResponse struct {
	XMLName            xml.Name           `xml:"http://replicon.com/ PublishDraftResponse"`
	PublishDraftResult *ProjectReference1 `xml:"http://replicon.com/ PublishDraftResult"`
}

type GetExpenseCodes struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetExpenseCodes"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetExpenseCodesResponse struct {
	XMLName               xml.Name                           `xml:"http://replicon.com/ GetExpenseCodesResponse"`
	GetExpenseCodesResult *ArrayOfProjectExpenseCodeDetails1 `xml:"http://replicon.com/ GetExpenseCodesResult"`
}

type GetExpenseCodesWhichCouldBeAllowingExpenseEntry struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetExpenseCodesWhichCouldBeAllowingExpenseEntry"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetExpenseCodesWhichCouldBeAllowingExpenseEntryResponse struct {
	XMLName                                               xml.Name                           `xml:"http://replicon.com/ GetExpenseCodesWhichCouldBeAllowingExpenseEntryResponse"`
	GetExpenseCodesWhichCouldBeAllowingExpenseEntryResult *ArrayOfProjectExpenseCodeDetails1 `xml:"http://replicon.com/ GetExpenseCodesWhichCouldBeAllowingExpenseEntryResult"`
}

type UpdateExpenseCodeAllowingExpenseEntry struct {
	XMLName        xml.Name `xml:"http://replicon.com/ UpdateExpenseCodeAllowingExpenseEntry"`
	ProjectUri     *string  `xml:"http://replicon.com/ projectUri"`
	ExpenseCodeUri *string  `xml:"http://replicon.com/ expenseCodeUri"`
	Allowed        *string  `xml:"http://replicon.com/ allowed"`
}

type UpdateExpenseCodeAllowingExpenseEntryResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateExpenseCodeAllowingExpenseEntryResponse"`
}

type PutExpenseCodesAllowingExpenseEntry struct {
	XMLName         xml.Name `xml:"http://replicon.com/ PutExpenseCodesAllowingExpenseEntry"`
	ProjectUri      *string  `xml:"http://replicon.com/ projectUri"`
	ExpenseCodeUris *string  `xml:"http://replicon.com/ expenseCodeUris"`
}

type PutExpenseCodesAllowingExpenseEntryResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ PutExpenseCodesAllowingExpenseEntryResponse"`
}

type UpdateEstimatedExpenses struct {
	XMLName    xml.Name         `xml:"http://replicon.com/ UpdateEstimatedExpenses"`
	ProjectUri *string          `xml:"http://replicon.com/ projectUri"`
	Estimate   *MoneyParameter1 `xml:"http://replicon.com/ estimate"`
}

type UpdateEstimatedExpensesResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateEstimatedExpensesResponse"`
}

type GetTimeEnteredSummary struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetTimeEnteredSummary"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetTimeEnteredSummaryResponse struct {
	XMLName                     xml.Name                    `xml:"http://replicon.com/ GetTimeEnteredSummaryResponse"`
	GetTimeEnteredSummaryResult *ProjectTimeEnteredSummary1 `xml:"http://replicon.com/ GetTimeEnteredSummaryResult"`
}

type GetBillableAmountSummary struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetBillableAmountSummary"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetBillableAmountSummaryResponse struct {
	XMLName                        xml.Name                 `xml:"http://replicon.com/ GetBillableAmountSummaryResponse"`
	GetBillableAmountSummaryResult *ProjectBillableSummary1 `xml:"http://replicon.com/ GetBillableAmountSummaryResult"`
}

type GetCostAmountSummary struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetCostAmountSummary"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetCostAmountSummaryResponse struct {
	XMLName                    xml.Name             `xml:"http://replicon.com/ GetCostAmountSummaryResponse"`
	GetCostAmountSummaryResult *ProjectCostSummary1 `xml:"http://replicon.com/ GetCostAmountSummaryResult"`
}

type GetTimeEnteredSeries struct {
	XMLName             xml.Name             `xml:"http://replicon.com/ GetTimeEnteredSeries"`
	ProjectUri          *string              `xml:"http://replicon.com/ projectUri"`
	DateRange           *DateRangeParameter1 `xml:"http://replicon.com/ dateRange"`
	PeriodResolutionUri *string              `xml:"http://replicon.com/ periodResolutionUri"`
}

type GetTimeEnteredSeriesResponse struct {
	XMLName                    xml.Name                   `xml:"http://replicon.com/ GetTimeEnteredSeriesResponse"`
	GetTimeEnteredSeriesResult *ProjectTimeEnteredSeries1 `xml:"http://replicon.com/ GetTimeEnteredSeriesResult"`
}

type GetBillableAmountSeries struct {
	XMLName             xml.Name             `xml:"http://replicon.com/ GetBillableAmountSeries"`
	ProjectUri          *string              `xml:"http://replicon.com/ projectUri"`
	DateRange           *DateRangeParameter1 `xml:"http://replicon.com/ dateRange"`
	PeriodResolutionUri *string              `xml:"http://replicon.com/ periodResolutionUri"`
}

type GetBillableAmountSeriesResponse struct {
	XMLName                       xml.Name                `xml:"http://replicon.com/ GetBillableAmountSeriesResponse"`
	GetBillableAmountSeriesResult *ProjectBillableSeries1 `xml:"http://replicon.com/ GetBillableAmountSeriesResult"`
}

type GetCostAmountSeries struct {
	XMLName             xml.Name             `xml:"http://replicon.com/ GetCostAmountSeries"`
	ProjectUri          *string              `xml:"http://replicon.com/ projectUri"`
	DateRange           *DateRangeParameter1 `xml:"http://replicon.com/ dateRange"`
	PeriodResolutionUri *string              `xml:"http://replicon.com/ periodResolutionUri"`
}

type GetCostAmountSeriesResponse struct {
	XMLName                   xml.Name            `xml:"http://replicon.com/ GetCostAmountSeriesResponse"`
	GetCostAmountSeriesResult *ProjectCostSeries1 `xml:"http://replicon.com/ GetCostAmountSeriesResult"`
}

type GetChargesByExpenseCodeSummary struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetChargesByExpenseCodeSummary"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetChargesByExpenseCodeSummaryResponse struct {
	XMLName                              xml.Name                             `xml:"http://replicon.com/ GetChargesByExpenseCodeSummaryResponse"`
	GetChargesByExpenseCodeSummaryResult *ProjectChargesByExpenseCodeSummary1 `xml:"http://replicon.com/ GetChargesByExpenseCodeSummaryResult"`
}

type GetUriFromSlug struct {
	XMLName     xml.Name `xml:"http://replicon.com/ GetUriFromSlug"`
	ProjectSlug *string  `xml:"http://replicon.com/ projectSlug"`
}

type GetUriFromSlugResponse struct {
	XMLName              xml.Name `xml:"http://replicon.com/ GetUriFromSlugResponse"`
	GetUriFromSlugResult *string  `xml:"http://replicon.com/ GetUriFromSlugResult"`
}

type GetProjectReferenceFromSlug struct {
	XMLName     xml.Name `xml:"http://replicon.com/ GetProjectReferenceFromSlug"`
	ProjectSlug *string  `xml:"http://replicon.com/ projectSlug"`
}

type GetProjectReferenceFromSlugResponse struct {
	XMLName                           xml.Name           `xml:"http://replicon.com/ GetProjectReferenceFromSlugResponse"`
	GetProjectReferenceFromSlugResult *ProjectReference1 `xml:"http://replicon.com/ GetProjectReferenceFromSlugResult"`
}

type Delete struct {
	XMLName    xml.Name `xml:"http://replicon.com/ Delete"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type DeleteResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ DeleteResponse"`
}

type BulkDelete struct {
	XMLName                    xml.Name `xml:"http://replicon.com/ BulkDelete"`
	ProjectUris                *string  `xml:"http://replicon.com/ projectUris"`
	ProjectBulkDeleteOptionUri *string  `xml:"http://replicon.com/ projectBulkDeleteOptionUri"`
}

type BulkDeleteResponse struct {
	XMLName          xml.Name                   `xml:"http://replicon.com/ BulkDeleteResponse"`
	BulkDeleteResult *ProjectBulkDeleteResults1 `xml:"http://replicon.com/ BulkDeleteResult"`
}

type GetEligibleProjectLeaders struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetEligibleProjectLeaders"`
}

type GetEligibleProjectLeadersResponse struct {
	XMLName                         xml.Name                        `xml:"http://replicon.com/ GetEligibleProjectLeadersResponse"`
	GetEligibleProjectLeadersResult *ArrayOfProjectLeaderReference1 `xml:"http://replicon.com/ GetEligibleProjectLeadersResult"`
}

type GetAllProjectLeadersAssignedToProjects struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetAllProjectLeadersAssignedToProjects"`
}

type GetAllProjectLeadersAssignedToProjectsResponse struct {
	XMLName                                      xml.Name                        `xml:"http://replicon.com/ GetAllProjectLeadersAssignedToProjectsResponse"`
	GetAllProjectLeadersAssignedToProjectsResult *ArrayOfProjectLeaderReference1 `xml:"http://replicon.com/ GetAllProjectLeadersAssignedToProjectsResult"`
}

type GetEligibleEstimationModes struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetEligibleEstimationModes"`
}

type GetEligibleEstimationModesResponse struct {
	XMLName                          xml.Name                                `xml:"http://replicon.com/ GetEligibleEstimationModesResponse"`
	GetEligibleEstimationModesResult *ArrayOfProjectEstimationModeReference1 `xml:"http://replicon.com/ GetEligibleEstimationModesResult"`
}

type UpdateEstimationMode struct {
	XMLName           xml.Name `xml:"http://replicon.com/ UpdateEstimationMode"`
	ProjectUri        *string  `xml:"http://replicon.com/ projectUri"`
	EstimationModeUri *string  `xml:"http://replicon.com/ estimationModeUri"`
}

type UpdateEstimationModeResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateEstimationModeResponse"`
}

type GetAllProjectTeamMembers struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetAllProjectTeamMembers"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type GetAllProjectTeamMembersResponse struct {
	XMLName                        xml.Name                            `xml:"http://replicon.com/ GetAllProjectTeamMembersResponse"`
	GetAllProjectTeamMembersResult *ArrayOfProjectTeamMemberReference1 `xml:"http://replicon.com/ GetAllProjectTeamMembersResult"`
}

type GetAllProjectTeamMemberDetails struct {
	XMLName    xml.Name `xml:"http://replicon.com/ GetAllProjectTeamMemberDetails"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
	AsOfDate   *Date1   `xml:"http://replicon.com/ asOfDate"`
}

type GetAllProjectTeamMemberDetailsResponse struct {
	XMLName                              xml.Name                          `xml:"http://replicon.com/ GetAllProjectTeamMemberDetailsResponse"`
	GetAllProjectTeamMemberDetailsResult *ArrayOfProjectTeamMemberDetails1 `xml:"http://replicon.com/ GetAllProjectTeamMemberDetailsResult"`
}

type GetProjectTeamMemberDetails struct {
	XMLName     xml.Name `xml:"http://replicon.com/ GetProjectTeamMemberDetails"`
	ProjectUri  *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUri *string  `xml:"http://replicon.com/ resourceUri"`
	AsOfDate    *Date1   `xml:"http://replicon.com/ asOfDate"`
}

type GetProjectTeamMemberDetailsResponse struct {
	XMLName                           xml.Name                   `xml:"http://replicon.com/ GetProjectTeamMemberDetailsResponse"`
	GetProjectTeamMemberDetailsResult *ProjectTeamMemberDetails1 `xml:"http://replicon.com/ GetProjectTeamMemberDetailsResult"`
}

type PutProjectTeamMemberAssignments struct {
	XMLName      xml.Name `xml:"http://replicon.com/ PutProjectTeamMemberAssignments"`
	ProjectUri   *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUris *string  `xml:"http://replicon.com/ resourceUris"`
}

type PutProjectTeamMemberAssignmentsResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ PutProjectTeamMemberAssignmentsResponse"`
}

type UpdateProjectTeamMemberAssignment struct {
	XMLName                              xml.Name `xml:"http://replicon.com/ UpdateProjectTeamMemberAssignment"`
	ProjectUri                           *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUri                          *string  `xml:"http://replicon.com/ resourceUri"`
	ProjectTeamMemberAssignmentOptionUri *string  `xml:"http://replicon.com/ projectTeamMemberAssignmentOptionUri"`
}

type UpdateProjectTeamMemberAssignmentResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectTeamMemberAssignmentResponse"`
}

type BulkUpdateProjectTeamMembersAssignment struct {
	XMLName                              xml.Name `xml:"http://replicon.com/ BulkUpdateProjectTeamMembersAssignment"`
	ProjectUri                           *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUri                          *string  `xml:"http://replicon.com/ resourceUri"`
	ProjectTeamMemberAssignmentOptionUri *string  `xml:"http://replicon.com/ projectTeamMemberAssignmentOptionUri"`
}

type BulkUpdateProjectTeamMembersAssignmentResponse struct {
	XMLName                                      xml.Name                                       `xml:"http://replicon.com/ BulkUpdateProjectTeamMembersAssignmentResponse"`
	BulkUpdateProjectTeamMembersAssignmentResult *ProjectTeamMemberBulkUpdateAssignmentResults1 `xml:"http://replicon.com/ BulkUpdateProjectTeamMembersAssignmentResult"`
}

type AssignResourceToProject struct {
	XMLName              xml.Name `xml:"http://replicon.com/ AssignResourceToProject"`
	ProjectUri           *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUri          *string  `xml:"http://replicon.com/ resourceUri"`
	ResourceToReplaceUri *string  `xml:"http://replicon.com/ resourceToReplaceUri"`
}

type AssignResourceToProjectResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ AssignResourceToProjectResponse"`
}

type GetProjectNameFormatForUser struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetProjectNameFormatForUser"`
	UserUri *string  `xml:"http://replicon.com/ userUri"`
}

type GetProjectNameFormatForUserResponse struct {
	XMLName                           xml.Name `xml:"http://replicon.com/ GetProjectNameFormatForUserResponse"`
	GetProjectNameFormatForUserResult *string  `xml:"http://replicon.com/ GetProjectNameFormatForUserResult"`
}

type UpdateProjectNameFormatForUser struct {
	XMLName              xml.Name `xml:"http://replicon.com/ UpdateProjectNameFormatForUser"`
	UserUri              *string  `xml:"http://replicon.com/ userUri"`
	ProjectNameFormatUri *string  `xml:"http://replicon.com/ projectNameFormatUri"`
}

type UpdateProjectNameFormatForUserResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectNameFormatForUserResponse"`
}

type GetProjectNameFormatForNewUsers struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetProjectNameFormatForNewUsers"`
}

type GetProjectNameFormatForNewUsersResponse struct {
	XMLName                               xml.Name `xml:"http://replicon.com/ GetProjectNameFormatForNewUsersResponse"`
	GetProjectNameFormatForNewUsersResult *string  `xml:"http://replicon.com/ GetProjectNameFormatForNewUsersResult"`
}

type UpdateProjectNameFormatForNewUsers struct {
	XMLName              xml.Name `xml:"http://replicon.com/ UpdateProjectNameFormatForNewUsers"`
	ProjectNameFormatUri *string  `xml:"http://replicon.com/ projectNameFormatUri"`
}

type UpdateProjectNameFormatForNewUsersResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectNameFormatForNewUsersResponse"`
}

type ChangeProjectToBeNonBillable struct {
	XMLName    xml.Name `xml:"http://replicon.com/ ChangeProjectToBeNonBillable"`
	ProjectUri *string  `xml:"http://replicon.com/ projectUri"`
}

type ChangeProjectToBeNonBillableResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ ChangeProjectToBeNonBillableResponse"`
}

type GetProjectCodeSettingsForNewProjects struct {
	XMLName xml.Name `xml:"http://replicon.com/ GetProjectCodeSettingsForNewProjects"`
}

type GetProjectCodeSettingsForNewProjectsResponse struct {
	XMLName                                    xml.Name                     `xml:"http://replicon.com/ GetProjectCodeSettingsForNewProjectsResponse"`
	GetProjectCodeSettingsForNewProjectsResult *ProjectCodeSettingsDetails1 `xml:"http://replicon.com/ GetProjectCodeSettingsForNewProjectsResult"`
}

type UpdateProjectCodeSettingForNewProjectsAsUserEntered struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectCodeSettingForNewProjectsAsUserEntered"`
}

type UpdateProjectCodeSettingForNewProjectsAsUserEnteredResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectCodeSettingForNewProjectsAsUserEnteredResponse"`
}

type UpdateProjectCodeSettingForNewProjectsAsAutoincrementingNumber struct {
	XMLName             xml.Name `xml:"http://replicon.com/ UpdateProjectCodeSettingForNewProjectsAsAutoincrementingNumber"`
	StartingProjectCode *string  `xml:"http://replicon.com/ startingProjectCode"`
}

type UpdateProjectCodeSettingForNewProjectsAsAutoincrementingNumberResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ UpdateProjectCodeSettingForNewProjectsAsAutoincrementingNumberResponse"`
}

type GetProjectActualsByTeamMember struct {
	XMLName                              xml.Name `xml:"http://replicon.com/ GetProjectActualsByTeamMember"`
	ProjectUri                           *string  `xml:"http://replicon.com/ projectUri"`
	ProjectActualsByTeamMemberOptionUris *string  `xml:"http://replicon.com/ projectActualsByTeamMemberOptionUris"`
}

type GetProjectActualsByTeamMemberResponse struct {
	XMLName                             xml.Name                            `xml:"http://replicon.com/ GetProjectActualsByTeamMemberResponse"`
	GetProjectActualsByTeamMemberResult *ProjectActualsByTeamMemberSummary1 `xml:"http://replicon.com/ GetProjectActualsByTeamMemberResult"`
}

type PutTask struct {
	XMLName xml.Name                 `xml:"http://replicon.com/ PutTask"`
	Project *ProjectTargetParameter1 `xml:"http://replicon.com/ project"`
	Task    *TaskParameter1          `xml:"http://replicon.com/ task"`
}

type PutTaskResponse struct {
	XMLName       xml.Name        `xml:"http://replicon.com/ PutTaskResponse"`
	PutTaskResult *TaskReference1 `xml:"http://replicon.com/ PutTaskResult"`
}

type PutTaskHierarchy struct {
	XMLName       xml.Name                        `xml:"http://replicon.com/ PutTaskHierarchy"`
	Project       *ProjectTargetParameter1        `xml:"http://replicon.com/ project"`
	TaskHierarchy *ArrayOfTaskHierarchyParameter1 `xml:"http://replicon.com/ taskHierarchy"`
}

type PutTaskHierarchyResponse struct {
	XMLName                xml.Name                         `xml:"http://replicon.com/ PutTaskHierarchyResponse"`
	PutTaskHierarchyResult *ArrayOfTaskHierarchyPutResults1 `xml:"http://replicon.com/ PutTaskHierarchyResult"`
}

type PutProjectInfo struct {
	XMLName     xml.Name                 `xml:"http://replicon.com/ PutProjectInfo"`
	Target      *ProjectTargetParameter1 `xml:"http://replicon.com/ target"`
	ProjectInfo *ProjectInfoParameter1   `xml:"http://replicon.com/ projectInfo"`
}

type PutProjectInfoResponse struct {
	XMLName              xml.Name           `xml:"http://replicon.com/ PutProjectInfoResponse"`
	PutProjectInfoResult *ProjectReference1 `xml:"http://replicon.com/ PutProjectInfoResult"`
}

type PutProject struct {
	XMLName xml.Name           `xml:"http://replicon.com/ PutProject"`
	Project *ProjectParameter1 `xml:"http://replicon.com/ project"`
}

type PutProjectResponse struct {
	XMLName          xml.Name           `xml:"http://replicon.com/ PutProjectResponse"`
	PutProjectResult *ProjectReference1 `xml:"http://replicon.com/ PutProjectResult"`
}

type PutTaskAssignmentsForResource struct {
	XMLName     xml.Name `xml:"http://replicon.com/ PutTaskAssignmentsForResource"`
	ProjectUri  *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUri *string  `xml:"http://replicon.com/ resourceUri"`
	TaskUris    *string  `xml:"http://replicon.com/ taskUris"`
}

type PutTaskAssignmentsForResourceResponse struct {
	XMLName xml.Name `xml:"http://replicon.com/ PutTaskAssignmentsForResourceResponse"`
}

type GetTaskAssignmentsForResource struct {
	XMLName     xml.Name `xml:"http://replicon.com/ GetTaskAssignmentsForResource"`
	ProjectUri  *string  `xml:"http://replicon.com/ projectUri"`
	ResourceUri *string  `xml:"http://replicon.com/ resourceUri"`
	AsOfDate    *Date1   `xml:"http://replicon.com/ asOfDate"`
}

type GetTaskAssignmentsForResourceResponse struct {
	XMLName                             xml.Name                               `xml:"http://replicon.com/ GetTaskAssignmentsForResourceResponse"`
	GetTaskAssignmentsForResourceResult *ArrayOfTaskResourceAssignmentDetails1 `xml:"http://replicon.com/ GetTaskAssignmentsForResourceResult"`
}

type ArrayOfProjectReference1 struct {
	ProjectReference1 []*ProjectReference1 `xml:"http://replicon.com/ ProjectReference1"`
}

type ProjectReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Slug        *string `xml:"http://replicon.com/ slug"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProjectDetails1 struct {
	BillingType                     *BillingTypeReference1             `xml:"http://replicon.com/ billingType"`
	Budget                          *BudgetDetails1                    `xml:"http://replicon.com/ budget"`
	Client                          *ClientReference1                  `xml:"http://replicon.com/ client"`
	Code                            *string                            `xml:"http://replicon.com/ code"`
	CostType                        *CostTypeDetails1                  `xml:"http://replicon.com/ costType"`
	CustomFields                    *ArrayOfCustomFieldValueDetails1   `xml:"http://replicon.com/ customFields"`
	Description                     *string                            `xml:"http://replicon.com/ description"`
	DisplayText                     *string                            `xml:"http://replicon.com/ displayText"`
	EstimatedCost                   *MoneyDetails1                     `xml:"http://replicon.com/ estimatedCost"`
	EstimatedExpenses               *MoneyDetails1                     `xml:"http://replicon.com/ estimatedExpenses"`
	EstimatedHours                  *TaskDuration1                     `xml:"http://replicon.com/ estimatedHours"`
	EstimationMode                  *ProjectEstimationModeReference1   `xml:"http://replicon.com/ estimationMode"`
	IsProjectLeaderApprovalRequired *string                            `xml:"http://replicon.com/ isProjectLeaderApprovalRequired"`
	IsTimeEntryAllowed              *string                            `xml:"http://replicon.com/ isTimeEntryAllowed"`
	Name                            *string                            `xml:"http://replicon.com/ name"`
	PercentCompleted                *int32                             `xml:"http://replicon.com/ percentCompleted"`
	Program                         *ProgramReference1                 `xml:"http://replicon.com/ program"`
	ProjectLeader                   *ProjectLeaderReference1           `xml:"http://replicon.com/ projectLeader"`
	Slug                            *string                            `xml:"http://replicon.com/ slug"`
	Status                          *ProjectStatusLabelReference1      `xml:"http://replicon.com/ status"`
	TimeAndExpenseEntryType         *TimeAndExpenseEntryTypeReference1 `xml:"http://replicon.com/ timeAndExpenseEntryType"`
	TimeEntryDateRange              *DateRangeDetails1                 `xml:"http://replicon.com/ timeEntryDateRange"`
	Uri                             *string                            `xml:"http://replicon.com/ uri"`
}

type BillingTypeReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type BudgetDetails1 struct {
	Capital     *MoneyDetails1 `xml:"http://replicon.com/ capital"`
	Operational *MoneyDetails1 `xml:"http://replicon.com/ operational"`
	Total       *MoneyDetails1 `xml:"http://replicon.com/ total"`
}

type MoneyDetails1 struct {
	Amount   *string             `xml:"http://replicon.com/ amount"`
	Currency *CurrencyReference1 `xml:"http://replicon.com/ currency"`
}

type CurrencyReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Symbol      *string `xml:"http://replicon.com/ symbol"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ClientReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Slug        *string `xml:"http://replicon.com/ slug"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type CostTypeDetails1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ArrayOfCustomFieldValueDetails1 struct {
	CustomFieldValueDetails1 []*CustomFieldValueDetails1 `xml:"http://replicon.com/ CustomFieldValueDetails1"`
}

type CustomFieldValueDetails1 struct {
	CustomField     *CustomFieldReference1     `xml:"http://replicon.com/ customField"`
	CustomFieldType *CustomFieldTypeReference1 `xml:"http://replicon.com/ customFieldType"`
	Date            *Date1                     `xml:"http://replicon.com/ date"`
	DropDownOption  *string                    `xml:"http://replicon.com/ dropDownOption"`
	Number          *string                    `xml:"http://replicon.com/ number"`
	Text            *string                    `xml:"http://replicon.com/ text"`
}

type CustomFieldReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	GroupUri    *string `xml:"http://replicon.com/ groupUri"`
	Name        *string `xml:"http://replicon.com/ name"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type CustomFieldTypeReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type Date1 struct {
	Day   *int32 `xml:"http://replicon.com/ day"`
	Month *int32 `xml:"http://replicon.com/ month"`
	Year  *int32 `xml:"http://replicon.com/ year"`
}

type TaskDuration1 struct {
	Hours   *int32 `xml:"http://replicon.com/ hours"`
	Minutes *int32 `xml:"http://replicon.com/ minutes"`
}

type ProjectEstimationModeReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProgramReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Slug        *string `xml:"http://replicon.com/ slug"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProjectLeaderReference1 struct {
	DisplayText *string         `xml:"http://replicon.com/ displayText"`
	Uri         *string         `xml:"http://replicon.com/ uri"`
	User        *UserReference1 `xml:"http://replicon.com/ user"`
}

type UserReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	LoginName   *string `xml:"http://replicon.com/ loginName"`
	Slug        *string `xml:"http://replicon.com/ slug"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProjectStatusLabelReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type TimeAndExpenseEntryTypeReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type DateRangeDetails1 struct {
	EndDate   *Date1 `xml:"http://replicon.com/ endDate"`
	StartDate *Date1 `xml:"http://replicon.com/ startDate"`
}

type UriError1 struct {
	Uri *string `xml:"http://replicon.com/ uri"`
}

type ArrayOfProjectDetails1 struct {
	ProjectDetails1 []*ProjectDetails1 `xml:"http://replicon.com/ ProjectDetails1"`
}

type ValidationError1 struct {
	Notifications *ArrayOfValidationNotificationSummary1 `xml:"http://replicon.com/ notifications"`
}

type ArrayOfValidationNotificationSummary1 struct {
	ValidationNotificationSummary1 []*ValidationNotificationSummary1 `xml:"http://replicon.com/ ValidationNotificationSummary1"`
}

type ValidationNotificationSummary1 struct {
	Context     *ValidationContextSummary1 `xml:"http://replicon.com/ context"`
	DisplayText *string                    `xml:"http://replicon.com/ displayText"`
	FailureUri  *string                    `xml:"http://replicon.com/ failureUri"`
	SeverityUri *string                    `xml:"http://replicon.com/ severityUri"`
}

type ValidationContextSummary1 struct {
	DisplayText            *string `xml:"http://replicon.com/ displayText"`
	ParameterCorrelationId *string `xml:"http://replicon.com/ parameterCorrelationId"`
	Uri                    *string `xml:"http://replicon.com/ uri"`
}

type DateRangeParameter1 struct {
	EndDate                   *Date1  `xml:"http://replicon.com/ endDate"`
	RelativeDateRangeAsOfDate *Date1  `xml:"http://replicon.com/ relativeDateRangeAsOfDate"`
	RelativeDateRangeUri      *string `xml:"http://replicon.com/ relativeDateRangeUri"`
	StartDate                 *Date1  `xml:"http://replicon.com/ startDate"`
}

type InvalidDateRangeError1 struct {
	Context                  *DateRangeParameter1 `xml:"http://replicon.com/ context"`
	DisplayText              *string              `xml:"http://replicon.com/ displayText"`
	InvalidDateRangeErrorUri *string              `xml:"http://replicon.com/ invalidDateRangeErrorUri"`
}

type ApplyNewClientToProjectError1 struct {
	AreExpensesAlreadyEnteredAgainstProject *string            `xml:"http://replicon.com/ areExpensesAlreadyEnteredAgainstProject"`
	CurrentClient                           *ClientReference1  `xml:"http://replicon.com/ currentClient"`
	IsNewClientInactive                     *string            `xml:"http://replicon.com/ isNewClientInactive"`
	IsTimeAlreadyEnteredAgainstProject      *string            `xml:"http://replicon.com/ isTimeAlreadyEnteredAgainstProject"`
	NewClient                               *ClientReference1  `xml:"http://replicon.com/ newClient"`
	Project                                 *ProjectReference1 `xml:"http://replicon.com/ project"`
}

type InvalidProjectCompletionPercentageError1 struct {
	PercentCompleteEntered *int32             `xml:"http://replicon.com/ percentCompleteEntered"`
	Project                *ProjectReference1 `xml:"http://replicon.com/ project"`
}

type MoneyParameter1 struct {
	Amount      string `xml:"http://replicon.com/ amount"`
	CurrencyUri string `xml:"http://replicon.com/ currencyUri"`
}

type ValidationResultsSummary1 struct {
	Notifications *ArrayOfValidationNotificationSummary1 `xml:"http://replicon.com/ notifications"`
}

type ArrayOfProjectExpenseCodeDetails1 struct {
	ProjectExpenseCodeDetails1 []*ProjectExpenseCodeDetails1 `xml:"http://replicon.com/ ProjectExpenseCodeDetails1"`
}

type ProjectExpenseCodeDetails1 struct {
	ExpenseCode                     *ExpenseCodeReference1 `xml:"http://replicon.com/ expenseCode"`
	IsExpenseEntryToThisCodeAllowed *string                `xml:"http://replicon.com/ isExpenseEntryToThisCodeAllowed"`
}

type ExpenseCodeReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProjectTimeEnteredSummary1 struct {
	TimeEnteredActual          *CalendarDayDuration1 `xml:"http://replicon.com/ timeEnteredActual"`
	TimeEnteredEstimatedToDate *CalendarDayDuration1 `xml:"http://replicon.com/ timeEnteredEstimatedToDate"`
	TimeEnteredTotalEstimated  *CalendarDayDuration1 `xml:"http://replicon.com/ timeEnteredTotalEstimated"`
}

type CalendarDayDuration1 struct {
	Hours   *int32 `xml:"http://replicon.com/ hours"`
	Minutes *int32 `xml:"http://replicon.com/ minutes"`
	Seconds *int32 `xml:"http://replicon.com/ seconds"`
}

type ProjectBillableSummary1 struct {
	BillableActual          *MoneyDetails1 `xml:"http://replicon.com/ billableActual"`
	BillableEstimatedToDate *MoneyDetails1 `xml:"http://replicon.com/ billableEstimatedToDate"`
	BillableTotalEstimated  *MoneyDetails1 `xml:"http://replicon.com/ billableTotalEstimated"`
	PaidTotal               *MoneyDetails1 `xml:"http://replicon.com/ paidTotal"`
	PendingTotal            *MoneyDetails1 `xml:"http://replicon.com/ pendingTotal"`
}

type ProjectCostSummary1 struct {
	CostActual          *MoneyDetails1 `xml:"http://replicon.com/ costActual"`
	CostEstimatedToDate *MoneyDetails1 `xml:"http://replicon.com/ costEstimatedToDate"`
	CostTotalEstimated  *MoneyDetails1 `xml:"http://replicon.com/ costTotalEstimated"`
}

type ProjectTimeEnteredSeries1 struct {
	DataPoints *ArrayOfProjectTimeEnteredSeriesDataPoint1 `xml:"http://replicon.com/ dataPoints"`
}

type ArrayOfProjectTimeEnteredSeriesDataPoint1 struct {
	ProjectTimeEnteredSeriesDataPoint1 []*ProjectTimeEnteredSeriesDataPoint1 `xml:"http://replicon.com/ ProjectTimeEnteredSeriesDataPoint1"`
}

type ProjectTimeEnteredSeriesDataPoint1 struct {
	Period               *PeriodDetails1       `xml:"http://replicon.com/ period"`
	TimeEnteredActual    *CalendarDayDuration1 `xml:"http://replicon.com/ timeEnteredActual"`
	TimeEnteredDelta     *CalendarDayDuration1 `xml:"http://replicon.com/ timeEnteredDelta"`
	TimeEnteredEstimated *CalendarDayDuration1 `xml:"http://replicon.com/ timeEnteredEstimated"`
}

type PeriodDetails1 struct {
	Day         *int32 `xml:"http://replicon.com/ day"`
	Month       *int32 `xml:"http://replicon.com/ month"`
	PeriodEnd   *Date1 `xml:"http://replicon.com/ periodEnd"`
	PeriodStart *Date1 `xml:"http://replicon.com/ periodStart"`
	Quarter     *int32 `xml:"http://replicon.com/ quarter"`
	Week        *int32 `xml:"http://replicon.com/ week"`
	Year        *int32 `xml:"http://replicon.com/ year"`
}

type ProjectBillableSeries1 struct {
	Currency   *CurrencyReference1                     `xml:"http://replicon.com/ currency"`
	DataPoints *ArrayOfProjectBillableSeriesDataPoint1 `xml:"http://replicon.com/ dataPoints"`
}

type ArrayOfProjectBillableSeriesDataPoint1 struct {
	ProjectBillableSeriesDataPoint1 []*ProjectBillableSeriesDataPoint1 `xml:"http://replicon.com/ ProjectBillableSeriesDataPoint1"`
}

type ProjectBillableSeriesDataPoint1 struct {
	BillableAmountActual    *string         `xml:"http://replicon.com/ billableAmountActual"`
	BillableAmountDelta     *string         `xml:"http://replicon.com/ billableAmountDelta"`
	BillableAmountEstimated *string         `xml:"http://replicon.com/ billableAmountEstimated"`
	Period                  *PeriodDetails1 `xml:"http://replicon.com/ period"`
}

type ProjectCostSeries1 struct {
	Currency   *CurrencyReference1                 `xml:"http://replicon.com/ currency"`
	DataPoints *ArrayOfProjectCostSeriesDataPoint1 `xml:"http://replicon.com/ dataPoints"`
}

type ArrayOfProjectCostSeriesDataPoint1 struct {
	ProjectCostSeriesDataPoint1 []*ProjectCostSeriesDataPoint1 `xml:"http://replicon.com/ ProjectCostSeriesDataPoint1"`
}

type ProjectCostSeriesDataPoint1 struct {
	CostAmountActual    *string         `xml:"http://replicon.com/ costAmountActual"`
	CostAmountDelta     *string         `xml:"http://replicon.com/ costAmountDelta"`
	CostAmountEstimated *string         `xml:"http://replicon.com/ costAmountEstimated"`
	Period              *PeriodDetails1 `xml:"http://replicon.com/ period"`
}

type ProjectChargesByExpenseCodeSummary1 struct {
	ExpenseCodeSummaries     *ArrayOfExpenseCodeChargeSummary1 `xml:"http://replicon.com/ expenseCodeSummaries"`
	Project                  *ProjectReference1                `xml:"http://replicon.com/ project"`
	TotalActualAmountCharged *MultiCurrencyMoneyDetails1       `xml:"http://replicon.com/ totalActualAmountCharged"`
}

type ArrayOfExpenseCodeChargeSummary1 struct {
	ExpenseCodeChargeSummary1 []*ExpenseCodeChargeSummary1 `xml:"http://replicon.com/ ExpenseCodeChargeSummary1"`
}

type ExpenseCodeChargeSummary1 struct {
	ActualAmountCharged *MultiCurrencyMoneyDetails1 `xml:"http://replicon.com/ actualAmountCharged"`
	ExpenseCode         *ExpenseCodeReference1      `xml:"http://replicon.com/ expenseCode"`
}

type MultiCurrencyMoneyDetails1 struct {
	BaseCurrencyValue         *MoneyDetails1        `xml:"http://replicon.com/ baseCurrencyValue"`
	BaseCurrencyValueAsOfDate *Date1                `xml:"http://replicon.com/ baseCurrencyValueAsOfDate"`
	MultiCurrencyValue        *ArrayOfMoneyDetails1 `xml:"http://replicon.com/ multiCurrencyValue"`
}

type ArrayOfMoneyDetails1 struct {
	MoneyDetails1 []*MoneyDetails1 `xml:"http://replicon.com/ MoneyDetails1"`
}

type SlugNotFoundError1 struct {
	Slug *string `xml:"http://replicon.com/ slug"`
}

type ProjectDeleteError1 struct {
	DisplayText *string            `xml:"http://replicon.com/ displayText"`
	Project     *ProjectReference1 `xml:"http://replicon.com/ project"`
}

type ProjectBulkDeleteResults1 struct {
	ArchivedUris *string                     `xml:"http://replicon.com/ archivedUris"`
	DeletedUris  *string                     `xml:"http://replicon.com/ deletedUris"`
	Errors       *ArrayOfProjectDeleteError1 `xml:"http://replicon.com/ errors"`
}

type ArrayOfProjectDeleteError1 struct {
	ProjectDeleteError1 []*ProjectDeleteError1 `xml:"http://replicon.com/ ProjectDeleteError1"`
}

type ArrayOfProjectLeaderReference1 struct {
	ProjectLeaderReference1 []*ProjectLeaderReference1 `xml:"http://replicon.com/ ProjectLeaderReference1"`
}

type ArrayOfProjectEstimationModeReference1 struct {
	ProjectEstimationModeReference1 []*ProjectEstimationModeReference1 `xml:"http://replicon.com/ ProjectEstimationModeReference1"`
}

type ArrayOfProjectTeamMemberReference1 struct {
	ProjectTeamMemberReference1 []*ProjectTeamMemberReference1 `xml:"http://replicon.com/ ProjectTeamMemberReference1"`
}

type ProjectTeamMemberReference1 struct {
	Project  *ProjectReference1  `xml:"http://replicon.com/ project"`
	Resource *ResourceReference1 `xml:"http://replicon.com/ resource"`
}

type ResourceReference1 struct {
	Department          *DepartmentReference1          `xml:"http://replicon.com/ department"`
	DisplayText         *string                        `xml:"http://replicon.com/ displayText"`
	ResourcePlaceholder *ResourcePlaceholderReference1 `xml:"http://replicon.com/ resourcePlaceholder"`
	ResourceType        *ResourceTypeDetails1          `xml:"http://replicon.com/ resourceType"`
	Slug                *string                        `xml:"http://replicon.com/ slug"`
	Uri                 *string                        `xml:"http://replicon.com/ uri"`
	User                *UserReference1                `xml:"http://replicon.com/ user"`
}

type DepartmentReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ResourcePlaceholderReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ResourceTypeDetails1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ArrayOfProjectTeamMemberDetails1 struct {
	ProjectTeamMemberDetails1 []*ProjectTeamMemberDetails1 `xml:"http://replicon.com/ ProjectTeamMemberDetails1"`
}

type ProjectTeamMemberDetails1 struct {
	BillingRatesAllowedForBillingTime *ArrayOfProjectBillingRateDetails1 `xml:"http://replicon.com/ billingRatesAllowedForBillingTime"`
	Project                           *ProjectReference1                 `xml:"http://replicon.com/ project"`
	Resource                          *ResourceReference1                `xml:"http://replicon.com/ resource"`
	Role                              *ProjectRoleReference1             `xml:"http://replicon.com/ role"`
}

type ArrayOfProjectBillingRateDetails1 struct {
	ProjectBillingRateDetails1 []*ProjectBillingRateDetails1 `xml:"http://replicon.com/ ProjectBillingRateDetails1"`
}

type ProjectBillingRateDetails1 struct {
	BillingRate                           *BillingRateReference1   `xml:"http://replicon.com/ billingRate"`
	EffectiveBillingRate                  *BillingRateAsOfDetails1 `xml:"http://replicon.com/ effectiveBillingRate"`
	IsAvailableForAssignmentToTeamMembers *string                  `xml:"http://replicon.com/ isAvailableForAssignmentToTeamMembers"`
}

type BillingRateReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type BillingRateAsOfDetails1 struct {
	AsOfDate                    *Date1         `xml:"http://replicon.com/ asOfDate"`
	BillingRateScheduleEntryUri *string        `xml:"http://replicon.com/ billingRateScheduleEntryUri"`
	EffectiveDate               *Date1         `xml:"http://replicon.com/ effectiveDate"`
	EndDate                     *Date1         `xml:"http://replicon.com/ endDate"`
	Value                       *MoneyDetails1 `xml:"http://replicon.com/ value"`
}

type ProjectRoleReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Name        *string `xml:"http://replicon.com/ name"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProjectTeamMemberAssignmentError1 struct {
	AssignedTasks     *ArrayOfTaskReference1       `xml:"http://replicon.com/ assignedTasks"`
	ProjectTeamMember *ProjectTeamMemberReference1 `xml:"http://replicon.com/ projectTeamMember"`
}

type ArrayOfTaskReference1 struct {
	TaskReference1 []*TaskReference1 `xml:"http://replicon.com/ TaskReference1"`
}

type TaskReference1 struct {
	DisplayText *string `xml:"http://replicon.com/ displayText"`
	Uri         *string `xml:"http://replicon.com/ uri"`
}

type ProjectTeamMemberBulkUpdateAssignmentResults1 struct {
	Errors  *ArrayOfProjectTeamMemberAssignmentError1 `xml:"http://replicon.com/ errors"`
	Updated *string                                   `xml:"http://replicon.com/ updated"`
}

type ArrayOfProjectTeamMemberAssignmentError1 struct {
	ProjectTeamMemberAssignmentError1 []*ProjectTeamMemberAssignmentError1 `xml:"http://replicon.com/ ProjectTeamMemberAssignmentError1"`
}

type AssignResourceToProjectError1 struct {
	DisplayText       *string             `xml:"http://replicon.com/ displayText"`
	FailureUri        *string             `xml:"http://replicon.com/ failureUri"`
	Project           *ProjectReference1  `xml:"http://replicon.com/ project"`
	Resource          *ResourceReference1 `xml:"http://replicon.com/ resource"`
	ResourceToReplace *ResourceReference1 `xml:"http://replicon.com/ resourceToReplace"`
}

type ProjectCodeSettingsDetails1 struct {
	NextProjectCode          *string `xml:"http://replicon.com/ nextProjectCode"`
	ProjectCodeSettingOption *string `xml:"http://replicon.com/ projectCodeSettingOption"`
}

type ProjectActualsByTeamMemberSummary1 struct {
	ActualBillingTotal  *MultiCurrencyMoneyDetails1              `xml:"http://replicon.com/ actualBillingTotal"`
	ActualCostTotal     *MultiCurrencyMoneyDetails1              `xml:"http://replicon.com/ actualCostTotal"`
	ActualHoursTotal    *CalendarDayDuration1                    `xml:"http://replicon.com/ actualHoursTotal"`
	Project             *ProjectReference1                       `xml:"http://replicon.com/ project"`
	TeamMemberSummaries *ArrayOfProjectTeamMemberActualsSummary1 `xml:"http://replicon.com/ teamMemberSummaries"`
}

type ArrayOfProjectTeamMemberActualsSummary1 struct {
	ProjectTeamMemberActualsSummary1 []*ProjectTeamMemberActualsSummary1 `xml:"http://replicon.com/ ProjectTeamMemberActualsSummary1"`
}

type ProjectTeamMemberActualsSummary1 struct {
	ActualBillingTotal        *MultiCurrencyMoneyDetails1 `xml:"http://replicon.com/ actualBillingTotal"`
	ActualCostTotal           *MultiCurrencyMoneyDetails1 `xml:"http://replicon.com/ actualCostTotal"`
	ActualHoursTotal          *CalendarDayDuration1       `xml:"http://replicon.com/ actualHoursTotal"`
	IsActiveProjectTeamMember *string                     `xml:"http://replicon.com/ isActiveProjectTeamMember"`
	Resource                  *ResourceReference1         `xml:"http://replicon.com/ resource"`
}

type ProjectTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type TaskParameter1 struct {
	AssignedResources  *ArrayOfResourceTargetParameter1   `xml:"http://replicon.com/ assignedResources"`
	Code               *string                            `xml:"http://replicon.com/ code"`
	CostTypeUri        *string                            `xml:"http://replicon.com/ costTypeUri"`
	CustomFieldValues  *ArrayOfCustomFieldValueParameter1 `xml:"http://replicon.com/ customFieldValues"`
	Description        *string                            `xml:"http://replicon.com/ description"`
	EstimatedCost      *MoneyParameter2                   `xml:"http://replicon.com/ estimatedCost"`
	EstimatedHours     *TaskDuration1                     `xml:"http://replicon.com/ estimatedHours"`
	IsClosed           *string                            `xml:"http://replicon.com/ isClosed"`
	IsTimeEntryAllowed *string                            `xml:"http://replicon.com/ isTimeEntryAllowed"`
	Name               *string                            `xml:"http://replicon.com/ name"`
	PercentCompleted   *int32                             `xml:"http://replicon.com/ percentCompleted"`
	Target             *TaskTargetParameter1              `xml:"http://replicon.com/ target"`
	TimeEntryDateRange *DateRangeParameter1               `xml:"http://replicon.com/ timeEntryDateRange"`
}

type ArrayOfResourceTargetParameter1 struct {
	ResourceTargetParameter1 []*ResourceTargetParameter1 `xml:"http://replicon.com/ ResourceTargetParameter1"`
}

type ResourceTargetParameter1 struct {
	Department                                *DepartmentTargetParameter1          `xml:"http://replicon.com/ department"`
	Placeholder                               *ResourcePlaceholderTargetParameter1 `xml:"http://replicon.com/ placeholder"`
	ResourcePlaceholderParameterCorrelationId *string                              `xml:"http://replicon.com/ resourcePlaceholderParameterCorrelationId"`
	Uri                                       *string                              `xml:"http://replicon.com/ uri"`
	User                                      *UserTargetParameter1                `xml:"http://replicon.com/ user"`
}

type DepartmentTargetParameter1 struct {
	Name                   *string                     `xml:"http://replicon.com/ name"`
	ParameterCorrelationId *string                     `xml:"http://replicon.com/ parameterCorrelationId"`
	Parent                 *DepartmentTargetParameter1 `xml:"http://replicon.com/ parent"`
	Uri                    *string                     `xml:"http://replicon.com/ uri"`
}

type ResourcePlaceholderTargetParameter1 struct {
	Index                  *int32                       `xml:"http://replicon.com/ index"`
	ParameterCorrelationId *string                      `xml:"http://replicon.com/ parameterCorrelationId"`
	ProjectRole            *ProjectRoleTargetParameter1 `xml:"http://replicon.com/ projectRole"`
	Uri                    *string                      `xml:"http://replicon.com/ uri"`
}

type ProjectRoleTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type UserTargetParameter1 struct {
	LoginName *string `xml:"http://replicon.com/ loginName"`
	Uri       *string `xml:"http://replicon.com/ uri"`
}

type ArrayOfCustomFieldValueParameter1 struct {
	CustomFieldValueParameter1 []*CustomFieldValueParameter1 `xml:"http://replicon.com/ CustomFieldValueParameter1"`
}

type CustomFieldValueParameter1 struct {
	CustomField    CustomFieldTargetParameter1                `xml:"http://replicon.com/ customField"`
	Date           *Date1                                     `xml:"http://replicon.com/ date"`
	DropDownOption *CustomFieldDropDownOptionTargetParameter1 `xml:"http://replicon.com/ dropDownOption"`
	Number         *string                                    `xml:"http://replicon.com/ number"`
	Text           *string                                    `xml:"http://replicon.com/ text"`
}

type CustomFieldTargetParameter1 struct {
	GroupUri *string `xml:"http://replicon.com/ groupUri"`
	Name     *string `xml:"http://replicon.com/ name"`
	Uri      *string `xml:"http://replicon.com/ uri"`
}

type CustomFieldDropDownOptionTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type MoneyParameter2 struct {
	Amount   string                   `xml:"http://replicon.com/ amount"`
	Currency CurrencyTargetParameter1 `xml:"http://replicon.com/ currency"`
}

type CurrencyTargetParameter1 struct {
	Name   *string `xml:"http://replicon.com/ name"`
	Symbol *string `xml:"http://replicon.com/ symbol"`
	Uri    *string `xml:"http://replicon.com/ uri"`
}

type TaskTargetParameter1 struct {
	Name                   *string               `xml:"http://replicon.com/ name"`
	ParameterCorrelationId *string               `xml:"http://replicon.com/ parameterCorrelationId"`
	Parent                 *TaskTargetParameter1 `xml:"http://replicon.com/ parent"`
	Uri                    *string               `xml:"http://replicon.com/ uri"`
}

type InvalidTaskTargetParameterError1 struct {
	DisplayText *string               `xml:"http://replicon.com/ displayText"`
	Target      *TaskTargetParameter1 `xml:"http://replicon.com/ target"`
}

type ArrayOfTaskHierarchyParameter1 struct {
	TaskHierarchyParameter1 []*TaskHierarchyParameter1 `xml:"http://replicon.com/ TaskHierarchyParameter1"`
}

type TaskHierarchyParameter1 struct {
	ChildTasks *ArrayOfTaskHierarchyParameter1 `xml:"http://replicon.com/ childTasks"`
	Task       *TaskParameter1                 `xml:"http://replicon.com/ task"`
}

type ArrayOfTaskHierarchyPutResults1 struct {
	TaskHierarchyPutResults1 []*TaskHierarchyPutResults1 `xml:"http://replicon.com/ TaskHierarchyPutResults1"`
}

type TaskHierarchyPutResults1 struct {
	ChildTasks             *ArrayOfTaskHierarchyPutResults1 `xml:"http://replicon.com/ childTasks"`
	ParameterCorrelationId *string                          `xml:"http://replicon.com/ parameterCorrelationId"`
	Task                   *TaskReference1                  `xml:"http://replicon.com/ task"`
}

type ProjectInfoParameter1 struct {
	Budget                          *BudgetParameter2                   `xml:"http://replicon.com/ budget"`
	Client                          *ClientTargetParameter1             `xml:"http://replicon.com/ client"`
	Code                            *string                             `xml:"http://replicon.com/ code"`
	CostTypeUri                     *string                             `xml:"http://replicon.com/ costTypeUri"`
	CustomFieldValues               *ArrayOfCustomFieldValueParameter1  `xml:"http://replicon.com/ customFieldValues"`
	Description                     *string                             `xml:"http://replicon.com/ description"`
	EstimatedCost                   *MoneyParameter2                    `xml:"http://replicon.com/ estimatedCost"`
	EstimatedExpenses               *MoneyParameter2                    `xml:"http://replicon.com/ estimatedExpenses"`
	EstimatedHours                  *TaskDuration1                      `xml:"http://replicon.com/ estimatedHours"`
	EstimationModeUri               *string                             `xml:"http://replicon.com/ estimationModeUri"`
	IsProjectLeaderApprovalRequired *string                             `xml:"http://replicon.com/ isProjectLeaderApprovalRequired"`
	IsTimeEntryAllowed              *string                             `xml:"http://replicon.com/ isTimeEntryAllowed"`
	Name                            *string                             `xml:"http://replicon.com/ name"`
	PercentCompleted                *int32                              `xml:"http://replicon.com/ percentCompleted"`
	Program                         *ProgramTargetParameter1            `xml:"http://replicon.com/ program"`
	ProjectLeader                   *UserTargetParameter1               `xml:"http://replicon.com/ projectLeader"`
	ProjectStatusLabel              *ProjectStatusLabelTargetParameter1 `xml:"http://replicon.com/ projectStatusLabel"`
	TimeEntryDateRange              *DateRangeParameter1                `xml:"http://replicon.com/ timeEntryDateRange"`
}

type BudgetParameter2 struct {
	CapitalBudget     *MoneyParameter2 `xml:"http://replicon.com/ capitalBudget"`
	OperationalBudget *MoneyParameter2 `xml:"http://replicon.com/ operationalBudget"`
	TotalBudget       *MoneyParameter2 `xml:"http://replicon.com/ totalBudget"`
}

type ClientTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type ProgramTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type ProjectStatusLabelTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type InvalidProjectDetailsParameterError1 struct {
	DisplayText *string                `xml:"http://replicon.com/ displayText"`
	Parameter   *ProjectInfoParameter1 `xml:"http://replicon.com/ parameter"`
}

type InvalidProjectTargetParameterError1 struct {
	DisplayText *string                  `xml:"http://replicon.com/ displayText"`
	Target      *ProjectTargetParameter1 `xml:"http://replicon.com/ target"`
}

type InvalidBudgetParameterError1 struct {
	DisplayText *string           `xml:"http://replicon.com/ displayText"`
	Parameter   *BudgetParameter1 `xml:"http://replicon.com/ parameter"`
}

type BudgetParameter1 struct {
	CapitalBudget     *MoneyParameter1 `xml:"http://replicon.com/ capitalBudget"`
	OperationalBudget *MoneyParameter1 `xml:"http://replicon.com/ operationalBudget"`
	TotalBudget       *MoneyParameter1 `xml:"http://replicon.com/ totalBudget"`
}

type ProjectParameter1 struct {
	Expenses    *ProjectExpensesParameter1      `xml:"http://replicon.com/ expenses"`
	ProjectInfo *ProjectInfoParameter1          `xml:"http://replicon.com/ projectInfo"`
	Target      *ProjectTargetParameter1        `xml:"http://replicon.com/ target"`
	Tasks       *ArrayOfTaskHierarchyParameter1 `xml:"http://replicon.com/ tasks"`
	Team        *ProjectTeamParameter1          `xml:"http://replicon.com/ team"`
}

type ProjectExpensesParameter1 struct {
	ExpenseCodes *ArrayOfExpenseCodeTargetParameter1 `xml:"http://replicon.com/ expenseCodes"`
}

type ArrayOfExpenseCodeTargetParameter1 struct {
	ExpenseCodeTargetParameter1 []*ExpenseCodeTargetParameter1 `xml:"http://replicon.com/ ExpenseCodeTargetParameter1"`
}

type ExpenseCodeTargetParameter1 struct {
	Name *string `xml:"http://replicon.com/ name"`
	Uri  *string `xml:"http://replicon.com/ uri"`
}

type ProjectTeamParameter1 struct {
	TeamMembers *ArrayOfProjectTeamMemberParameter1 `xml:"http://replicon.com/ teamMembers"`
}

type ArrayOfProjectTeamMemberParameter1 struct {
	ProjectTeamMemberParameter1 []*ProjectTeamMemberParameter1 `xml:"http://replicon.com/ ProjectTeamMemberParameter1"`
}

type ProjectTeamMemberParameter1 struct {
	Resource                       *ResourceTargetParameter1    `xml:"http://replicon.com/ resource"`
	ResourcePlaceholderProjectRole *ProjectRoleTargetParameter1 `xml:"http://replicon.com/ resourcePlaceholderProjectRole"`
}

type ArrayOfTaskResourceAssignmentDetails1 struct {
	TaskResourceAssignmentDetails1 []*TaskResourceAssignmentDetails1 `xml:"http://replicon.com/ TaskResourceAssignmentDetails1"`
}

type TaskResourceAssignmentDetails1 struct {
	Assignments *ArrayOfProjectTeamMemberDetails1 `xml:"http://replicon.com/ assignments"`
	TaskUri     *string                           `xml:"http://replicon.com/ taskUri"`
}

type ArrayOfanyURI struct {
	AnyURI []*string `xml:"http://schemas.microsoft.com/2003/10/Serialization/Arrays anyURI"`
}
