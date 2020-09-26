package command

import (
	"fmt"
	"github.com/tickstep/cloudpan189-api/cloudpan"
	"github.com/tickstep/cloudpan189-go/internal/config"
	"github.com/tickstep/library-go/logger"
	"path"
)

var (
	panCommandVerbose = logger.New("PANCOMMAND", config.EnvVerbose)
)

// GetFileInfoByPaths 获取指定文件路径的文件详情信息
func GetFileInfoByPaths(paths ...string) (fileInfoList []*cloudpan.FileEntity, failedPaths []string, error error) {
	if len(paths) <= 0 {
		return nil, nil, fmt.Errorf("请指定文件路径")
	}
	activeUser := GetActiveUser()

	for idx := 0; idx < len(paths); idx++ {
		absolutePath := path.Clean(activeUser.PathJoin(0, paths[idx]))
		fe, err := activeUser.PanClient().FileInfoByPath(absolutePath)
		if err != nil {
			failedPaths = append(failedPaths, absolutePath)
			continue
		}
		fileInfoList = append(fileInfoList, fe)
	}
	return
}

func matchPathByShellPattern(patterns ...string) (panpaths []string, err error) {
	acUser := GetActiveUser()
	for k := range patterns {
		ps := acUser.PathJoin(0, patterns[k])
		panpaths = append(panpaths, ps)
	}
	return panpaths, nil
}

func IsFamilyCloud(familyId int64) bool {
	return familyId > 0
}

func GetFamilyCloudMark(familyId int64) string {
	if familyId > 0 {
		return "家庭云"
	}
	return "个人云"
}