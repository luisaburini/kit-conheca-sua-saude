package resources

import "fyne.io/fyne/v2"

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

func Words() []*fyne.StaticResource {
	return []*fyne.StaticResource{
		resourceAnticoncepcionalPng,
		resourceBocaPng,
		resourceCancerDePenisPng,
		resourceCancerDeVulvaPng,
		resourceCoracaoPng,
		resourceDedoPng,
		resourceGonorreiaPng,
		resourceGravidezPng,
		resourceSaudeMentalPng,
		resourceSaudePng,
		resourceSexoPng,
		resourceVulvaPng,
	}
}
