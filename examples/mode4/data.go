package main

import "github.com/janezkenda/tinygba/display"

var (
	palette = []display.Color{
		0x0010,
		0x0000,
		0x7FE0,
		0x0300,
		0x0000,
		0x7C18,
	}

	frontBitmap = []uint32{
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x03030301, 0x01030303, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,

		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01030301, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x03010301, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,

		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010301, 0x01010103, 0x01010101, 0x01010101, 0x01010101, 0x01030303, 0x01010101, 0x03010101,
		0x01030303, 0x01010101, 0x01030303, 0x01010101, 0x03030101, 0x01010103, 0x01010101, 0x01030303,
		0x01010101, 0x03030101, 0x01010103, 0x01010101, 0x01030303, 0x01010101, 0x03030101, 0x01010103,
		0x03030301, 0x01010101, 0x01010101, 0x01010101, 0x03030301, 0x03030303, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010301, 0x01010301, 0x01010101, 0x01010101,
		0x03010101, 0x03030103, 0x01010101, 0x03030301, 0x01010101, 0x03010101, 0x03030103, 0x01010101,

		0x01030301, 0x01010303, 0x03010101, 0x03030103, 0x01010101, 0x01030301, 0x01010303, 0x03010101,
		0x03030103, 0x01010101, 0x01030301, 0x01010303, 0x03030101, 0x01010101, 0x01010101, 0x01010101,
		0x01030101, 0x03010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010301,
		0x01010101, 0x01030101, 0x01010101, 0x01010101, 0x03030101, 0x03010101, 0x01010103, 0x01030301,
		0x01010101, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101,
		0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301,
		0x03030101, 0x01010101, 0x01010101, 0x01010101, 0x01030101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010301, 0x01010101, 0x03010101, 0x03030303, 0x01030303,

		0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01010101, 0x03030101, 0x03010101, 0x01010103,
		0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101,
		0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03030301, 0x01010101, 0x01010101,
		0x01030101, 0x01010301, 0x01030101, 0x01010303, 0x01030303, 0x01030101, 0x01010303, 0x01030303,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x03030101, 0x03010101, 0x01010103, 0x03010303,
		0x01010303, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101,
		0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301,
		0x03030101, 0x03010103, 0x01010103, 0x01010101, 0x03030101, 0x01010303, 0x03030301, 0x03010101,

		0x03010101, 0x03030301, 0x01030101, 0x01010301, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x03030101, 0x03010101, 0x01010103, 0x01030303, 0x01030301, 0x03030101, 0x03010101, 0x01010103,
		0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101,
		0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010101,
		0x01030101, 0x01010301, 0x01030101, 0x03010101, 0x03010101, 0x01030101, 0x01030101, 0x01010301,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x03030101, 0x03010101, 0x01010103, 0x01010303,
		0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101,
		0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301,

		0x03030101, 0x03010101, 0x01010103, 0x01010101, 0x01030101, 0x01010101, 0x01030101, 0x03010101,
		0x03010101, 0x01030101, 0x01030101, 0x01010301, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103,
		0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101,
		0x03010101, 0x01010103, 0x01010303, 0x01030301, 0x03030101, 0x03010101, 0x01010103, 0x01010101,
		0x01030101, 0x01010101, 0x01030101, 0x03010101, 0x03010101, 0x01030101, 0x01030101, 0x01010301,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x03010101, 0x03030103, 0x01010101, 0x01030301,
		0x01030301, 0x03010101, 0x03030103, 0x01010101, 0x01030301, 0x01010303, 0x03010101, 0x03030103,

		0x01010101, 0x01030301, 0x01010303, 0x03010101, 0x03030103, 0x01010101, 0x01030301, 0x01010303,
		0x03030101, 0x03010101, 0x01010103, 0x01010101, 0x01030101, 0x01010101, 0x01030101, 0x03010101,
		0x03010101, 0x01030101, 0x01030101, 0x03010301, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01030303, 0x01010101, 0x03030101, 0x01010303, 0x01010101, 0x01030303, 0x01010101,
		0x03030101, 0x01010103, 0x01010101, 0x01030303, 0x01010101, 0x03030101, 0x01010103, 0x01010101,
		0x01030303, 0x01010101, 0x03030101, 0x01010103, 0x03030301, 0x03030103, 0x01010303, 0x01010101,
		0x03030301, 0x01010101, 0x03030301, 0x01010101, 0x01030303, 0x03030301, 0x03030301, 0x01030301,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,

		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
		0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101, 0x01010101,
	}

	backBitmap = []uint32{
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x04040402, 0x02040404, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,

		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02040402, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x04020402, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,

		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020402, 0x02020204, 0x02020202, 0x02020202, 0x02020202, 0x02040404, 0x02020202, 0x04020202,
		0x02040404, 0x02020202, 0x02040404, 0x02020202, 0x04040202, 0x02020204, 0x02020202, 0x02020202,
		0x02020202, 0x04040202, 0x02020204, 0x02020202, 0x02040404, 0x02020202, 0x04040202, 0x02020204,
		0x04040402, 0x02020202, 0x02020202, 0x02020202, 0x04040402, 0x02020404, 0x02020202, 0x02020202,
		0x02020202, 0x04020202, 0x02020202, 0x02020202, 0x02020402, 0x02020402, 0x02020202, 0x02020202,
		0x04020202, 0x04040204, 0x02020202, 0x04040402, 0x02020202, 0x04020202, 0x04040204, 0x02020202,

		0x02040402, 0x02020404, 0x02020202, 0x02020202, 0x02020202, 0x02040402, 0x02020404, 0x04020202,
		0x04040204, 0x02020202, 0x02040402, 0x02020404, 0x04040202, 0x02020202, 0x02020202, 0x02020202,
		0x02040202, 0x02040202, 0x02020202, 0x02020202, 0x02020202, 0x04040202, 0x02020202, 0x02020202,
		0x02020202, 0x02040202, 0x02020202, 0x02020202, 0x04040202, 0x04020202, 0x02020204, 0x02040402,
		0x02020202, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x02020202, 0x02020202,
		0x02020202, 0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402,
		0x04040202, 0x02020202, 0x02020202, 0x02020202, 0x02040202, 0x02040202, 0x02020202, 0x02020202,
		0x02020202, 0x04020202, 0x02020202, 0x02020202, 0x02020202, 0x04020202, 0x04040404, 0x02040404,

		0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02020202, 0x04040202, 0x04020202, 0x02020204,
		0x02020404, 0x02040402, 0x04020202, 0x04040404, 0x02020202, 0x02020404, 0x02040402, 0x04040202,
		0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04040202, 0x04040402, 0x02020202, 0x02020202,
		0x02040202, 0x02040202, 0x04040202, 0x02020204, 0x04040402, 0x04020202, 0x04040202, 0x02020204,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x04040202, 0x04020202, 0x02020204, 0x04020404,
		0x02020404, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x02020202, 0x04020202,
		0x02020204, 0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402,
		0x04040202, 0x04020204, 0x02020204, 0x02020202, 0x04040202, 0x02040404, 0x02020402, 0x02020402,

		0x02020204, 0x04020204, 0x02040202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x04040202, 0x04020202, 0x02020204, 0x02040404, 0x02040402, 0x04040202, 0x04020202, 0x02020204,
		0x02020404, 0x02040402, 0x02020202, 0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04040202,
		0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020202,
		0x02040202, 0x04020202, 0x02020202, 0x02020402, 0x02020204, 0x04020202, 0x02020402, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x04040202, 0x04020202, 0x02020204, 0x02020404,
		0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04020202, 0x04040404,
		0x02020204, 0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402,

		0x04040202, 0x04020202, 0x02020204, 0x02020202, 0x02040202, 0x04020202, 0x04040202, 0x02020404,
		0x02020204, 0x04020202, 0x02020404, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204,
		0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04040202,
		0x04020202, 0x02020204, 0x02020404, 0x02040402, 0x04040202, 0x04020202, 0x02020204, 0x02020202,
		0x02040202, 0x04020202, 0x02020402, 0x02020402, 0x02020204, 0x04020202, 0x02040202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x04020202, 0x04040204, 0x02020202, 0x02040402,
		0x02040402, 0x04020202, 0x04040204, 0x02020202, 0x02040402, 0x02020404, 0x04040202, 0x04040202,

		0x02020204, 0x02040402, 0x02020404, 0x04020202, 0x04040204, 0x02020202, 0x02040402, 0x02020404,
		0x04040202, 0x04020202, 0x02020204, 0x02020202, 0x02040202, 0x04020202, 0x02020402, 0x02020402,
		0x02020204, 0x04020204, 0x04020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02040404, 0x02020202, 0x04040202, 0x02020404, 0x02020202, 0x02040404, 0x02020202,
		0x04040202, 0x02020204, 0x04020202, 0x04020404, 0x02020404, 0x04040202, 0x02020204, 0x02020202,
		0x02040404, 0x02020202, 0x04040202, 0x02020204, 0x04040402, 0x04040204, 0x02020404, 0x02020202,
		0x04040402, 0x02040404, 0x04040202, 0x02040404, 0x04040402, 0x04040202, 0x04040204, 0x02020204,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,

		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
		0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202, 0x02020202,
	}
)
