package resources

import "fyne.io/fyne/v2"

func GetIconFalar() *fyne.StaticResource {
	return resourceFalarPng
}

func GetIconMinhasFrases() *fyne.StaticResource {
	return resourceMinhasFrasesPng
}

func GetIconSave() *fyne.StaticResource {
	return resourceDiskettePng
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
