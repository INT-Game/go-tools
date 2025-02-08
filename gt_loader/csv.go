package gt_loader

import (
	"context"
	"encoding/csv"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"os"
)

// LoadCSVCallback 加载CSV回调函数
// @param index 行索引, 从0开始
// @param row 行数据
type LoadCSVCallback func(index int, row []string) (err error)

// LoadCSV 加载CSV文件
// @param ctx 上下文
// @param path 文件路径
// @param skipRows 跳过的行数, 如果需要跳过1行, 则传入1
// @param callback 回调函数
func LoadCSV(ctx context.Context, path string, skipRows int, callback LoadCSVCallback) (err error) {
	csvLogger.CInfo(ctx, "开始加载: %s", path)
	f, err := os.Open(path)
	if err != nil {
		return
	}

	defer func() {
		if err = f.Close(); err != nil {
			return
		}
	}()

	reader := csv.NewReader(transform.NewReader(f, simplifiedchinese.GBK.NewDecoder()))
	rows, err := reader.ReadAll()
	if err != nil {
		return
	}

	for i, row := range rows {
		if i >= skipRows {
			if err = callback(i, row); err != nil {
				csvLogger.CError(ctx, "处理%d行失败, row: %v, err: %v", i, row, err)
				return
			}
		}
	}

	csvLogger.CInfo(ctx, "加载完成: %s", path)
	return
}
