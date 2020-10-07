package prompts

import (
	"github.com/manifoldco/promptui"
)

func GetListValue(items []string, promptxt string) (string, error) {
	selectList := promptui.Select{
		Label: promptxt,
		Items: items,
	}

	_, result, err := selectList.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
