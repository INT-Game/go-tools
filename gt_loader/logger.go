package gt_loader

import "github.com/INT-Game/go-tools/slog"

var (
	cfgLogger      = slog.NewSLogger("[cfg] %s")
	csvLogger      = slog.NewSLogger("[csv] %s")
	exlLogger      = slog.NewSLogger("[exl] %s")
	exlSheetLogger = slog.NewSLogger("[exl-sheet] %s")
)
