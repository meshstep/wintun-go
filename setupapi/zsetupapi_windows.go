// Code generated by 'go generate'; DO NOT EDIT.

package setupapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modsetupapi = windows.NewLazySystemDLL("setupapi.dll")

	procSetupDiCreateDeviceInfoListExW    = modsetupapi.NewProc("SetupDiCreateDeviceInfoListExW")
	procSetupDiGetDeviceInfoListDetailW   = modsetupapi.NewProc("SetupDiGetDeviceInfoListDetailW")
	procSetupDiCreateDeviceInfoW          = modsetupapi.NewProc("SetupDiCreateDeviceInfoW")
	procSetupDiEnumDeviceInfo             = modsetupapi.NewProc("SetupDiEnumDeviceInfo")
	procSetupDiDestroyDeviceInfoList      = modsetupapi.NewProc("SetupDiDestroyDeviceInfoList")
	procSetupDiBuildDriverInfoList        = modsetupapi.NewProc("SetupDiBuildDriverInfoList")
	procSetupDiCancelDriverInfoSearch     = modsetupapi.NewProc("SetupDiCancelDriverInfoSearch")
	procSetupDiEnumDriverInfoW            = modsetupapi.NewProc("SetupDiEnumDriverInfoW")
	procSetupDiGetSelectedDriverW         = modsetupapi.NewProc("SetupDiGetSelectedDriverW")
	procSetupDiSetSelectedDriverW         = modsetupapi.NewProc("SetupDiSetSelectedDriverW")
	procSetupDiGetDriverInfoDetailW       = modsetupapi.NewProc("SetupDiGetDriverInfoDetailW")
	procSetupDiDestroyDriverInfoList      = modsetupapi.NewProc("SetupDiDestroyDriverInfoList")
	procSetupDiGetClassDevsExW            = modsetupapi.NewProc("SetupDiGetClassDevsExW")
	procSetupDiCallClassInstaller         = modsetupapi.NewProc("SetupDiCallClassInstaller")
	procSetupDiOpenDevRegKey              = modsetupapi.NewProc("SetupDiOpenDevRegKey")
	procSetupDiGetDeviceRegistryPropertyW = modsetupapi.NewProc("SetupDiGetDeviceRegistryPropertyW")
	procSetupDiSetDeviceRegistryPropertyW = modsetupapi.NewProc("SetupDiSetDeviceRegistryPropertyW")
	procSetupDiGetDeviceInstallParamsW    = modsetupapi.NewProc("SetupDiGetDeviceInstallParamsW")
	procSetupDiGetClassInstallParamsW     = modsetupapi.NewProc("SetupDiGetClassInstallParamsW")
	procSetupDiSetDeviceInstallParamsW    = modsetupapi.NewProc("SetupDiSetDeviceInstallParamsW")
	procSetupDiSetClassInstallParamsW     = modsetupapi.NewProc("SetupDiSetClassInstallParamsW")
	procSetupDiClassNameFromGuidExW       = modsetupapi.NewProc("SetupDiClassNameFromGuidExW")
	procSetupDiClassGuidsFromNameExW      = modsetupapi.NewProc("SetupDiClassGuidsFromNameExW")
	procSetupDiGetSelectedDevice          = modsetupapi.NewProc("SetupDiGetSelectedDevice")
	procSetupDiSetSelectedDevice          = modsetupapi.NewProc("SetupDiSetSelectedDevice")
)

func setupDiCreateDeviceInfoListEx(classGUID *windows.GUID, hwndParent uintptr, machineName *uint16, reserved uintptr) (handle DevInfo, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiCreateDeviceInfoListExW.Addr(), 4, uintptr(unsafe.Pointer(classGUID)), uintptr(hwndParent), uintptr(unsafe.Pointer(machineName)), uintptr(reserved), 0, 0)
	handle = DevInfo(r0)
	if handle == DevInfo(windows.InvalidHandle) {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetDeviceInfoListDetail(deviceInfoSet DevInfo, deviceInfoSetDetailData *_SP_DEVINFO_LIST_DETAIL_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetDeviceInfoListDetailW.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoSetDetailData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiCreateDeviceInfo(deviceInfoSet DevInfo, DeviceName *uint16, classGUID *windows.GUID, DeviceDescription *uint16, hwndParent uintptr, CreationFlags DICD, deviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetupDiCreateDeviceInfoW.Addr(), 7, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(DeviceName)), uintptr(unsafe.Pointer(classGUID)), uintptr(unsafe.Pointer(DeviceDescription)), uintptr(hwndParent), uintptr(CreationFlags), uintptr(unsafe.Pointer(deviceInfoData)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiEnumDeviceInfo(deviceInfoSet DevInfo, memberIndex uint32, deviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiEnumDeviceInfo.Addr(), 3, uintptr(deviceInfoSet), uintptr(memberIndex), uintptr(unsafe.Pointer(deviceInfoData)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiDestroyDeviceInfoList(deviceInfoSet DevInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiDestroyDeviceInfoList.Addr(), 1, uintptr(deviceInfoSet), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiBuildDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, driverType SPDIT) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiBuildDriverInfoList.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiCancelDriverInfoSearch(deviceInfoSet DevInfo) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiCancelDriverInfoSearch.Addr(), 1, uintptr(deviceInfoSet), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiEnumDriverInfo(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, driverType SPDIT, memberIndex uint32, driverInfoData *SP_DRVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiEnumDriverInfoW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType), uintptr(memberIndex), uintptr(unsafe.Pointer(driverInfoData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, driverInfoData *SP_DRVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetSelectedDriverW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(driverInfoData)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiSetSelectedDriver(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, driverInfoData *SP_DRVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetSelectedDriverW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(driverInfoData)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetDriverInfoDetail(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, driverInfoData *SP_DRVINFO_DATA, driverInfoDetailData *_SP_DRVINFO_DETAIL_DATA, driverInfoDetailDataSize uint32, requiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiGetDriverInfoDetailW.Addr(), 6, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(driverInfoData)), uintptr(unsafe.Pointer(driverInfoDetailData)), uintptr(driverInfoDetailDataSize), uintptr(unsafe.Pointer(requiredSize)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiDestroyDriverInfoList(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, driverType SPDIT) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiDestroyDriverInfoList.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(driverType))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetClassDevsEx(classGUID *windows.GUID, Enumerator *uint16, hwndParent uintptr, Flags DIGCF, deviceInfoSet DevInfo, machineName *uint16, reserved uintptr) (handle DevInfo, err error) {
	r0, _, e1 := syscall.Syscall9(procSetupDiGetClassDevsExW.Addr(), 7, uintptr(unsafe.Pointer(classGUID)), uintptr(unsafe.Pointer(Enumerator)), uintptr(hwndParent), uintptr(Flags), uintptr(deviceInfoSet), uintptr(unsafe.Pointer(machineName)), uintptr(reserved), 0, 0)
	handle = DevInfo(r0)
	if handle == DevInfo(windows.InvalidHandle) {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiCallClassInstaller(installFunction DI_FUNCTION, deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiCallClassInstaller.Addr(), 3, uintptr(installFunction), uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiOpenDevRegKey(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, Scope DICS_FLAG, HwProfile uint32, KeyType DIREG, samDesired uint32) (key windows.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procSetupDiOpenDevRegKey.Addr(), 6, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(Scope), uintptr(HwProfile), uintptr(KeyType), uintptr(samDesired))
	key = windows.Handle(r0)
	if key == windows.InvalidHandle {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, property SPDRP, propertyRegDataType *uint32, propertyBuffer *byte, propertyBufferSize uint32, requiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procSetupDiGetDeviceRegistryPropertyW.Addr(), 7, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(property), uintptr(unsafe.Pointer(propertyRegDataType)), uintptr(unsafe.Pointer(propertyBuffer)), uintptr(propertyBufferSize), uintptr(unsafe.Pointer(requiredSize)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiSetDeviceRegistryProperty(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, property SPDRP, propertyBuffer *byte, propertyBufferSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiSetDeviceRegistryPropertyW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(property), uintptr(unsafe.Pointer(propertyBuffer)), uintptr(propertyBufferSize), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, deviceInstallParams *_SP_DEVINSTALL_PARAMS) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetDeviceInstallParamsW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(deviceInstallParams)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiGetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, classInstallParams *SP_CLASSINSTALL_HEADER, classInstallParamsSize uint32, requiredSize *uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiGetClassInstallParamsW.Addr(), 5, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(classInstallParams)), uintptr(classInstallParamsSize), uintptr(unsafe.Pointer(requiredSize)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiSetDeviceInstallParams(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, deviceInstallParams *_SP_DEVINSTALL_PARAMS) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetDeviceInstallParamsW.Addr(), 3, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(deviceInstallParams)))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiSetClassInstallParams(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA, classInstallParams *SP_CLASSINSTALL_HEADER, classInstallParamsSize uint32) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiSetClassInstallParamsW.Addr(), 4, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), uintptr(unsafe.Pointer(classInstallParams)), uintptr(classInstallParamsSize), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiClassNameFromGuidEx(classGUID *windows.GUID, className *uint16, classNameSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiClassNameFromGuidExW.Addr(), 6, uintptr(unsafe.Pointer(classGUID)), uintptr(unsafe.Pointer(className)), uintptr(classNameSize), uintptr(unsafe.Pointer(requiredSize)), uintptr(unsafe.Pointer(machineName)), uintptr(reserved))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiClassGuidsFromNameEx(className *uint16, classGuidList *windows.GUID, classGuidListSize uint32, requiredSize *uint32, machineName *uint16, reserved uintptr) (err error) {
	r1, _, e1 := syscall.Syscall6(procSetupDiClassGuidsFromNameExW.Addr(), 6, uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(classGuidList)), uintptr(classGuidListSize), uintptr(unsafe.Pointer(requiredSize)), uintptr(unsafe.Pointer(machineName)), uintptr(reserved))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func setupDiGetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiGetSelectedDevice.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func SetupDiSetSelectedDevice(deviceInfoSet DevInfo, deviceInfoData *SP_DEVINFO_DATA) (err error) {
	r1, _, e1 := syscall.Syscall(procSetupDiSetSelectedDevice.Addr(), 2, uintptr(deviceInfoSet), uintptr(unsafe.Pointer(deviceInfoData)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
