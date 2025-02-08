package gt_loader

import (
	"context"
	"github.com/xuri/excelize/v2"
)

// LoadExcelCallback
// @param f 文件对象
type LoadExcelCallback func(f *excelize.File) (err error)

// LoadExcel
// @param ctx 上下文
// @param path 文件路径
// @param callback 回调函数
func LoadExcel(ctx context.Context, path string, callback LoadExcelCallback) (err error) {
	exlLogger.CInfo(ctx, "开始加载: %s", path)
	f, err := excelize.OpenFile(path)
	if err != nil {
		return
	}

	defer func() {
		if err = f.Close(); err != nil {
			return
		}
	}()

	if err = callback(f); err != nil {
		return
	}

	exlLogger.CInfo(ctx, "加载完成: %s", path)
	return
}

// LoadExcelSheetCallback
// @param index 行索引, 从0开始
// @param row 行数据
type LoadExcelSheetCallback func(index int, row []string) (err error)

// LoadExcelSheet
// @param ctx 上下文
// @param f 文件对象
// @param sheet 表名
// @param skipRows 跳过的行数, 如果需要跳过1行, 则传入1
// @param callback 回调函数
func LoadExcelSheet(ctx context.Context, f *excelize.File, sheet string, skipRows int, callback LoadExcelSheetCallback) (err error) {
	exlSheetLogger.CInfo(ctx, "开始加载: %s, sheet: %s", f.Path, sheet)
	rows, err := f.Rows(sheet)
	if err != nil {
		return
	}

	defer func() {
		if err = rows.Close(); err != nil {
			return
		}
	}()

	rowIdx := 0
	for rows.Next() {
		if rowIdx >= skipRows {
			var row []string
			row, err = rows.Columns()
			if err != nil {
				return
			}

			if err = callback(rowIdx, row); err != nil {
				return
			}
		}
		rowIdx++
	}

	exlSheetLogger.CInfo(ctx, "加载完成: %s, sheet: %s", f.Path, sheet)
	return
}
