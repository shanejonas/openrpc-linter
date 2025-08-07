package reporters

import (
	"fmt"
	"io"

	"github.com/shanejonas/openrpc-linter/types"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

type TextReporter struct{}

var (
	errorStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5555")).Bold(true)
	ruleHeaderStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#BDBBFE")).Bold(true)
	messageStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFD700"))
	successStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#50FA7B")).Bold(true)
)

func (r *TextReporter) Format(results []types.RuleFunctionResult, totalRules int, output io.Writer) error {
	errorCount := 0
	ruleErrors := make(map[string][]string)

	// Collect errors
	for _, result := range results {
		if result.Message != "" {
			ruleErrors[result.RuleID] = append(ruleErrors[result.RuleID], result.Message)
			errorCount++
		}
	}

	// Print errors grouped by rule with styling, using lipgloss/list
	for ruleId, messages := range ruleErrors {
		ruleHeader := ruleHeaderStyle.Render(fmt.Sprintf("Rule: %s", ruleId))
		if _, err := fmt.Fprintf(output, "%s\n", ruleHeader); err != nil {
			return err
		}
		// Use lipgloss/list to render the list of messages with custom styling
		// Convert []string to []interface{} for lipgloss list
		items := make([]interface{}, len(messages))
		for i, msg := range messages {
			items[i] = msg
		}
		msgList := list.New(items...).
			Enumerator(list.Bullet).
			ItemStyle(messageStyle)
		if _, err := fmt.Fprintf(output, "%s\n", msgList.String()); err != nil {
			return err
		}
	}

	if errorCount == 0 {
		msg := successStyle.Render(fmt.Sprintf("✅ All %d rules passed!", totalRules))
		if _, err := fmt.Fprintf(output, "\n%s\n", msg); err != nil {
			return err
		}
	} else {
		rulesWithErrors := len(ruleErrors)
		summary := errorStyle.Render(fmt.Sprintf("❌ %d error(s) found in %d rules", errorCount, rulesWithErrors))
		if _, err := fmt.Fprintf(output, "\n%s\n", summary); err != nil {
			return err
		}
	}

	return nil
}
