package theme

import "fyne.io/fyne"

var pointerDefault = &fyne.StaticResource{
	StaticName: "pointer.png",
	StaticContent: []byte{
		137, 80, 78, 71, 13, 10, 26, 10, 0, 0, 0, 13, 73, 72, 68, 82, 0, 0, 0, 48, 0, 0, 0, 48, 8, 6, 0, 0, 0, 87, 2, 249, 135, 0, 0, 0, 6, 98, 75, 71, 68, 0, 255, 0, 255, 0, 255, 160, 189, 167, 147, 0, 0, 0, 9, 112, 72, 89, 115, 0, 0, 46, 35, 0, 0, 46, 35, 1, 120, 165, 63, 118, 0, 0, 0, 7, 116, 73, 77, 69, 7, 226, 6, 24, 14, 33, 16, 22, 3, 178, 116, 0, 0, 0, 25, 116, 69, 88, 116, 67, 111, 109, 109, 101, 110, 116, 0, 67, 114, 101, 97, 116, 101, 100, 32, 119, 105, 116, 104, 32, 71, 73, 77, 80, 87, 129, 14, 23, 0, 0, 1, 0, 73, 68, 65, 84, 104, 222, 213, 218, 193, 18, 131, 32, 12, 69, 209, 112, 199, 255, 255, 101, 186, 233, 56, 109, 209, 138, 72, 66, 158, 59, 133, 100, 222, 81, 71, 23, 80, 204, 172, 86, 179, 98, 162, 7, 102, 102, 197, 172, 74, 3, 148, 17, 124, 158, 40, 34, 248, 189, 160, 134, 216, 1, 181, 86, 73, 196, 215, 19, 80, 68, 52, 175, 144, 26, 130, 163, 139, 74, 8, 206, 6, 84, 16, 252, 27, 84, 64, 112, 53, 33, 59, 130, 158, 73, 153, 17, 244, 78, 204, 138, 224, 206, 228, 140, 8, 238, 22, 100, 67, 48, 82, 148, 9, 193, 104, 97, 22, 4, 79, 138, 51, 32, 120, 218, 96, 53, 130, 25, 77, 86, 34, 152, 213, 104, 21, 130, 153, 205, 86, 32, 152, 221, 48, 26, 129, 71, 211, 72, 4, 94, 141, 163, 16, 120, 222, 157, 8, 4, 222, 239, 168, 55, 130, 136, 47, 133, 39, 34, 4, 224, 137, 8, 3, 120, 33, 66, 1, 30, 136, 109, 164, 168, 148, 121, 203, 9, 229, 225, 250, 4, 43, 195, 207, 120, 18, 219, 104, 248, 44, 171, 58, 40, 135, 239, 6, 100, 13, 223, 5, 200, 28, 254, 18, 144, 61, 252, 95, 128, 66, 248, 83, 128, 74, 248, 67, 128, 82, 248, 6, 160, 22, 254, 253, 19, 108, 255, 130, 74, 91, 15, 80, 14, 223, 0, 20, 55, 125, 160, 28, 126, 7, 40, 111, 183, 121, 1, 249, 98, 132, 49, 112, 207, 8, 124, 0, 0, 0, 0, 73, 69, 78, 68, 174, 66, 96, 130}}

var batteryLight = &fyne.StaticResource{
	StaticName: "battery_light.svg",
	StaticContent: []byte{
		60, 115, 118, 103, 32, 120, 109, 108, 110, 115, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 119, 51, 46, 111, 114, 103, 47, 50, 48, 48, 48, 47, 115, 118, 103, 34, 32, 119, 105, 100, 116, 104, 61, 34, 50, 52, 34, 32, 104, 101, 105, 103, 104, 116, 61, 34, 50, 52, 34, 32, 118, 105, 101, 119, 66, 111, 120, 61, 34, 48, 32, 48, 32, 50, 52, 32, 50, 52, 34, 62, 60, 112, 97, 116, 104, 32, 100, 61, 34, 77, 48, 32, 48, 104, 50, 52, 118, 50, 52, 72, 48, 122, 34, 32, 102, 105, 108, 108, 61, 34, 110, 111, 110, 101, 34, 47, 62, 60, 112, 97, 116, 104, 32, 100, 61, 34, 77, 49, 53, 46, 54, 55, 32, 52, 72, 49, 52, 86, 50, 104, 45, 52, 118, 50, 72, 56, 46, 51, 51, 67, 55, 46, 54, 32, 52, 32, 55, 32, 52, 46, 54, 32, 55, 32, 53, 46, 51, 51, 118, 49, 53, 46, 51, 51, 67, 55, 32, 50, 49, 46, 52, 32, 55, 46, 54, 32, 50, 50, 32, 56, 46, 51, 51, 32, 50, 50, 104, 55, 46, 51, 51, 99, 46, 55, 52, 32, 48, 32, 49, 46, 51, 52, 45, 46, 54, 32, 49, 46, 51, 52, 45, 49, 46, 51, 51, 86, 53, 46, 51, 51, 67, 49, 55, 32, 52, 46, 54, 32, 49, 54, 46, 52, 32, 52, 32, 49, 53, 46, 54, 55, 32, 52, 122, 34, 47, 62, 60, 47, 115, 118, 103, 62}}

var batteryDark = &fyne.StaticResource{
	StaticName: "battery_dark.svg",
	StaticContent: []byte{
		60, 115, 118, 103, 32, 120, 109, 108, 110, 115, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 119, 51, 46, 111, 114, 103, 47, 50, 48, 48, 48, 47, 115, 118, 103, 34, 32, 119, 105, 100, 116, 104, 61, 34, 50, 52, 34, 32, 104, 101, 105, 103, 104, 116, 61, 34, 50, 52, 34, 32, 118, 105, 101, 119, 66, 111, 120, 61, 34, 48, 32, 48, 32, 50, 52, 32, 50, 52, 34, 62, 60, 112, 97, 116, 104, 32, 100, 61, 34, 77, 48, 32, 48, 104, 50, 52, 118, 50, 52, 72, 48, 122, 34, 32, 102, 105, 108, 108, 61, 34, 110, 111, 110, 101, 34, 47, 62, 60, 112, 97, 116, 104, 32, 102, 105, 108, 108, 61, 34, 35, 102, 102, 102, 102, 102, 102, 34, 32, 100, 61, 34, 77, 49, 53, 46, 54, 55, 32, 52, 72, 49, 52, 86, 50, 104, 45, 52, 118, 50, 72, 56, 46, 51, 51, 67, 55, 46, 54, 32, 52, 32, 55, 32, 52, 46, 54, 32, 55, 32, 53, 46, 51, 51, 118, 49, 53, 46, 51, 51, 67, 55, 32, 50, 49, 46, 52, 32, 55, 46, 54, 32, 50, 50, 32, 56, 46, 51, 51, 32, 50, 50, 104, 55, 46, 51, 51, 99, 46, 55, 52, 32, 48, 32, 49, 46, 51, 52, 45, 46, 54, 32, 49, 46, 51, 52, 45, 49, 46, 51, 51, 86, 53, 46, 51, 51, 67, 49, 55, 32, 52, 46, 54, 32, 49, 54, 46, 52, 32, 52, 32, 49, 53, 46, 54, 55, 32, 52, 122, 34, 47, 62, 60, 47, 115, 118, 103, 62, 10}}

var brightnessLight = &fyne.StaticResource{
	StaticName: "brightness_light.svg",
	StaticContent: []byte{
		60, 115, 118, 103, 32, 120, 109, 108, 110, 115, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 119, 51, 46, 111, 114, 103, 47, 50, 48, 48, 48, 47, 115, 118, 103, 34, 32, 119, 105, 100, 116, 104, 61, 34, 50, 52, 34, 32, 104, 101, 105, 103, 104, 116, 61, 34, 50, 52, 34, 32, 118, 105, 101, 119, 66, 111, 120, 61, 34, 48, 32, 48, 32, 50, 52, 32, 50, 52, 34, 62, 60, 112, 97, 116, 104, 32, 100, 61, 34, 77, 48, 32, 48, 104, 50, 52, 118, 50, 52, 72, 48, 122, 34, 32, 102, 105, 108, 108, 61, 34, 110, 111, 110, 101, 34, 47, 62, 60, 112, 97, 116, 104, 32, 100, 61, 34, 77, 50, 49, 32, 51, 72, 51, 99, 45, 49, 46, 49, 32, 48, 45, 50, 32, 46, 57, 45, 50, 32, 50, 118, 49, 52, 99, 48, 32, 49, 46, 49, 46, 57, 32, 50, 32, 50, 32, 50, 104, 49, 56, 99, 49, 46, 49, 32, 48, 32, 50, 45, 46, 57, 32, 50, 45, 50, 86, 53, 99, 48, 45, 49, 46, 49, 45, 46, 57, 45, 50, 45, 50, 45, 50, 122, 109, 48, 32, 49, 54, 46, 48, 49, 72, 51, 86, 52, 46, 57, 57, 104, 49, 56, 118, 49, 52, 46, 48, 50, 122, 77, 56, 32, 49, 54, 104, 50, 46, 53, 108, 49, 46, 53, 32, 49, 46, 53, 32, 49, 46, 53, 45, 49, 46, 53, 72, 49, 54, 118, 45, 50, 46, 53, 108, 49, 46, 53, 45, 49, 46, 53, 45, 49, 46, 53, 45, 49, 46, 53, 86, 56, 104, 45, 50, 46, 53, 76, 49, 50, 32, 54, 46, 53, 32, 49, 48, 46, 53, 32, 56, 72, 56, 118, 50, 46, 53, 76, 54, 46, 53, 32, 49, 50, 32, 56, 32, 49, 51, 46, 53, 86, 49, 54, 122, 109, 52, 45, 55, 99, 49, 46, 54, 54, 32, 48, 32, 51, 32, 49, 46, 51, 52, 32, 51, 32, 51, 115, 45, 49, 46, 51, 52, 32, 51, 45, 51, 32, 51, 86, 57, 122, 34, 47, 62, 60, 47, 115, 118, 103, 62}}

var brightnessDark = &fyne.StaticResource{
	StaticName: "brightness_dark.svg",
	StaticContent: []byte{
		60, 115, 118, 103, 32, 120, 109, 108, 110, 115, 61, 34, 104, 116, 116, 112, 58, 47, 47, 119, 119, 119, 46, 119, 51, 46, 111, 114, 103, 47, 50, 48, 48, 48, 47, 115, 118, 103, 34, 32, 119, 105, 100, 116, 104, 61, 34, 50, 52, 34, 32, 104, 101, 105, 103, 104, 116, 61, 34, 50, 52, 34, 32, 118, 105, 101, 119, 66, 111, 120, 61, 34, 48, 32, 48, 32, 50, 52, 32, 50, 52, 34, 62, 60, 112, 97, 116, 104, 32, 100, 61, 34, 77, 48, 32, 48, 104, 50, 52, 118, 50, 52, 72, 48, 122, 34, 32, 102, 105, 108, 108, 61, 34, 110, 111, 110, 101, 34, 47, 62, 60, 112, 97, 116, 104, 32, 102, 105, 108, 108, 61, 34, 35, 102, 102, 102, 102, 102, 102, 34, 32, 100, 61, 34, 77, 50, 49, 32, 51, 72, 51, 99, 45, 49, 46, 49, 32, 48, 45, 50, 32, 46, 57, 45, 50, 32, 50, 118, 49, 52, 99, 48, 32, 49, 46, 49, 46, 57, 32, 50, 32, 50, 32, 50, 104, 49, 56, 99, 49, 46, 49, 32, 48, 32, 50, 45, 46, 57, 32, 50, 45, 50, 86, 53, 99, 48, 45, 49, 46, 49, 45, 46, 57, 45, 50, 45, 50, 45, 50, 122, 109, 48, 32, 49, 54, 46, 48, 49, 72, 51, 86, 52, 46, 57, 57, 104, 49, 56, 118, 49, 52, 46, 48, 50, 122, 77, 56, 32, 49, 54, 104, 50, 46, 53, 108, 49, 46, 53, 32, 49, 46, 53, 32, 49, 46, 53, 45, 49, 46, 53, 72, 49, 54, 118, 45, 50, 46, 53, 108, 49, 46, 53, 45, 49, 46, 53, 45, 49, 46, 53, 45, 49, 46, 53, 86, 56, 104, 45, 50, 46, 53, 76, 49, 50, 32, 54, 46, 53, 32, 49, 48, 46, 53, 32, 56, 72, 56, 118, 50, 46, 53, 76, 54, 46, 53, 32, 49, 50, 32, 56, 32, 49, 51, 46, 53, 86, 49, 54, 122, 109, 52, 45, 55, 99, 49, 46, 54, 54, 32, 48, 32, 51, 32, 49, 46, 51, 52, 32, 51, 32, 51, 115, 45, 49, 46, 51, 52, 32, 51, 45, 51, 32, 51, 86, 57, 122, 34, 47, 62, 60, 47, 115, 118, 103, 62, 10}}
