package msg

import (
	"strings"
)

// escapeMarkdownV2 escapes special characters for Telegram MarkdownV2
func escapeMarkdownV2(s string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(s)
}

func Msg(message string) string {
	if message == "" {
		return "*⚠️ No message defined!*"
	}

	lower := strings.ToLower(message)

	switch {
	case strings.Contains(lower, "alert"):
		clean := removeWord(message, "alert:")
		return "*🚨 ALERT:* _" + escapeMarkdownV2(clean) + "_"

	case strings.Contains(lower, "success"):
		clean := removeWord(message, "success:")
		return "*✅ SUCCESS:* _" + escapeMarkdownV2(clean) + "_"

	case strings.Contains(lower, "error"):
		clean := removeWord(message, "error:")
		return "*❌ ERROR:* _" + escapeMarkdownV2(clean) + "_"

	case strings.Contains(lower, "warn"):
		clean := removeWord(message, "warn:")
		return "*⚠️ WARNING:* _" + escapeMarkdownV2(clean) + "_"

	default:
		return "*📢 Message:* _" + escapeMarkdownV2(message) + "_"
	}
}

func removeWord(input, word string) string {
	index := strings.Index(strings.ToLower(input), word)
	if index == -1 {
		return input
	}
	before := strings.TrimSpace(input[:index])
	after := strings.TrimSpace(input[index+len(word):])
	return strings.TrimSpace(before + " " + after)
}
