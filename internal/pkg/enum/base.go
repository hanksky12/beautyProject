package enum

type Base struct {
	Number      int
	Name        string
	ChineseName string
}

func (b *Base) GetChineseName() string {
	return b.ChineseName
}

func (b *Base) GetName() string {
	return b.Name
}

func (b *Base) GetNumber() int {
	return b.Number
}

// /////////////////////
// 搭配 Map 使用
func GetEnumByName[T any](name string, m map[string]T) (T, bool) {
	// 返回自訂的常數
	val, exists := m[name]
	return val, exists
}

func GetChineseNameByName[T interface{ GetChineseName() string }](name string, m map[string]T) string {
	if val, exists := m[name]; exists {
		return val.GetChineseName()
	}
	return ""
}

func ReplaceChineseNameToName[T interface {
	GetChineseName() string
	GetName() string
}](conditions map[string]any, m map[string]T) {
	for key, value := range conditions {
		for _, item := range m {
			if value == item.GetChineseName() {
				conditions[key] = item.GetName()
			}
		}
	}
}
