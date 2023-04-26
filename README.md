# Kit Conheça Sua Saúde
Kit Conheça Sua Saúde para pessoas com deficiência e dificuldade de vocalização

# Dependências

- Instale o Go a partir das instruções indicadas no [site oficial](https://go.dev/doc/install)
- Instale o fyne 

```
go get -u fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest
```
- Instale o gomobile

```
go get -u golang.org/x/mobile/cmd/gomobile
go install golang.org/x/mobile/cmd/gomobile@latest
```

- Instale o [Android Studio](https://developer.android.com/studio?gclid=CjwKCAjw9J2iBhBPEiwAErwpebXq0FBhXqHl31GT0I3iap_P7QUwcb9LBByaPrUI5BjT0T90DRkxORoCG8cQAvD_BwE&gclsrc=aw.ds)
- Faça as instalações no SDK Manager do Android Studio
- Defina as variáveis $GOPATH, $ANDROID_HOME, $ANDROID_NDK_HOME 


```
export GOPATH=$HOME/go/bin/
export ANDROID_HOME=$HOME/Android/Sdk
export ANDROID_NDK_HOME=$ANDROID_HOME/ndk/25.2.9519653
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin:$ANDROID_HOME/platform-tools/
```


# Geração do apk

```
fyne package -os android -appID conheca.sua.saude -icon assets/icons/icon.png
```
