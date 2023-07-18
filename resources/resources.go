package resources

import "fyne.io/fyne/v2"

func GetAudioResource(word string) *fyne.StaticResource {
	return resourceAnticoncepcionalMp3
}

func IconBack() *fyne.StaticResource {
	return resourceVoltarPng
}

func IconSpeak() *fyne.StaticResource {
	return resourceFalarPng
}

func IconMySentences() *fyne.StaticResource {
	return resourceMinhasFrasesPng
}

func IconSave() *fyne.StaticResource {
	return resourceDiskettePng
}

func IconClear() *fyne.StaticResource {
	return resourceVassouraPng
}

func IconDatabase() *fyne.StaticResource {
	return resourceCardGamesPng
}

func Collection() map[string]fyne.Resource {
	iconCollection := make(map[string]fyne.Resource)
	iconCollection["Anticoncepcional"] = resourceAnticoncepcionalPng
	iconCollection["Boca"] = resourceBocaPng
	iconCollection["Câncer de Pênis"] = resourceCancerDePenisPng
	iconCollection["Câncer de Vulva"] = resourceCancerDeVulvaPng
	iconCollection["Coração"] = resourceCoracaoPng
	iconCollection["Dedo"] = resourceDedoPng
	iconCollection["Gonorréia"] = resourceGonorreiaPng
	iconCollection["Gravidez"] = resourceGravidezPng
	iconCollection["Saúde Mental"] = resourceSaudeMentalPng
	iconCollection["Saúde"] = resourceSaudePng
	iconCollection["Sexo"] = resourceSexoPng
	iconCollection["Vulva"] = resourceVulvaPng
	return iconCollection
	// resourceAnticoncepcionalPng,
	// resourceBocaPng,
	// resourceCancerDePenisPng,
	// resourceCancerDeVulvaPng,
	// resourceCoracaoPng,
	// resourceDedoPng,
	// resourceGonorreiaPng,
	// resourceGravidezPng,
	// resourceSaudeMentalPng,
	// resourceSaudePng,
	// resourceSexoPng,
	// resourceVulvaPng,
}
