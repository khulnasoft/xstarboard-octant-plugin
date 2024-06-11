package configaudit

import (
	"strconv"

	"github.com/khulnasoft/starboard/pkg/apis/khulnasoft/v1alpha1"
	"github.com/vmware-tanzu/octant/pkg/view/component"
)

func NewSummarySections(summary *v1alpha1.ConfigAuditSummary) []component.SummarySection {
	if summary == nil {
		return []component.SummarySection{}
	}
	return []component.SummarySection{
		{Header: "Passing Audit Checks", Content: component.NewText(strconv.Itoa(summary.PassCount))},
		{Header: "Warning Audit Checks", Content: component.NewText(strconv.Itoa(summary.WarningCount))},
		{Header: "Dangerous Audit Checks", Content: component.NewText(strconv.Itoa(summary.DangerCount))},
	}
}
