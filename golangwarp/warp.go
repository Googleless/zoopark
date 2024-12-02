package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	url := "https://raw.githubusercontent.com/xSlam3/warp_amneziawg_generator/main/wwgcf.sh"
	fileName := "wwgcf.sh"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при загрузке файла:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	fmt.Println("Содержимое скрипта:n", string(content))

	err = runAsAdmin(fileName)
	if err != nil {
		fmt.Println("Ошибка при выполнении скрипта:", err)
		return
	}

	fmt.Println("Скрипт выполнен успешно.")
}

func runAsAdmin(script string) error {
	cmd := exec.Command("powershell", "-Command", fmt.Sprintf("Start-Process wsl -ArgumentList 'bash %s' -Verb RunAs", script))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("ошибка: %v, вывод: %s", err, output)
	}
	return nil
}
