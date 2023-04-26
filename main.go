package main

import (
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func main() {
	a := app.New()
	w := a.NewWindow("Kit Conheca Sua Saude")
	iconUrls := []string{
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRvOu-MpDMO_DQ2J_Gf8Onss3zJMlIu0QecHQ&usqp=CAU",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSCBg0I8YJIUzVmCO0JxoWPc6Sy4w22DG-5mDEq2TpptNm7Wa2ZzJJVPGsMumCLauVdlqs&usqp=CAU",
		"https://static.vecteezy.com/system/resources/previews/015/101/034/original/female-sex-health-icon-outline-sexual-education-vector.jpg",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRoMAmt8w8Iz_hzsZt1YEVP3JiZae4DyIANsw&usqp=CAU",
		"https://cdn-icons-png.flaticon.com/512/7077/7077410.png",
		"https://cdn-icons-png.flaticon.com/512/4773/4773193.png",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcT5tZCEl-7hyZD9OkLT5hAhEGbSWCejtQHWNFoXcoE7T4pqRjmQ_yjoq1CqCkdAtXIRTNU&usqp=CAU",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQXOgu6RVZSYcCmBsmctRPvjOcuEKWDlhFZCg&usqp=CAU",
		"https://img1.pnghut.com/6/8/18/ui3icEyR0y/health-food-orange-cooking-dish-vegetable.jpg",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTfGKHXqh_3UFYj_BNPpiRqF-JGNW6Yq9wi-w&usqp=CAU",
		"https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcS_ddxdSmIgIkwuQKY30-2UmtpOfYYDgor48g&usqp=CAU",
		"https://images.squarespace-cdn.com/content/v1/5b6601fe9f8770774444ea60/1618606924463-VYXWWB5S61UQ80YCLAFA/ESW-Website-Icons_Sexual-Health.png?format=500w",
	}

	grid := container.NewAdaptiveGrid(4)
	for _, url := range iconUrls {
		res, err := fyne.LoadResourceFromURLString(url)
		if err != nil {
			res, _ = fyne.LoadResourceFromURLString("https://cdn-icons-png.flaticon.com/512/5741/5741333.png")
		}
		grid.Add(widget.NewCard("", "", canvas.NewImageFromResource(res)))
	}
	w.SetContent(container.NewMax(grid))
	w.ShowAndRun()
}

// func fileNameWithoutExtSliceNotation(fileName string) string {
// 	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
// }

// func main() {
// 	a := app.New()
// 	w := a.NewWindow("Kit Conheca Sua Saude")

// 	files, err := ioutil.ReadDir("./assets/icons/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	grid := container.NewAdaptiveGrid(3)
// 	for _, file := range files {
// 		grid.Add(widget.NewCard("", fileNameWithoutExtSliceNotation(file.Name()), canvas.NewImageFromFile("./assets/icons/"+file.Name())))
// 	}
// 	w.SetContent(container.NewMax(grid))
// 	w.ShowAndRun()
// }
