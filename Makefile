NDK=/home/luisa/Android/Sdk/ndk/23.2.8568313
AL_PATH=~/Documents/openal-soft/
APK_TMP=/tmp/_saude_apk_/
ANDROID_HOME=~/Android/Sdk/
CMAKE_PATH=~/Downloads/cmake-3.22.1-linux-x86_64/bin
#CMAKE_PATH=/usr/bin/cmake	
GOMOBILE_PATH=~/go/bin/bin/gomobile

studio: android
	unzip saude.apk -d $(APK_TMP)
	cp -r $(APK_TMP)/lib/* AndroidProject/app/src/main/jniLibs/
	rm -rf $(APK_TMP)

openal:
	PATH="$(CMAKE_PATH):$(PATH)" ANDROID_PLATFORM=android-19 ANDROID_HOME=$(ANDROID_HOME) ANDROID_NDK_HOME=$(NDK) $(GOMOBILE_PATH) init -openal $(AL_PATH)

init:
	ANDROID_NDK_HOME=$(NDK) gomobile init -openal $(AL_PATH) -ldflags="-w" -gcflags="-w -I $(AL_PATH)/include/" 

android:
	ANDROID_NDK_HOME=$(NDK) $(GOMOBILE_PATH) build -target android -androidapi=19 -o saude.apk -ldflags="-w" -gcflags="-w -I $(AL_PATH)/include/" 

bind:
	ANDROID_HOME=$(NDK) ANDROID_NDK_HOME=$(NDK) gomobile bind -target android -androidapi=19 -o saude.apk -ldflags="-w" -gcflags="-w -I $(AL_PATH)/include/" 

android_install:
	ANDROID_NDK_HOME=$(NDK) $(GOMOBILE_PATH) install -target android -androidapi=19 -o saude.apk -ldflags="-w" -gcflags="-I $(AL_PATH)/include/" -x

linux:
	go build .
