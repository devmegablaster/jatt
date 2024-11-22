package styles

import "github.com/charmbracelet/lipgloss"

var (
	LogoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF9933")).Bold(true).Padding(1, 1)
	DebugStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#3B82F6"))
	SuccessStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#22C55E")).Bold(true)
	ErrorStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
	WarningStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF9933"))
	StatsStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#EAB308")).Bold(true)
)
