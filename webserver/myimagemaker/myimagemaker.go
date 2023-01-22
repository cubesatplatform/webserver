package myimagemaker

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

// func MakeImage(irdata []byte) {
func MakeIRImage(irdata []byte) image.Image {

	//irdata := "0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789"

	//32 x 24

	width_pixels := 640
	height_pixels := 480
	size_block := 20

	myimage := image.NewRGBA(image.Rect(0, 0, width_pixels, height_pixels))
	colors := make(map[int]color.RGBA, 256)

	//https://components.ai/theme/7JnLZLghNSsT34IF969D
	/*
			colors[0] = color.RGBA{68, 1, 84, 255}   // green
			colors[1] = color.RGBA{72, 40, 120, 255} // limegreen
			colors[2] = color.RGBA{62, 74, 137, 255}
			colors[3] = color.RGBA{49, 104, 142, 255}
			colors[4] = color.RGBA{28, 130, 142, 255}
			colors[5] = color.RGBA{31, 158, 137, 255}
			colors[6] = color.RGBA{53, 183, 121, 255}
			colors[7] = color.RGBA{109, 205, 189, 255}
			colors[8] = color.RGBA{180, 222, 44, 255}
			colors[9] = color.RGBA{253, 231, 37, 255}

		colors[0] = color.RGBA{0, 32, 77, 255}
		colors[1] = color.RGBA{0, 33, 78, 255}
		colors[2] = color.RGBA{0, 34, 80, 255}
		colors[3] = color.RGBA{0, 34, 82, 255}
		colors[4] = color.RGBA{0, 35, 83, 255}
		colors[5] = color.RGBA{0, 36, 85, 255}
		colors[6] = color.RGBA{0, 37, 87, 255}
		colors[7] = color.RGBA{0, 37, 88, 255}
		colors[8] = color.RGBA{0, 38, 90, 255}
		colors[9] = color.RGBA{0, 39, 92, 255}
		colors[10] = color.RGBA{0, 39, 94, 255}
		colors[11] = color.RGBA{0, 40, 96, 255}
		colors[12] = color.RGBA{0, 41, 97, 255}
		colors[13] = color.RGBA{0, 42, 99, 255}
		colors[14] = color.RGBA{0, 42, 101, 255}
		colors[15] = color.RGBA{0, 43, 103, 255}
		colors[16] = color.RGBA{0, 44, 105, 255}
		colors[17] = color.RGBA{0, 44, 106, 255}
		colors[18] = color.RGBA{0, 45, 108, 255}
		colors[19] = color.RGBA{0, 46, 110, 255}
		colors[20] = color.RGBA{0, 46, 111, 255}
		colors[21] = color.RGBA{0, 47, 111, 255}
		colors[22] = color.RGBA{0, 47, 111, 255}
		colors[23] = color.RGBA{0, 48, 111, 255}
		colors[24] = color.RGBA{0, 48, 111, 255}
		colors[25] = color.RGBA{0, 49, 111, 255}
		colors[26] = color.RGBA{0, 50, 111, 255}
		colors[27] = color.RGBA{0, 51, 111, 255}
		colors[28] = color.RGBA{0, 51, 111, 255}
		colors[29] = color.RGBA{0, 52, 111, 255}
		colors[30] = color.RGBA{0, 53, 110, 255}
		colors[31] = color.RGBA{1, 54, 110, 255}
		colors[32] = color.RGBA{6, 54, 110, 255}
		colors[33] = color.RGBA{11, 55, 110, 255}
		colors[34] = color.RGBA{15, 56, 110, 255}
		colors[35] = color.RGBA{18, 56, 109, 255}
		colors[36] = color.RGBA{21, 57, 109, 255}
		colors[37] = color.RGBA{24, 58, 109, 255}
		colors[38] = color.RGBA{26, 59, 109, 255}
		colors[39] = color.RGBA{29, 59, 109, 255}
		colors[40] = color.RGBA{31, 60, 109, 255}
		colors[41] = color.RGBA{33, 61, 109, 255}
		colors[42] = color.RGBA{35, 62, 108, 255}
		colors[43] = color.RGBA{36, 62, 108, 255}
		colors[44] = color.RGBA{38, 63, 108, 255}
		colors[45] = color.RGBA{40, 64, 108, 255}
		colors[46] = color.RGBA{42, 64, 108, 255}
		colors[47] = color.RGBA{43, 65, 108, 255}
		colors[48] = color.RGBA{45, 66, 108, 255}
		colors[49] = color.RGBA{46, 67, 108, 255}
		colors[50] = color.RGBA{48, 67, 108, 255}
		colors[51] = color.RGBA{49, 68, 107, 255}
		colors[52] = color.RGBA{50, 69, 107, 255}
		colors[53] = color.RGBA{52, 69, 107, 255}
		colors[54] = color.RGBA{53, 70, 107, 255}
		colors[55] = color.RGBA{54, 71, 107, 255}
		colors[56] = color.RGBA{56, 72, 107, 255}
		colors[57] = color.RGBA{57, 72, 107, 255}
		colors[58] = color.RGBA{58, 73, 107, 255}
		colors[59] = color.RGBA{59, 74, 107, 255}
		colors[60] = color.RGBA{61, 74, 107, 255}
		colors[61] = color.RGBA{62, 75, 107, 255}
		colors[62] = color.RGBA{63, 76, 107, 255}
		colors[63] = color.RGBA{64, 77, 107, 255}
		colors[64] = color.RGBA{65, 77, 107, 255}
		colors[65] = color.RGBA{66, 78, 107, 255}
		colors[66] = color.RGBA{67, 79, 107, 255}
		colors[67] = color.RGBA{68, 79, 107, 255}
		colors[68] = color.RGBA{70, 80, 107, 255}
		colors[69] = color.RGBA{71, 81, 107, 255}
		colors[70] = color.RGBA{72, 82, 107, 255}
		colors[71] = color.RGBA{73, 82, 107, 255}
		colors[72] = color.RGBA{74, 83, 107, 255}
		colors[73] = color.RGBA{75, 84, 108, 255}
		colors[74] = color.RGBA{76, 84, 108, 255}
		colors[75] = color.RGBA{77, 85, 108, 255}
		colors[76] = color.RGBA{78, 86, 108, 255}
		colors[77] = color.RGBA{79, 87, 108, 255}
		colors[78] = color.RGBA{80, 87, 108, 255}
		colors[79] = color.RGBA{81, 88, 108, 255}
		colors[80] = color.RGBA{82, 89, 108, 255}
		colors[81] = color.RGBA{83, 89, 108, 255}
		colors[82] = color.RGBA{84, 90, 108, 255}
		colors[83] = color.RGBA{85, 91, 109, 255}
		colors[84] = color.RGBA{86, 92, 109, 255}
		colors[85] = color.RGBA{87, 92, 109, 255}
		colors[86] = color.RGBA{88, 93, 109, 255}
		colors[87] = color.RGBA{89, 94, 109, 255}
		colors[88] = color.RGBA{89, 95, 109, 255}
		colors[89] = color.RGBA{90, 95, 109, 255}
		colors[90] = color.RGBA{91, 96, 110, 255}
		colors[91] = color.RGBA{92, 97, 110, 255}
		colors[92] = color.RGBA{93, 97, 110, 255}
		colors[93] = color.RGBA{94, 98, 110, 255}
		colors[94] = color.RGBA{95, 99, 110, 255}
		colors[95] = color.RGBA{96, 100, 111, 255}
		colors[96] = color.RGBA{97, 100, 111, 255}
		colors[97] = color.RGBA{98, 101, 111, 255}
		colors[98] = color.RGBA{99, 102, 111, 255}
		colors[99] = color.RGBA{100, 102, 111, 255}
		colors[100] = color.RGBA{100, 103, 112, 255}
		colors[101] = color.RGBA{101, 104, 112, 255}
		colors[102] = color.RGBA{102, 105, 112, 255}
		colors[103] = color.RGBA{103, 105, 112, 255}
		colors[104] = color.RGBA{104, 106, 113, 255}
		colors[105] = color.RGBA{105, 107, 113, 255}
		colors[106] = color.RGBA{106, 108, 113, 255}
		colors[107] = color.RGBA{107, 108, 113, 255}
		colors[108] = color.RGBA{108, 109, 114, 255}
		colors[109] = color.RGBA{108, 110, 114, 255}
		colors[110] = color.RGBA{109, 110, 114, 255}
		colors[111] = color.RGBA{110, 111, 115, 255}
		colors[112] = color.RGBA{111, 112, 115, 255}
		colors[113] = color.RGBA{112, 113, 115, 255}
		colors[114] = color.RGBA{113, 113, 116, 255}
		colors[115] = color.RGBA{114, 114, 116, 255}
		colors[116] = color.RGBA{114, 115, 116, 255}
		colors[117] = color.RGBA{115, 116, 117, 255}
		colors[118] = color.RGBA{116, 116, 117, 255}
		colors[119] = color.RGBA{117, 117, 117, 255}
		colors[120] = color.RGBA{118, 118, 118, 255}
		colors[121] = color.RGBA{119, 119, 118, 255}
		colors[122] = color.RGBA{120, 119, 119, 255}
		colors[123] = color.RGBA{120, 120, 119, 255}
		colors[124] = color.RGBA{121, 121, 119, 255}
		colors[125] = color.RGBA{122, 122, 120, 255}
		colors[126] = color.RGBA{123, 122, 120, 255}
		colors[127] = color.RGBA{124, 123, 120, 255}
		colors[128] = color.RGBA{125, 124, 120, 255}
		colors[129] = color.RGBA{126, 125, 120, 255}
		colors[130] = color.RGBA{127, 125, 120, 255}
		colors[131] = color.RGBA{128, 126, 121, 255}
		colors[132] = color.RGBA{129, 127, 121, 255}
		colors[133] = color.RGBA{130, 128, 121, 255}
		colors[134] = color.RGBA{131, 128, 121, 255}
		colors[135] = color.RGBA{132, 129, 121, 255}
		colors[136] = color.RGBA{132, 130, 121, 255}
		colors[137] = color.RGBA{133, 131, 121, 255}
		colors[138] = color.RGBA{134, 131, 121, 255}
		colors[139] = color.RGBA{135, 132, 121, 255}
		colors[140] = color.RGBA{136, 133, 121, 255}
		colors[141] = color.RGBA{137, 134, 121, 255}
		colors[142] = color.RGBA{138, 135, 121, 255}
		colors[143] = color.RGBA{139, 135, 121, 255}
		colors[144] = color.RGBA{140, 136, 121, 255}
		colors[145] = color.RGBA{141, 137, 121, 255}
		colors[146] = color.RGBA{142, 138, 121, 255}
		colors[147] = color.RGBA{143, 138, 121, 255}
		colors[148] = color.RGBA{144, 139, 121, 255}
		colors[149] = color.RGBA{145, 140, 120, 255}
		colors[150] = color.RGBA{146, 141, 120, 255}
		colors[151] = color.RGBA{147, 142, 120, 255}
		colors[152] = color.RGBA{148, 142, 120, 255}
		colors[153] = color.RGBA{149, 143, 120, 255}
		colors[154] = color.RGBA{150, 144, 120, 255}
		colors[155] = color.RGBA{151, 145, 120, 255}
		colors[156] = color.RGBA{152, 146, 120, 255}
		colors[157] = color.RGBA{153, 146, 120, 255}
		colors[158] = color.RGBA{154, 147, 119, 255}
		colors[159] = color.RGBA{155, 148, 119, 255}
		colors[160] = color.RGBA{156, 149, 119, 255}
		colors[161] = color.RGBA{157, 150, 119, 255}
		colors[162] = color.RGBA{158, 150, 119, 255}
		colors[163] = color.RGBA{159, 151, 119, 255}
		colors[164] = color.RGBA{160, 152, 119, 255}
		colors[165] = color.RGBA{161, 153, 118, 255}
		colors[166] = color.RGBA{162, 154, 118, 255}
		colors[167] = color.RGBA{163, 154, 118, 255}
		colors[168] = color.RGBA{164, 155, 118, 255}
		colors[169] = color.RGBA{165, 156, 118, 255}
		colors[170] = color.RGBA{166, 157, 117, 255}
		colors[171] = color.RGBA{168, 158, 117, 255}
		colors[172] = color.RGBA{169, 159, 117, 255}
		colors[173] = color.RGBA{170, 159, 117, 255}
		colors[174] = color.RGBA{171, 160, 116, 255}
		colors[175] = color.RGBA{172, 161, 116, 255}
		colors[176] = color.RGBA{173, 162, 116, 255}
		colors[177] = color.RGBA{174, 163, 116, 255}
		colors[178] = color.RGBA{175, 164, 115, 255}
		colors[179] = color.RGBA{176, 164, 115, 255}
		colors[180] = color.RGBA{177, 165, 115, 255}
		colors[181] = color.RGBA{178, 166, 114, 255}
		colors[182] = color.RGBA{179, 167, 114, 255}
		colors[183] = color.RGBA{180, 168, 114, 255}
		colors[184] = color.RGBA{181, 169, 113, 255}
		colors[185] = color.RGBA{182, 169, 113, 255}
		colors[186] = color.RGBA{183, 170, 113, 255}
		colors[187] = color.RGBA{184, 171, 112, 255}
		colors[188] = color.RGBA{185, 172, 112, 255}
		colors[189] = color.RGBA{186, 173, 112, 255}
		colors[190] = color.RGBA{187, 174, 111, 255}
		colors[191] = color.RGBA{188, 175, 111, 255}
		colors[192] = color.RGBA{190, 175, 111, 255}
		colors[193] = color.RGBA{191, 176, 110, 255}
		colors[194] = color.RGBA{192, 177, 110, 255}
		colors[195] = color.RGBA{193, 178, 109, 255}
		colors[196] = color.RGBA{194, 179, 109, 255}
		colors[197] = color.RGBA{195, 180, 109, 255}
		colors[198] = color.RGBA{196, 181, 108, 255}
		colors[199] = color.RGBA{197, 181, 108, 255}
		colors[200] = color.RGBA{198, 182, 107, 255}
		colors[201] = color.RGBA{199, 183, 107, 255}
		colors[202] = color.RGBA{200, 184, 106, 255}
		colors[203] = color.RGBA{201, 185, 106, 255}
		colors[204] = color.RGBA{203, 186, 105, 255}
		colors[205] = color.RGBA{204, 187, 105, 255}
		colors[206] = color.RGBA{205, 188, 104, 255}
		colors[207] = color.RGBA{206, 188, 104, 255}
		colors[208] = color.RGBA{207, 189, 103, 255}
		colors[209] = color.RGBA{208, 190, 103, 255}
		colors[210] = color.RGBA{209, 191, 102, 255}
		colors[211] = color.RGBA{210, 192, 102, 255}
		colors[212] = color.RGBA{211, 193, 101, 255}
		colors[213] = color.RGBA{212, 194, 100, 255}
		colors[214] = color.RGBA{214, 195, 100, 255}
		colors[215] = color.RGBA{215, 196, 99, 255}
		colors[216] = color.RGBA{216, 197, 99, 255}
		colors[217] = color.RGBA{217, 197, 98, 255}
		colors[218] = color.RGBA{218, 198, 97, 255}
		colors[219] = color.RGBA{219, 199, 97, 255}
		colors[220] = color.RGBA{220, 200, 96, 255}
		colors[221] = color.RGBA{221, 201, 95, 255}
		colors[222] = color.RGBA{222, 202, 95, 255}
		colors[223] = color.RGBA{224, 203, 94, 255}
		colors[224] = color.RGBA{225, 204, 93, 255}
		colors[225] = color.RGBA{226, 205, 92, 255}
		colors[226] = color.RGBA{227, 206, 92, 255}
		colors[227] = color.RGBA{228, 207, 91, 255}
		colors[228] = color.RGBA{229, 208, 90, 255}
		colors[229] = color.RGBA{230, 209, 89, 255}
		colors[230] = color.RGBA{232, 210, 89, 255}
		colors[231] = color.RGBA{233, 211, 88, 255}
		colors[232] = color.RGBA{234, 211, 87, 255}
		colors[233] = color.RGBA{235, 212, 86, 255}
		colors[234] = color.RGBA{236, 213, 85, 255}
		colors[235] = color.RGBA{237, 214, 84, 255}
		colors[236] = color.RGBA{239, 215, 83, 255}
		colors[237] = color.RGBA{240, 216, 82, 255}
		colors[238] = color.RGBA{241, 217, 81, 255}
		colors[239] = color.RGBA{242, 218, 80, 255}
		colors[240] = color.RGBA{243, 219, 79, 255}
		colors[241] = color.RGBA{244, 220, 78, 255}
		colors[242] = color.RGBA{246, 221, 77, 255}
		colors[243] = color.RGBA{247, 222, 76, 255}
		colors[244] = color.RGBA{248, 223, 75, 255}
		colors[245] = color.RGBA{249, 224, 74, 255}
		colors[246] = color.RGBA{250, 225, 73, 255}
		colors[247] = color.RGBA{251, 226, 72, 255}
		colors[248] = color.RGBA{253, 227, 70, 255}
		colors[249] = color.RGBA{254, 228, 69, 255}
		colors[250] = color.RGBA{255, 229, 68, 255}
		colors[251] = color.RGBA{255, 230, 66, 255}
		colors[252] = color.RGBA{255, 231, 66, 255}
		colors[253] = color.RGBA{255, 232, 67, 255}
		colors[254] = color.RGBA{255, 233, 68, 255}
		colors[255] = color.RGBA{255, 234, 70, 255}



			camColors := [] byte {0x480F,
				0x400F,0x400F,0x400F,0x4010,0x3810,0x3810,0x3810,0x3810,0x3010,0x3010,
				0x3010,0x2810,0x2810,0x2810,0x2810,0x2010,0x2010,0x2010,0x1810,0x1810,
				0x1811,0x1811,0x1011,0x1011,0x1011,0x0811,0x0811,0x0811,0x0011,0x0011,
				0x0011,0x0011,0x0011,0x0031,0x0031,0x0051,0x0072,0x0072,0x0092,0x00B2,
				0x00B2,0x00D2,0x00F2,0x00F2,0x0112,0x0132,0x0152,0x0152,0x0172,0x0192,
				0x0192,0x01B2,0x01D2,0x01F3,0x01F3,0x0213,0x0233,0x0253,0x0253,0x0273,
				0x0293,0x02B3,0x02D3,0x02D3,0x02F3,0x0313,0x0333,0x0333,0x0353,0x0373,
				0x0394,0x03B4,0x03D4,0x03D4,0x03F4,0x0414,0x0434,0x0454,0x0474,0x0474,
				0x0494,0x04B4,0x04D4,0x04F4,0x0514,0x0534,0x0534,0x0554,0x0554,0x0574,
				0x0574,0x0573,0x0573,0x0573,0x0572,0x0572,0x0572,0x0571,0x0591,0x0591,
				0x0590,0x0590,0x058F,0x058F,0x058F,0x058E,0x05AE,0x05AE,0x05AD,0x05AD,
				0x05AD,0x05AC,0x05AC,0x05AB,0x05CB,0x05CB,0x05CA,0x05CA,0x05CA,0x05C9,
				0x05C9,0x05C8,0x05E8,0x05E8,0x05E7,0x05E7,0x05E6,0x05E6,0x05E6,0x05E5,
				0x05E5,0x0604,0x0604,0x0604,0x0603,0x0603,0x0602,0x0602,0x0601,0x0621,
				0x0621,0x0620,0x0620,0x0620,0x0620,0x0E20,0x0E20,0x0E40,0x1640,0x1640,
				0x1E40,0x1E40,0x2640,0x2640,0x2E40,0x2E60,0x3660,0x3660,0x3E60,0x3E60,
				0x3E60,0x4660,0x4660,0x4E60,0x4E80,0x5680,0x5680,0x5E80,0x5E80,0x6680,
				0x6680,0x6E80,0x6EA0,0x76A0,0x76A0,0x7EA0,0x7EA0,0x86A0,0x86A0,0x8EA0,
				0x8EC0,0x96C0,0x96C0,0x9EC0,0x9EC0,0xA6C0,0xAEC0,0xAEC0,0xB6E0,0xB6E0,
				0xBEE0,0xBEE0,0xC6E0,0xC6E0,0xCEE0,0xCEE0,0xD6E0,0xD700,0xDF00,0xDEE0,
				0xDEC0,0xDEA0,0xDE80,0xDE80,0xE660,0xE640,0xE620,0xE600,0xE5E0,0xE5C0,
				0xE5A0,0xE580,0xE560,0xE540,0xE520,0xE500,0xE4E0,0xE4C0,0xE4A0,0xE480,
				0xE460,0xEC40,0xEC20,0xEC00,0xEBE0,0xEBC0,0xEBA0,0xEB80,0xEB60,0xEB40,
				0xEB20,0xEB00,0xEAE0,0xEAC0,0xEAA0,0xEA80,0xEA60,0xEA40,0xF220,0xF200,
				0xF1E0,0xF1C0,0xF1A0,0xF180,0xF160,0xF140,0xF100,0xF0E0,0xF0C0,0xF0A0,
				0xF080,0xF060,0xF040,0xF020,0xF800,};

	*/

	colors[255] = color.RGBA{253, 231, 37, 255}
	colors[254] = color.RGBA{251, 231, 35, 255}
	colors[253] = color.RGBA{248, 230, 33, 255}
	colors[252] = color.RGBA{246, 230, 32, 255}
	colors[251] = color.RGBA{244, 230, 30, 255}
	colors[250] = color.RGBA{241, 229, 29, 255}
	colors[249] = color.RGBA{239, 229, 28, 255}
	colors[248] = color.RGBA{236, 229, 27, 255}
	colors[247] = color.RGBA{234, 229, 26, 255}
	colors[246] = color.RGBA{231, 228, 25, 255}
	colors[245] = color.RGBA{229, 228, 25, 255}
	colors[244] = color.RGBA{226, 228, 24, 255}
	colors[243] = color.RGBA{223, 227, 24, 255}
	colors[242] = color.RGBA{221, 227, 24, 255}
	colors[241] = color.RGBA{218, 227, 25, 255}
	colors[240] = color.RGBA{216, 226, 25, 255}
	colors[239] = color.RGBA{213, 226, 26, 255}
	colors[238] = color.RGBA{210, 226, 27, 255}
	colors[237] = color.RGBA{208, 225, 28, 255}
	colors[236] = color.RGBA{205, 225, 29, 255}
	colors[235] = color.RGBA{202, 225, 31, 255}
	colors[234] = color.RGBA{200, 224, 32, 255}
	colors[233] = color.RGBA{197, 224, 33, 255}
	colors[232] = color.RGBA{194, 223, 35, 255}
	colors[231] = color.RGBA{192, 223, 37, 255}
	colors[230] = color.RGBA{189, 223, 38, 255}
	colors[229] = color.RGBA{186, 222, 40, 255}
	colors[228] = color.RGBA{184, 222, 41, 255}
	colors[227] = color.RGBA{181, 222, 43, 255}
	colors[226] = color.RGBA{178, 221, 45, 255}
	colors[225] = color.RGBA{176, 221, 47, 255}
	colors[224] = color.RGBA{173, 220, 48, 255}
	colors[223] = color.RGBA{170, 220, 50, 255}
	colors[222] = color.RGBA{168, 219, 52, 255}
	colors[221] = color.RGBA{165, 219, 54, 255}
	colors[220] = color.RGBA{162, 218, 55, 255}
	colors[219] = color.RGBA{160, 218, 57, 255}
	colors[218] = color.RGBA{157, 217, 59, 255}
	colors[217] = color.RGBA{155, 217, 60, 255}
	colors[216] = color.RGBA{152, 216, 62, 255}
	colors[215] = color.RGBA{149, 216, 64, 255}
	colors[214] = color.RGBA{147, 215, 65, 255}
	colors[213] = color.RGBA{144, 215, 67, 255}
	colors[212] = color.RGBA{142, 214, 69, 255}
	colors[211] = color.RGBA{139, 214, 70, 255}
	colors[210] = color.RGBA{137, 213, 72, 255}
	colors[209] = color.RGBA{134, 213, 73, 255}
	colors[208] = color.RGBA{132, 212, 75, 255}
	colors[207] = color.RGBA{129, 211, 77, 255}
	colors[206] = color.RGBA{127, 211, 78, 255}
	colors[205] = color.RGBA{124, 210, 80, 255}
	colors[204] = color.RGBA{122, 209, 81, 255}
	colors[203] = color.RGBA{119, 209, 83, 255}
	colors[202] = color.RGBA{117, 208, 84, 255}
	colors[201] = color.RGBA{115, 208, 86, 255}
	colors[200] = color.RGBA{112, 207, 87, 255}
	colors[199] = color.RGBA{110, 206, 88, 255}
	colors[198] = color.RGBA{108, 205, 90, 255}
	colors[197] = color.RGBA{105, 205, 91, 255}
	colors[196] = color.RGBA{103, 204, 92, 255}
	colors[195] = color.RGBA{101, 203, 94, 255}
	colors[194] = color.RGBA{99, 203, 95, 255}
	colors[193] = color.RGBA{96, 202, 96, 255}
	colors[192] = color.RGBA{94, 201, 98, 255}
	colors[191] = color.RGBA{92, 200, 99, 255}
	colors[190] = color.RGBA{90, 200, 100, 255}
	colors[189] = color.RGBA{88, 199, 101, 255}
	colors[188] = color.RGBA{86, 198, 103, 255}
	colors[187] = color.RGBA{84, 197, 104, 255}
	colors[186] = color.RGBA{82, 197, 105, 255}
	colors[185] = color.RGBA{80, 196, 106, 255}
	colors[184] = color.RGBA{78, 195, 107, 255}
	colors[183] = color.RGBA{76, 194, 108, 255}
	colors[182] = color.RGBA{74, 193, 109, 255}
	colors[181] = color.RGBA{72, 193, 110, 255}
	colors[180] = color.RGBA{70, 192, 111, 255}
	colors[179] = color.RGBA{68, 191, 112, 255}
	colors[178] = color.RGBA{66, 190, 113, 255}
	colors[177] = color.RGBA{64, 189, 114, 255}
	colors[176] = color.RGBA{63, 188, 115, 255}
	colors[175] = color.RGBA{61, 188, 116, 255}
	colors[174] = color.RGBA{59, 187, 117, 255}
	colors[173] = color.RGBA{58, 186, 118, 255}
	colors[172] = color.RGBA{56, 185, 119, 255}
	colors[171] = color.RGBA{55, 184, 120, 255}
	colors[170] = color.RGBA{53, 183, 121, 255}
	colors[169] = color.RGBA{52, 182, 121, 255}
	colors[168] = color.RGBA{50, 182, 122, 255}
	colors[167] = color.RGBA{49, 181, 123, 255}
	colors[166] = color.RGBA{47, 180, 124, 255}
	colors[165] = color.RGBA{46, 179, 124, 255}
	colors[164] = color.RGBA{45, 178, 125, 255}
	colors[163] = color.RGBA{44, 177, 126, 255}
	colors[162] = color.RGBA{42, 176, 127, 255}
	colors[161] = color.RGBA{41, 175, 127, 255}
	colors[160] = color.RGBA{40, 174, 128, 255}
	colors[159] = color.RGBA{39, 173, 129, 255}
	colors[158] = color.RGBA{38, 173, 129, 255}
	colors[157] = color.RGBA{37, 172, 130, 255}
	colors[156] = color.RGBA{37, 171, 130, 255}
	colors[155] = color.RGBA{36, 170, 131, 255}
	colors[154] = color.RGBA{35, 169, 131, 255}
	colors[153] = color.RGBA{34, 168, 132, 255}
	colors[152] = color.RGBA{34, 167, 133, 255}
	colors[151] = color.RGBA{33, 166, 133, 255}
	colors[150] = color.RGBA{33, 165, 133, 255}
	colors[149] = color.RGBA{32, 164, 134, 255}
	colors[148] = color.RGBA{32, 163, 134, 255}
	colors[147] = color.RGBA{31, 162, 135, 255}
	colors[146] = color.RGBA{31, 161, 135, 255}
	colors[145] = color.RGBA{31, 161, 136, 255}
	colors[144] = color.RGBA{31, 160, 136, 255}
	colors[143] = color.RGBA{31, 159, 136, 255}
	colors[142] = color.RGBA{31, 158, 137, 255}
	colors[141] = color.RGBA{30, 157, 137, 255}
	colors[140] = color.RGBA{30, 156, 137, 255}
	colors[139] = color.RGBA{30, 155, 138, 255}
	colors[138] = color.RGBA{31, 154, 138, 255}
	colors[137] = color.RGBA{31, 153, 138, 255}
	colors[136] = color.RGBA{31, 152, 139, 255}
	colors[135] = color.RGBA{31, 151, 139, 255}
	colors[134] = color.RGBA{31, 150, 139, 255}
	colors[133] = color.RGBA{31, 149, 139, 255}
	colors[132] = color.RGBA{31, 148, 140, 255}
	colors[131] = color.RGBA{32, 147, 140, 255}
	colors[130] = color.RGBA{32, 146, 140, 255}
	colors[129] = color.RGBA{32, 146, 140, 255}
	colors[128] = color.RGBA{33, 145, 140, 255}
	colors[127] = color.RGBA{33, 144, 141, 255}
	colors[126] = color.RGBA{33, 143, 141, 255}
	colors[125] = color.RGBA{33, 142, 141, 255}
	colors[124] = color.RGBA{34, 141, 141, 255}
	colors[123] = color.RGBA{34, 140, 141, 255}
	colors[122] = color.RGBA{34, 139, 141, 255}
	colors[121] = color.RGBA{35, 138, 141, 255}
	colors[120] = color.RGBA{35, 137, 142, 255}
	colors[119] = color.RGBA{35, 136, 142, 255}
	colors[118] = color.RGBA{36, 135, 142, 255}
	colors[117] = color.RGBA{36, 134, 142, 255}
	colors[116] = color.RGBA{37, 133, 142, 255}
	colors[115] = color.RGBA{37, 132, 142, 255}
	colors[114] = color.RGBA{37, 131, 142, 255}
	colors[113] = color.RGBA{38, 130, 142, 255}
	colors[112] = color.RGBA{38, 130, 142, 255}
	colors[111] = color.RGBA{38, 129, 142, 255}
	colors[110] = color.RGBA{39, 128, 142, 255}
	colors[109] = color.RGBA{39, 127, 142, 255}
	colors[108] = color.RGBA{39, 126, 142, 255}
	colors[107] = color.RGBA{40, 125, 142, 255}
	colors[106] = color.RGBA{40, 124, 142, 255}
	colors[105] = color.RGBA{41, 123, 142, 255}
	colors[104] = color.RGBA{41, 122, 142, 255}
	colors[103] = color.RGBA{41, 121, 142, 255}
	colors[102] = color.RGBA{42, 120, 142, 255}
	colors[101] = color.RGBA{42, 119, 142, 255}
	colors[100] = color.RGBA{42, 118, 142, 255}
	colors[99] = color.RGBA{43, 117, 142, 255}
	colors[98] = color.RGBA{43, 116, 142, 255}
	colors[97] = color.RGBA{44, 115, 142, 255}
	colors[96] = color.RGBA{44, 114, 142, 255}
	colors[95] = color.RGBA{44, 113, 142, 255}
	colors[94] = color.RGBA{45, 113, 142, 255}
	colors[93] = color.RGBA{45, 112, 142, 255}
	colors[92] = color.RGBA{46, 111, 142, 255}
	colors[91] = color.RGBA{46, 110, 142, 255}
	colors[90] = color.RGBA{46, 109, 142, 255}
	colors[89] = color.RGBA{47, 108, 142, 255}
	colors[88] = color.RGBA{47, 107, 142, 255}
	colors[87] = color.RGBA{48, 106, 142, 255}
	colors[86] = color.RGBA{48, 105, 142, 255}
	colors[85] = color.RGBA{49, 104, 142, 255}
	colors[84] = color.RGBA{49, 103, 142, 255}
	colors[83] = color.RGBA{49, 102, 142, 255}
	colors[82] = color.RGBA{50, 101, 142, 255}
	colors[81] = color.RGBA{50, 100, 142, 255}
	colors[80] = color.RGBA{51, 99, 141, 255}
	colors[79] = color.RGBA{51, 98, 141, 255}
	colors[78] = color.RGBA{52, 97, 141, 255}
	colors[77] = color.RGBA{52, 96, 141, 255}
	colors[76] = color.RGBA{53, 95, 141, 255}
	colors[75] = color.RGBA{53, 94, 141, 255}
	colors[74] = color.RGBA{54, 93, 141, 255}
	colors[73] = color.RGBA{54, 92, 141, 255}
	colors[72] = color.RGBA{55, 91, 141, 255}
	colors[71] = color.RGBA{55, 90, 140, 255}
	colors[70] = color.RGBA{56, 89, 140, 255}
	colors[69] = color.RGBA{56, 88, 140, 255}
	colors[68] = color.RGBA{57, 86, 140, 255}
	colors[67] = color.RGBA{57, 85, 140, 255}
	colors[66] = color.RGBA{58, 84, 140, 255}
	colors[65] = color.RGBA{58, 83, 139, 255}
	colors[64] = color.RGBA{59, 82, 139, 255}
	colors[63] = color.RGBA{59, 81, 139, 255}
	colors[62] = color.RGBA{60, 80, 139, 255}
	colors[61] = color.RGBA{60, 79, 138, 255}
	colors[60] = color.RGBA{61, 78, 138, 255}
	colors[59] = color.RGBA{61, 77, 138, 255}
	colors[58] = color.RGBA{62, 76, 138, 255}
	colors[57] = color.RGBA{62, 74, 137, 255}
	colors[56] = color.RGBA{62, 73, 137, 255}
	colors[55] = color.RGBA{63, 72, 137, 255}
	colors[54] = color.RGBA{63, 71, 136, 255}
	colors[53] = color.RGBA{64, 70, 136, 255}
	colors[52] = color.RGBA{64, 69, 136, 255}
	colors[51] = color.RGBA{65, 68, 135, 255}
	colors[50] = color.RGBA{65, 66, 135, 255}
	colors[49] = color.RGBA{66, 65, 134, 255}
	colors[48] = color.RGBA{66, 64, 134, 255}
	colors[47] = color.RGBA{66, 63, 133, 255}
	colors[46] = color.RGBA{67, 62, 133, 255}
	colors[45] = color.RGBA{67, 61, 132, 255}
	colors[44] = color.RGBA{68, 59, 132, 255}
	colors[43] = color.RGBA{68, 58, 131, 255}
	colors[42] = color.RGBA{68, 57, 131, 255}
	colors[41] = color.RGBA{69, 56, 130, 255}
	colors[40] = color.RGBA{69, 55, 129, 255}
	colors[39] = color.RGBA{69, 53, 129, 255}
	colors[38] = color.RGBA{70, 52, 128, 255}
	colors[37] = color.RGBA{70, 51, 127, 255}
	colors[36] = color.RGBA{70, 50, 126, 255}
	colors[35] = color.RGBA{70, 48, 126, 255}
	colors[34] = color.RGBA{71, 47, 125, 255}
	colors[33] = color.RGBA{71, 46, 124, 255}
	colors[32] = color.RGBA{71, 45, 123, 255}
	colors[31] = color.RGBA{71, 44, 122, 255}
	colors[30] = color.RGBA{71, 42, 122, 255}
	colors[29] = color.RGBA{72, 41, 121, 255}
	colors[28] = color.RGBA{72, 40, 120, 255}
	colors[27] = color.RGBA{72, 38, 119, 255}
	colors[26] = color.RGBA{72, 37, 118, 255}
	colors[25] = color.RGBA{72, 36, 117, 255}
	colors[24] = color.RGBA{72, 35, 116, 255}
	colors[23] = color.RGBA{72, 33, 115, 255}
	colors[22] = color.RGBA{72, 32, 113, 255}
	colors[21] = color.RGBA{72, 31, 112, 255}
	colors[20] = color.RGBA{72, 29, 111, 255}
	colors[19] = color.RGBA{72, 28, 110, 255}
	colors[18] = color.RGBA{72, 27, 109, 255}
	colors[17] = color.RGBA{72, 26, 108, 255}
	colors[16] = color.RGBA{72, 24, 106, 255}
	colors[15] = color.RGBA{72, 23, 105, 255}
	colors[14] = color.RGBA{72, 22, 104, 255}
	colors[13] = color.RGBA{72, 20, 103, 255}
	colors[12] = color.RGBA{71, 19, 101, 255}
	colors[11] = color.RGBA{71, 17, 100, 255}
	colors[10] = color.RGBA{71, 16, 99, 255}
	colors[9] = color.RGBA{71, 14, 97, 255}
	colors[8] = color.RGBA{71, 13, 96, 255}
	colors[7] = color.RGBA{70, 11, 94, 255}
	colors[6] = color.RGBA{70, 10, 93, 255}
	colors[5] = color.RGBA{70, 8, 92, 255}
	colors[4] = color.RGBA{70, 7, 90, 255}
	colors[3] = color.RGBA{69, 5, 89, 255}
	colors[2] = color.RGBA{69, 4, 87, 255}
	colors[1] = color.RGBA{68, 2, 86, 255}
	colors[0] = color.RGBA{68, 1, 84, 255}

	size_board := 32

	for i := 0; i < len(irdata); i++ {
		x := i % size_board
		y := i / size_board
		//val := int(irdata[i])
		val := int(irdata[i])

		//fmt.Println(x, y, val)
		//valcolor := colors[i%10]
		//valcolor := colors[val%10]
		//valcolor := color.RGBA{val, val, val, 255}

		zz := int(val)
		z := int(val - 161)

		if z < 0 {
			z = 0
		}
		if z > 9 {
			z = 9
		}
		valcolor := colors[val]
		fmt.Println("x=", x, "y=", y, "zz=", zz, "valcolor=", valcolor, "i=", i)

		draw.Draw(myimage, image.Rect(x*size_block, y*size_block, x*size_block+size_block, y*size_block+size_block), &image.Uniform{valcolor}, image.ZP, draw.Src)
	}
	return myimage

	/*
		myfile, err := os.Create(new_png_file)
		if err != nil {
			panic(err.Error())
		}
		defer myfile.Close()
		png.Encode(myfile, myimage)                 // ... save image
		fmt.Println("new ir immage ", new_png_file) // view image issue : firefox  /tmp/chessboard.png
	*/
}
