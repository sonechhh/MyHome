package device

type Devices struct {
	DeviceArr []Device
}

func (devices Devices) DevicesInfo() string {
	resString := ""
	if len(devices.DeviceArr) > 0 {
		resString = "\t\tУстройства:" + "\n"
	}
	for i, device := range devices.DeviceArr {
		resString += device.getDeviceString()
		if (i + 1) != len(devices.DeviceArr) {
			resString += "\t\t\t------------------------\n"
		}
	}
	return resString
}
